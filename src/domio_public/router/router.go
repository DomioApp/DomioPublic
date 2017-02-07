package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "domio_public/routes"
)

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)

    for _, route := range routes.RoutesList {
        var handler http.Handler
        handler = route.HandlerFunc

        router.
        Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }

    return router
}