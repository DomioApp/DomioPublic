package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "domio_public/routes"
    "domio_public/components/config"
)

func NewRouter() *mux.Router {
    conf := config.Config

    router := mux.NewRouter().StrictSlash(true)

    if (conf.Env == "development") {
    }

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