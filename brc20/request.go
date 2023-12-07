package brc20

type RequestTickList struct {
	Limit  int
	Offset int
}

type RequestTickInfo struct {
	Tick string
}

type RequestBalance struct {
	Limit  int
	Offset int

	// At least one address/tick is not empty
	Address string `json:"address"`
	Tick    string `json:"tick"`
}

type RequestActivities struct {
	Limit   int
	Offset  int
	Address string
	Tick    string
	Order   Order
	Event   Event
}

type RequestTransferableInscriptions struct {
	Limit   int
	Offset  int
	Address string
	Tick    string
}
