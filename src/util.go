package ojigi

import (
    "os"
    "bufio"
    "strings"
    "fmt"
)

const (
DirPath = "/etc/ojigi"
FilePath = DirPath+"/ojigi_note"
PasswdPath = DirPath+"/passwd"
NewFilePath = DirPath+"/new_ojigi_note"
)

func GetPasswdFromService(service string, optionalPath ...string) string {
    path := FilePath
    if len(optionalPath) > 0 {
        path = optionalPath[0]
    }

    file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            return ""
        }
        line := strings.Split(sc.Text(), ":")
        if line[0] == service {
            return line[1]
        }
    }
    return ""
}
