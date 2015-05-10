package main

import (
	"bytes"
	"errors"
	"strings"
	"text/template"
)

type Artifact struct {
	Group, Id, Version, Class, Ext string
}

func ParseArtifact(a string) (Artifact, error) {
	tokens := strings.Split(a, ":")
	if len(tokens) == 3 {
		return Artifact{tokens[0], tokens[1], tokens[2], "", "jar"}, nil
	} else if len(tokens) == 4 {
		return Artifact{tokens[0], tokens[1], tokens[2], tokens[3], "jar"}, nil
	} else if len(tokens) == 5 {
		return Artifact{tokens[0], tokens[1], tokens[2], tokens[3], tokens[4]}, nil
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
