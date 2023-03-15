package ojigi

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

func Modify(service string, passwd []byte, key []byte, command string) {

    complate := false
    oldFile, oldFileErr := os.OpenFile(filePaths.file, os.O_RDONLY|os.O_CREATE, 0600)
    if oldFileErr != nil {
        fmt.Println("Faild: can not open password file")
        return
    }

    newFile, newFileErr := os.OpenFile(filePaths.newfile, os.O_WRONLY|os.O_CREATE, 0600)
    if newFileErr != nil {
        fmt.Println("Faild: can not open password file")
        return
    }

    defer oldFile.Close()
    defer newFile.Close()

    sc := bufio.NewScanner(oldFile)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            os.Remove(filePaths.newfile)
            fmt.Println("Faild: can not update password file")
            break
        }
        if l := strings.Split(sc.Text(), ":"); l[0] == service {
            if command == "modify" {
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

    os.Remove(filePaths.file)
    os.Rename(filePaths.newfile, filePaths.file)
    if complate {
        fmt.Printf("Success to %s password\n", command)
    } else {
        fmt.Printf("Fail to %s, password not found\n", command)
    }
}
