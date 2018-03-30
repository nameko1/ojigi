package ojigi

import (
    "os"
)

func Run(opts *Options) {
    if _, err := os.Stat(PasswdPath); err != nil {
        Init()
        return
    }

    if opts.action == "help" {
        Usage()
    }

    Authentication()

    switch opts.action {
    case "show":
        Show(opts.service)
    case "copy":
        Copy(opts.service)
    case "register":
        Register(opts.service, opts.passwd)
    case "modify":
        Modify(opts.service, opts.passwd, opts.action)
    case "delete":
        Modify(opts.service, "", opts.action)
    default:
    }
}
