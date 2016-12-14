# Bonito

## Description

Bonito lives in the Digital Ocean. Some services on DO may be needed temporarily, e.g. during working hours, 
and not in the weekend or overnight. 
Bonito helps you to shutdown a droplet, snapshot the data and release it,
paying no cloud costs until you need them again. Start a new, possibly bigger droplet from the snapshotnext day.

All this can be done on the digital Oceans' Dashboard, 
with a few clicks and a lot of patience - the worse is waiting up to 15 minutes for a snapshot to be taken,
then delete the droplet. It's easy to loose patience and forget do perform the key action - destorying the droplet.
Bonito waits untill the snapshot is ready and not loose patience.
Besides, I wanted to do something in Go.

## Usage

BE CAREFUL, this is in development. Unless --force is specified, all destructive actions ask for confirmation. 

Start with `bonito --help`

    NAME:
       bonito
    
    USAGE:
       bonito [global options] command [command options] [arguments...]
    
    VERSION:
       0.1.0
    
    COMMANDS:
         list     lists all snapshots, droplets or keys
         letgo    Let a droplet go -shutdown, snapshot and destroy
         up       starts a droplet from a snapshot
         help, h  Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
       --token value, -t value  Authentication token. Must be provided here or as [$DO_TOKEN_BONITO]
       --help, -h               show help
       --version, -v            print the version
      
First, you need a Digital Ocean authorization token. See 
[https://cloud.digitalocean.com/settings/api/tokens](https://cloud.digitalocean.com/settings/api/tokens) dor details.
Put the token in environment variable DO_TOKEN_BONITO (`export DO_TOKEN_BONITO=A7f9...`) or provide in the command line
with the option --token (`bonito --token A7f9... ...`).

Each command has its own help, e.g.:

    rife:bonito assen$ ./bonito up --help

To start a server from an existing snapshot:

    bonito up --snapshot jenkins [--latest] --name jenkins [--verbose] -keys assen
 
    ... work ...
    [bonito] created from image [jenkins-2016-12-12-12:33] statuse=[active] ip=188.166.5.244
    
To shut it down:

    bonito letgo --name jenkins --nosnapshot --force
    
Besides:
 
     bonito list droplets|snapshots|keys
      
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
