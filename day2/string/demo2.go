package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

/**
	字符串常用操作
 */
const ENSTRING  = "go lange study"
const CNSTRING  = "go 学习总结"

//计算字符串长度
func strLen()  {
	fmt.Println(len(ENSTRING))
	fmt.Println(len(CNSTRING))
	fmt.Println(len([]rune(CNSTRING))) //将类型转换成rune，这样就不用考虑中英文字节长度问题了
	fmt.Println(utf8.RuneCountInString(CNSTRING))//对于中文字符，就需要使用utf8.RuneCountInString()的这个方法来统计
}

//截取字符串长度(包括中英文)
func getSubString(str string,length int)  {
	rs := []rune(str)  //类型rune，使用它完全不用考虑unicode字节问题，一个中文就只站一个数组下标
	fmt.Println(len(rs))
	fmt.Println(string(rs[:length]))
}

//是否存在某个字符或子串
func hasContains()  {
	fmt.Println(strings.Contains(ENSTRING,"st"))
	fmt.Println(strings.ContainsAny(ENSTRING,"cn du"))
	fmt.Println(strings.ContainsAny(ENSTRING,"c f"))
	fmt.Println(strings.ContainsRune(ENSTRING,'s'))
	fmt.Println(strings.ContainsRune(CNSTRING,'习'))
	fmt.Println(strings.Contains(CNSTRING,"学"))
}

//子串出现次数
func strCount()  {
	fmt.Println(strings.Count("abbabaab","ab"))
	fmt.Println(strings.Count("five","")) //当第二个参数为空时，Count 的返回值是：utf8.RuneCountInString(s) + 1
}

//字符串分割，下面的这些函数所得到的都是一个字符串切片类型：[]string
func strSplit()  {
	fmt.Println(strings.Fields(ENSTRING))   //Fields()用一个或多个连续的空格分隔字符串 s，返回子字符串的数组（slice）
	fmt.Println(strings.FieldsFunc(ENSTRING,unicode.IsSpace))     //strings.Fields()其实就是通过这种方式实现的。
	fmt.Println(strings.Split("java_php_go_python","_")) //使用指定的分隔符来进行字符串的拆分
	fmt.Println(strings.SplitAfter("java_php_go_python","_")) //这个方法和上面的方法区别，只是在于分割后的内容是否带上分隔符
	fmt.Println(strings.SplitN("java,php,python,go,c#,c++",",",3))
	fmt.Println(strings.SplitAfterN("java,php,python,go,c#,c++",",",3))
}
//下面的这个方法可以传入strings.FieldsFunc()函数的第二个参数。可以根据自己的业务来实现里面的逻辑
func diyFiledFunc(r rune) bool {
	return false
}
//strings.HasPrefix()和strings.HasSuffix()返回值为bool类型
func hasPrefixSuffix()  {
	fmt.Println(strings.HasPrefix("jh_article","jh"))
	fmt.Println(strings.HasPrefix("jh_article","zh"))
	fmt.Println(strings.HasSuffix("demo.png","png"))
	fmt.Println(strings.HasSuffix("demo.png","jpg"))
}

//检索字符或字符串出现的位置。注意，这里获取的值下标是从0开始的。如果没有找到的话，返回值是-1
func strIndex()  {
	fmt.Println(strings.Index(ENSTRING,"g"))
	fmt.Println(strings.Index(CNSTRING,"总"))  //如果是中文，得到的位置数是该字符位置的字节数，并不是实际的按一个个的字符来数的位置
	fmt.Println(strings.IndexRune(CNSTRING,'总'))
	fmt.Println(strings.IndexAny(ENSTRING,"d a cd"))
	fmt.Println(strings.IndexFunc(ENSTRING,unicode.IsSpace))

	fmt.Println(strings.LastIndex(ENSTRING,"s"))
	fmt.Println(strings.LastIndexAny(ENSTRING,"d a c"))
	fmt.Println(strings.LastIndexByte(ENSTRING,'e'))
	fmt.Println(strings.LastIndexFunc(ENSTRING,unicode.IsSpace))
}

//字符串的联合操作
func strJoin()  {
	fmt.Println(strings.Join([]string{"php","java","python"},ENSTRING))
}

//字符串重复次数
func strRepeat()  {
	fmt.Println(strings.Repeat("ab",5))//将参数1重复参数2的次数
}

//字符串替换
func strReplace()  {
	//参数1：要被替换的字符串
	//参数2：被替换的字符字串
	//参数3：被替换的新的字符串
	//参数4：替换次数，小于0的情况下，就是全部替换
	fmt.Println(strings.Replace(ENSTRING,"go","GO",-1))

	r := strings.NewReplacer("<","<",">",">")
	fmt.Println(r.Replace("this is <b>HTML</b>!"))
}

//大小写转换
func strUpLower()  {
	fmt.Println(strings.ToLower(ENSTRING))
	fmt.Println(strings.ToUpper(ENSTRING))
	var sc unicode.SpecialCase
	fmt.Println(strings.ToLowerSpecial(sc,ENSTRING))  //该函数把s字符串里面的每个单词转化为大写，但是调用的是unicode.SpecialCase的ToUpper方法
	// 定义转换规则
	var _MyCase = unicode.SpecialCase{
		// 将半角逗号替换为全角逗号，ToTitle 不处理
		unicode.CaseRange{',', ',', [unicode.MaxCase]rune{'，' - ',', '，' - ',', 0}},
		// 将半角句号替换为全角句号，ToTitle 不处理
		unicode.CaseRange{'.', '.', [unicode.MaxCase]rune{'。' - '.', '。' - '.', 0}},
		// 将 ABC 分别替换为全角的 ＡＢＣ、ａｂｃ，ToTitle 不处理
		unicode.CaseRange{'A', 'C', [unicode.MaxCase]rune{'Ａ' - 'A', 'ａ' - 'A', 0}},
	}
	fmt.Println(strings.ToLowerSpecial(_MyCase,"ABCDEF,abcdef."))
}

//去除空格
func strTrim()  {
	fmt.Println(strings.TrimSpace("  this is a demo  "))
	fmt.Println(strings.Trim("  this is a demo  "," "))
	fmt.Println(strings.Trim("__php_java_c++_go","__"))
	fmt.Println(strings.TrimLeft("  this is a demo----"," "))
	fmt.Println(strings.TrimRight("  this is a demo----   ",""))
	fmt.Println(strings.TrimPrefix("http://www.baidu.com","http://"))
	fmt.Println(strings.TrimSuffix("demo.jpg",".jpg"))
	fmt.Println(strings.TrimFunc("  this is a demo  ",unicode.IsSpace))
	fmt.Println(strings.TrimLeftFunc("  this is a demo  ",unicode.IsSpace))
	fmt.Println(strings.TrimRightFunc("  this is a demo  ",unicode.IsSpace))
}

func main() {
	//strLen()
	//getSubString(ENSTRING,5)
	//getSubString(CNSTRING,5)
	//hasContains()
	//fmt.Println(unicode.ToLower('c'),unicode.ToLower('c'))
	//strCount()
	//strSplit()
	//hasPrefixSuffix()
	//strIndex()
	//strJoin()
	//strRepeat()
	//strReplace()
	//strUpLower()
	strTrim()
}