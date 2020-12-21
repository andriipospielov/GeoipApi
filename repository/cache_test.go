package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net"
	"testing"
)

func TestNewCachedResults(t *testing.T) {
	sut := NewCachedResults()
	assert.IsType(t, &CachedResults{}, sut)
	assert.IsType(t, make(map[string]map[string]interface{}), sut.v)
}

type MutexMock struct {
	mock.Mock
}

func (m *MutexMock) Lock() {
	m.Called()
	return
}

func (m *MutexMock) Unlock() {
	m.Called()
	return
}

func TestCachedResults_Set(t *testing.T) {
	var mutex = new(MutexMock)
	mutex.On("Lock")
	mutex.On("Unlock")

	sut := CachedResults{
		mu: mutex,
		v:  make(map[string]map[string]interface{}),
	}

	sut.Set("123", net.IP{}, 123)
	mutex.AssertExpectations(t)

}
