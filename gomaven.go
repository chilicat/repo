package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"path"
	"strings"
	"sync"
)

func main() {
	app := cli.NewApp()
	app.Name = "gomaven"
	app.Usage = "to deal with maven repositories"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:      "download",
			ShortName: "d",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "artifact,a",
					Usage: "Defines artifact coordinates to download (e.g.: commons-io:commons-io:2.4)",
					Value: "",
				},
				cli.StringFlag{
					Name:   "destination,d",
					Usage:  "Defines download destination.",
					EnvVar: "MVN_DEST",
					Value:  ".",
				},
				cli.StringFlag{
					Name:   "baseUrl,b",
					Usage:  "Defines maven base repo url",
					EnvVar: "MVN_BASE_URL",
					Value:  "http://repo1.maven.org/maven2",
				},
				cli.StringFlag{
					Name:   "template,t",
					Usage:  "Defines file template",
					EnvVar: "MVN_FILE_TMPL",
					Value:  "{{.Id}}-{{.Version}}.{{.Ext}}",
				},
				cli.BoolFlag{
					Name:  "verbose,v",
					Usage: "verbose output.",
				},
			},
			Usage:  "Downloads artifact from maven repository.",
			Action: DownloadCommand,
		},
	}
	app.Run(os.Args)
}

type ArtifactInfo struct {
	artifact Artifact
	url      string
	destFile string
}

type ErrorMsg struct {
	err  error
	kill bool
}

func prepareArtifact(c *cli.Context) []ArtifactInfo {
	art := c.String("a")
	dest := c.String("d")
	base := c.String("b")
	tmp := c.String("t")
	if art == "" {
		fatal("Please define artifact: -a commons-io:commons-io:2.4")
	}
	if c.Bool("verbose") {
		fmt.Println("[INFO] Artifact: " + art)
		fmt.Println("[INFO] Destination: " + dest)
		fmt.Println("[INFO] Template: " + tmp)
		fmt.Println("[INFO] Base URL: " + base)
	}
	artifacts := make([]ArtifactInfo, 0)
	for _, el := range strings.Split(art, ",") {
		artifact, err := ParseArtifact(el)
		checkError(err)
		filename, err := ToFileName(artifact, tmp)
		checkError(err)
		artifacts = append(artifacts, ArtifactInfo{
			artifact,
			ToUrl(base, artifact),
			path.Join(dest, filename),
		})
	}
	return artifacts
}

func DownloadCommand(c *cli.Context) {
	artifacts := prepareArtifact(c)

	var wg sync.WaitGroup
	wg.Add(len(artifacts))

	errorQueue := make(chan ErrorMsg, len(artifacts)+1)

	for _, el := range artifacts {
		go func(el ArtifactInfo) {
			defer wg.Done()
			err := Download(el.url, el.destFile)
			if err != nil {
				errorQueue <- ErrorMsg{err, false}
			}
		}(el)
	}

	// Wait for all downloads and send kill pill
	wg.Wait()
	errorQueue <- ErrorMsg{nil, true}

	var failed bool
	for err := range errorQueue {
		// stop loop if kill pile has been send.
		if err.kill {
			break
		}
		if failed == false {
			fmt.Println("----")
		}
		fmt.Printf("[ERR] %s\n", err.err)
		failed = true
	}
	if failed {
		os.Exit(1)
	}
}
