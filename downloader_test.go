package main

import (
  "testing"
)

func TestDownloader(t *testing.T) {
    conf := Config {
        "http://central.maven.org/maven2",
        "/tmp/down-test",
        "{{.Id}}-{{.Version}}.{{.Ext}}",
        true,
        5,
    }

    down := NewDownloader(conf)
    asset(t, 5, down.num, "num")

    consume := make(chan DownloadResult, 100)

    req1 := DownloadRequest { 
    	Artifact { "commons-lang", "commons-lang", "2.4", "jar", "jar"},
        false,
    	consume,
    }

    req2 := DownloadRequest { 
        Artifact { "commons-lang", "commons-lang", "2.4", "pom", "pom"},
        false,
        consume,
    }

    down.Request(req1)
    down.Request(req2)

    num := 0
    for res := range consume {
		assetError(t, res.err, "")
		num++
		if num == 2 {
			close(consume)
		}
    }
}



