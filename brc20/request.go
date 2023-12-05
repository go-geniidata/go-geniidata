package brc20

type RequestActivities struct {
	Limit   int32
	Offset  int32
	Address string
	Tick    string
	Order   Order
	Event   Event
}
