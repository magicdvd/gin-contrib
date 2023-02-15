# validator

## 替换默认validator from v8 => v9

```golang
import "github.com/gin-gonic/gin/binding"
import "gitlab.com/silenthill/gin-contrib/validatoroverriding"
import "gopkg.in/go-playground/validator.v9"

func main() {
    //不带自定义函数的初始化 都用默认的
    binding.Validator = validatoroverriding.NewValidatorV9()

    binding.Validator = validatoroverriding.NewValidatorV9(func(validator *validate.Validate){
        //自定义tag, 处理国际化等等等等
    })
    // regular gin logic
}
```

## 国际化(i18n) 目前定制化只支持中文

```golang

    //替换v9的地方做绑定
    binding.Validator = validatoroverriding.NewValidatorV9(func(validator *validate.Validate){
        validatoroverriding.BindValidator(validator)
    })

    //显示错误的地方
    type QueryStatistic struct {
        Name `form:"name" trans:"名字"` //定义绑定结构的时候,tag里使用"trans"
    }

    q := QueryStatistic{}
    if err := c.ShouldBind(&q); err != nil {
        c.JSON(http.StatusBadRequest, validatoroverriding.TranslateFirst(err))
        return
    }

```
