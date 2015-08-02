# repo
repo  is a command line tool to download artifacts from a maven or nexus repository. 

## Get started

Download the precompiled binary for your platform here: https://github.com/chilicat/repo/releases/tag/v0.1.0

repo has a self self-documented command line interface. The best way to figure out what you can do with the tool is calling:

```
repo help

NAME:
   repo - to deal with maven repositories

USAGE:
   repo [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   pull, d  Downloads artifact from maven repository.
   help, h  Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

And from there you can drill into subcommands:


```
repo help pull

```


## Download a artifact 

repo makes it easy to download artifacts from a maven repository by providing the typical artifact coordinations GROUP_ID:ID:VERSION. The example below downlods the commons-io jar file in version 2.4:

```
repo pull commons-io:commons-io:2.4

Download: http://repo1.maven.org/maven2/commons-io/commons-io/2.4/commons-io-2.4.jar -> ./commons-io-2.4.jar
```

repo checks also the md5 checksum of the file before downloading. If you execute the command above a second time you can see that the tool actually skipps the download:


```
repo pull commons-io:commons-io:2.4

Download Skipped (up-to-date) -> ./commons-io-2.4.jar
```

You can also download multiple artifacts:


```
repo pull commons-io:commons-io:2.4 commons-lang:commons-lang:2.4

```


See help for options:

```
repo help pull

NAME:
   pull - Downloads artifact from maven repository.

USAGE:
   command pull [command options] [arguments...]

OPTIONS:
   --destination, -d "."            Defines download destination. [$MVN_DEST]
   --baseUrl, -b "http://central.maven.org/maven2" Defines maven base repo url [$MVN_BASE_URL]
   --template, -t "{{.Id}}-{{.Version}}.{{.Ext}}"  Defines file template [$MVN_FILE_TMPL]
   --verbose, -v              verbose output.
   --recursive, -r               Recursive download of pom dependencies
   --extension, -e "jar"            Default extension
```

## Build

In order to build the project you have to get codegangsta cli

```
go get github.com/codegangsta/cli
```

## Thanks to:

- codegangsta: https://github.com/codegangsta/cli
