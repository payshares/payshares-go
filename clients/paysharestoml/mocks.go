package paysharestoml

import "github.com/stretchr/testify/mock"

// MockClient is a mockable paysharestoml client.
type MockClient struct {
	mock.Mock
}

// GetPaysharesToml is a mocking a method
func (m *MockClient) GetPaysharesToml(domain string) (*Response, error) {
	a := m.Called(domain)
	return a.Get(0).(*Response), a.Error(1)
}

// GetPaysharesTomlByAddress is a mocking a method
func (m *MockClient) GetPaysharesTomlByAddress(address string) (*Response, error) {
	a := m.Called(address)
	return a.Get(0).(*Response), a.Error(1)
}
