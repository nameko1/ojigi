package ojigi
import (
    "fmt"
)

func Show(service string, key []byte) {
    cipherPasswd, length := GetPasswdFromService(service)
    passwd := DecodePasswd(cipherPasswd, key, length)
    if len(passwd) != 0 {
        fmt.Println("\n"+string(passwd))
    } else {
        fmt.Println("\nPassword not registered")
    }
}
