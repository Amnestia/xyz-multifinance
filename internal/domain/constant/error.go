package constant

// LoginFailedError error on failed login
type LoginFailedError struct{}

func (LoginFailedError) Error() string {
	return "Unknown user credentials"
}

// OverlimitError error on failed login
type OverlimitError struct{}

func (OverlimitError) Error() string {
	return "Limit for this transaction is not enough"
}
