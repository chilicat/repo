# gomaven

gomaven is a command line tool to download artifacts from a maven or nexus repository. 

## Get started

Download the precompiled binary for your platform here: https://github.com/chilicat/gomaven/releases/tag/v0.1.0

gomaven has a self self-documented command line interface. The best way to figure out what you can do with the tool is calling:

```
gomaven help

NAME:
   gomaven - to deal with maven repositories

USAGE:
   gomaven [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   download, d	Downloads artifact from maven repository.
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

And from there you can drill into subcommands:


```
gomaven help download

```


## Download a artifact 

gomaven makes it easy to download artifacts from a maven repository by providing the typical artifact coordinations GROUP_ID:ID:VERSION. The example below downlods the commons-io jar file in version 2.4:

```
gomaven download -a commons-io:commons-io:2.4

Download: http://repo1.maven.org/maven2/commons-io/commons-io/2.4/commons-io-2.4.jar -> ./commons-io-2.4.jar
```

gomaven checks also the md5 checksum of the file before downloading. If you execute the command above a second time you can see that the tool actually skipps the download:


```
gomaven download -a commons-io:commons-io:2.4

Download Skipped (up-to-date) -> ./commons-io-2.4.jar
```

You can also download multiple artifacts:


```
gomaven download -a -a commons-io:commons-io:2.4,commons-lang:commons-lang:2.4

```


See help for options:

```
gomaven download help

USAGE:
   command download [command options] [arguments...]

OPTIONS:
   --artifact, -a 					Defines artifact coordinates to download (e.g.: commons-io:commons-io:2.4)
   --destination, -d "."				Defines download destination. [$MVN_DEST]
   --baseUrl, -b "http://repo1.maven.org/maven2"	Defines maven base repo url [$MVN_BASE_URL]
   --template, -t "{{.Id}}-{{.Version}}.{{.Ext}}"	Defines file template [$MVN_FILE_TMPL]
   --verbose, -v					Enables verbose output.

```

## Build

In order to build the project you have to get codegangsta cli

```
go get github.com/codegangsta/cli
```

## Roadmap

- Add support for downloading latest artifact 
- Add support for downloading recursive artifacts.
- Add support for uploading artifacts

## Thanks to:

- codegangsta: https://github.com/codegangsta/cli
