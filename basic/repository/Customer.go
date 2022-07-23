package repository

type repoCustomer struct {
}

func newCustomer() *repoCustomer {
	return &repoCustomer{}
}

func (rc *repoCustomer) ApplyForWork(customerID, workID int) error {
	return nil
}
