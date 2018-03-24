package ojigi

import (
    "os"
)

type Options struct {
    action string
    service string
    passwd string
    encrypt bool
}

func defaultOptions() *Options{
    return &Options{
        action: "help",
        service: "",
        passwd: "",
        encrypt: true} 
}

func parseOptions(opts *Options, args []string) {
    for i := 0; i < len(args); i++ {
        arg := args[i]
        switch arg {
        case "show":
            opts.action="show"
        case "register":
            opts.action="register"
        case "-s", "--service":
            opts.service=args[i+1]
        case "-p", "--passwd":
            opts.passwd=args[i+1]
        case "--no-encrypt":
            opts.encrypt=false
        default:
        }
    } 
}

func ParseOptions() *Options{
    opts := defaultOptions()
    parseOptions(opts, os.Args[1:])
    return opts
}
