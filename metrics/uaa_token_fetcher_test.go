package metrics_test

import (
    "github.com/signalfx/cloudfoundry-bridge/metrics"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/signalfx/cloudfoundry-bridge/testhelpers"
)

var _ = Describe("UaaTokenFetcher", func() {
    var tokenFetcher *metrics.UAATokenFetcher
    var fakeUAA *testhelpers.FakeUAA
    var fakeToken string

    BeforeEach(func() {
        fakeUAA = testhelpers.NewFakeUAA("bearer", "123456789")
        fakeToken = fakeUAA.AuthToken()
        fakeUAA.Start()

        tokenFetcher = &metrics.UAATokenFetcher{
            UaaUrl: fakeUAA.URL(),
        }
    })

    It("fetches a token from the UAA", func() {
        receivedAuthToken := tokenFetcher.FetchAuthToken()
        Expect(fakeUAA.Requested()).To(BeTrue())
        Expect(receivedAuthToken).To(Equal(fakeToken))
    })
})
