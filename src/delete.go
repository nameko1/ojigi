package ojigi

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

func Delete(service string) {
    if len(service) == 0 {
        Usage()
        return
    }

    oldFile, oldFileErr := os.OpenFile(FilePath, os.O_RDONLY|os.O_CREATE, 0600)
    if oldFileErr != nil {
        fmt.Println("faild: can not open password file")
        return
    }

    newFile, newFileErr := os.OpenFile(NewFilePath, os.O_WRONLY|os.O_CREATE, 0600)
    if newFileErr != nil {
        fmt.Println("faild: can not open password file")
        return
    }

    defer oldFile.Close()
    defer newFile.Close()

    sc := bufio.NewScanner(oldFile)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            os.Remove(NewFilePath)
            fmt.Println("faild: can not update password file")
            break 
        }
        if l := strings.Split(sc.Text(), ":"); l[0] != service {
            newFile.Write([]byte(sc.Text()+"\n"))
        }
    }

    os.Remove(FilePath)
    os.Rename(NewFilePath, FilePath)
}
