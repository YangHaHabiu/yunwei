package reportLogDb

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/myip"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/api/internal/svc"
)

type RecordOperationLogMiddleware struct {
	svcCtx *svc.ServiceContext
}

func NewRecordOperationLogMiddleware(svc *svc.ServiceContext) *RecordOperationLogMiddleware {
	return &RecordOperationLogMiddleware{
		svcCtx: svc,
	}
}

func (l *RecordOperationLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := r.Context()
		var (
			result    string
			operation string
		)
		split := strings.Split(r.RequestURI, "?")
		if len(split) > 0 {
			operation = split[0]
		}
		opers := fmt.Sprintf("/%s/%s/%s", strings.Split(operation, "/")[1], strings.Split(operation, "/")[2], strings.Split(operation, "/")[3])

		if !tool.StrInArr(opers, tool.TopThreeArray) && r.Method != "GET" {
			// 开始的时间
			startTime := time.Now()
			robots, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				fmt.Println(err)
			}
			if string(robots) == "" && len(split) > 1 {
				result = strings.Join(split[1:], "")
			} else {
				result = string(robots)
			}
			r.Body = ioutil.NopCloser(bytes.NewBuffer(robots))

			next(w, r)
			sub := time.Now().Sub(startTime)
			compile := regexp.MustCompile(`&_t=\d+|_t=\d+`)
			result = compile.ReplaceAllString(result, "")
			l.svcCtx.AdminRpc.SysLogAdd(m, &adminclient.SysLogAddReq{
				UserName:  ctxdata.GetUnameFromCtx(m),
				Ip:        myip.GetCurrentIP(r),
				Params:    result,
				Method:    r.Method,
				Operation: operation,
				Time:      float32(sub.Seconds()),
			})
		} else {
			next(w, r)
		}
	}
}
