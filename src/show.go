package ojigi
import (
    "fmt"
)

func Show(service string) {
    passwd := GetPasswdFromService(service)
    if len(passwd) != 0 {
        fmt.Println(passwd)
    } else {
        fmt.Println("Password not registered") 
    }
}
