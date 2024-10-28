package entity

// Account entity
type Account struct {
	ID       int
	ClientID int
	Balance  int
	CardPAN  string
	// other...
}

// TODO: GenFakeAccount, for testing service
