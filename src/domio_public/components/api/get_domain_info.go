package api

import (
    "domio_public/components/config"
    "net/http"
    "time"
    "log"
)

func GetDomainInfo(domainName string, token string) DomainJson {

    var domain DomainJson

    url := config.Config.ApiUrl


    log.Print(url + "/domain/" + domainName)

    req, requestErr := http.NewRequest("GET", url + "/domain/" + domainName, nil)

    req.Header.Set("Content-Type", "application/json")

    timeout := time.Duration(5 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    req.Header.Add("Authorization", "Bearer " + token)

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return domain
    }
    defer resp.Body.Close()

    decodeError := DecodeJsonRequestBody(resp, &domain)

    if (decodeError != nil) {
        log.Print(decodeError)
        return domain
    }

    return domain
}
