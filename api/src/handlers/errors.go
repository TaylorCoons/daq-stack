package handlers

type MalformedBasicAuth struct{}

func (MalformedBasicAuth) Error() string {
	return "malformed HTTP basic authorization header."
}

type NotAuthorized struct{}

func (NotAuthorized) Error() string {
	return "basic authorization failed"
}
