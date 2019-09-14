package service

const (
	AsnAppTypeApp = iota
	AsnAppTypeBook
)

var appTypeName = map[string]int{
	"app":  AsnAppTypeApp,
	"book": AsnAppTypeBook,
}

func AsnAppType(t string) int {
	if typeCode, ok := appTypeName[t]; ok {
		return typeCode
	}
	return 0
}
