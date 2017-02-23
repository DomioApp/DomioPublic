package api

import (
    "net/http"
    "encoding/json"
    "domio_public/errors"
)

func DecodeJsonRequestBody(resp *http.Response, obj interface{}) *errors.DomioError {

    decoder := json.NewDecoder(resp.Body)

    decodeError := decoder.Decode(&obj)
    if (decodeError != nil) {
        return &errors.JsonDecodeError
    }

    return nil
}