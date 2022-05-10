package validatil

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestExecValid(t *testing.T) {
	fmt.Println(len("地球"))
	fmt.Println(len("go"))
	fmt.Println(strings.Count("地球", "") - 1)
	fmt.Println("是否满足最小值要求", ExecValid("100", "min", "120"))
	fmt.Println("是否满足最大值要求", ExecValid("100", "max", "90"))
	fmt.Println("是否满足最大长度要求", ExecValid("123sdaswerew", "maxlen", "5"))
	fmt.Println("是否满足最小长度要求", ExecValid("21312dsafadf", "minlen", "50"))
	fmt.Println("是否满足指定长度", ExecValid("13fadfwerwr", "len", "50"))
	fmt.Println("邮箱格式验证", ExecValid("1231das#qq.com", "email"), ExecValid("1231das@qq.com", "email"))
	fmt.Println("座机", ExecValid("0771-6772237", "tel"))
	fmt.Println("座机", ExecValid("0771-677223711", "tel"))
	fmt.Println("手机号码", ExecValid("13687717717", "mobile"))
	fmt.Println("手机号码", ExecValid("53687717717", "mobile"))
	fmt.Println("enum", ExecValid("123", "enum", "1234", "457"))
	fmt.Println("range", ExecValid("100", "range", "1", "10"))
	fmt.Println("邮政编码", ExecValid("518000", "zipcode"))
	fmt.Println("邮政编码", ExecValid("5180001", "zipcode"))
	fmt.Println("IP", ExecValid("127.0.0.1", "ip"))
	fmt.Println("IP", ExecValid("1809.1.1.1", "ip"))
	fmt.Println("字母验证", ExecValid("12312asdasda", "alpha"))
	fmt.Println("字母验证", ExecValid("asADASDdasda", "alpha"))
	fmt.Println("是否是数字", ExecValid("123123", "numeric"))
	fmt.Println("是否是数字", ExecValid("12312.23", "numeric"))
	fmt.Println("数字字母", ExecValid("1231212asdfasdfAASDAS", "alphanumeric"))
	fmt.Println("数字字母", ExecValid("1231212asd..fasdfAASDAS", "alphanumeric"))
	fmt.Println("数字字母横线", ExecValid("1231212asd|\asfasdfAASDAS", "alphadash"))
	fmt.Println("数字字母横线", ExecValid("1231212-_dfAASDAS", "alphadash"))
	fmt.Println("正则规则验证", ExecValid("pe111?sach", "regexp", "p([a-z]+)ch"))

	var vals url.Values
	rules := map[string][]string{
		"Name":    []string{"len:10", "alpha", "required"},
		"Age":     []string{"range:0:100", "int", "required"},
		"Address": []string{"required"},
	}
	params, errs := Valid(vals, rules)
	fmt.Println(params, errs)

}
