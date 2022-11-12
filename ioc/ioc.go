package ioc

import (
	"reflect"
	"sync"
)

var container *sync.Map = &sync.Map{}

func Set[T any](singletonFunc func() T) {
	singletonObj := singletonFunc()
	singletonType := reflect.TypeOf(&singletonObj)
	container.Store(singletonType.String(), &singletonObj)
}

func Get[T any](in T) *T {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	singletonType := reflect.TypeOf(&in)
	singletonObjAny, ok := container.Load(singletonType.String())
	if !ok {
		return nil
	}

	singletonObj, ok := singletonObjAny.(*T)
	return singletonObj
}

func Reset() {
	container = &sync.Map{}
}
