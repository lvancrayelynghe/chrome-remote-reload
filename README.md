# Chrome Remote Debugging Page Reload

Simple tool to reload a Chrome tab from command line, using Chrome's remote debugging protocol

## Usage

```
$ chrome-remote-reload --help

Usage: chrome-remote-reload [-h=<host>] [-p=<port>]

Chrome Remote Debugging Page Reload tool

First, run Chrome remote debugger with: chrome --remote-debugging-port=9222


Options:
  -h, --host="localhost"   Chrome remote debugger host
  -p, --port="9222"        Chrome remote debugger port
```


## Installation from binaries

See the [GitHub releases](https://github.com/Benoth/chrome-remote-reload/releases)


## Installation from sources

### Install Go

```
sudo apt-get install curl git mercurial make binutils bison gcc build-essential
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source ~/.gvm/scripts/gvm
gvm install go1.7.4 --prefer-binary
gvm use go1.7.4 --default
```

### Install packages

```
go get github.com/jawher/mow.cli
go get golang.org/x/net/websocket
```

### Cross compile

Use gox :

```
go get github.com/mitchellh/gox
gox
```
