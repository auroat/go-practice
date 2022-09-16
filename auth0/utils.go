package auth0

import (
	"github.com/kelseyhightower/envconfig"
)

type QueensEnvironment struct {
	AuthZeroAudience string `required:"true"`
	AuthZeroIss      string `required:"true"`
	AuthZeroPemURI   string `required:"true"`
}

func ParseEnvironmentVariables() (*QueensEnvironment, error) {
	env := &QueensEnvironment{}
	if err := envconfig.Process("queens", env); err != nil {
		return env, err
	}

	return env, nil
}
