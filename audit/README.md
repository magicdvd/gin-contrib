# 使用方法

## Middleware

```golang
//一般用法 （在需要加审计日志的地方加上）
r := gin.New()
logger := log.NewStdLogger("debug")
r.GET("/path/", audit.MWAuditlog("操作名称"), handler.YOURROUTEHANDLER)


//自定义审计日志
func YOURROUTEHANDLER(ctx *gin.Context) {
    ...
    err := audit.Customize().
        SetCondition("asdfasdfasdf").
        SetResult("asdfasdfasdf").
        Do(ctx)
    ...
}
