# Configr

## Introduction

On a system, there are several configuration files and they are all stored in different locations. In order to version these files, the user would need to perform several manual steps or create a shell script to copy them to a location. The purpose of `configr` is to provide a simple binary which allows the user several tools in order to effectively manage and version their configuration files, and share them across multiple systems via a git repo.

## Dependencies

Currently there is only one dependency, which is `vim`. This is the default editor that is used by the application. In time, there will be other editors supported. The application can still be used without `vim` being installed on the system, though all `file edit` subcommands will fail.

## Usage

```
NAME:
   configr - Configuration file management utility

USAGE:
   configr [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   file, f        Interact with files in the inventory
   collection, c  interact with the collection
   config, conf   Edit configr configuration settings
   edit           Open a file in vim
   inv            Edit the inventory file
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Features

### Inventory

The inventory file is a JSON file which the user can specify the configuration files managed by the system. The inventory editted by running `configr inv`.

The user may add aliases to files by editing the inventory file. Aliases on files can be viewed by running `configr file ls`. 

The user may add, remove, and list files by running the `configr file` subcommand and selecting an option from there. 

The user may edit a file managed in the inventory by running `configr edit` or `configr file edit` and then supplying either the alias to the file, the full path to the file, or just the filename.

**NOTE**: Files that end in `config`, `config.json`, and `config.yml` must also contain the preceding directory name if editing with just the filename.

### Collection

The user is able to specify a directory for their collection in their configr main configuration file (which can be accessed by running `configr config`). 

Collections are intended to be paths to `git` repos containing dotfiles/configuration files. This is intended to allow the user to effectively version their configuration files.

The collection subcommand does not directly interface with `git`, so pushes, pulls, branches, and other git features will need to be run by the user directly.

