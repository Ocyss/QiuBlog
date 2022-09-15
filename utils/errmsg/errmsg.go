package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000  User      用户表

	// code = 2000  Tags      标签表

	// code = 3000  Article   文章表

	// code = 4000  Category/Menuchild 分类/菜单子项表
	ERROR_PARM_SO = 4001
)

var codemsg = map[int]string{
	SUCCESS:       "OK",
	ERROR:         "ERROR",
	ERROR_PARM_SO: "缺少必要参数!",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
