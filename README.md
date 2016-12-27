# Chrome Remote Debugging Page Reload

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
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source ~/.gvm/scripts/gvm
sudo apt-get install bison
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.7 --prefer-binary
gvm use go1.7 --default
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
