# Bonito

## Description

Bonito lives in the Digital Ocean. Some services may be needed temporarily, e.g. during working hours, 
and not in the weekend or overnight. Bonito helps you to shutdown the servers when you do not need them
and spin them up when needed again, paying no cloud costs in the meanwhile. 

All this can be done on the digital Oceans' 
Dashboard, 
with a few clicks and some patience. Bonito automates this. Besides, I wanted to do something in Go.

## Usage

DON'T USE, this is in development.

Start with `bonito --help`

    COMMANDS:
        sometest  used during develpment for random tests. Ignore.
        list      list all snapshots or droplets
        shutdown
        up        starts a droplet from a snapshot
        help, h   Shows a list of commands or help for one command
   
    GLOBAL OPTIONS:
      --token value, -t value  Authentication token. Must be provided here or as [$DO_TOKEN_BONITO]
      --help, -h               show help
      --version, -v            print the version
      
First, you need a Digital Ocean authorization token. See 
[https://cloud.digitalocean.com/settings/api/tokens](https://cloud.digitalocean.com/settings/api/tokens) dor details.
Put the token in environment variable DO_TOKEN_BONITO (`export DO_TOKEN_BONIto=A7f9...`) or provide in the command line
with the option --token (`bonito --token A7f9... ...`).

To start a server from an existing snapshot:

    bonito up --template {regex} [--latest] [--name mydroplet] [--verbose] -keys mykeyname
    
To shut it down:

    bonito down --template={regex}
      
## Install

To install, use `go get`:

```bash
$ go get -d github.com/kolov/sardine
```

## Contribution

1. Fork ([https://github.com/kolov/bonito/fork](https://github.com/kolov/bonito/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[kolov](https://github.com/kolov)
