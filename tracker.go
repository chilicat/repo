package main

import (
	"sync/atomic"
)

var counter int64 = 0

type Tracker struct {
	conf    Config
	wait    chan error
	request chan Artifact
}

func (t Tracker) Request(artifact Artifact) {
	up()
	t.request <- artifact
}

func (t Tracker) Wait() error {
	for err := range t.wait {
		return err
	}
	return nil
}

func up() int64 {
	return atomic.AddInt64(&counter, 1)
}

func down() int64 {
	return atomic.AddInt64(&counter, -1)
}

func NewTracker(conf Config) Tracker {
	wait := make(chan error, 1)
	requests := make(chan Artifact, 100)

	pomResults := make(chan DownloadResult, 100)
	results := make(chan DownloadResult, 100)

	t := Tracker{conf, wait, requests}

	downloader := NewDownloader(conf)

	go func() {
		for res := range pomResults {
			err, pom := ParsePomFromString(res.content)
			if err != nil {
				wait <- res.err
			}
			for _, a := range pom.Dependencies.ToArtifacts("compile") {
				t.Request(a)
			}

			if down() <= 0 {
				wait <- nil
			}
		}
	}()

	go func() {
		set := make(map[string]bool)
		for artifact := range requests {
			if !set[artifact.String()] {
				set[artifact.String()] = true
				if !artifact.IsPom() {
					up()
					pomArtifact := artifact.Pom()
					downloader.Request(DownloadRequest{
						pomArtifact,
						true,
						pomResults,
					})
				}
				downloader.Request(DownloadRequest{
					artifact,
					false,
					results,
				})
			} else {
				down()
			}
		}
	}()

	go func() {
		for res := range results {
			if res.err != nil {
				wait <- res.err
			} else if down() <= 0 {
				wait <- nil
			}
		}
	}()

	return t
}
