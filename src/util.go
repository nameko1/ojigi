package ojigi

import (
    "os"
    "bufio"
    "encoding/hex"
    "crypto/sha1"
    "strings"
    "fmt"
    "syscall"
    "golang.org/x/crypto/ssh/terminal"
    "io"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
)

const (
DirPath = "/etc/ojigi"
FilePath = DirPath+"/ojigi_note"
PasswdPath = DirPath+"/passwd"
NewFilePath = DirPath+"/new_ojigi_note"
)

func DecodePasswd(hoge string, key []byte) []byte {
    encPasswd, decodeErr := hex.DecodeString(hoge)
    if decodeErr != nil {
        fmt.Printf("Error: %s\n", decodeErr)
        os.Exit(0)
    }

    block, aesErr := aes.NewCipher(key)
    if aesErr != nil {
        fmt.Printf("Error: %s\n", aesErr)
        os.Exit(0)
    }

    iv := encPasswd[:aes.BlockSize]
    //size問題
    passwd := make([]byte, 50)

    stream := cipher.NewCTR(block, iv)
    stream.XORKeyStream(passwd, encPasswd[aes.BlockSize:])
    return passwd
}

func EncodePasswd(passwd []byte, key []byte) string {
    block, aesErr := aes.NewCipher(key)
    if aesErr != nil {
        fmt.Println("\nFail: fail to register password")
        return ""
    }

    cipherPasswd := make([]byte, aes.BlockSize+len(passwd))

    iv := cipherPasswd[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        fmt.Println("\nFail: fail to register password")
        return ""
    }
    stream := cipher.NewCTR(block, iv)
    stream.XORKeyStream(cipherPasswd[aes.BlockSize:], passwd)
    return hex.EncodeToString(cipherPasswd)
}

func GetPasswdFromService(service string, optionalPath ...string) string {
    path := FilePath
    if len(optionalPath) > 0 {
        path = optionalPath[0]
    }

    file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            return ""
        }
        line := strings.Split(sc.Text(), ":")
        if line[0] == service {
            return line[1]
        }
    }
    return ""
}


func Sha1Sum(data []byte) string {
    hash := sha1.Sum(data)
    return hex.EncodeToString(hash[:])
}

func PasswdScanf(msg string, error func()) []byte {
    fmt.Print(msg);
    passwd, err := terminal.ReadPassword(int(syscall.Stdin))
    if err != nil {
        error()
    }
    //TODO 
    // fmt.Println("")
    return passwd
}

func VerifyPasswdScanf(firstMsg string, secondMsg string, error func()) []byte {
    passwd := PasswdScanf(firstMsg, error)
    verify := PasswdScanf(secondMsg, error)

    if string(passwd) != string(verify) {
        error()
    }
    return passwd
}
