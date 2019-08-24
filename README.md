Scooba
======

Scooba lets you explore the history of a new codebase starting with the earliest available commit and navigating forward at your own pace. It is meant to be used for educational purposes or for onboarding onto a new project, giving you the smallest possible incremental changes to go through one chunk at a time, along with the related technical decisions that were made.

## Installation

You will need to have **Docker** and **make** on your host.

Running `make` will set up the Docker environment and produce the `scooba` binary. 

## Running

Scooba depends on a couple of shared libaries that have been shipped within the repository for easier development. To force the linker to use the version in the `lib/` directory, you need to call the binary like so:

````LD_LIBRARY_PATH=`pwd`/lib/:$LD_LIBRARY_PATH ./scooba````

## Command reference

`scooba dive [-c <commit_hash>]` - Checks out the oldest historical commit and sets the stage for everything else

- `-c, --commit <commit_hash>` : Start from a target commit instead of the chronologically oldest