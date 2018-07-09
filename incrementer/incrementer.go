package incrementer

import "sync"

// Incrementer store value
type Incrementer struct {
	value uint
	max   uint
	mu    *sync.RWMutex
}

// New create instance Incrementer
func New() *Incrementer {
	i := new(Incrementer)
	i.max = ^uint(0)
	i.mu = new(sync.RWMutex)
	return i
}

// GetNumber return Incrementer value
func (i *Incrementer) GetNumber() uint {
	defer i.mu.RUnlock()
	i.mu.RLock()
	return i.value
}

// IncrementNumber adds one if provided no more than the maximum value
func (i *Incrementer) IncrementNumber() {
	defer i.mu.Unlock()
	i.mu.Lock()
	i.value++
	if i.value >= i.max {
		i.value = 0
	}
}

// SetMaxValue sets the maximum value
// and resets the value to zero if it is greater than the maximum
func (i *Incrementer) SetMaxValue(max uint) {
	defer i.mu.Unlock()
	i.mu.Lock()
	if i.value >= max {
		i.value = 0
	}
	i.max = max
}
