# modulegen

Generate modulefiles for use with lmod.

## Installation

To install go-style, first set your ```GOPATH``` environment variable, e.g:

    export GOPATH=$HOME/go

Then use:

    go get github.com/JIC-CSB/modulegen

The resulting binary will be:

    $GOPATH/bin/modulegen

## Use

Point at a linuxbrew install path, up to the version number, and modulegen will spit out a module file to stdout:

    $ ./modulegen /common/software/linuxbrew/Cellar/samtools/1.3/
    
