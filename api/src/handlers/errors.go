package handlers

type MalformedBasicAuth struct{}

func (MalformedBasicAuth) Error() string {
	return "malformed HTTP basic authorization header."
}
