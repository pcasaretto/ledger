package suite

import (
	"context"
	"net/http"

	. "github.com/numary/ledger/it/internal"
	. "github.com/numary/ledger/it/internal/otlpinterceptor"
	ledgerclient "github.com/numary/numary-sdk-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Scenario("Server API", func() {
	When("reading server configuration", func() {
		var (
			response     ledgerclient.ConfigInfoResponse
			httpResponse *http.Response
			err          error
		)
		BeforeEach(func() {
			response, httpResponse, err = Client().ServerApi.
				GetInfo(context.Background()).
				Execute()
			Expect(err).To(BeNil())
		})
		It("should respond with the correct configuration", func() {
			Expect(response.Data).To(Equal(ledgerclient.ConfigInfo{
				Config: ledgerclient.Config{
					Storage: ledgerclient.LedgerStorage{
						Driver:  "postgres",
						Ledgers: []string{},
					},
				},
				Server:  "numary-ledger",
				Version: "develop",
			}))
		})
		It("should register a trace", func() {
			Expect(httpResponse).To(HaveTrace(NewTrace("/_info").
				WithAttributes(HTTPStandardAttributes(http.MethodGet, "/_info", "/_info"))))
		})
	})
})
