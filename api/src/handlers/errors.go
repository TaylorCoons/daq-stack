package handlers

type MalformedBasicAuth struct{}

func (MalformedBasicAuth) Error() string {
	return "malformed HTTP basic authorization header."
}

type NotAuthorized struct{}

func (NotAuthorized) Error() string {
	return "basic authorization failed."
}

type NoApiKeyProvided struct{}

func (NoApiKeyProvided) Error() string {
	return "no api key provided."
}

type TokenNotAuthorized struct{}

func (TokenNotAuthorized) Error() string {
	return "not authorized: token not valid."
}
