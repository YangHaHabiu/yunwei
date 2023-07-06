package comm

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

//定义校验结构体
type InputArg struct {
	FileConfig  string `validate:"required" reg_error_info:"缺少参数-f (传入kubeconfig认证文件)"`
	Actions     string `validate:"oneof=create query update redeployment delete" reg_error_info:"缺少参数-a (create query update rolloutRestart delete 其中一个)"`
	Works       string `validate:"oneof=deployment statefullset" reg_error_info:"缺少参数-w (deployment statefullset 其中一个)"`
	YamlFile    string
	NameSpace   string
	ServiceName string
}

//打印错误日志
func ProcessErr(u interface{}, err error) string {
	if err == nil { //如果为nil 阐明校验通过
		return ""
	}
	invalid, ok := err.(*validator.InvalidValidationError) //如果是输出参数有效，则间接返回输出参数谬误
	if ok {
		return "输出参数谬误：" + invalid.Error()
	}
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field()                    //获取是哪个字段不合乎格局
		field, ok := reflect.TypeOf(u).FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("reg_error_info") //获取field对应的reg_error_info tag值
			return fieldName + ":" + errorInfo           //返回谬误
		} else {
			return "缺失reg_error_info"
		}
	}
	return ""
}
