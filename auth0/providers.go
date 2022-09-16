package auth0

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/form3tech-oss/jwt-go"
	log "github.com/sirupsen/logrus"
)

type Auth0ProviderInterface interface {
	Authenticate(token string) error
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type auth0Provider struct {
	audience string
	iss      string
	pemUri   string
}

func NewAuth0Provider() Auth0ProviderInterface {
	env, _ := ParseEnvironmentVariables()
	return &auth0Provider{
		audience: env.AuthZeroAudience,
		iss:      env.AuthZeroIss,
		pemUri:   env.AuthZeroPemURI,
	}
}

func (ap *auth0Provider) Authenticate(token string) error {

	log.Info("start authentication with token")

	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(jwtToken *jwt.Token) (interface{}, error) {
		cert, err := ap.getPemCert(jwtToken)
		if err != nil {
			return errors.New("access denid"), err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	})

	if err != nil {
		log.Error(err.Error())
		return errors.New("access denid")
	}

	aud := parsedToken.Claims.(jwt.MapClaims).VerifyAudience(ap.audience, false)

	if !aud {
		return errors.New("access denid")
	}

	checkIss := parsedToken.Claims.(jwt.MapClaims).VerifyIssuer(ap.iss, false)
	if !checkIss {
		return errors.New("access denied")
	}

	return nil
}

func (ap *auth0Provider) getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	env, _ := ParseEnvironmentVariables()
	resp, err := http.Get(env.AuthZeroPemURI)

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		log.Errorf("error in parse public key: %s", err.Error())
		return cert, err
	}

	return cert, nil
}
