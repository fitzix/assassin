package service

const (
	AsnAppTypeApp         = 0
	AsnAppTypeBook        = 1
	AsnAppStatusUnPublish = 0
	AsnAppStatusPublish   = 1
	AsnAppOrderHot        = 1
	AsnUploadTypeApp
	AsnUploadTypeArticle
)

var appTypeName = map[string]int8{
	"app":   AsnAppTypeApp,         // app类型 app
	"book":  AsnAppTypeBook,        // 书籍
	"hot":   AsnAppOrderHot,        // 排序 热度
	"pub":   AsnAppStatusPublish,   // app 状态 发布
	"unpub": AsnAppStatusUnPublish, // 未发布
}

func AsnTypeExist(t string) (exist bool) {
	_, exist = appTypeName[t]
	return
}

func AsnType(t string) int8 {
	rsp, ok := appTypeName[t]
	if ok {
		return rsp
	}
	return -1
}
