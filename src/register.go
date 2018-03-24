package ojigi
import (
    "fmt"
)

func register(service string, passwd string) {
    fmt.Println("call register")
    fmt.Println(service, passwd)
}
