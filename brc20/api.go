package brc20

import (
	"context"
	"github.com/go-geniidata/go-geniidata/utils"
)

type Brc20 struct {
	apiKey string
}

func New(apiKey string) *Brc20 {
	return &Brc20{
		apiKey: apiKey,
	}
}

func (b *Brc20) TickList(ctx context.Context, req RequestTickList) (ResponseTickList, error) {
	var resp ResponseTickList
	url := UrlTickList(req)
	err := utils.GetWithKey(ctx, url, b.apiKey, &resp)
	return resp, err
}

func (b *Brc20) TickInfo(ctx context.Context, req RequestTickInfo) (ResponseTickInfo, error) {
	var resp ResponseTickInfo
	url := UrlTickInfo(req)
	err := utils.GetWithKey(ctx, url, b.apiKey, &resp)
	return resp, err
}

func (b *Brc20) Activities(ctx context.Context, req RequestActivities) (ResponseActivities, error) {
	var resp ResponseActivities
	url := UrlActivities(req)
	err := utils.GetWithKey(ctx, url, b.apiKey, &resp)
	return resp, err
}
