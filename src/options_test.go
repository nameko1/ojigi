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
    scenario7 = [...] string{"-s", "service"}
    scenario8 = [...] string{"list"}
    scenario9 = [...] string{"show", "-h"}
)

func TestParseOptions(t *testing.T) {
    opts := defaultOptions()
    parseOptions(opts, scenario1[0:])
    if opts.help != "help" {
        t.Error("command should be help when set option with no argument")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario2[0:])
    if opts.command != scenario2[0] && opts.service != scenario2[2] {
        t.Error("should get correct command and arguments")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario3[0:])
    if opts.command != scenario3[0] && opts.service != scenario3[2] && string(opts.passwd) != scenario3[4] {
        t.Error("should get correct command and arguments")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario4[0:])
    if opts.command != scenario4[0] && opts.service != scenario4[2] {
        t.Error("should get correct command and arguments")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario5[0:])
    if opts.help != "help" {
        t.Error("command should be help when set option with no argument")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario6[0:])
    if opts.help != "help" {
        t.Error("no multi command")
    }

    opts = defaultOptions()
    parseOptions(opts, scenario7[0:])
    if opts.command != "-s" {
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
    opts.passwd = nil
    if validateOptions(opts) {
        t.Error("register command require a service name and password")
    }
    opts.command = "register"
    opts.service = ""
    opts.passwd = []byte("pass")
    if validateOptions(opts) {
        t.Error("register command require a service name")
    }
    opts.command = "register"
    opts.service = "service"
    opts.passwd = nil
    if validateOptions(opts) {
        t.Error("register command require a password")
    }

    //modify
    opts.command = "modify"
    opts.service = ""
    opts.passwd = nil
    if validateOptions(opts) {
        t.Error("modify command require a service name and password")
    }
    opts.command = "modify"
    opts.service = ""
    opts.passwd = []byte("pass")
    if validateOptions(opts) {
        t.Error("modify command require a service name")
    }
    opts.command = "modify"
    opts.service = "service"
    opts.passwd = nil
    if validateOptions(opts) {
        t.Error("modify command require a password")
    }
}
