package web_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/alpine-hodler/driver/web/transport"
	"github.com/alpine-hodler/driver/x/web"
	"github.com/joho/godotenv"
)

func TestExamples(t *testing.T) {
	// defer tools.Quiet()()

	godotenv.Load(".test.env")
	os.Setenv("CB_PRO_URL", "https://api-public.sandbox.exchange.coinbase.com") // safety check

	t.Run("Coinbase Pro Client", func(t *testing.T) { ExampleNewClient_cbp() })
	t.Run("Fetch Coinbase Pro Accounts", func(t *testing.T) { ExampleFetch_cbpAccounts() })
}

func ExampleNewClient_cbp() {
	cbpurl := os.Getenv("CB_PRO_URL")
	passphrase := os.Getenv("CB_PRO_ACCESS_PASSPHRASE")
	key := os.Getenv("CB_PRO_ACCESS_KEY")
	secret := os.Getenv("CB_PRO_SECRET")

	client, err := web.NewClient(context.TODO(), transport.NewAPIKey().
		SetKey(key).
		SetPassphrase(passphrase).
		SetSecret(secret).
		SetURL(cbpurl))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("A Coinbase Pro client: %v", client)
}

func ExampleFetch_cbpAccounts() {
	// Read credentials from environment variables.
	cbpurl := os.Getenv("CB_PRO_URL")
	passphrase := os.Getenv("CB_PRO_ACCESS_PASSPHRASE")
	key := os.Getenv("CB_PRO_ACCESS_KEY")
	secret := os.Getenv("CB_PRO_SECRET")

	// Get a new client using an API Key for authentication.
	client, err := web.NewClient(context.TODO(), transport.NewAPIKey().
		SetKey(key).
		SetPassphrase(passphrase).
		SetSecret(secret).
		SetURL(cbpurl))

	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}

	u, err := url.JoinPath(cbpurl, "accounts")
	parsedURL, _ := url.Parse(u)

	cfg := &web.FetchConfig{
		Client: client,
		Method: http.MethodGet,
		URL:    parsedURL,
	}

	_, err = web.Fetch(context.TODO(), cfg)
	if err != nil {
		log.Fatalf("error fetching accounts: %v", err)
	}
}