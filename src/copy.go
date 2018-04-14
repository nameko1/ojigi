package ojigi

import (
    "fmt"
    "os/exec"
)

func Copy(service string, key []byte) {
    cipherPasswd, length := GetPasswdFromService(service)
    passwd := DecodePasswd(cipherPasswd, key, length)
    if len(passwd) != 0 {
        copyCmd := exec.Command("pbcopy")
        in, err := copyCmd.StdinPipe()
        if err != nil {
            return
        }
        if err := copyCmd.Start(); err != nil {
            return
        }
        if _, err := in.Write(passwd); err != nil {
            return
        }
        if err := in.Close(); err != nil {
            return
        }
        fmt.Println("Copied password to your clipboard")
    } else {
        fmt.Printf("Password of %s not registered", service)
    }
}
