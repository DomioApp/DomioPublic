package api

import (
    "fmt"
    "net/http"
    "log"
)

type AppStatusInfo struct {
    Version       string `json:"app_version"`
    BuildAgo      string `json:"app_buildago"`
    Buildstamp    string `json:"app_buildstamp"`
    BuildTimeDate string `json:"app_builddatetime"`
    Hash          string `json:"app_hash"`
}


// GetAPIStatus comment
func GetAPIStatus() AppStatusInfo {
    url := "https://api.domio.in"
    //fmt.Println("URL:> ", url)

    req, err := http.NewRequest("GET", url, nil)

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    var appStatusInfo AppStatusInfo

    decodeError := DecodeJsonRequestBody(resp, &appStatusInfo)

    if (decodeError != nil) {
        log.Print(decodeError)
        return appStatusInfo
    }

    log.Print(appStatusInfo.BuildTimeDate)

    return appStatusInfo
}
