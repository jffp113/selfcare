package operation

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
)

func getClient() *http.Client {
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Jar: jar,
	}

	return client
}
