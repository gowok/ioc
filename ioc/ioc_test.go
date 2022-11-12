package ioc

import (
	"reflect"
	"testing"

	"github.com/golang-must/must"
)

type testUserRepository struct {
}

func (repo *testUserRepository) GetUsers() []string {
	return []string{"Alex Under", "John Thor"}
}

type testUserServive struct {
	userRepository testUserRepository
}

func TestSet(t *testing.T) {
	must := must.New(t)

	Set(func() *testUserRepository {
		return &testUserRepository{}
	})

	singletonType := reflect.TypeOf(&testUserRepository{})
	singletonObj, ok := container[singletonType.String()]

	must.NotNil(singletonObj)
	must.True(ok)
}

func TestGet(t *testing.T) {
	must := must.New(t)

	singletonObj := testUserRepository{}
	singletonType := reflect.TypeOf(&singletonObj)
	container[singletonType.String()] = &singletonObj

	singletonObjFromContainer := Get(singletonObj)
	must.NotNil(singletonObjFromContainer)
}

func TestReset(t *testing.T) {
	must := must.New(t)

	Set(func() *testUserRepository {
		return &testUserRepository{}
	})

	Reset()

	singletonObj := Get(testUserRepository{})
	must.Nil(singletonObj)
}
