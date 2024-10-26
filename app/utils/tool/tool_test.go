package tool

import (
	"testing"
)

func TestSendMessage(t *testing.T) {
	title, content := "QiuBlogBot:您有一条新的留言", "昵称: 嘤嘤嘤\nQQ: 阿巴巴巴\n邮箱: 阿啦啦啦啦\n内容: 寄"
	t.Errorf("err:%v", Send(title, content))
}
