package validatoroverriding

import (
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type ValidatorV9 struct {
	once     sync.Once
	validate *validator.Validate
	initFunc func(validate *validator.Validate) *validator.Validate
}

func NewValidatorV9(initFunc ...func(validate *validator.Validate) *validator.Validate) *ValidatorV9 {
	fn := func(validate *validator.Validate) *validator.Validate { return validate }
	if len(initFunc) > 0 {
		fn = initFunc[0]
	}
	a := new(ValidatorV9)
	a.initFunc = fn
	return a
}

var _ binding.StructValidator = &ValidatorV9{}

func (v *ValidatorV9) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *ValidatorV9) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *ValidatorV9) lazyinit() {
	v.once.Do(func() {
		v.validate = v.initFunc(validator.New())
		v.validate.SetTagName("binding")
		// //调用配置的初始化函数
		// fmt.Println("init")

	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
