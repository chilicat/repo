package main

import (
	"testing"
)

func TestMetadataParse(t *testing.T) {
	data := `
        <metadata modelVersion="1.1.0">
  <groupId>org.hibernate</groupId>
  <artifactId>hibernate-core</artifactId>
  <version>4.3.9-SNAPSHOT</version>
  <versioning>
    <snapshot>
      <timestamp>20150408.141831</timestamp>
      <buildNumber>1</buildNumber>
    </snapshot>
    <lastUpdated>20150408141831</lastUpdated>
  </versioning>
</metadata>
`

	err, v := ParseMDFromString(data)
	if err != nil {
		t.Errorf("Parsing failed: %q", err)
	}

	asset(t, "org.hibernate", v.GroupId, "MD.Group")
	asset(t, "hibernate-core", v.ArtifactId, "MD.ArtifactId")
	asset(t, "4.3.9-SNAPSHOT", v.Version, "MD.Version")
  asset(t, "20150408.141831", v.Snapshot.Timestamp, "MD.Snapshot.Version")

	/*art := v.Dependencies.ToArtifacts("compile")

	asset(t, 2, len(art), "Dependency count")

	logging := art[0]
	asset(t, "org.jboss.logging", logging.Group, "Group")
	asset(t, "jboss-logging", logging.Id, "Id")
	asset(t, "3.1.3.GA", logging.Version, "Version")
  */
}
