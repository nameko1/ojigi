package ojigi 

func Run(opts *Options) {
    switch opts.action {
        case "show":
            show(opts.service)
        case "register":
            register(opts.service, opts.passwd)
        default:
    }

}
