package ojigi

import (
    "fmt"
    "os"
)
func list() {
    fmt.Println(`
Usage:
  ojigi list

Show service list

Options:
    `)
}

func show() {
    fmt.Println(`
Usage:
  ojigi show SERVICE

Show the password for service

Options:
    `)
}

func copy() {
    fmt.Println(`
Usage:
  ojigi copy SERVICE

Copy the password to clipboard
    `)
}

func register() {
    fmt.Println(`
Usage:
  ojigi register SERVICE [options]

Register the password for service

Options:
  -p --password           Register the password for service
    `)
}

func modify() {
    fmt.Println(`
Usage:
  ojigi modify SERVICE [options]

Modify the password for service

Options:
  -p --password           New password for service
    `)
}

func delete() {
    fmt.Println(`
Usage:
  ojigi delete SERVICE

Delete the password for service

Options:
    `)
}

func help() {
    fmt.Println(`
Usage:
  ojigi <COMMAND> [options]

Commands:
  list                    Show service list
  show                    Show the password for service
  copy                    Copy the password to clipboard
  register                Register the password for service
  modify                  Modify the password for service
  delete                  Delete the password for service
  help                    Show usage

Run ojigi COMMAND --help, more infomation on a command
    `)
}

func Usage(command string) {
    switch command {
    case "list":
        list()
    case "show":
        show()
    case "copy":
        copy()
    case "register":
        register()
    case "modify":
        modify()
    case "delete":
        delete()
    default:
        help()
    }
    os.Exit(0)
}
