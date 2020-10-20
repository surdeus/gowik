package markdown

import(
	bf "gopkg.in/russross/blackfriday.v2"
)

func
Process(input []byte) []byte {
	return bf.Run(input)
}
