package ojigi 

import 
    "os"

const DirPath = "/etc/ojigi"
const FilePath = DirPath+"/passwd"

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
    default:
    }
}
