package main

import (
	"strings"
)

func ToPath(a Artifact) string {
	return strings.Replace(a.Group, ".", "/", 0) + "/" + a.Id + "/" + a.Version
}

func ToUrl(base string, a Artifact) string {
	return base + "/" + ToPath(a) + "/" + a.Id + "-" + a.Version + "." + a.Ext
}
