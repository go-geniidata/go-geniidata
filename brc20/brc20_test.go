package brc20_test

import (
	"context"
	"github.com/go-geniidata/go-geniidata/brc20"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

var b *brc20.Brc20

func TestMain(m *testing.M) {
	apiKey := os.Getenv("APIKEY")
	b = brc20.New(apiKey)

	os.Exit(m.Run())
}

func TestBrc20_Balance(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()
	req := brc20.RequestBalance{
		Limit:   10,
		Offset:  0,
		Tick:    "seed",
		Address: "bc1pxgpyskz606rlpfjsg3jsluuqedxflyrnfdscftcc7slesqlgqxfsa2mvcf",
	}

	resp, err := b.Balance(ctx, req)
	require.NoError(t, err)
	for i, e := range resp.Data.List {
		t.Logf("%d: %v", i, e)
	}
}

func TestBrc20_Activities(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()
	req := brc20.RequestActivities{
		Limit:  10,
		Offset: 0,
		Tick:   "ordi",
		Order:  brc20.Asc,
		Event:  brc20.InscribeTransfer,
	}

	resp, err := b.Activities(ctx, req)
	require.NoError(t, err)
	for i, e := range resp.Data.List {
		t.Logf("%d: %v", i, e)
	}
}

func TestBrc20_TransferableInscriptions(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()
	req := brc20.RequestTransferableInscriptions{
		Limit:   10,
		Offset:  0,
		Tick:    "ordi",
		Address: "bc1pv00lg0mj34g2uwgznd4slp7ezsace5kvjlyrfe9hpxzahflgnemqny7r8x",
	}

	resp, err := b.TransferableInscriptions(ctx, req)
	require.NoError(t, err)
	for i, e := range resp.Data.List {
		t.Logf("%d: %v", i, e)
	}
}
