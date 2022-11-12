package ioc

import (
	"reflect"
)

type IOC map[string]any

var container IOC = IOC{}

func Set[T any](singletonFunc func() T) {
	singletonObj := singletonFunc()
	singletonType := reflect.TypeOf(singletonObj)
	container[singletonType.String()] = singletonObj
}

func Get[T any](in T) T {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	singletonType := reflect.TypeOf(in)
	singletonObj := container[singletonType.String()]
	return singletonObj.(T)
}
