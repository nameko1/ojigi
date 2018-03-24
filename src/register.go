package ojigi
import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

const filePath = "/tmp/hoge"

func isExistPasswd(service string) bool {
    file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println(err) 
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            break 
        }
        name := strings.Split(sc.Text(), ":")[0]
        if name == service {
            return true 
        }
    }
    return false
}

func writePasswd(service string, passwd string) bool{
    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return false
    }
    defer file.Close()

    wr := bufio.NewWriter(file)
    wr.WriteString(service + ":" + passwd + "\n")
    wr.Flush()
    return true
}

func Register(service string, passwd string) {
    if len(service) == 0 || len(passwd) == 0 {
        Usage()
        return
    }

    if isExistPasswd(service) {
        fmt.Printf("ojigi: %s password is already registered\n", service) 
        return
    }

    if writePasswd(service, passwd) {
       fmt.Println("Success to register password!!")  
    } else {
       fmt.Println("Fail to resister password")  
    }
}
