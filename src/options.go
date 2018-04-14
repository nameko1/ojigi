package ojigi

import (
    "os"
)

type Options struct {
    action string
    service string
    passwd []byte
}

func defaultOptions() *Options {
    return &Options{
        action: "",
        service: "",
        passwd: nil}

}
func validateOptions(opts *Options) bool {
    if len(opts.service) == 0 && opts.action != "list" {
        return false
    }
    return true
}

func parseOptions(opts *Options, args []string) {
    for i := 0; i < len(args); i++ {
        arg := args[i]
        switch arg {
        case "list":
            opts.action = opts.action + arg
        case "show":
            opts.action = opts.action + arg
        case "register":
            opts.action = opts.action + arg
        case "modify":
            opts.action = opts.action + arg
        case "delete":
            opts.action = opts.action + arg
        case "copy":
            opts.action = opts.action + arg
        case "-s", "--service":
            if i + 1 >= len(args) {
                opts.action = "help"
                return
            }
            i++
            opts.service=args[i]
        case "-p", "--passwd":
            if i + 1 >= len(args) {
                opts.action = "help"
                return
            }
            i++
            opts.passwd=[]byte(args[i])
        default:
            opts.action = "help"
        }
    }
    // if faild validation than no action
    if !validateOptions(opts) {
        opts.action = "help"
    }
}

func ParseOptions() *Options{
    opts := defaultOptions()
    parseOptions(opts, os.Args[1:])
    return opts
}
