package arguments

import (
    "flag"
    "os"
    "fmt"
    "log"
    //"domio_public/components/config"
    //"domio_public/components/server"
    //"domio_public/db"
    "domio_public/components/config"
    "domio_public/components/server"
)

type Arguments struct {
    Command string
}

func ProcessArguments() (Arguments, error) {
    var args = Arguments{}

    if len(os.Args) == 1 {
        showHelpAndExit()
    }

    switch os.Args[1] {

    case "init":
        args = processInitArguments()

    case "start":
        args = processStartArguments()

    default:
        fmt.Printf("%q is not valid command.\n", os.Args[1])
        os.Exit(2)
    }

    return args, nil
}

func showHelpAndExit() {
    fmt.Println("usage: domio <command> [<args>]")
    fmt.Println("Commands are: ")
    fmt.Println(" init   Init with new config file")
    fmt.Println(" start  Start server")
    os.Exit(1)

}

func processInitArguments() Arguments {
    var args = Arguments{
        Command: "init",
    }

    initCommand := flag.NewFlagSet("init", flag.ExitOnError)
    filenameFlag := initCommand.String("file", "config.json", "config file absolute path")
    awsAccessKeyIdFlag := initCommand.String("aws-access-key-id", "", "AWS Access Key ID")
    awsSecretAccessKeyFlag := initCommand.String("aws-secret-access-key", "", "AWS Secret Access Key")
    dbNameFlag := initCommand.String("db-name", "", "DB name")
    dbUserFlag := initCommand.String("db-user", "", "DB user name")
    dbPasswordFlag := initCommand.String("db-password", "", "DB password")
    webPortFlag := initCommand.Uint("port", 8080, "Port for the HTTP server to run on")
    envFlag := initCommand.String("env", "development", "Environment name: development, testing, production")

    log.Print(*filenameFlag)
    log.Print(*awsAccessKeyIdFlag)
    log.Print(*awsSecretAccessKeyFlag)
    log.Print(*awsSecretAccessKeyFlag)
    log.Print(*dbNameFlag)
    log.Print(*dbUserFlag)
    log.Print(*dbPasswordFlag)
    log.Print(*dbPasswordFlag)
    log.Print(*webPortFlag)
    log.Print(*envFlag)

    initCommand.Parse(os.Args[2:])

    if initCommand.Parsed() {
        config.InitConfigFile(filenameFlag, awsAccessKeyIdFlag, awsSecretAccessKeyFlag, dbNameFlag, dbUserFlag, dbPasswordFlag, webPortFlag, envFlag)
    }
    return args
}

func processStartArguments() Arguments {
    var args = Arguments{
        Command: "start",
    }

    startCommand := flag.NewFlagSet("start", flag.ExitOnError)
    //recipientFlag := sendCommand.String("recipient", "", "Recipient of your message")
    //messageFlag := sendCommand.String("message", "", "Text message")
    log.Print(*startCommand)

    startCommand.Parse(os.Args[2:])

    config.LoadConfig()

    if startCommand.Parsed() {
        server.Start()
    }
    return args
}
