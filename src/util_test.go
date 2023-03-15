package ojigi

import (
    "testing"
)

var testFile = "./testFile.txt"
var hexpasswd = "fa8711ad764f00e5229a81b9b141239bf96ecb7d8df76627"
func TestDecodePasswd(t *testing.T) {
    passwd := string(DecodePasswd(hexpasswd, []byte("31f30ddbcb1bf8446576f0e64aa4c88a9f055e3c"), 8))
    if passwd != "password" {
        t.Errorf("%s is uncorrect password", passwd)
    }
}

func TestGetPasswdFromService(t *testing.T) {
    str, length := GetPasswdFromService("test", testFile)
    if str != hexpasswd || length != 8 {
        t.Error("password is uncorrect")
    }

    str, _ = GetPasswdFromService("hoge", testFile)
    if str != "" {
        t.Errorf("password not registered: %s", str)
    }

    str, _ = GetPasswdFromService(":", testFile)
    if str != "" {
        t.Errorf("password not registered: %s", str)
    }

    str, _ = GetPasswdFromService("", testFile)
    if str != "" {
        t.Errorf("password not registered: %s", str)
    }

    str, _ = GetPasswdFromService(".*", testFile)
    if str != "" {
        t.Errorf("password not registered: %s", str)
    }
}
