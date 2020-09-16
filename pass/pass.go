package pass

import(
	"bytes"
	"io/ioutil"
	"github.com/surdeus/gowik/str"
	"github.com/surdeus/gowik/path"
)

func
IsCorrect(user, passToCheck string) bool {
	passOnServerBytes, err := ioutil.ReadFile(path.HashFile(user))
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
