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
	for {
		req := brc20.RequestTickList{
			Limit:  limit,
			Offset: offset,
		}
		resp, err := b.TickList(ctx, req)
		if err != nil {
			log.Fatalf("http request error: %s", err)
		}
		//fmt.Printf("count: %d\n", resp.Data.Count)

		for _, t := range resp.Data.List {
			fmt.Printf(
				"%s,%s,%s,%d,%s,%s,%d,%d,%s,%s\n",
				t.Tick, t.Deployer, time.Unix(t.DeployTime, 0),
				t.Decimals, t.Max, t.Limit, t.Holders, t.Transactions,
				t.Minted, t.MintProgress,
			)
		}

		offset += limit
		if offset >= resp.Data.Count {
			break
		}
	}
}
