package main

import (
    "fmt"
    "encoding/json"
    "net"
    "net/http"
    "strings"
    "time"
)

type timeHandler struct {
    // Timezone caching can go here if needed.
}


func timeErrorHandle(w http.ResponseWriter, statusCode int, message string) {
    w.WriteHeader(statusCode)
    fmt.Fprintln(w, message)
    // When logging is added, add err parameter and do that here.
}


func getTzNameList(r *http.Request) []string {
    var tzNameList = []string{}

    var query = r.URL.Query()
    var tzName, hasTZName = query["tz"]
    if hasTZName {
        tzNameList = strings.Split(tzName[0], ",")
    }
    return tzNameList
}


func getTzLocMap(tzNameList []string) (map[string]*time.Location, string, error) {
    tzLocMap := map[string]*time.Location {}

    if len(tzNameList) == 0 {
        tzLocMap["current_time"] = time.UTC
    } else if len(tzNameList) == 1 {
        // Only one timezone, key is same as no timezone.
        tzName := tzNameList[0]
        tzLoc, tzLocErr := time.LoadLocation(tzName)
        if tzLocErr != nil {
            return nil, "", tzLocErr
        }
        tzLocMap["current_time"] = tzLoc
    } else {
        // Multiple time zones, key by timezone name.
        for _, tzName := range tzNameList {
            tzLoc, tzLocErr := time.LoadLocation(tzName)
            if tzLocErr != nil {
                return nil, ": " + tzName, tzLocErr
            }
            tzLocMap[tzName] = tzLoc
        }
    }
    return tzLocMap, "", nil
}


func getTimeMap(tzMap map[string]*time.Location) map[string]string {
    timeMap := map[string]string {}

    now := time.Now()
    for tzName, tzLoc := range tzMap {
        timeString := now.In(tzLoc).Format("2006-01-02 15:04:05 -0700")
        timeMap[tzName] = timeString
    }
    return timeMap
}


func (h *timeHandler) timeGetHandle(w http.ResponseWriter, r *http.Request) {
    // Timezones from parameters (if any).
    var tzNameList = getTzNameList(r)

    // Map timezone names to timezone locations.
    tzMap, tzName, tzMapErr := getTzLocMap(tzNameList)
    if tzMapErr != nil {
        timeErrorHandle(w, http.StatusNotFound, "invalid timezone" + tzName)
        return
    }

    // Map timezone names to formatted times.
    timeMap := getTimeMap(tzMap)

    // Convert timezones to JSON message.
    trJson, trJsonErr := json.MarshalIndent(timeMap, "", "  ")
    if trJsonErr != nil {
        timeErrorHandle(w, http.StatusInternalServerError, "Could not encode time.")
        return
    }

    // Write response.
    w.WriteHeader(http.StatusOK)
    w.Write(trJson)
    fmt.Fprintln(w, "")  // newline
}


func (h *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch {
    case r.Method == "GET":
        h.timeGetHandle(w, r)
    }
}


func main() {
    // Don't know yet what port this is assigned to,
    // let OS allocate one and display it.
    listener, err := net.Listen("tcp", ":0")
    if err != nil {
        panic(err)
    }
    fmt.Println("Port:", listener.Addr().(*net.TCPAddr).Port)

    http.Handle("/api/time", new(timeHandler))
    http.Serve(listener, nil)
}
