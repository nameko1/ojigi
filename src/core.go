package ojigi 

import 
    "os"

const FilePath = "/etc/ojigi/passwd"

func Run(opts *Options) {
    if _, err := os.Stat(FilePath); err != nil {
        if !Init() {
            return 
        }
    }

    switch opts.action {
    case "show":
        Show(opts.service)
    case "register":
        Register(opts.service, opts.passwd)
    default:
    }
}
