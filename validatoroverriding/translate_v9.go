package validatoroverriding

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"

	validator "gopkg.in/go-playground/validator.v9"
	zhtranslations "gopkg.in/go-playground/validator.v9/translations/zh"
)

type transBindingv9 struct {
	uni    *ut.UniversalTranslator
	trans  ut.Translator
	locale string
}

func (c *transBindingv9) setLocale(locale string) {
	if c.locale != locale {
		c.locale = locale
		c.trans, _ = c.uni.GetTranslator(c.locale)
	}
}

func (c *transBindingv9) translateFirst(es validator.ValidationErrors) string {
	return es[0].Translate(c.trans)
}

func (c *transBindingv9) bindValidator(validate *validator.Validate) error {
	err := zhtranslations.RegisterDefaultTranslations(validate, c.trans)
	if err != nil {
		return err
	}
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tag := fld.Name
		name := fld.Tag.Get("trans")
		if name == "" {
			return tag
		}
		return name
	})
	return nil
}

func (c *transBindingv9) registerCustomTranslate(tag string, text string, validate *validator.Validate) error {
	return validate.RegisterTranslation(tag, c.trans, func(ut ut.Translator) error {
		return ut.Add(tag, text, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
}
