package brc20

import "fmt"

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

func UrlActivities(req RequestActivities) string {
	// https://api.geniidata.com/api/1/brc20/activities?limit=20&offset=0&tick=ordi&order=asc
	return fmt.Sprintf(
		"%s/api/1/brc20/activities?limit=%d&offset=%d&tick=%s&order=%s&event=%s",
		DefaultServer, req.Limit, req.Offset, req.Tick, req.Order, req.Event,
	)
}
