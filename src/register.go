package ojigi

import (
    "fmt"
    "os"
    "bufio"
)

func isExistPasswd(service string) bool {
    if len(GetPasswdFromService(service)) != 0 {
        return true
    }
    return false
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
    wr.WriteString(service + ":" + cipherPasswd + "\n")
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
