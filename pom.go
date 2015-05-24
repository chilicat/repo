package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type PomDependency struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
	Scope      string `xml:"scope"`
}

func (p PomDependency) String() string {
	return fmt.Sprintf("%s:%s:%s", p.GroupId, p.ArtifactId, p.Version)
}

type PomDependencies struct {
	XMLName    xml.Name        `xml:"dependencies"`
	Dependency []PomDependency `xml:"dependency"`
}

type Pom struct {
	XMLName      xml.Name        `xml:"project"`
	GroupId      string          `xml:"groupId"`
	ArtifactId   string          `xml:"artifactId"`
	Version      string          `xml:"version"`
	Dependencies PomDependencies `xml:"dependencies"`
}

func (p Pom) String() string {
	return fmt.Sprintf("%s:%s:%s -> %s", p.GroupId, p.ArtifactId, p.Version, p.Dependencies)
}

func (p PomDependencies) byScope(scope string) []PomDependency {
	found := make([]PomDependency, 0)
	for _, d := range p.Dependency {
		if d.Scope == scope {
			found = append(found, d)
		}
	}
	return found
}

func (p PomDependencies) ToArtifacts(scope string) []Artifact {
	deps := p.byScope(scope)
	list := make([]Artifact, len(deps), len(deps))
	for i, dep := range deps {
		list[i] = toArtifact(dep)
	}
	return list
}

func (p PomDependencies) ToArtifactsAll() []Artifact {
	deps := p.Dependency
	list := make([]Artifact, len(deps), len(deps))
	for i, dep := range deps {
		list[i] = toArtifact(dep)
	}
	return list
}

func toArtifact(dep PomDependency) Artifact {
	return Artifact{dep.GroupId, dep.ArtifactId, dep.Version, "", "jar"}
}

func ParsePomFromString(xmlStr string) (error, Pom) {
	v := Pom{}
	err := xml.NewDecoder(strings.NewReader(xmlStr)).Decode(&v)
	return err, v
}

func ParsePomFromFile(file string) (error, Pom) {
	v := Pom{}
	xmlFile, err := os.Open(file)
	if err == nil {
		defer xmlFile.Close()
		err = xml.NewDecoder(xmlFile).Decode(&v)
	}
	return err, v
}
