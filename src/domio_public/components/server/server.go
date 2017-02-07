package server

import (
    "domio_public/router"
    "log"
    "domio_public/components/config"
    "net/http"
    "fmt"
)

func StartRouter() {
    log.Print("Starting router...")
    domiorouter := router.NewRouter()
    conf := config.Config

    if (conf.Env == "development") {
        log.Print("Development environment, handling static files by Go...")
        domiorouter.Path("/style.css").Handler(http.FileServer(http.Dir("/usr/local/domio_client")))
        domiorouter.Path("/bundle.js").Handler(http.FileServer(http.Dir("/usr/local/domio_client")))
        domiorouter.PathPrefix("/app").Handler(http.FileServer(http.Dir("/Users/sergeibasharov/WebstormProjects/DomioClient/src")))
        //domiorouter.PathPrefix("/app").Handler(http.FileServer(http.Dir("/Users/sbasharov/WebstormProjects/DomioClient/src")))
    }


    log.Printf("Web server is running on http://localhost:%d", config.Config.Port)
    err := http.ListenAndServe(fmt.Sprintf(":%v", config.Config.Port), domiorouter)
    if (err != nil) {
        msg := fmt.Sprintf("Failed to start web server on port %d", config.Config.Port)
        log.Fatal(msg)
    }

}
func Start() {
    fmt.Print("Starting app...")
    StartRouter()
}
