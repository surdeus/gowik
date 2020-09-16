package file
import(
	"os"
)

func
Exist(file string) bool {
	if _, err os.Stat(path) ; err==nil {
		return true
	} else {
		return false
	}
}