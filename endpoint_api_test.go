package coinbase

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Initialize the client without mock mode enabled on rpc.
// All calls hit the coinbase API and tests focus on checking
// format of the response and validity of sent requests
func initClient() Client {
	return ApiKeyClient(os.Getenv("COINBASE_KEY"), os.Getenv("COINBASE_SECRET"))
}

// About Endpoint Tests:
//All Endpoint Tests actually call the Coinbase API and check the return values
// with type assertions. This was done because of the varying specific values
// returned depending on the API Key and Secret pair used when running the tests.
// Endpoint tests do not include tests that could be run an arbitrary amount of times
// i.e buy, sell, etc...

func TestGetBalanceEndpoint(t *testing.T) {
	c := initClient()
	amount, err := c.GetBalance()
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, 1.1, amount)
}

func TestGetReceiveAddressEndpoint(t *testing.T) {
	t.Skip("Skipping GetReceiveAddressEndpoint")
	fmt.Println("Skipped to avoid generating many addresses")
	c := initClient()
	params := &ReceiveAddressParams{
		Address: &AddressParams{
			Callback_url: "http://www.wealthlift.com",
			Label:        "My Test Address",
		},
	}
	address, err := c.GenerateReceiveAddress(params)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", address)
}

func TestGetAllAddressesEndpoint(t *testing.T) {
	c := initClient()
	params := &AddressesParams{
		Page:  1,
		Limit: 5,
	}
	addresses, err := c.GetAllAddresses(params)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", addresses.Addresses[0].Created_at)
	assert.IsType(t, "string", addresses.Addresses[1].Address)
}

func TestGetButtonEndpoint(t *testing.T) {
	c := initClient()
	params := &ButtonParams{
		Button: &button{
			Name:               "test",
			Type:               "buy_now",
			Subscription:       false,
			Price_string:       "1.23",
			Price_currency_iso: "USD",
			Custom:             "Order123",
			Callback_url:       "http://www.example.com/my_custom_button_callback",
			Description:        "Sample Description",
			Style:              "custom_large",
			Include_email:      true,
		},
	}
	data, err := c.GetButton(params)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", data.Embed_html)
	assert.IsType(t, "string", data.Type)
}

func TestGetExchangeRatesEndpoint(t *testing.T) {
	c := initClient()
	data, err := c.GetExchangeRates()
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", data["btc_to_usd"])
}

func TestGetExchangeRateEndpoint(t *testing.T) {
	c := initClient()
	data, err := c.GetExchangeRate("btc", "usd")
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, 0.0, data)
}

func TestGetTransactionsEndpoint(t *testing.T) {
	c := initClient()
	data, err := c.GetTransactions(1)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, 1, data.Total_count)
	assert.IsType(t, "string", data.Transactions[0].Hsh)
}

func TestGetBuyPriceEndpoint(t *testing.T) {
	c := initClient()
	data, err := c.GetBuyPrice(1)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", data.Subtotal.Currency)
	assert.IsType(t, "string", data.Total.Amount)
}

func TestGetSellPriceEndpoint(t *testing.T) {
	c := initClient()
	data, err := c.GetSellPrice(1)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", data.Subtotal.Currency)
	assert.IsType(t, "string", data.Total.Amount)
}

func TestGetTransactionEndpoint(t *testing.T) {
	c := initClient()
	_, err := c.GetTransaction("5446968682a19ab940000004")
	if err != nil {
		assert.IsType(t, "string", err.Error())
	}
}
