package file

import(
	"os"
)

func
RegularExist(p string) bool {
	if fi, e := os.Stat(p) ; e==nil  {
		if fi.Mode().IsRegular() {
			return true	
		}
	}
	return false
}

func
DirExist(p string) bool {
	if fi, e := os.Stat(p) ; e==nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
	
}
