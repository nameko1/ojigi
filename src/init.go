package ojigi

import (
    "os"
    "fmt"
    "bufio"
)

func faildCreatePass() {
    os.Remove(PasswdPath)
    fmt.Println("\nSorry try again")
    os.Exit(0)
}

func createPasswdFile() {
    file, err := os.OpenFile(PasswdPath, os.O_WRONLY|os.O_CREATE, 0400)
    if err != nil {
        return
    }
    defer file.Close()

    fmt.Println("Setting your password first!")
    passwd := VerifyPasswdScanf("Enter your password: ", "\nVerify: ", faildCreatePass)
    if string(passwd) == "" {
        faildCreatePass()
    }

    passwdStr := Sha1Sum(passwd)
    wr := bufio.NewWriter(file)
    wr.WriteString(passwdStr)
    wr.Flush()
    fmt.Println("\nRegister new password!!")
    return
}

func Init() {
    if _, err := os.Stat(DirPath); err != nil {
        if err := os.Mkdir(DirPath, 0600); err != nil {
            fmt.Println(err)
            return
        }
    }
    createPasswdFile()
    return
}
