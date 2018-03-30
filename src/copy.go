package ojigi
import (
    "fmt"
    "os/exec"
)

func Copy(service string) {
    passwd := GetPasswdFromService(service)
    if len(passwd) != 0 {
        copyCmd := exec.Command("pbcopy")
        in, err := copyCmd.StdinPipe()
        if err != nil {
            return
        }
        if err := copyCmd.Start(); err != nil {
            return
        }
        if _, err := in.Write([]byte(passwd)); err != nil {
            return
        }
        if err := in.Close(); err != nil {
            return
        }
        fmt.Println("Already copied password to your clipboard")
    } else {
        fmt.Println("Password not registered")
    }
}
