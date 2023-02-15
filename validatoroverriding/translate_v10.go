package validatoroverriding

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"

	validator "github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

type transBindingv10 struct {
	uni    *ut.UniversalTranslator
	trans  ut.Translator
	locale string
}

func (c *transBindingv10) setLocale(locale string) {
	if c.locale != locale {
		c.locale = locale
		c.trans, _ = c.uni.GetTranslator(c.locale)
	}
}

func (c *transBindingv10) translateFirst(es validator.ValidationErrors) string {
	return es[0].Translate(c.trans)
}

func (c *transBindingv10) bindValidator(validate *validator.Validate) error {
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

func (c *transBindingv10) registerCustomTranslate(tag string, text string, validate *validator.Validate) error {
	return validate.RegisterTranslation(tag, c.trans, func(ut ut.Translator) error {
		return ut.Add(tag, text, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
}
