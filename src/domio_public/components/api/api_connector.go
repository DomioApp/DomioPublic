package api

import (
    "net/http"
    "log"
    "domio_public/components/config"
    "fmt"
    "github.com/fatih/color"
    "time"
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

    var appStatusInfo AppStatusInfo

    url := config.Config.ApiUrl

    color.Set(color.FgGreen)
    fmt.Println("URL:> ", url)
    color.Unset()

    req, requestErr := http.NewRequest("GET", url, nil)

    req.Header.Set("Content-Type", "application/json")

    timeout := time.Duration(5 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return appStatusInfo
    }
    defer resp.Body.Close()

    decodeError := DecodeJsonRequestBody(resp, &appStatusInfo)

    if (decodeError != nil) {
        log.Print(decodeError)
        return appStatusInfo
    }

    return appStatusInfo
}
