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

	Set(func() testUserRepository {
		return testUserRepository{}
	})

	singletonType := reflect.TypeOf(testUserRepository{})
	singletonObj, ok := container[singletonType.String()]
	must.NotNil(singletonObj)
	must.True(ok)
}

func TestGet(t *testing.T) {
	must := must.New(t)

	singletonType := reflect.TypeOf(testUserRepository{})
	container[singletonType.String()] = testUserRepository{}

	singletonObj := Get(testUserRepository{})
	must.NotNil(singletonObj)
}
