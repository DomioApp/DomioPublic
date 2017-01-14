package routes

import (
    "net/http"
    "domio_public/handlers/index_handler"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesList = Routes{
    Route{
        "Index",
        "GET",
        "/",
        index_handler.IndexHandler,
    },
}