package domain

type Account struct {
	Address Address
}

func (a *Account) updateAddress(address Address) {
	a.Address = address
}
