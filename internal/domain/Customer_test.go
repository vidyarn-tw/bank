package domain

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestShouldUpdateCustomerAddress(t *testing.T) {
	customer := Customer{
		Accounts: []*Account{{
			Address: Address{city: "Blr"},
		},
			{Address: Address{city: "Blr"}},
		},
		Address: Address{city: "Blr"},
	}

	newAddress := Address{city: "Hyd"}
	customer.updateAddress(newAddress)
	assert.Equal(t, customer.Address.city, newAddress.city)
	assert.Equal(t, customer.Accounts[0].Address.city, newAddress.city)
}
