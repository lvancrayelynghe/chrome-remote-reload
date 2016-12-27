package main

import (
    "os"
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "github.com/jawher/mow.cli"
    "golang.org/x/net/websocket"
)


func main() {
    app := cli.App("chrome-remote-reload", "Chrome Remote Debugging Page Reload tool\n\nFirst, run Chrome remote debugger with: chrome --remote-debugging-port=9222\n")
    app.Spec = "[-h=<host>] [-p=<port>]"

    var (
        host = app.StringOpt("h host", "localhost", "Chrome remote debugger host")
        port = app.StringOpt("p port", "9222",      "Chrome remote debugger port")
    )

    app.Action = func() {
        log.Printf("Init listing ...")

        wsUrl := getWebSocketDebuggerUrl(*host, *port)
        log.Printf("Connect on WS ... %s", wsUrl)

        response := refreshTab(*host, wsUrl)
        log.Printf("Response ... %s", response)
    }

    app.Run(os.Args)
}

func refreshTab(host string, tabWsUrl string) string {
    origin := "http://" + host + "/"
    ws, err := websocket.Dial(tabWsUrl, "", origin)
    if err != nil {
        log.Fatal(err)
    }

    if _, err := ws.Write([]byte("{\"id\": 0, \"method\": \"Page.reload\"}")); err != nil {
        log.Fatal(err)
    }

    var msg = make([]byte, 512)
    if _, err := ws.Read(msg); err != nil {
        log.Fatal(err)
    }

    return string(msg[:])
}

func getWebSocketDebuggerUrl(host string, port string) string {
    res, err := http.Get("http://" + host + ":" + port + "/json")
    if err != nil {
        log.Fatal(err)
    }

    jsonBytes, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }

    type Tab struct {
        Description, DevtoolsFrontendUrl, FaviconUrl, Id, Title, Type, Url, WebSocketDebuggerUrl string
    }
    var tabs []Tab

    err = json.Unmarshal(jsonBytes, &tabs)
    if err != nil {
        log.Fatal(err)
    }

    firstTab := tabs[0]

    return firstTab.WebSocketDebuggerUrl
}
