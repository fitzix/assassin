package service

// A-B-CC
const (
	StatusSuccess                           = 0
	StatusWebBadRequest                     = 3000
	StatusWebOperateErr                     = 3001
	StatusWebParamErr                       = 3002
	StatusWebResourceNotInTenant            = 3003
	StatusWebLabelNotExist                  = 3100 // 标签 01
	StatusWebLabelLinkedFace                = 3101
	StatusWebFaceNotExit                    = 3200 // 人员 02
	StatusWebFaceImgNotSupport              = 3201
	StatusWebFaceLinkedDevice               = 3202
	StatusWebDevice                         = 3300 // 设备 03
	StatusWebRoleNotExist                   = 3400 // 角色 04
	StatusWebRoleLinkedUser                 = 3401
	StatusWebAuthUnauthorized               = 3500 // 权限 06
	StatusWebAuthWrongPwd                   = 3501
	StatusWebAuthExpired                    = 3502
	StatusWebAuthErrToken                   = 3503
	StatusWebAuthWrongPermissionList        = 3504
	StatusWebAuthCannotModifySelfPermission = 3505
	StatusWebUserNotExist                   = 3601 // 用户
	StatusWebUserCannotDeleteSelf           = 3602
	StatusWebUserNameExist                  = 3603
)

var statusText = map[int]string{
	StatusSuccess:                           "ok",
	StatusWebBadRequest:                     "请求出错",
	StatusWebOperateErr:                     "操作失败",
	StatusWebParamErr:                       "参数错误",
	StatusWebResourceNotInTenant:            "非法操作",  // 不可访问非本租户资源
	StatusWebLabelNotExist:                  "标签不存在", // 标签
	StatusWebLabelLinkedFace:                "该标签已和用户关联，请先解除所有关联",
	StatusWebFaceNotExit:                    "人员不存在", // 人脸
	StatusWebFaceImgNotSupport:              "图片格式或大小有误",
	StatusWebFaceLinkedDevice:               "该人员已和设备关联, 请先解除设备",
	StatusWebRoleNotExist:                   "角色不存在", // 角色
	StatusWebRoleLinkedUser:                 "该角色已和用户关联，请先解除所有关联",
	StatusWebAuthUnauthorized:               "没有权限",
	StatusWebAuthWrongPwd:                   "用户名或密码错误",
	StatusWebAuthErrToken:                   "签名令牌无效",
	StatusWebAuthWrongPermissionList:        "所选权限列表有误",
	StatusWebAuthCannotModifySelfPermission: "不能修改自己的权限",
	StatusWebUserNotExist:                   "用户不存在", // 用户
	StatusWebUserCannotDeleteSelf:           "不能删除自己",
	StatusWebUserNameExist:                  "用户名已存在",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func AsnStatusText(code int) string {
	return statusText[code]
}
