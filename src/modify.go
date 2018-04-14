package ojigi

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

func Modify(service string, passwd []byte, key []byte, action string) {

    complate := false
    oldFile, oldFileErr := os.OpenFile(FilePath, os.O_RDONLY|os.O_CREATE, 0600)
    if oldFileErr != nil {
        fmt.Println("Faild: can not open password file")
        return
    }

    newFile, newFileErr := os.OpenFile(NewFilePath, os.O_WRONLY|os.O_CREATE, 0600)
    if newFileErr != nil {
        fmt.Println("Faild: can not open password file")
        return
    }

    defer oldFile.Close()
    defer newFile.Close()

    sc := bufio.NewScanner(oldFile)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            os.Remove(NewFilePath)
            fmt.Println("Faild: can not update password file")
            break
        }
        if l := strings.Split(sc.Text(), ":"); l[0] == service {
            if action == "modify" {
                if len(passwd) == 0 {
                    passwd = VerifyPasswdScanf("Enter service password: ", "Verify: ", faildPassword)
                }
                cipherPasswd := EncodePasswd(passwd, key)
                newFile.Write([]byte(l[0]+":"+strconv.Itoa(len(passwd))+":"+cipherPasswd+"\n"))
            }
            complate = true
        } else {
            newFile.Write([]byte(sc.Text()+"\n"))
        }
    }

    os.Remove(FilePath)
    os.Rename(NewFilePath, FilePath)
    if complate {
        fmt.Printf("Success to %s password\n", action)
    } else {
        fmt.Printf("Fail to %s, password not found\n", action)
    }
}
