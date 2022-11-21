package ioc

import (
	"reflect"
	"testing"

	"github.com/golang-must/must"
)

type testUserRepository struct {
	count int
}

func (repo *testUserRepository) GetUsers() []string {
	return []string{"Alex Under", "John Thor"}
}

type testUserService struct {
	UserRepository *testUserRepository `ioc:"inject"`
}

func TestSet(t *testing.T) {
	must := must.New(t)

	Set(func() testUserRepository {
		return testUserRepository{}
	})

	singletonType := reflect.TypeOf(&testUserRepository{})
	singletonObj, ok := container.Load(singletonType.String())

	must.NotNil(singletonObj)
	must.True(ok)
}

func TestGet(t *testing.T) {
	must := must.New(t)

	singletonObj := testUserRepository{}
	singletonType := reflect.TypeOf(&singletonObj)
	container.Store(singletonType.String(), &singletonObj)

	singletonObjFromContainer := Get(singletonObj)
	must.NotNil(singletonObjFromContainer)
}

func TestReset(t *testing.T) {
	must := must.New(t)

	Set(func() testUserRepository {
		return testUserRepository{}
	})

	Reset()

	singletonObj := Get(testUserRepository{})
	must.Nil(singletonObj)
}

func TestInject(t *testing.T) {
	must := must.New(t)

	expected := 10
	Set(func() testUserRepository {
		return testUserRepository{expected}
	})

	service := testUserService{
		UserRepository: &testUserRepository{},
	}

	Inject(service)

	must.Equal(expected, service.UserRepository.count)
}
