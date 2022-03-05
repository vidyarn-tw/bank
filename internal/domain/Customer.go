package domain

type Customer struct {
	Accounts []*Account
	Address  Address
}

func (c *Customer) updateAddress(address Address) {
	c.Address = address
	for _, account := range c.Accounts {
		account.Address = address
	}
}
