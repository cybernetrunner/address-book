package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	name    = "John"
	address = "Baker str 221"
	phone   = "8(800)555 35 35"

	stubModel = &AddressField{
		name,
		address,
		phone,
	}
)

func TestNewAddressField(t *testing.T) {
	assert.Equal(t, NewAddressField(name, address, phone), stubModel)
}
