package ask

import (
	"qiublog/utils/errmsg"
)

func ErrParam() (int, any) {
	code := errmsg.ERROR_PARAM
	return code, nil
}
