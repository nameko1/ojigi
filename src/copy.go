package ojigi

import (
    "fmt"
    "os/exec"
)

func Copy(service string, key []byte) {
    cipherPasswd, length := GetPasswdFromService(service)
    if length != 0 {
        passwd := DecodePasswd(cipherPasswd, key, length)
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
        fmt.Printf("Don't know \"%s\" password\n", service)
    }
}
