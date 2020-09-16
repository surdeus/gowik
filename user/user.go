package user

import(
	"github.com/surdeus/gowik/file"
)

func
Exist(user string) bool {
	return file.Exist(path.HashFile(user))
}
