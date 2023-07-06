package captcha

import (
	"context"
	"encoding/json"
	"fmt"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"ywadmin-v3/common/constant"
	img "ywadmin-v3/common/tool/image"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/common/tool/util"

	"github.com/TestsLing/aj-captcha-go/model/vo"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetWordCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

const (
	TEXT = "的一了是我不在人们有来他这上着个地到大里说就去子得也和那要下看天时过出小么起你都把好还多没为又可家学只以主会样年想生同老中十从自面前头道它后然走很像见两用她国动进成回什边作对开而己些现山民候经发工向事命给长水几义三声于高手知理眼志点心战二问但身方实吃做叫当住听革打呢真全才四已所敌之最光产情路分总条白话东席次亲如被花口放儿常气五第使写军吧文运再果怎定许快明行因别飞外树物活部门无往船望新带队先力完却站代员机更九您每风级跟笑啊孩万少直意夜比阶连车重便斗马哪化太指变社似士者干石满日决百原拿群究各六本思解立河村八难早论吗根共让相研今其书坐接应关信觉步反处记将千找争领或师结块跑谁草越字加脚紧爱等习阵怕月青半火法题建赶位唱海七女任件感准张团屋离色脸片科倒睛利世刚且由送切星导晚表够整认响雪流未场该并底深刻平伟忙提确近亮轻讲农古黑告界拉名呀土清阳照办史改历转画造嘴此治北必服雨穿内识验传业菜爬睡兴形量咱观苦体众通冲合破友度术饭公旁房极南枪读沙岁线野坚空收算至政城劳落钱特围弟胜教热展包歌类渐强数乡呼性音答哥际旧神座章帮啦受系令跳非何牛取入岸敢掉忽种装顶急林停息句区衣般报叶压慢叔背细"
)

func NewGetWordCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWordCaptchaLogic {
	return &GetWordCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWordCaptchaLogic) GetWordCaptcha() (resp *types.GetWordCaptchaResp, err error) {
	// 初始化背景图片
	backgroundImage := img.GetClickBackgroundImage()

	pointList, wordList, err := l.getImageData(backgroundImage)
	if err != nil {
		return nil, err
	}

	originalImageBase64 := backgroundImage.Base64()

	if err != nil {
		return nil, err
	}
	token := util.GetUuid()
	codeKey := fmt.Sprintf(constant.CodeKeyPrefix, token)
	jsonPoint, err := json.Marshal(pointList)
	if err != nil {
		return nil, xerr.NewErrMsg("获取验证码失败，原因：" + err.Error())
	}
	//设置redis超时时间
	err = l.svcCtx.RedisClient.Setex(codeKey, string(jsonPoint), 3*60)
	if err != nil {
		return nil, xerr.NewErrMsg("缓存验证码信息失败，原因：" + err.Error())
	}
	resp = new(types.GetWordCaptchaResp)
	resp.ClientUid = token
	resp.SecretKey = pointList[0].SecretKey
	resp.WordList = wordList
	resp.OriginalImageBase64 = originalImageBase64

	return
}

func (l *GetWordCaptchaLogic) getImageData(image *util.ImageUtil) ([]vo.PointVO, []string, error) {
	wordCount := 4

	// 某个字不参与校验
	num := util.RandomInt(1, wordCount)
	currentWord := l.getRandomWords(wordCount)

	var pointList []vo.PointVO
	var wordList []string

	i := 0

	// 构建本次的 secret
	key := util.RandString(16)

	for _, s := range currentWord {
		point := l.randomWordPoint(image.Width, image.Height, i, wordCount)
		point.SetSecretKey(key)
		// 随机设置文字 TODO 角度未设置
		image.SetArtText(s, 18, point)
		// if err != nil {
		// 	return nil, nil, err
		// }

		if (num - 1) != i {
			pointList = append(pointList, point)
			wordList = append(wordList, s)
		}
		i++
	}
	return pointList, wordList, nil
}

// getRandomWords 获取随机文件
func (l *GetWordCaptchaLogic) getRandomWords(count int) []string {
	runesArray := []rune(TEXT)
	size := len(runesArray)

	set := make(map[string]bool)
	var wordList []string

	for {
		word := runesArray[util.RandomInt(0, size-1)]
		set[string(word)] = true
		if len(set) >= count {
			for str, _ := range set {
				wordList = append(wordList, str)
			}
			break
		}
	}
	return wordList
}

func (l *GetWordCaptchaLogic) randomWordPoint(width int, height int, i int, count int) vo.PointVO {
	avgWidth := width / (count + 1)
	fontSizeHalf := 18 / 2

	var x, y int
	if avgWidth < fontSizeHalf {
		x = util.RandomInt(1+fontSizeHalf, width)
	} else {
		if i == 0 {
			x = util.RandomInt(1+fontSizeHalf, avgWidth*(i+1)-fontSizeHalf)
		} else {
			x = util.RandomInt(avgWidth*i+fontSizeHalf, avgWidth*(i+1)-fontSizeHalf)
		}
	}
	y = util.RandomInt(18, height-fontSizeHalf)
	return vo.PointVO{X: x, Y: y}
}
