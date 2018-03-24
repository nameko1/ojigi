package ojigi

import (
    "os"
    "fmt"
)

func Init() bool {
    if err := os.Mkdir(DirPath, 0600); err != nil {
        fmt.Println(err) 
        return false
    }
    return true
}
