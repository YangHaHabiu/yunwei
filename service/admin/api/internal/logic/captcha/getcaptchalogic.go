package captcha

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TestsLing/aj-captcha-go/model/vo"
	"math"

	//"github.com/TestsLing/aj-captcha-go/util"
	"golang.org/x/image/colornames"
	"image/color"
	"ywadmin-v3/common/constant"
	"ywadmin-v3/common/tool/util"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	img "ywadmin-v3/common/tool/image"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	point  vo.PointVO
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.GetCaptchaResp, err error) {
	// 初始化背景图片
	backgroundImage := img.GetBackgroundImage()
	// 为背景图片设置水印
	backgroundImage.SetText(l.svcCtx.Config.VerificationCodeWatermark, 20, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	// 初始化模板图片
	templateImage := img.GetTemplateImage()
	l.pictureTemplatesCut(backgroundImage, templateImage)
	token := util.GetUuid()
	codeKey := fmt.Sprintf(constant.CodeKeyPrefix, token)
	jsonPoint, err := json.Marshal(l.point)
	if err != nil {
		return nil, xerr.NewErrMsg("获取验证码失败，原因：" + err.Error())
	}
	//设置redis超时时间
	err = l.svcCtx.RedisClient.Setex(codeKey, string(jsonPoint), 3*60)
	if err != nil {
		return nil, xerr.NewErrMsg("缓存验证码信息失败，原因：" + err.Error())
	}

	resp = new(types.GetCaptchaResp)
	resp.ClientUid = token
	resp.JigsawImageBase64 = templateImage.Base64()
	resp.OriginalImageBase64 = backgroundImage.Base64()
	resp.SecretKey = l.point.SecretKey

	return
}

func (l *GetCaptchaLogic) pictureTemplatesCut(backgroundImage *util.ImageUtil, templateImage *util.ImageUtil) {
	// 生成拼图坐标点
	l.generateJigsawPoint(backgroundImage, templateImage)
	// 裁剪模板图
	l.cutByTemplate(backgroundImage, templateImage, l.point.X, 0)

	// 插入干扰图
	for {
		newTemplateImage := img.GetTemplateImage()
		if newTemplateImage.Src != templateImage.Src {
			offsetX := util.RandomInt(0, backgroundImage.Width-newTemplateImage.Width-5)
			if math.Abs(float64(newTemplateImage.Width-offsetX)) > float64(newTemplateImage.Width/2) {
				l.interferenceByTemplate(backgroundImage, newTemplateImage, offsetX, l.point.Y)
				break
			}
		}
	}
}

// 插入干扰图
func (l *GetCaptchaLogic) interferenceByTemplate(backgroundImage *util.ImageUtil, templateImage *util.ImageUtil, x1 int, y1 int) {
	xLength := templateImage.Width
	yLength := templateImage.Height

	for x := 0; x < xLength; x++ {
		for y := 0; y < yLength; y++ {
			// 如果模板图像当前像素点不是透明色 copy源文件信息到目标图片中
			isOpacity := templateImage.IsOpacity(x, y)

			// 当前模板像素在背景图中的位置
			backgroundX := x + x1
			backgroundY := y + y1

			// 当不为透明时
			if !isOpacity {
				// 背景图区域模糊
				backgroundImage.VagueImage(backgroundX, backgroundY)
			}

			//防止数组越界判断
			if x == (xLength-1) || y == (yLength-1) {
				continue
			}

			rightOpacity := templateImage.IsOpacity(x+1, y)
			downOpacity := templateImage.IsOpacity(x, y+1)

			//描边处理，,取带像素和无像素的界点，判断该点是不是临界轮廓点,如果是设置该坐标像素是白色
			if (isOpacity && !rightOpacity) || (!isOpacity && rightOpacity) || (isOpacity && !downOpacity) || (!isOpacity && downOpacity) {
				backgroundImage.RgbaImage.SetRGBA(backgroundX, backgroundY, colornames.White)
			}
		}
	}
}

func (l *GetCaptchaLogic) cutByTemplate(backgroundImage *util.ImageUtil, templateImage *util.ImageUtil, x1, y1 int) {
	xLength := templateImage.Width
	yLength := templateImage.Height

	for x := 0; x < xLength; x++ {
		for y := 0; y < yLength; y++ {
			// 如果模板图像当前像素点不是透明色 copy源文件信息到目标图片中
			isOpacity := templateImage.IsOpacity(x, y)

			// 当前模板像素在背景图中的位置
			backgroundX := x + x1
			backgroundY := y + y1

			// 当不为透明时
			if !isOpacity {
				// 获取原图像素
				backgroundRgba := backgroundImage.RgbaImage.RGBAAt(backgroundX, backgroundY)
				// 将原图的像素扣到模板图上
				templateImage.SetPixel(backgroundRgba, x, y)
				// 背景图区域模糊
				backgroundImage.VagueImage(backgroundX, backgroundY)
			}

			//防止数组越界判断
			if x == (xLength-1) || y == (yLength-1) {
				continue
			}

			rightOpacity := templateImage.IsOpacity(x+1, y)
			downOpacity := templateImage.IsOpacity(x, y+1)

			//描边处理，,取带像素和无像素的界点，判断该点是不是临界轮廓点,如果是设置该坐标像素是白色
			if (isOpacity && !rightOpacity) || (!isOpacity && rightOpacity) || (isOpacity && !downOpacity) || (!isOpacity && downOpacity) {
				templateImage.RgbaImage.SetRGBA(x, y, colornames.White)
				backgroundImage.RgbaImage.SetRGBA(backgroundX, backgroundY, colornames.White)
			}
		}
	}
}

// 生成模板图在背景图中的随机坐标点
func (l *GetCaptchaLogic) generateJigsawPoint(backgroundImage *util.ImageUtil, templateImage *util.ImageUtil) {
	widthDifference := backgroundImage.Width - templateImage.Width
	heightDifference := backgroundImage.Height - templateImage.Height

	x, y := 0, 0

	if widthDifference <= 0 {
		x = 5
	} else {
		x = util.RandomInt(100, widthDifference-100)
	}
	if heightDifference <= 0 {
		y = 5
	} else {
		y = util.RandomInt(5, heightDifference)
	}
	point := vo.PointVO{X: x, Y: y}
	point.SetSecretKey(util.RandString(16))
	l.point = point
}
