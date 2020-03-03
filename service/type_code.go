package service

const (
	AsnAppTypeApp = iota
	AsnAppTypeBook

	AsnAppStatusUnPublish = iota
	AsnAppStatusPublish

	AsnAppOrderNew = iota
	AsnAppOrderHot

	AsnUploadTypeApp = iota
	AsnUploadTypeArticle
)

var appTypeName = map[string]int{
	"app":   AsnAppTypeApp,
	"book":  AsnAppTypeBook,
	"hot":   AsnAppOrderHot,
	"new":   AsnAppOrderNew,
	"pub":   AsnAppStatusPublish,
	"unpub": AsnAppStatusUnPublish,
}

func AsnType(t string) int {
	if typeCode, ok := appTypeName[t]; ok {
		return typeCode
	}
	return -1
}
