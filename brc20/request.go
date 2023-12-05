package brc20

type RequestTickList struct {
	Limit  int
	Offset int
}

type RequestTickInfo struct {
	Tick string
}

type RequestActivities struct {
	Limit   int
	Offset  int
	Address string
	Tick    string
	Order   Order
	Event   Event
}
