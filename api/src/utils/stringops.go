package utils

import "fmt"

func Quote(s string) string {
	return fmt.Sprintf("\"%v\"", s)
}
