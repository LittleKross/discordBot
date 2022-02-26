package config
import()

type config struct {
	token string
}

func New(t string) *config {
	c := &config {
		token: t,
	}
	return c
}