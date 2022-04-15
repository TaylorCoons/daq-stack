package auth

type TokenNotAuthorized struct{}

func (TokenNotAuthorized) Error() string {
	return "not authorized: token not valid"
}
