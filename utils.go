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

func ToSnapshotUrl(base string, a Artifact, timestamp string) string {
	v := strings.Replace(a.Version, "-SNAPSHOT", "", -1)
	s := base + "/" + ToPath(a) + "/" + a.Id + "-" + v + "." + timestamp + "." +  a.Ext
	return s
}

func ToMetadataUrl(base string, a Artifact) string {
	return base + "/" + ToPath(a) + "/maven-metadata.xml" 
}
