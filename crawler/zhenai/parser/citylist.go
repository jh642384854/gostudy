package parser

import (
	"dev/crawler/engine"
	"regexp"
)
/**
	这个功能是用来解析城市地址的
 */
//返回的结果是Request，这个对象里面包含了request和items
func ParseCityList(content []byte) engine.ParseResult  {
	regExpression := `<a href="(http://www.zhenai.com/zhenghun/[a-z]+)"[^>]*>([^<]*)</a>`
	reg := regexp.MustCompile(regExpression)
	citys := reg.FindAllSubmatch(content,-1)
	result := engine.ParseResult{}
	for _, city := range citys {
		name := city[2]
		result.Items = append(result.Items,string(name))//城市名称
		result.Requests = append(result.Requests,engine.Request{
			Url:string(city[1]),
			ParserFun:ParseCity,
		})
	}
	return result
}