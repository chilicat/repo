package main

import (
	"strings"
)

func ToPath(a Artifact) string {
	groupPath := strings.Join(strings.Split(a.Group, "."), "/")
	return groupPath + "/" + a.Id + "/" + a.Version
}

func ToUrl(base string, a Artifact) string {
	return base + "/" + ToPath(a) + "/" + a.Id + "-" + a.Version + "." + a.Ext
}
