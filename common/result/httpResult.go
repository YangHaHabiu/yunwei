package result

import (
	"fmt"
	"net/http"
	"ywadmin-v3/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回

		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试，原因：" + err.Error()
		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		errcodeMap := map[uint32]bool{
			100003: true,
			100012: true,
			100013: true,
			200003: true,
		}
		errmsg = fmt.Sprintf("errcode:%d,errmsg:%s", errcode, errmsg)
		if _, ok := errcodeMap[errcode]; ok {
			httpx.WriteJson(w, 200, Error(401, errmsg))
		} else {
			httpx.WriteJson(w, 200, Error(400, errmsg))
		}
	}
}

// http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("errcode:%d,errmsg:%s,%s", xerr.REUQEST_PARAM_ERROR, xerr.MapErrMsg(xerr.REUQEST_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, 200, Error(400, errMsg))
}
