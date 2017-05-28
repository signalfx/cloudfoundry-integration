package metrics

import (
    "log"

    "github.com/cloudfoundry-incubator/uaago"
)

type UAATokenFetcher struct {
    UaaUrl                string
    Username              string
    Password              string
    InsecureSSLSkipVerify bool
}

func (uaa *UAATokenFetcher) FetchAuthToken() string {
    uaaClient, err := uaago.NewClient(uaa.UaaUrl)
    if err != nil {
        log.Fatalf("Error creating uaa client: %s", err.Error())
    }

    var authToken string
    log.Printf("Getting UAA token...")
    authToken, err = uaaClient.GetAuthToken(uaa.Username, uaa.Password, uaa.InsecureSSLSkipVerify)
    if err != nil {
        log.Fatalf("Error getting oauth token: %s. Please check your username and password.", err.Error())
    }
    return authToken
}
