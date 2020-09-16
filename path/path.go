package path

var(
	DataDir string
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
HashFile(u string) string {
	return HashDir+"/"+u
}
