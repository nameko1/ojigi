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
    file, fileErr := os.OpenFile(FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if fileErr != nil {
        fmt.Println("\nFail: can't open password file")
        return
    }
    defer file.Close()

    cipherPasswd := EncodePasswd(passwd, key)

    wr := bufio.NewWriter(file)
    wr.WriteString(service + ":" + strconv.Itoa(len(passwd)) + ":" + cipherPasswd + "\n")
    wr.Flush()
    fmt.Println("\nSuccess to register password!!")
}

func faildPassword() {
    fmt.Println("\nSorry try again")
    os.Exit(0)
}

func Register(service string, passwd []byte, key []byte) {
    if isExistPasswd(service) {
        fmt.Printf("\nPassword of %s is already registered\n", service)
        return
    }
    if passwd == nil {
        passwd = VerifyPasswdScanf("\nEnter service password: ", "\nVerify: ", faildPassword)
    }

    writePasswd(service, passwd, key)
}
