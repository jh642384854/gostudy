package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/hero"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {

	app := iris.New()
	//定义宏方式一：
	app.Macros().Get("string").RegisterFunc("has", func(validNames []string) func(string) bool {
		return func(paramValue string) bool {
			for _, validName := range validNames {
				if validName == paramValue{
					return  true
				}
			}
			return  false
		}
	})

	//定义宏方式二：
	app.Macros().Register("slice","",false,true, func(paramValue string) (i interface{}, b bool) {
		return strings.Split(paramValue,"/"),true
	}).RegisterFunc("contains", func(expectedItems []string) func(paramValue []string) bool {
		sort.Strings(expectedItems)
		return func(paramValue []string) bool {
			if len(paramValue) != len(expectedItems){
				return  false
			}
			sort.Strings(paramValue)
			for i:=0; i<len(paramValue) ;i++  {
				if paramValue[i] != expectedItems[i]{
					return false
				}
			}
			return  true
		}
	})

	//ParamResolvers是特定go std或自定义类型的参数类型的全局参数解析器。
	context.ParamResolvers[reflect.TypeOf([]string{})] = func(paramIndex int) interface{} {
		return func(ctx context.Context) []string {
			//当您希望检索具有默认不支持的值类型的参数时，例如ctx.Params()。然后您可以使用' GetEntry '或' GetEntryAt '并将其下划线' ValueRaw '转换为所需类型。
			// 类型应该与宏的求值器函数相同(宏#寄存器上的最后一个参数)返回值。
			// When you want to retrieve a parameter with a value type that it is not supported by-default, such as ctx.Params().GetInt
			// then you can use the `GetEntry` or `GetEntryAt` and cast its underline `ValueRaw` to the desired type.
			// The type should be the same as the macro's evaluator function (last argument on the Macros#Register) return value.
			return ctx.Params().GetEntryAt(paramIndex).ValueRaw.([]string)
		}
	}
	// http://localhost:8080/static_validation/python
	app.Get("/static_validation/{name:string has([java,php,python])}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("hello  %s | the name should be 'java' or 'php' or 'python'",name)
	})

	// http://localhost:8080/test_slice_contains/value1/value2
	//这个Handler里面定义的myparam []string这个参数，是iris默认不支持的，这样就需要重新定义上面的context.ParamResolvers这个方法来实现[]string这个类型的定义
	app.Get("/test_slice_hero/{myparam:slice}",hero.Handler(func(myparam []string) string{
		return fmt.Sprintf("myparam value is :%#v \n",myparam)
	}))

	// http://localhost:8080/test_slice_contains/value1/value2
	app.Get("/test_slice_contains/{myparam:slice contains([value1,value2])}", func(ctx context.Context) {
		// When it is not a builtin function available to retrieve your value with the type you want, such as ctx.Params().GetInt
		// then you can use the `GetEntry.ValueRaw` to get the real value, which is set-ed by your macro above.
		myparam := ctx.Params().GetEntry("myparam").ValueRaw.([]string)
		ctx.Writef("myparam's value (a trailing path parameter type) is: %#v\n", myparam)
	})
	

	app.Run(iris.Addr(":8080"))
}
