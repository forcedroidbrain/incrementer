package incrementer

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	i := New()
	assert.Equal(t, &Incrementer{value: 0, max: ^uint(0), mu: new(sync.RWMutex)}, i)
}

func TestIncrementer_GetNumber(t *testing.T) {
	i := New()
	assert.Equal(t, uint(0), i.GetNumber())
	i.IncrementNumber()
	assert.Equal(t, uint(1), i.GetNumber())
}

func TestIncrementer_IncrementNumber(t *testing.T) {
	i := New()
	i.IncrementNumber()
	assert.Equal(t, uint(1), i.value)
	i.SetMaxValue(2)
	i.IncrementNumber()
	assert.Equal(t, uint(0), i.value)
}

func TestIncrementer_SetMaxValue(t *testing.T) {
	i := New()
	i.SetMaxValue(3)
	assert.Equal(t, uint(3), i.max)
}

func TestIncrementer_SetMaxValue2(t *testing.T) {
	i := New()
	i.IncrementNumber()
	i.IncrementNumber()
	i.SetMaxValue(2)
	assert.Equal(t, uint(0), i.value)
}
