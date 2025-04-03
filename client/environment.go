package client

type Env string

const (
	Environment Env = "https://api.magichour.ai"
	MockServer  Env = "https://api.sideko.dev/v1/mock/magichour/magic-hour/0.12.1"
)

// String returns the environment as a string
func (e Env) String() string {
	return string(e)
}
