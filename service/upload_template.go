package service

const (
	// tmplDefault = ""
	tmplAppDesc = " "
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

// func GetTmplContent(t int, data interface{}) []byte {
// 	tmplText := "> 应用介绍"
// 	if content, ok := tmpl[t]; ok {
// 		tmplText = content
// 	}
//
// 	tmpl , err := template.New("md-tmpl").Parse(tmplText)
// 	if err != nil {
// 		zapLogger.Sugar().Errorf("make md tmpl err: %s", err)
// 		return []byte(tmplText)
// 	}
// 	var resp bytes.Buffer
// 	if err:=tmpl.Execute(&resp, data); err != nil {
// 		zapLogger.Sugar().Errorf("make md tmpl err: %s", err)
// 		return []byte(tmplText)
// 	}
// 	return resp.Bytes()
// }
