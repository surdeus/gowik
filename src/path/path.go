package path

var(
	WebDir string
	DataDir string
	StaticDir string
	TmplDir string
	PageDir string
	HashDir string
	SaltFile string
)

func
PageFile(u string, t string) string {
	return PageDir+"/"+u+"/"+t 
}

func
UserDir(u string) string {
	return PageDir+"/"+u
}

func
HashFile(u string) string {
	return HashDir+"/"+u
}
