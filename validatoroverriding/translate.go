package validatoroverriding

import (
	"errors"
	"sync"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validatorv10 "github.com/go-playground/validator/v10"
	validatorv9 "gopkg.in/go-playground/validator.v9"
)

var once sync.Once
var tc *transCenter

type transCenter struct {
	v9  *transBindingv9
	v10 *transBindingv10
}

//TranslateFirst 把第一条错误翻译输出
func TranslateFirst(err error) string {
	switch s := err.(type) {
	case validatorv9.ValidationErrors:
		return getTransBinding().v9.translateFirst(s)
	case validatorv10.ValidationErrors:
		return getTransBinding().v10.translateFirst(s)
	default:
		return err.Error()
	}
}

//BindValidator 使验证器支持中文
func BindValidator(v interface{}) error {
	switch s := v.(type) {
	case *validatorv9.Validate:
		return getTransBinding().v9.bindValidator(s)
	case *validatorv10.Validate:
		return getTransBinding().v10.bindValidator(s)
	default:
		return errors.New("unsupported validate(only v9,v10)")
	}
}

func getTransBinding() *transCenter {
	once.Do(func() {
		tc = new(transCenter)
		zhlocale := zh.New()
		enlocale := en.New()
		tc.v9 = new(transBindingv9)
		tc.v9.uni = ut.New(enlocale, enlocale, zhlocale)
		tc.v9.setLocale("zh")
		tc.v10 = new(transBindingv10)
		tc.v10.uni = ut.New(enlocale, enlocale, zhlocale)
		tc.v10.setLocale("zh")
	})
	return tc
}

func RegisterCustomTranslate(tag string, text string, validate interface{}) error {
	switch s := validate.(type) {
	case *validatorv9.Validate:
		return getTransBinding().v9.registerCustomTranslate(tag, text, s)
	case *validatorv10.Validate:
		return getTransBinding().v10.registerCustomTranslate(tag, text, s)
	default:
		return errors.New("RegisterCustomTranslate unsupported validate(v9,v10")
	}
}
