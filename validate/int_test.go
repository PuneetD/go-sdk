package validate

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestIntMin(t *testing.T) {
	assert := assert.New(t)

	var verr error
	val := 10
	verr = Int(&val).Min(1)()
	assert.Nil(verr)

	val = 10
	verr = Int(&val).Min(10)()
	assert.Nil(verr)

	verr = Int(nil).Min(10)()
	assert.NotNil(verr)
	assert.Nil(Value(verr))
	assert.Equal(ErrIntMin, Cause(verr))

	val = 1
	verr = Int(&val).Min(10)()
	assert.NotNil(verr)
	assert.NotNil(Value(verr))
	assert.Equal(ErrIntMin, Cause(verr))

	val = 10
	verr = Int(&val).Min(10)()
	assert.Nil(verr)
}

func TestIntMax(t *testing.T) {
	assert := assert.New(t)

	var verr error
	val := 1
	verr = Int(&val).Max(10)()
	assert.Nil(verr)

	verr = Int(nil).Max(10)()
	assert.Nil(verr)

	val = 10
	verr = Int(&val).Max(10)()
	assert.Nil(verr)

	val = 11
	verr = Int(&val).Max(10)()
	assert.NotNil(verr)
	assert.NotNil(Value(verr))
	assert.Equal(ErrIntMax, Cause(verr))
}

func TestIntBetween(t *testing.T) {
	assert := assert.New(t)

	var verr error
	val := 5
	verr = Int(&val).Between(1, 10)()
	assert.Nil(verr)

	verr = Int(nil).Between(5, 10)()
	assert.NotNil(verr)
	assert.Nil(Value(verr))
	assert.Equal(ErrIntMin, Cause(verr))

	val = 1
	verr = Int(&val).Between(5, 10)()
	assert.NotNil(verr)
	assert.Equal(1, Value(verr))
	assert.Equal(ErrIntMin, Cause(verr))

	val = 5
	verr = Int(&val).Between(5, 10)()
	assert.Nil(verr)

	val = 10
	verr = Int(&val).Between(5, 10)()
	assert.Nil(verr)

	val = 11
	verr = Int(&val).Between(5, 10)()
	assert.NotNil(verr)
	assert.Equal(11, Value(verr))
	assert.Equal(ErrIntMax, Cause(verr))
}

func TestIntPositive(t *testing.T) {
	assert := assert.New(t)

	var verr error
	val := 5
	verr = Int(&val).Positive()()
	assert.Nil(verr)

	verr = Int(nil).Positive()()
	assert.NotNil(verr)
	assert.Nil(Value(verr))
	assert.Equal(ErrIntPositive, Cause(verr))

	val = -5
	verr = Int(&val).Positive()()
	assert.NotNil(verr)
	assert.NotNil(Value(verr))
	assert.Equal(ErrIntPositive, Cause(verr))
}

func TestIntNegative(t *testing.T) {
	assert := assert.New(t)

	var verr error
	val := -5
	verr = Int(&val).Negative()()
	assert.Nil(verr)

	verr = Int(nil).Negative()()
	assert.NotNil(verr)
	assert.Nil(Value(verr))
	assert.Equal(ErrIntNegative, Cause(verr))

	val = 5
	verr = Int(&val).Negative()()
	assert.NotNil(verr)
	assert.NotNil(Value(verr))
	assert.Equal(ErrIntNegative, Cause(verr))
}
