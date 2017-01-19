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
        router.PathPrefix("/style.css").Handler(http.FileServer(http.Dir("/usr/local/domio_client/style.css")))
        router.PathPrefix("/public.bundle.js").Handler(http.FileServer(http.Dir("/usr/local/domio_client/client.public.bundle.js")))
        router.PathPrefix("/user.bundle.js").Handler(http.FileServer(http.Dir("/usr/local/domio_client/client.user.bundle.js")))
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