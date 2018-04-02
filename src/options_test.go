package ojigi

import (
    "testing"
)

var(
    scenario1 = [...] string{"show", "-s"}
    scenario2 = [...] string{"show", "-s", "service"}
    scenario3 = [...] string{"register", "-s", "service", "-p", "passwd"}
    scenario4 = [...] string{"register", "-s", "service"}
    scenario5 = [...] string{"register", "-s", "-p", "service"}
    scenario6 = [...] string{"show", "copy", "-s", "service"}
)

func TestParseOptions(t *testing.T) {
    opts := defaultOptions()
    parseOptions(opts, scenario1[0:])
    if opts.action != "help" {
        t.Error("action should be help when set option with no argument")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario2[0:])
    if opts.action != scenario2[0] && opts.service != scenario2[2] {
        t.Error("should get correct action and arguments")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario3[0:])
    if opts.action != scenario3[0] && opts.service != scenario3[2] && opts.passwd != scenario3[4] {
        t.Error("should get correct action and arguments")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario4[0:])
    if opts.action != scenario4[0] && opts.service != scenario4[2] {
        t.Error("should get correct action and arguments")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario5[0:])
    if opts.action != "help" {
        t.Error("action should be help when set option with no argument")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario6[0:])
    if opts.action != "help" {
        t.Error("no multi action")
    }
}

func TestValidateOptions(t *testing.T) {
    opts := defaultOptions()
    //show
    opts.action = "show"
    opts.service = ""
    if validateOptions(opts) {
        t.Error("show command require a service name")
    }

    //copy
    opts.action = "copy"
    opts.service = ""
    if validateOptions(opts) {
        t.Error("copy command require a service name")
    }

    //delete
    opts.action = "delete"
    opts.service = ""
    if validateOptions(opts) {
        t.Error("delete command require a service name")
    }

    //register
    opts.action = "register"
    opts.service = ""
    opts.passwd = ""
    if validateOptions(opts) {
        t.Error("register command require a service name and password")
    }
    opts.action = "register"
    opts.service = ""
    opts.passwd = "pass"
    if validateOptions(opts) {
        t.Error("register command require a service name")
    }
    opts.action = "register"
    opts.service = "service"
    opts.passwd = ""
    if validateOptions(opts) {
        t.Error("register command require a password")
    }

    //modify
    opts.action = "modify"
    opts.service = ""
    opts.passwd = ""
    if validateOptions(opts) {
        t.Error("modify command require a service name and password")
    }
    opts.action = "modify"
    opts.service = ""
    opts.passwd = "pass"
    if validateOptions(opts) {
        t.Error("modify command require a service name")
    }
    opts.action = "modify"
    opts.service = "service"
    opts.passwd = ""
    if validateOptions(opts) {
        t.Error("modify command require a password")
    }
}
