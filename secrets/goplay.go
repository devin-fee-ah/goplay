package secrets

import (
	"encoding/json"
)

// GoPlay secret
type GoPlay struct {
	DatabaseURL string `json:"DATABASE_URL"`
}

// GetGoPlay from SecretsManager
func (s *Secrets) GetGoPlay() (secret *GoPlay, err error) {
	secret = s.goPlay
	if s.goPlay == nil {
		var value string
		secret = &GoPlay{}

		value, err = s.awsToolbelt.GetSecret(s.env.GoPlaySecretName)
		if err != nil {
			return secret, err
		}

		err = json.Unmarshal([]byte(value), secret)
		if err != nil {
			return secret, err
		}

		s.goPlay = secret
	}
	return
}
