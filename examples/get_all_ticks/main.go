package main

import (
	"context"
	"fmt"
	"github.com/go-geniidata/go-geniidata/brc20"
	"log"
	"os"
	"time"
)

func main() {
	apiKey := os.Getenv("APIKEY")
	b := brc20.New(apiKey)
	limit := 100
	offset := 0

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	fmt.Println("tick,deployer,deployTime,decimals,max,limit,holders,transactions,minted,mintProgress")
	count, err := doWork(ctx, b, limit, offset)
	if err != nil {
		log.Fatalf("http request error: %s", err)
	}
	offset += limit
	for offset < count {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
			count, err = doWork(ctx, b, limit, offset)
			if err != nil {
				log.Fatalf("http request error: %s", err)
			}
			offset += limit
		}
	}
}

func doWork(ctx context.Context, b *brc20.Brc20, limit, offset int) (int, error) {
	req := brc20.RequestTickList{
		Limit:  limit,
		Offset: offset,
	}
	resp, err := b.TickList(ctx, req)
	if err != nil {
		return 0, err
	}

	for _, t := range resp.Data.List {
		fmt.Printf(
			"\"%s\",%s,%s,%d,%s,%s,%d,%d,%s,%s\n",
			t.Tick, t.Deployer, time.Unix(t.DeployTime, 0),
			t.Decimals, t.Max, t.Limit, t.Holders, t.Transactions,
			t.Minted, t.MintProgress,
		)
	}
	return resp.Data.Count, nil
}
