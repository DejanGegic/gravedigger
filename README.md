# Gravedigger - Dig up dead Go code

## Intro

**Gravedigger** is a small piece of software that parses your Go source code and reports unused functions. As I could not find a simple and reliable tool, I decided to write my own.  

## Installation

I have pre-compiled binaries for Windows, macOS and Linux. Please rename them to `gravedigger` after proceeding if you do not plan on compiling them yourself.

<!-- ## Go Installation (Required to have Go installed on your system)

This works on all platforms.  
Run `go install github.com/dejangegic/gravedigger@latest` in the terminal.  
Make sure that the go bin is in your PATH. -->

### Run binaries directly

Download binaries from "releases" on this github page and run it locally as `./gravedigger`.

### Copy to PATH (no Go installation required)

#### Linux

Run `sudo cp gravedigger /usr/local/bin/gravedigger` to install it to the path.

#### macOS

Run `sudo cp gravedigger /usr/local/bin/gravedigger` to install it to the path.  
***Note**: This has not been tested as I do not own a mac. But it should work*

#### Windows

Copy the `.exe` to Program Files and add them it to your PATH. If you do not know how to do this, then maybe you shouldn't run unverified binaries on your machine 🤷.

## How to use?

You can use it in the same directory by running it with no arguments.  
`gravedigger`

Or use it on a specific directory by providing the path as an argument.  
`gravedigger /path/to/project`

## Limitations, Bugs, and Feature Requests

1. **Gravedigger** Should NOT be used on parent directory that consists of multiple projects.
Scan one project at a time or it may not work properly. This is fine for now, and I do not plan on changing it in the future.

2. Having the option to display the number of scanned functions and lines i sa simple feature that will be added soon in the future. But I'll keep it "clean" for now.

3. The ability to pass function names to ignore, as well as certain files, and even whole directories, would be a very nice feature. As Gravedigger is a tool developed only for my personal use, I'll add it if I see that it's useful to at least a couple of developers besides myself.
