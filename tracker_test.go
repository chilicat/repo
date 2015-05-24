package main

import (
  "testing"
  "fmt"
  "time"
)

func TestTracker(t *testing.T) {
    fmt.Printf("Tracker\n")
    conf := Config {
        "http://central.maven.org/maven2",
        "/tmp/down-test",
        "{{.Id}}-{{.Version}}.{{.Ext}}",
        true,
        5,
    }

    tracker := NewTracker(conf)
    tracker.Request(Artifact { "commons-lang", "commons-lang", "2.4", "jar", "jar"})
    tracker.Request(Artifact { "org.hibernate", "hibernate-core", "4.3.9.Final", "jar", "jar" } )
    tracker.Request(Artifact { "org.apache.ant", "ant" ,"1.9.4", "jar", "jar" } )
    tracker.Request(Artifact { "org.springframework", "spring-core", "4.1.6.RELEASE", "jar", "jar" } )

    go func() {
        time.Sleep(100)
        //tracker.Request(Artifact { "commons-io", "commons-io", "2.6", "jar", "jar"})
        tracker.Request(Artifact { "commons-lang", "commons-lang", "2.4", "jar", "jar"})
    }()

    err := tracker.Wait()
    if err != nil {
        fmt.Printf("[ERROR] %s\n", err)
    }
}






