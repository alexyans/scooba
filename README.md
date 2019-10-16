Scooba
======

Scooba lets you explore the history of a new codebase starting with the earliest available commit and navigating forward at your own pace. It is meant to be used for educational purposes or for onboarding onto a new project, giving you the smallest possible incremental changes to go through one chunk at a time, along with the related technical decisions that were made.

You can start from the oldest commit by default, or specify a target commit as the starting point. From there, you can move forward or backward in history. At every step, the changes introduced since the last visited commit appear in git's index, which makes them `git diff`-able, and allows them to be piped into visual diff tools or other code visualization tools. 

## Installation

You will need to have **Docker** and **make** on your host.

Running `make` will set up the Docker environment and produce the `scooba` binary. 

## Running

Scooba depends on a couple of shared libaries that are produced during the `make build` step for easier development. To force the linker to use the version in the `lib/` directory, you need to call the binary like so:

````LD_LIBRARY_PATH=/path/to/lib/:$LD_LIBRARY_PATH scooba````

## Command reference

`scooba [dive | d] [-c <commit_hash>]` - Checks out the oldest historical commit and sets the stage for everything else

- `-c, --commit <commit_hash>` : Start from a target commit instead of the chronologically oldest

`scooba [forward | f]` - Visit the next commit in topological order (children of the current commit)

`scooba [backward | b]` - Visit the previous commit in topological order (parent of the current commit)

## Roadmap
- Associate commit range to PR and Issue threads
- Extend navigation
- Handle repository migration
- Make default step-through smarter (for instance, skip over commits that fix typos or edit the readme)
