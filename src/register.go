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

func writePasswd(service string, passwd string) bool{
    file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
