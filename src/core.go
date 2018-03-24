package ojigi 

func Run(opts *Options) {
    switch opts.action {
        case "show":
            Show(opts.service)
        case "register":
            Register(opts.service, opts.passwd)
        default:
    }

}
