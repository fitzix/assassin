package service

const (
	tmplDefault = ""
	tmplAppDesc = ``
	tmplArticle = ``
)

var tmpl = map[int]string{
	AsnUploadTypeApp: tmplAppDesc,
	AsnUploadTypeArticle: tmplArticle,
}

func GetTmplContent(t int) []byte {
	if content, ok := tmpl[t]; ok {
		return []byte(content)
	}
	return []byte{}
}
