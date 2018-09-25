# grafctl

[![Travis CI](https://travis-ci.org/dimitrovvlado/grafctl.svg?branch=master)](https://travis-ci.org/dimitrovvlado/grafctl)

Command-line utility for managing Grafana

## Installation

### Linux

```sh
curl -L https://github.com/dimitrovvlado/grafctl/releases/download/$(curl -s https://raw.githubusercontent.com/dimitrovvlado/grafctl/master/VERSION.txt)/grafctl-linux-amd64 -o /usr/local/bin/grafctl && chmod +x /usr/local/bin/grafctl
```

### MacOS

```sh
curl -L https://github.com/dimitrovvlado/grafctl/releases/download/$(curl -s https://raw.githubusercontent.com/dimitrovvlado/grafctl/master/VERSION.txt)/grafctl-darwin-amd64 -o /usr/local/bin/grafctl && chmod +x /usr/local/bin/grafctl
```

## Features

* Written in portable go, binary is free of dependencies
* Import/export of dashboards
* Create/Read/Delete datasources
* Read orgs, teams, users

## Usage

```
Usage:
  grafctl [flags]
  grafctl [command]

Available Commands:
  config      Configure this command line tool
  create      Create a resource
  delete      Delete a resource by name or id
  get         Display one or many resources
  help        Help about any command
  version     Print the version number of grafctl

Flags:
  -h, --help      help for grafctl
  -v, --verbose   Verbose output

Use "grafctl [command] --help" for more information about a command.
```
