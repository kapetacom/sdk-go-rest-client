package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructToQueryParams(t *testing.T) {
	t.Run("should return error if input data is not a struct", func(t *testing.T) {
		_, err := StructToQueryParams("not a struct")
		assert.Error(t, err)
	})
	t.Run("should return query params", func(t *testing.T) {
		type input struct {
			Name string `query:"name"`
			Age  int    `query:"age"`
		}
		data := input{
			Name: "john",
			Age:  25,
		}
		got, _ := StructToQueryParams(data)
		expected := "age=25&name=john"
		assert.Equal(t, expected, got)
	})
	t.Run("should return error when query params input is a string", func(t *testing.T) {
		_, err := StructToQueryParams("test")
		assert.Error(t, err)
	})
}
