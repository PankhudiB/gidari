package repository

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/alpine-hodler/driver/data/proto"
	db "github.com/alpine-hodler/driver/data/storage"
	"github.com/alpine-hodler/driver/tools"
	"github.com/alpine-hodler/driver/web/transport"
	"github.com/alpine-hodler/driver/x/web"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestWebIntegration(t *testing.T) {
	err := godotenv.Load(".test.env")
	require.NoError(t, err)

	os.Setenv("CB_PRO_URL", "https://api-public.sandbox.exchange.coinbase.com") // safety check

	ctx := context.Background()
	dns, _ := tools.MongoURI("mongo-coinbasepro", "", "", "27017", "coinbasepro")

	stg, err := db.New(ctx, dns)
	require.NoError(t, err)

	repo := NewCoinbasePro(ctx, stg)

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
		log.Fatalf("error creating client: %v", err)
	}

	u, err := url.JoinPath(cbpurl, "accounts")
	parsedURL, _ := url.Parse(u)

	cfg := &web.FetchConfig{
		Client: client,
		Method: http.MethodGet,
		URL:    parsedURL,
	}

	bytes, err := web.Fetch(context.TODO(), cfg)
	if err != nil {
		log.Fatalf("error fetching accounts: %v", err)
	}

	rsp := new(proto.CreateResponse)
	err = repo.UpsertAccountsJSON(ctx, bytes, rsp)
	require.NoError(t, err)
}