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
