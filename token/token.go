package token

// Stolen from https://github.com/cloudfoundry-incubator/datadog-firehose-nozzle/blob/master/uaatokenfetcher/uaa_token_fetcher.go

import (
  "fmt"
  "os"
	"github.com/cloudfoundry-incubator/uaago"
)

type UAATokenFetcher struct {
	UaaUrl                string
	Username              string
	Password              string
	InsecureSSLSkipVerify bool
}

func (uaa *UAATokenFetcher) FetchAuthToken() (string, error) {

	uaaClient, err := uaago.NewClient(uaa.UaaUrl)
	if err != nil {
		return "", err
	}

	var authToken string
	authToken, err = uaaClient.GetAuthToken(uaa.Username, uaa.Password, uaa.InsecureSSLSkipVerify)
	if err != nil {
		return "", err
	}
	return authToken, nil
}

func (uaa *UAATokenFetcher) RefreshAuthToken() (string, error) {
  fmt.Fprintf(os.Stdout, "Refreshing authorization token\n")
	return uaa.FetchAuthToken()
}
