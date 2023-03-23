package ioc

import (
	"log"
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

func Resolve[T any](obj T) T {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ioc: %s\n", err)
		}
	}()

	valueof := reflect.ValueOf(obj)
	typeof := valueof.Type()

	max := typeof.NumField()
	for i := 0; i < max; i++ {
		field := typeof.Field(i)
		if field.Tag.Get("ioc") == "inject" {
			singletonObjAny, ok := container.Load(field.Type.String())
			if !ok {
				panic("not found")
			}

			fieldElem := valueof.Field(i).Elem()
			if fieldElem.CanSet() {
				fieldElem.Set(reflect.ValueOf(singletonObjAny).Elem())
			}
		}
	}

	return obj
}
