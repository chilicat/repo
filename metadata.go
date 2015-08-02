package main

import (
	"encoding/xml"
	"os"
	"strings"
)

type Snapshot struct {
	XMLName    xml.Name        `xml:"snapshot"`
	Timestamp string `xml:"timestamp"`
}

type Metadata struct {
	XMLName      xml.Name        `xml:"metadata"`
	GroupId      string          `xml:"groupId"`
	ArtifactId   string          `xml:"artifactId"`
	Version      string          `xml:"version"`
	Snapshot Snapshot `xml:"versioning>snapshot"`
}

func ParseMDFromString(xmlStr string) (error, Metadata) {
	v := Metadata{}
	err := xml.NewDecoder(strings.NewReader(xmlStr)).Decode(&v)
	return err, v
}

func ParseMDFromFile(file string) (error, Metadata) {
	v := Metadata{}
	xmlFile, err := os.Open(file)
	if err == nil {
		defer xmlFile.Close()
		err = xml.NewDecoder(xmlFile).Decode(&v)
	}
	return err, v
}
