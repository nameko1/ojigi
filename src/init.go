package ojigi

import (
    "os"
    "fmt"
    "bufio"
)

func faildCreatePass() {
    os.Remove(filePaths.passwd)
    fmt.Println("Sorry try again")
    os.Exit(0)
}

func createPasswdFile() {
    file, err := os.OpenFile(filePaths.passwd, os.O_WRONLY|os.O_CREATE, 0400)
    if err != nil {
        faildCreatePass()
    }
    defer file.Close()

    fmt.Println("Not found Password\nSetting your password")
    passwd := VerifyPasswdScanf("Enter your password: ", "Verify: ", faildCreatePass)
    if string(passwd) == "" {
        faildCreatePass()
    }

    passwdStr := Sha1Sum(passwd)
    wr := bufio.NewWriter(file)
    wr.WriteString(passwdStr)
    wr.Flush()
    fmt.Println("Register new password!!")
}

func Init() {
    if filePaths.directory == "" {
        fmt.Println("$OJIGIPATH not found!! \nNeet to set Password file path into $OJIGIPATH")
        os.Exit(0)
    }
    if _, err := os.Stat(filePaths.directory); err != nil {
        if err := os.Mkdir(filePaths.directory, 0700); err != nil {
            fmt.Println(err)
            return
        }
    }
    createPasswdFile()
}
