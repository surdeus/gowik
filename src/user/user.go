package user

import(
	"os"
	"bytes"
	"io/ioutil"
	"github.com/surdeus/gowik/src/str"
	"github.com/surdeus/gowik/src/file"
	"github.com/surdeus/gowik/src/path"
)

func
Exist(u string) bool {
	return file.RegularExist(path.HashFile(u)) && file.DirExist(path.UserDir(u))
}

func
Create(u, p string) error {
	return ioutil.WriteFile(path.HashFile(u), []byte(p), 0644)
}

func
ChangePassword(u, p string) error {
	return ioutil.WriteFile(path.HashFile(u), []byte(p), 0644)
}

func
Delete(u string) error {
	if e := os.RemoveAll(path.UserDir(u)) ; e != nil {
		return e
	}
	if e := os.RemoveAll(path.HashFile(u)) ; e != nil {
		return e
	}
	return nil
}

func
IsPasswordCorrect(u, passToCheck string) bool {
	passOnServerBytes, err := ioutil.ReadFile(path.HashFile(u))
	if err!=nil {
		return false	
	}
	passOnServer := str.Chop(string(passOnServerBytes), 1)
	
	if bytes.Equal([]byte(passOnServer), []byte(passToCheck)) {
		return true
	} else {
		return false
	}
}
