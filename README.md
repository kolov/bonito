# Bonito

## Description

Bonito lives in the Digital Ocean. Some services on DO may be needed temporarily, e.g. during working hours, 
and not in the weekend or overnight. 
Bonito helps you to shutdown a droplet, snapshot the data and destoy it,
paying no cloud costs until the next droplet is started rom the last snapshot.   

All this can be done on the digital Oceans' Dashboard, 
with a few clicks and a lot of patience - the worse is waiting up to 15 minutes for a snapshot to be taken,
before deletin the droplet. It's easy to loose patience and forget to destory the droplet at the end.
Bonito waits untill as long as needed and does not loose patience.
Besides, I was cusious to do something in Go.

## Usage

BE CAREFUL, this is in development. Unless --force is specified, all destructive actions ask for confirmation. I use 
it myself for a couple of days now. 

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
$ go get -d github.com/kolov/bonito
```
