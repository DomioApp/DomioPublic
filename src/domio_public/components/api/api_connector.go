package api

import (
    "net/http"
    "log"
    "domio_public/components/config"
)

type AppStatusInfo struct {
    Version       string `json:"app_version"`
    BuildAgo      string `json:"app_buildago"`
    Buildstamp    string `json:"app_buildstamp"`
    BuildDateTime string `json:"app_builddatetime"`
    Hash          string `json:"app_hash"`
}


// GetAPIStatus comment
func GetAPIStatus() AppStatusInfo {
    url := config.Config.ApiUrl
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

    return appStatusInfo
}
