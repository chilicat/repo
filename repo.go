package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "repo"
	app.Usage = "to deal with maven repositories"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:      "pull",
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
					Value:  "http://central.maven.org/maven2",
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

func prepareArtifact(c *cli.Context) []Artifact {
	art := c.String("a")
	if art == "" {
		fatal("Please define artifact: -a commons-io:commons-io:2.4")
	}
	if c.Bool("verbose") {
		fmt.Println("[INFO] Artifact: " + art)
		fmt.Println("[INFO] Destination: " + c.String("d"))
		fmt.Println("[INFO] Template: " + c.String("t"))
		fmt.Println("[INFO] Base URL: " + c.String("b"))
	}
	artifacts := make([]Artifact, 0)
	for _, el := range strings.Split(art, ",") {
		artifact, err := ParseArtifact(el)
		checkError(err)
		artifacts = append(artifacts, artifact)
	}
	return artifacts
}

func DownloadCommand(c *cli.Context) {
	conf := Config{
		c.String("b"),
		c.String("d"),
		c.String("t"),
		false,
		5,
	}
	tracker := NewTracker(conf)
	for _, a := range prepareArtifact(c) {
		tracker.Request(a)
	}
	tracker.Wait()
}
