package auth

type NotAuthorized struct{}

func (NotAuthorized) Error() string {
	return "basic authorization failed"
}
