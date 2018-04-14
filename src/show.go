package ojigi
import (
    "fmt"
)

func Show(service string, key []byte) {
    cipherPasswd, length := GetPasswdFromService(service)
    if length != 0 {
        passwd := DecodePasswd(cipherPasswd, key, length)
        fmt.Println(string(passwd))
    } else {
        fmt.Printf("Don't know \"%s\" password\n", service)
    }
}
