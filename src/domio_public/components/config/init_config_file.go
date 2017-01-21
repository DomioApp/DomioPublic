package config

import (
    "fmt"
    "os"
    "log"
    "encoding/json"
    "io/ioutil"
    "path"
)

func InitConfigFile(filenameFlag *string, templatesFolder *string, webPortFlag *uint, envFlag *string, apiUrl *string) error {

    argsErr := false

    if *filenameFlag == "" {
        fmt.Println("Please supply the filename --file option.")
        argsErr = true
    }

    if *templatesFolder == "" {
        fmt.Println("Please supply the --templates-folder option.")
        argsErr = true
    }

    if argsErr {
        fmt.Println("-----------------------------------------")
        fmt.Println("Please provide required options.")
        os.Exit(1)
    }
    //dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    conf := Configuration{
        TemplatesFolder: *templatesFolder,
        Port: *webPortFlag,
        Env: *envFlag,
        ApiUrl: *apiUrl,
    }

    if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
        log.Print("Creating config folder...")
        log.Print(ConfigPath)
        os.MkdirAll(ConfigPath, 0660)
    }

    jsonConfig, _ := json.MarshalIndent(conf, "", "    ")
    err := ioutil.WriteFile(path.Join(ConfigPath, *filenameFlag), jsonConfig, 0660)
    if (err != nil) {
        log.Println(err)
        os.Exit(1)
    }
    fmt.Printf("Config file saved to %s\n", path.Join(ConfigPath, *filenameFlag))
    return nil

}
