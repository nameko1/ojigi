package ojigi

import (
    "testing"
)

var(
    scenario1 = [...] string{}
    scenario2 = [...] string{"show", "service"}
    scenario3 = [...] string{"register", "service", "-p", "passwd"}
    scenario4 = [...] string{"register", "service"}
    scenario5 = [...] string{"register", "-p", "password"}
    scenario6 = [...] string{"show", "copy", "service"}
    scenario7 = [...] string{"-p", "password"}
    scenario8 = [...] string{"list"}
    scenario9 = [...] string{"show", "-h"}
    scenario10 = [...] string{"register", "-p"}
    // scenario11 = [...] string{"regggg", "service", "-p", "password"}
    // scenario12 = [...] string{"regggg", "service", "-q", "password"}
    scenario13 = [...] string{"register", "service", "-q", "password"}
)

func TestParseOptions(t *testing.T) {

    opts := defaultOptions()
    parseOptions(opts, scenario1[0:])
    if opts.help != "help" {
        t.Error("should show usage when ojigi run with no command")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario2[0:])
    if opts.command != scenario2[0] && opts.service != scenario2[1] {
        t.Error("should get correct command and arguments")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario3[0:])
    if opts.command != scenario3[0] && opts.service != scenario3[1] && string(opts.passwd) != scenario3[3] {
        t.Error("should get correct command and arguments")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario4[0:])
    if opts.command != scenario4[0] && opts.service != scenario4[1] {
        t.Error("should get correct command and arguments")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario5[0:])
    if opts.help != "register" {
        t.Error("register command must set serivce")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario6[0:])
    if opts.help != "show" {
        t.Error("no multi command")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario7[0:])
    if opts.command != "-p" {
        t.Error("no command is not allowed")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario8[0:])
    if opts.command != "list" {
        t.Error("list command require no args")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario9[0:])
    if opts.help != "show" {
        t.Error("should support command help")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario10[0:])
    if opts.help != "register" {
        t.Error("register command must set serivce")
    }

    // show usage by core.go
    // opts = defaultOptions()
    // parseOptions(opts, scenario11[0:])
    // if opts.help != "help" {
    //     t.Error("should show help when input bad command")
    // }

    // opts = defaultOptions()
    // parseOptions(opts, scenario12[0:])
    // if opts.help != "help" {
    //     t.Error("should show help when input bad command")
    // }

    opts = defaultOptions()
    parseOptions(opts, scenario13[0:])
    if opts.help != "register" {
        t.Error("should show help when input bad options")
    }
}

func TestValidateOptions(t *testing.T) {
    opts := defaultOptions()
    //show
    opts.command = "show"
    opts.service = ""
    if validateOptions(opts) {
        t.Error("show command require a service name")
    }

    //copy
    opts.command = "copy"
    opts.service = ""
    if validateOptions(opts) {
        t.Error("copy command require a service name")
    }

    //delete
    opts.command = "delete"
    opts.service = ""
    if validateOptions(opts) {
        t.Error("delete command require a service name")
    }

    //register
    opts.command = "register"
    opts.service = ""
    opts.passwd = []byte("pass")
    if validateOptions(opts) {
        t.Error("register command require a service name")
    }

    //modify
    opts.command = "modify"
    opts.service = ""
    opts.passwd = []byte("pass")
    if validateOptions(opts) {
        t.Error("modify command require a service name")
    }
}
