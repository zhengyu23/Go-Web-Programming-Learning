package __4_1_Go与Cookie

import "time"

// 代码清单4-12	Cookie结构的定义

type Cookie struct {
	Name    string
	Value   string
	Path    string
	Domain  string
	Expires time.Time // 没有设置EXpires字段的cookie通常称为会话cookie，在浏览器关闭时会自动被移除
	// 相对的，设置了Expires字段的cookie被称为持久cookie，会一直存在直到过期或被删除
	// 明确指出什么时候会过期
	RawExpires string
	MaxAge     int // 执明被浏览器创建之后能存活多少秒
	Secure     bool
	HttpOnly   bool
	Raw        string
	Unparsed   []string
}
