package ask

import (
	"qiublog/utils/errmsg"
)

func ErrParam() (int, any) {
	return errmsg.ERROR_PARAM, nil
}
