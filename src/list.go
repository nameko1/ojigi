package ojigi

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func List(service string) {

    file, err := os.OpenFile(FilePath, os.O_RDONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            fmt.Printf("Error: %s\n", err)
        }
        line := strings.Split(sc.Text(), ":")
        if strings.Contains(line[0], service) {
            fmt.Printf("%s\n",line[0])
        }
    }
}
