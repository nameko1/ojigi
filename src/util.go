package ojigi

import (
    "os"
    "bufio"
    "encoding/hex"
    "crypto/sha1"
    "strings"
    "strconv"
    "fmt"
    "syscall"
    "golang.org/x/crypto/ssh/terminal"
    "io"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
)

func DecodePasswd(hexPasswd string, key []byte, length int) []byte {
    encPasswd, decodeErr := hex.DecodeString(hexPasswd)
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
    passwd := make([]byte, length)

    stream := cipher.NewCTR(block, iv)
    stream.XORKeyStream(passwd, encPasswd[aes.BlockSize:])
    return passwd
}

func EncodePasswd(passwd []byte, key []byte) string {
    block, aesErr := aes.NewCipher(key)
    if aesErr != nil {
        fmt.Println("Fail: fail to register password")
        return ""
    }

    cipherPasswd := make([]byte, aes.BlockSize+len(passwd))

    iv := cipherPasswd[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        fmt.Println("Fail: fail to register password")
        return ""
    }
    stream := cipher.NewCTR(block, iv)
    stream.XORKeyStream(cipherPasswd[aes.BlockSize:], passwd)
    return hex.EncodeToString(cipherPasswd)
}

func GetPasswdFromService(service string, optionalPath ...string) (string, int) {
    path := filePaths.file
    if len(optionalPath) > 0 {
        path = optionalPath[0]
    }

    file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return "", 0
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            return "", 0
        }
        line := strings.Split(sc.Text(), ":")
        if line[0] == service {
            length, _ := strconv.Atoi(line[1])
            return line[2], length
        }
    }
    return "", 0
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
    fmt.Println("")
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
