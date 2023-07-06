package log

import (
	"context"
	"strconv"
	"strings"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xdate"

	"io/ioutil"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TerminalListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTerminalListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TerminalListLogic {
	return &TerminalListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TerminalListLogic) TerminalList(req *types.ListTerminalLogReq) (*types.ListTerminalResp, error) {

	var (
		files    string
		username string
	)

	if ctxdata.GetUnameFromCtx(l.ctx) == globalkey.SuperUserName {
		files = l.svcCtx.Config.RecPath.FullPath

	} else {
		files = l.svcCtx.Config.RecPath.FullPath + "/" + ctxdata.GetUnameFromCtx(l.ctx)
		username = ctxdata.GetUnameFromCtx(l.ctx)
	}

	tmp, err := getAllFile(files, username, req)
	if err != nil {
		return nil, xerr.NewErrMsg("递归查询失败，原因：" + err.Error())
	}

	return &types.ListTerminalResp{
		tmp,
	}, nil
}

// 递归获取指定目录下的所有文件及文件夹
func getAllFile(pathname, username string, req *types.ListTerminalLogReq) ([]*types.ListTerminalData, error) {
	dates := make([]string, 0)
	if req.DateRange != "" {
		object := strings.Split(req.DateRange, ",")
		dates = xdate.GetBetweenDates(object[0], object[1])
	} else {
		start, end := xdate.GetWeekDate()
		dates = xdate.GetBetweenDates(start, end)
	}

	result := make([]*types.ListTerminalData, 0)
	newresult := make([]*types.ListTerminalData, 0)
	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		logx.Errorf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}
	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp := make([]*types.ListTerminalData, 0)
			n, _ := strconv.Atoi(fi.Name())
			if n > 0 {
				if tool.StrInArr(fi.Name(), dates) {
					temp, err = getAllFile(fullname, "", req)
				}
			} else {
				temp, err = getAllFile(fullname, "", req)
			}
			if err != nil {
				logx.Errorf("读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err
			}
			if len(temp) > 0 {
				result = append(result, &types.ListTerminalData{
					Label:    fi.Name(),
					Children: temp,
				})

			}
		}
		if strings.HasSuffix(fi.Name(), ".cast") {
			result = append(result, &types.ListTerminalData{
				Label: fi.Name(),
				Value: fi.Name(),
			})
		}

	}
	if username != "" {
		newresult = append(newresult, &types.ListTerminalData{
			Label:    username,
			Children: result,
		})
	} else {
		newresult = result
	}

	return newresult, nil
}
