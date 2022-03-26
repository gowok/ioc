package ioc

import (
	"testing"

	"github.com/golang-must/must"
)

func TestBind(t *testing.T) {

	t.Run("should save same value (primitive)", func(t *testing.T) {
		must := must.New(t)

		expectedKey := "secret"
		expectedVal := 260

		Bind(Key(expectedKey), func() interface{} {
			return expectedVal
		})

		actualVal, ok := container[Key(expectedKey)]
		must.True(ok)
		must.Equal(expectedVal, actualVal)
	})

	t.Run("should save same value (struct)", func(t *testing.T) {
		must := must.New(t)

		expectedKey := "thekey"
		expectedVal := struct {
			Name string
		}{
			"Alex Under",
		}

		Bind(Key(expectedKey), func() interface{} {
			return &expectedVal
		})

		actualVal, ok := container[Key(expectedKey)]
		must.True(ok)
		must.Equal(&expectedVal, actualVal)
	})

	t.Run("can re-use bound dependencies", func(t *testing.T) {
		must := must.New(t)

		boundKey := "name"
		boundVal := "Alex Under"
		expectedKey := "thekey"

		Bind(Key(boundKey), func() interface{} {
			return boundVal
		})

		type User struct {
			Name string
		}

		Bind(Key(expectedKey), func() interface{} {
			name, ok := Use(Key(boundKey))
			must.True(ok)
			return User{
				name.(string),
			}
		})

		actualVal, ok := container[Key(expectedKey)]
		must.True(ok)
		must.Equal(boundVal, actualVal.(User).Name)
	})

}

func TestUse(t *testing.T) {

	t.Run("should returns same value (primitive)", func(t *testing.T) {
		must := must.New(t)

		expectedKey := "secret"
		expectedVal := 260

		Bind(Key(expectedKey), func() interface{} {
			return expectedVal
		})

		actualVal, ok := Use(Key(expectedKey))
		must.True(ok)
		must.Equal(expectedVal, actualVal)
	})

	t.Run("should returns same value (struct)", func(t *testing.T) {
		must := must.New(t)

		expectedKey := "thekey"
		expectedVal := struct {
			Name string
		}{
			"Alex Under",
		}

		Bind(Key(expectedKey), func() interface{} {
			return &expectedVal
		})

		actualVal, ok := Use(Key(expectedKey))
		must.True(ok)
		must.Equal(&expectedVal, actualVal)
	})

}
