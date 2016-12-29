package main

import (
	"encoding/json"
	"github.com/jawher/mow.cli"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Tab struct {
	Description, DevtoolsFrontendUrl, FaviconUrl, Id, Title, Type, Url, WebSocketDebuggerUrl string
}

func main() {
	app := cli.App("chrome-remote-reload", "Chrome Remote Debugging Page Reload tool\n\nFirst, run Chrome remote debugger with: chrome --remote-debugging-port=9222\n")
	app.Spec = "[-h=<host>] [-p=<port>]"

	var (
		host = app.StringOpt("h host", "localhost", "Chrome remote debugger host")
		port = app.StringOpt("p port", "9222", "Chrome remote debugger port")
	)

	app.Action = func() {
		log.Printf("Init listing ...")

		tab := getFirstTab(*host, *port)
		log.Printf("Connect on WS ... %s", tab.WebSocketDebuggerUrl)

		response := refreshTab(*host, tab)
		log.Printf("Response ... %s", response)

		os.Exit(0)
	}

	app.Run(os.Args)
}

func refreshTab(host string, tab Tab) string {
	var msg = make([]byte, 1024)
	var origin = "http://" + host + "/"

	ws, err := websocket.Dial(tab.WebSocketDebuggerUrl, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	var cmd = ""
	if tab.Type == "page" {
		cmd = "{\"id\": 0, \"method\": \"Page.reload\"}"
	} else {
		cmd = "{\"id\": 0, \"method\": \"Runtime.evaluate\", \"params\": {\"expression\": \"WebInspector.Main._reloadPage(true)\"}}"
	}

	if _, err := ws.Write([]byte(cmd)); err != nil {
		log.Fatal(err)
	}

	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}

	return string(msg[:n])
}

func getFirstTab(host string, port string) Tab {
	var tabs []Tab

	res, err := http.Get("http://" + host + ":" + port + "/json")
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonBytes, &tabs)
	if err != nil {
		log.Fatal(err)
	}

	if len(tabs) <= 0 {
		log.Fatal("No tab found")
	}

	return tabs[0]
}
