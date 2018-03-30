package ojigi
import (
    "fmt"
)

func Show(service string) {
    passwd := GetPasswdFromService(service)
    if len(passwd) != 0 {
        fmt.Println("\n"+passwd)
    } else {
        fmt.Println("\nPassword not registered")
    }
}
