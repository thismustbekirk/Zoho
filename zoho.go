package zoho

import (
	"net"
	"net/http"
	"time"
)

// New initializes a Zoho structure
func New() *Zoho {
	z := Zoho{
		client: &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 5 * time.Second,
			},
		},
	}
	return &z
}

// Zoho is for accessing all APIs. It is used by subpackages to simplify passing authentication
// values between API subpackages.
type Zoho struct {
	user      string
	password  string
	scope     string
	authtoken string
	orgID     string

	client *http.Client
}
