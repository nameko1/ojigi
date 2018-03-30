package ojigi

import (
    "testing"
)

var testFile = "./testFile.txt"

func TestGetPasswdFromService(t *testing.T) {
    str := GetPasswdFromService("hoge", testFile)
    if str != "hogepasswd" {
        t.Error("password is uncorrect")
    }

    str = GetPasswdFromService("namekohoge", testFile)
    if str != "" {
        t.Errorf("password not registered: %s", str)
    }

    str = GetPasswdFromService("namekopass", testFile)
    if str != "" {
        t.Errorf("password not registered: %s", str)
    }

    str = GetPasswdFromService("", testFile)
    if str != "" {
        t.Errorf("password not registered: %s", str)
    }

    str = GetPasswdFromService(".*", testFile)
    if str != "" {
        t.Errorf("password not registered: %s", str)
    }
}
