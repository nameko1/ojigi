package ojigi

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func isExistPasswd(service string) bool {
    _, length := GetPasswdFromService(service)
    return length != 0
}

func writePasswd(service string, passwd []byte, key []byte) {
    file, fileErr := os.OpenFile(filePaths.file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if fileErr != nil {
        fmt.Println("Fail: can't open password file")
        return
    }
    defer file.Close()

    cipherPasswd := EncodePasswd(passwd, key)

    wr := bufio.NewWriter(file)
    wr.WriteString(service + ":" + strconv.Itoa(len(passwd)) + ":" + cipherPasswd + "\n")
    wr.Flush()
    fmt.Println("Success to register password!!")
}

func faildPassword() {
    fmt.Println("Sorry try again")
    os.Exit(0)
}

func Register(service string, passwd []byte, key []byte) {
    if isExistPasswd(service) {
        fmt.Printf("Password of %s is already registered\n", service)
        return
    }
    if passwd == nil {
        passwd = VerifyPasswdScanf("Enter service password: ", "Verify: ", faildPassword)
    }

    writePasswd(service, passwd, key)
}
