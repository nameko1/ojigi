package ojigi
import (
    "os"
    "bufio"
    "strings"
    "fmt"
)

const DirPath = "/etc/ojigi"
const FilePath = DirPath+"/passwd"
const NewFilePath = DirPath+"/newpasswd"

func GetPasswdFromService(service string) string {
    file, err := os.OpenFile(FilePath, os.O_RDONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println(err) 
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            break 
        }
        line := strings.Split(sc.Text(), ":")
        if line[0] == service {
            return line[1] 
        }
    }
    return ""
}
