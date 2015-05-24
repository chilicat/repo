package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"
)

type Artifact struct {
	Group, Id, Version, Class, Ext string
}

type ArtifactInfo struct {
	artifact Artifact
	url      string
	destFile string
}

func (p Artifact) String() string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", p.Group, p.Id, p.Version, p.Class, p.Ext)
}

func (a Artifact) Pom() Artifact {
	return Artifact{a.Group, a.Id, a.Version, "pom", "pom"}
}

func (a Artifact) IsPom() bool {
	return a.Class == "pom" && a.Ext == "pom"
}

func ParseArtifact(a string) (Artifact, error) {
	first := strings.Split(a, "@")
	tokens := strings.Split(first[0], ":")

	ext := "jar"
	if len(first) > 1 {
		ext = first[1]
	}

	if len(tokens) == 3 {
		return Artifact{tokens[0], tokens[1], tokens[2], "", ext}, nil
	} else if len(tokens) == 4 {
		return Artifact{tokens[0], tokens[1], tokens[2], tokens[3], ext}, nil
	}
	return Artifact{}, errors.New("Artifact description is insufficent, minium: <group_id>:<id>:<version>")
}

func ToFileName(a Artifact, tmp string) (string, error) {
	tmpl, err := template.New("artifact-file").Parse(tmp)
	if err == nil {
		var b bytes.Buffer
		err = tmpl.Execute(&b, a)
		return b.String(), err
	}
	return "", err
}
