package ioc

type Key string
type BindFunc func() interface{}

var container = map[Key]interface{}{}

func Bind(key Key, bf BindFunc) {
	container[key] = bf()
}

func Use(key Key) (interface{}, bool) {
	obj, ok := container[key]
	return obj, ok
}
