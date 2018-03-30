package ojigi

import (
    "os"
)

type Options struct {
    action string
    service string
    passwd string
}

func defaultOptions() *Options {
    return &Options{
        action: "help",
        service: "",
        passwd: ""}

}
func validOptions(opts *Options) bool {
    switch opts.action {
    case "copy","show","delete":
        if len(opts.service) == 0 {
            return false
        }
    case "register", "modify":
        if len(opts.service) == 0 || len(opts.passwd) == 0 {
            return false
        }
    default:
    }
    return true
}

func parseOptions(opts *Options, args []string) {
    for i := 0; i < len(args); i++ {
        arg := args[i]
        switch arg {
        case "show":
            opts.action = "show"
        case "register":
            opts.action = "register"
        case "modify":
            opts.action = "modify"
        case "delete":
            opts.action = "delete"
        case "copy":
            opts.action = "copy"
        case "-s", "--service":
            if i + 1 >= len(args) {
                opts.action = "help"
                return
            }
            opts.service=args[i+1]
        case "-p", "--passwd":
            if i + 1 >= len(args) {
                opts.action = "help"
                return
            }
            opts.passwd=args[i+1]
        default:
        }
    } 
    // if faild validation than no action
    if !validOptions(opts) {
        opts.action = "help"
    }
}

func ParseOptions() *Options{
    opts := defaultOptions()
    parseOptions(opts, os.Args[1:])
    return opts
}
