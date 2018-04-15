package ojigi

import (
    "os"
    "fmt"
    "golang.org/x/crypto/scrypt"
)

var BUFSIZE=40

func getUserPasswd() string {
    file, err := os.OpenFile(filePaths.passwd, os.O_RDONLY, 0400)
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
    fmt.Println("Authenticaion faild")
    os.Exit(0)
}

func Authentication() []byte {
    passwd := getUserPasswd()
    inputPasswd := PasswdScanf("Enter your password: ", faild)

    if passwd != Sha1Sum(inputPasswd) {
        faild()
    }
    //passwd convert 32bit key
    key, err := scrypt.Key(inputPasswd, []byte("nameko"), 16384, 8, 1, 32)
    if err != nil {
        fmt.Printf("Error: %s \n", err)
        os.Exit(0)
    }
    return key
}
