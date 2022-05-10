package mdtil

import (
	"fmt"
	"testing"
)

func TestMd2html(t *testing.T) {
	var md string;
	md="```bash\n\n```\n\n若看到了“hello, world”信息，那么Go已被正确安装。\n\n# 4. 卸载 Go\n\n卸载Go，其实就是将前面安装Go的东西全部删除：\n\n- 1.删除 go 目录：\n\n\n\n```bash\nsudo rm -rf /usr/local/go\n```\n\n- 2.删除软链接：\n\n\n\n```go\nsudo rm -rf /usr/bin/go\n```\n\n# 5. 升级 Go 版本\n\n升级 Go 版本其实就是:\n\n1. 卸载之前安装的旧版本Go，\n2. 再安装新版本的Go。\n\n------\n\n**参考文章：**\n\n- 起步 - Go 编程语言: http://docscn.studygolang.com/doc/install\n- 如何使用Go编程: http://docscn.studygolang.com/doc/code.html\n- Ubuntu16.04下部署golang配置环境: http://www.aweb.cc/article/detail/id/583.html\n\n\n\n作者：AlbertGou\n链接：https://www.jianshu.com/p/f2a237de8f07\n来源：简书\n著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。"
	fmt.Print(Md2html(md))
}
