package ojigi

import (
    "os"
    "fmt"
)

var BUFSIZE=40

func getPasswd() string {
    file, err := os.OpenFile(PasswdPath, os.O_RDONLY, 0400)
    if err != nil {
        faild()
    }
    defer file.Close()

    passwd := make([]byte, BUFSIZE)
    n, readErr := file.Read(passwd)
    if readErr != nil || n != 40 {
        faild()
    }

    return string(passwd)
}

func faild () {
    fmt.Println("\nAuthenticaion faild")
    os.Exit(0)
}

func Authentication() []byte {
    passwd := getPasswd()
    input := PasswdScanf("Enter your password: ", faild)

    if passwd != Sha1Sum(input) {
        faild()
    }
    return input
}
