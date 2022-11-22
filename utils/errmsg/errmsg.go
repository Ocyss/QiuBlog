package errmsg

//goland:noinspection ALL
const (
	SUCCESS     = 200
	ERROR       = 500
	ERROR_PARAM = 601
	// code = 1000  User      用户表
	ERROR_PWDERR_WRONG     = 1001
	ERROR_MANYERR          = 1002
	ERROR_REPEAT           = 1003
	ERROR_TOKEN_EXIST      = 1101
	ERROR_TOKEN_RUNTIME    = 1102
	ERROR_TOKEN_WRONG      = 1103
	ERROR_TOKEN_TYPE_WRONG = 1104
	// code = 2000  Tags      标签表

	// code = 3000  Article   文章表

	// code = 4000  Category/Menuchild 分类/菜单子项表
	ERROR_PARM_SO  = 4001
	ERROR_SET_TYPE = 4002
	ERROR_SET_NEW  = 4003
	ERROR_SET_UP   = 4004
	ERROR_SET_RE   = 4005
	ERROR_SET_SO   = 4006
	//code = 5000 Upload 上传类
	ERROR_FILE_WRONG  = 5001
	ERROR_CLASS_WRONG = 5002
	//code = 6000 Message 消息类
	ERROR_DE_APPROVE = 6001
)

var codemsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "ERROR",
	ERROR_PARAM:            "参数不对",
	ERROR_PWDERR_WRONG:     "用户名或者密码有误",
	ERROR_MANYERR:          "错误次数过多",
	ERROR_REPEAT:           "注册失败：重复的用户名",
	ERROR_TOKEN_EXIST:      "Token不存在",
	ERROR_TOKEN_RUNTIME:    "Token过期",
	ERROR_TOKEN_WRONG:      "Token不正确",
	ERROR_TOKEN_TYPE_WRONG: "Token格式错误",
	ERROR_PARM_SO:          "缺少必要参数!",
	ERROR_FILE_WRONG:       "上传文件/格式有误",
	ERROR_CLASS_WRONG:      "上传类型不对",
	ERROR_SET_TYPE:         "设置菜单子项的类型不对",
	ERROR_SET_NEW:          "设置子菜单，新建错误",
	ERROR_SET_UP:           "设置子菜单，更新错误",
	ERROR_SET_RE:           "设置子菜单，删除错误",
	ERROR_SET_SO:           "设置子菜单，排序错误",
	ERROR_DE_APPROVE:       "不可反审核！",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
