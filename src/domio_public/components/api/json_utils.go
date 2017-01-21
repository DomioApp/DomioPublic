package api

import (
    "net/http"
    "encoding/json"
)

func DecodeJsonRequestBody(resp *http.Response, obj interface{}) (error) {

    decoder := json.NewDecoder(resp.Body)

    decodeError := decoder.Decode(&obj)

    return decodeError
}