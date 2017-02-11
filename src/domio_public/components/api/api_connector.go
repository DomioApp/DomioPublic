package api

import (
    "net/http"
    "log"
    "domio_public/components/config"
    "time"
)

type DomainJson struct {
    Name          string `json:"name"`
    Owner         string `json:"owner"`
    PricePerMonth uint64 `json:"price_per_month"`
    IsVisible     bool `json:"is_visible"`
    ZoneId        string `json:"zone_id"`
    NS1           string `json:"ns1"`
    NS2           string `json:"ns2"`
    NS3           string `json:"ns3"`
    NS4           string `json:"ns4"`
}

type Subscription struct {
    Id string `json:"id"`
}

type AppStatusInfo struct {
    Version       string `json:"app_version"`
    BuildAgo      string `json:"app_buildago"`
    Buildstamp    string `json:"app_buildstamp"`
    BuildDateTime string `json:"app_builddatetime"`
    Hash          string `json:"app_hash"`
}

func GetAPIStatus() AppStatusInfo {

    var appStatusInfo AppStatusInfo

    url := config.Config.ApiUrl

    //color.Set(color.FgGreen)
    //fmt.Println("URL:> ", url)
    //color.Unset()

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

func GetUserDomains(token string) []DomainJson {

    var userDomains []DomainJson

    url := config.Config.ApiUrl

    req, requestErr := http.NewRequest("GET", url + "/domains/user", nil)

    req.Header.Set("Content-Type", "application/json")

    timeout := time.Duration(5 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    req.Header.Add("Authorization", "Bearer " + token)

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return userDomains
    }
    defer resp.Body.Close()

    decodeError := DecodeJsonRequestBody(resp, &userDomains)

    if (decodeError != nil) {
        log.Print(decodeError)
        return userDomains
    }

    return userDomains
}

func GetDomains() []DomainJson {

    var domains []DomainJson

    url := config.Config.ApiUrl

    req, requestErr := http.NewRequest("GET", url + "/domains/available", nil)

    req.Header.Set("Content-Type", "application/json")

    timeout := time.Duration(5 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return domains
    }
    defer resp.Body.Close()

    decodeError := DecodeJsonRequestBody(resp, &domains)

    if (decodeError != nil) {
        log.Print(decodeError)
        return domains
    }

    return domains
}
