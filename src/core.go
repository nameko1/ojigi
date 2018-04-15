package ojigi

import (
    "os"
)

var filePaths = GetPaths()

func Run(opts *Options) {
    if _, err := os.Stat(filePaths.passwd); err != nil {
        Init()
        return
    }

    if opts.action == "help" {
        Usage()
    }
    key := Authentication()

    switch opts.action {
    case "list":
        List(opts.service)
    case "show":
        Show(opts.service, key)
    case "copy":
        Copy(opts.service, key)
    case "register":
        Register(opts.service, opts.passwd, key)
    case "modify":
        Modify(opts.service, opts.passwd, key, opts.action)
    case "delete":
        Modify(opts.service, nil, key, opts.action)
    default:
        Usage()
    }
}
