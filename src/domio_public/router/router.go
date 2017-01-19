package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "domio_public/routes"
    //"domio_public/components/config"
)

func NewRouter() *mux.Router {
    //conf := config.Config

    router := mux.NewRouter().StrictSlash(true)

    styleSheet := http.FileServer(http.Dir("/usr/local/domio_client/"))
    http.Handle("/", styleSheet)


    /*
    if (conf.Env == "development") {
        styleSheet := http.FileServer(http.Dir("/style.css"))
        http.Handle("/usr/local/domio_client/style.css", styleSheet)

        scriptFile := http.FileServer(http.Dir("/public.bundle.js"))
        http.Handle("/usr/local/domio_client/client.public.bundle.js", scriptFile)
    }
    */

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