package app

type AppNotFoundError struct{}

func (AppNotFoundError) Error() string {
	return "app not found."
}
