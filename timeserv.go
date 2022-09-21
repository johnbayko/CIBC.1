package main

import (
    "fmt"
    "net"
    "net/http"
    "strings"
)

type timeHandler struct {
    // Might need caching here.
}

func (h *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println("r.Method", r.Method)  // debug
    switch {
    case r.Method == "GET":
       var tzNameList = []string{}
       {
           var query = r.URL.Query()
           var tzName, hasTZName = query["tz"]
           if hasTZName {
               tzNameList = strings.Split(tzName[0], ",")
           }
       }
       for tzNameIdx, tzName := range tzNameList {
           fmt.Println("tzNameIdx", tzNameIdx, "tzName", tzName)  // debug
       }
    }
}

func main() {
    fmt.Println("start")
    listener, err := net.Listen("tcp", ":0")
    if err != nil {
        panic(err)
    }
    fmt.Println("Port:", listener.Addr().(*net.TCPAddr).Port)

    http.Handle("/api/time", new(timeHandler))
    http.Serve(listener, nil)
}
