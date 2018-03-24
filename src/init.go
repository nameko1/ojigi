package ojigi

import (
    "os"
    "fmt"
)

func Init() bool {
    if err := os.Mkdir("/etc/ojigi", 0600); err != nil {
        fmt.Println(err) 
        return false
    }
    file, err := os.Create(FilePath)
    if err != nil {
        fmt.Println(err)
        return false
    }
    defer file.Close()
    return true
}
