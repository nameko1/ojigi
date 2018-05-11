package ojigi

import (
    "os"
)

type Options struct {
    command string
    service string
    help string
    passwd []byte
}

type Paths struct {
    directory string
    passwd string
    file string
    newfile string
}

func GetPaths() *Paths {
    dirPath := os.ExpandEnv("$OJIGIPATH")
    return &Paths{
    directory: dirPath,
    passwd: dirPath + "/passwd",
    file: dirPath + "/ojigi_note",
    newfile: dirPath + "/new_ojigi_note"}
}

func defaultOptions() *Options {
    return &Options{
        command: "",
        service: "",
        help: "",
        passwd: nil}

}
func validateOptions(opts *Options) bool {
    if opts.help == "" && len(opts.service) == 0 && opts.command != "list" {
        return false
    }
    return true
}

func parseOptions(opts *Options, args []string) {
    //command is only first arguments
    opts.command = args[0]
    for i := 1; i < len(args); i++ {
        arg := args[i]
        switch arg {
        //deprecated
        case "-s", "--service":
            if i + 1 >= len(args) {
                opts.help = "help"
                return
            }
            i++
            opts.service=args[i]
        case "-p", "--passwd":
            if i + 1 >= len(args) {
                opts.help = "help"
                return
            }
            i++
            opts.passwd=[]byte(args[i])
        case "-h", "--help":
            opts.help = opts.command
        default:
            opts.help = "help"
        }
    }
    // if faild validation than no command
    if !validateOptions(opts) {
        opts.help = "help"
    }
}

func ParseOptions() *Options{
    opts := defaultOptions()
    parseOptions(opts, os.Args[1:])
    return opts
}
