package simple_factory

/**
go中没有构造函数,一般用newXX表示返回一个类,该函数可以是唯一对包外可见的,隐藏具体的实现细节。
*/

type API interface {
	say(s string) string
}

type Hi struct{}

func (h *Hi) say(words string) string {
	return "hi " + words
}

type Hello struct{}

func (h *Hello) say(words string) string {
	return "hello " + words
}

func NewApi(t int) API {
	switch t {
	case 1:
		return &Hi{}
	case 2:
		return &Hello{}
	}
	return nil
}
