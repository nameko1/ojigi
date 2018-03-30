package ojigi

import (
    "os"
)

func Run(opts *Options) {
    if _, err := os.Stat(DirPath); err != nil {
        if !Init() {
            return
        }
    }

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
        Usage()
    }
}
