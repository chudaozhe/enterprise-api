package core

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

//gin > 1.4.0

//将验证器错误翻译成中文

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func InitTrans(locale string) {
	//注册翻译器
	uni = ut.New(zh.New(), en.New())
	trans, _ = uni.GetTranslator(locale)

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	//注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

// Translate 翻译错误信息
/**
func Translate(err error) map[string][]string {
	var result = make(map[string][]string)
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	return result
}
*/

func Translate(err error) string {
	var result []string
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result = append(result, err.Translate(trans))
	}
	return strings.Join(result, "; ")
}
