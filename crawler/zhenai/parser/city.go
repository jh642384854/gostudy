package parser

import (
	"dev/crawler/engine"
	"regexp"
)

/**
	这个是根据城市地址来抓取城市页面的用户信息
 */

//返回的结果是Request，这个对象里面包含了request和items
func ParseCity(content []byte) engine.ParseResult  {
	regExpression := `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
	reg := regexp.MustCompile(regExpression)
	citys := reg.FindAllSubmatch(content,-1)
	result := engine.ParseResult{}
	for _, city := range citys {
		name := city[2]
		result.Items = append(result.Items,string(name))//城市名称
		result.Requests = append(result.Requests,engine.Request{
			Url:string(city[1]),
			ParserFun:ParseUser,
		})
	}
	return result
}