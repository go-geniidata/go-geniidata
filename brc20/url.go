package brc20

import (
	"fmt"
	"net/url"
)

const DefaultServer = "https://api.geniidata.com"

type Order string

const (
	Asc  Order = "asc"
	Desc Order = "desc"
)

type Event string

const (
	InscribeMint     Event = "inscribe-mint"
	InscribeTransfer Event = "inscribe-transfer"
	Transfer         Event = "transfer"
)

func UrlTickList(req RequestTickList) string {
	v := url.Values{}
	if req.Limit > 0 {
		v.Set("limit", fmt.Sprintf("%d", req.Limit))
	}
	if req.Offset >= 0 {
		v.Set("offset", fmt.Sprintf("%d", req.Offset))
	}
	return buildUrl("/api/1/brc20/ticks", v)
}

func UrlTickInfo(req RequestTickInfo) string {
	return fmt.Sprintf("/api/1/brc20/tickinfo/%s", req.Tick)
}

func UrlBalance(req RequestBalance) string {
	v := url.Values{}
	if req.Limit > 0 {
		v.Set("limit", fmt.Sprintf("%d", req.Limit))
	}
	if req.Offset >= 0 {
		v.Set("offset", fmt.Sprintf("%d", req.Offset))
	}
	if req.Address != "" {
		v.Set("address", req.Address)
	}
	if req.Tick != "" {
		v.Set("tick", req.Tick)
	}
	return buildUrl("/api/1/brc20/balance", v)
}

func UrlActivities(req RequestActivities) string {
	// https://api.geniidata.com/api/1/brc20/activities?limit=20&offset=0&tick=ordi&order=asc
	v := url.Values{}
	if req.Limit > 0 {
		v.Set("limit", fmt.Sprintf("%d", req.Limit))
	}
	if req.Offset >= 0 {
		v.Set("offset", fmt.Sprintf("%d", req.Offset))
	}
	if req.Address != "" {
		v.Set("address", req.Address)
	}
	if req.Tick != "" {
		v.Set("tick", req.Tick)
	}
	if req.Order != "" {
		v.Set("order", string(req.Order))
	}
	if req.Event != "" {
		v.Set("event", string(req.Event))
	}
	return buildUrl("/api/1/brc20/activities", v)
}

func UrlTransferableInscriptions(req RequestTransferableInscriptions) string {
	v := url.Values{}
	if req.Limit > 0 {
		v.Set("limit", fmt.Sprintf("%d", req.Limit))
	}
	if req.Offset >= 0 {
		v.Set("offset", fmt.Sprintf("%d", req.Offset))
	}

	return buildUrl(fmt.Sprintf("/api/1/brc20/address/%s/tick/%s/transferableInscriptions", req.Address, req.Tick), v)
}

func buildUrl(path string, v url.Values) string {
	u, _ := url.Parse("https://api.geniidata.com")
	u = u.JoinPath(path)
	u.RawQuery = v.Encode()
	return u.String()
}
