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
    case "register":
        Register(opts.service, opts.passwd)
    case "delete":
        Delete(opts.service)
    default:
    }
}
