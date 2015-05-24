package main

import (
	"testing"
)

func TestPomParse(t *testing.T) {
	data := `
        <?xml version="1.0" encoding="UTF-8"?>
        <project xmlns="http://maven.apache.org/POM/4.0.0" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
          <modelVersion>4.0.0</modelVersion>
          <groupId>org.hibernate</groupId>
          <artifactId>hibernate-core</artifactId>
          <version>4.3.9.Final</version>
          <dependencies>
            <dependency>
              <groupId>org.jboss.logging</groupId>
              <artifactId>jboss-logging</artifactId>
              <version>3.1.3.GA</version>
              <scope>compile</scope>
            </dependency>
            <dependency>
              <groupId>org.jboss.logging</groupId>
              <artifactId>jboss-logging-annotations</artifactId>
              <version>1.2.0.Beta1</version>
              <scope>compile</scope>
            </dependency>
          </dependencies>
        </project>
        `

	err, v := ParsePomFromString(data)
	if err != nil {
		t.Errorf("Parsing failed: %q", err)
	}

	asset(t, "org.hibernate", v.GroupId, "Pom.Group")
	asset(t, "hibernate-core", v.ArtifactId, "Pom.ArtifactId")
	asset(t, "4.3.9.Final", v.Version, "Pom.Version")

	art := v.Dependencies.ToArtifacts("compile")

	asset(t, 2, len(art), "Dependency count")

	logging := art[0]
	asset(t, "org.jboss.logging", logging.Group, "Group")
	asset(t, "jboss-logging", logging.Id, "Id")
	asset(t, "3.1.3.GA", logging.Version, "Version")
}
