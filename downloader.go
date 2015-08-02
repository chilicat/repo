package main

import (
	"path"
	"fmt"
)

type DownloadRequest struct {
	artifact Artifact
	inMemory bool
	consumer chan DownloadResult
}

type DownloadResult struct {
	info    ArtifactInfo
	err     error
	content string
}

type Downloader struct {
	num     int
	conf    Config
	request chan DownloadRequest
}

func (d Downloader) Request(r DownloadRequest) {
	d.request <- r
}

func NewDownloader(conf Config) Downloader {
	num := conf.WorkersCount
	request := make(chan DownloadRequest, num)
	d := Downloader{num, conf, request}
	for i := 0; i < num; i++ {
		go func() {
			for req := range d.request {
				filename, _ := ToFileName(req.artifact, d.conf.Template)
				info := ArtifactInfo{
					req.artifact,
					ToUrl(d.conf.BaseUrl, req.artifact),
					path.Join(d.conf.Dest, filename),
				}
				var err error = nil
				var content string = ""
				if !d.conf.DryRun {
					if req.inMemory {
						content, err = DownloadAsString(info.url)
					} else {
						if info.artifact.IsSnapshot() {
							content, err = DownloadAsString(ToMetadataUrl(d.conf.BaseUrl, req.artifact))
							err, md := ParseMDFromString(content)
							if err == nil {
								info.url = ToSnapshotUrl(d.conf.BaseUrl, req.artifact, md.Snapshot.Timestamp)
								fmt.Printf("======= %s\n", info.url)
							} 
						}
						err = Download(info.url, info.destFile)
					}
				}
				
				result := DownloadResult{
					info,
					err,
					content,
				}
				req.consumer <- result
			}
		}()
	}
	return d
}
