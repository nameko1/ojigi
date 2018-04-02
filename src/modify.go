package ojigi

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

func Modify(service string, passwd []byte, key []byte, action string) {

    oldFile, oldFileErr := os.OpenFile(FilePath, os.O_RDONLY|os.O_CREATE, 0600)
    if oldFileErr != nil {
        fmt.Println("\nfaild: can not open password file")
        return
    }

    newFile, newFileErr := os.OpenFile(NewFilePath, os.O_WRONLY|os.O_CREATE, 0600)
    if newFileErr != nil {
        fmt.Println("\nfaild: can not open password file")
        return
    }

    defer oldFile.Close()
    defer newFile.Close()

    cipherPasswd := EncodePasswd(passwd, key)

    sc := bufio.NewScanner(oldFile)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            os.Remove(NewFilePath)
            fmt.Println("\nfaild: can not update password file")
            break
        }
        if l := strings.Split(sc.Text(), ":"); l[0] == service {
            if action == "modify" {
                newFile.Write([]byte(l[0]+":"+cipherPasswd+"\n"))
            }
        } else {
            newFile.Write([]byte(sc.Text()+"\n"))
        }
    }

    os.Remove(FilePath)
    os.Rename(NewFilePath, FilePath)
}
