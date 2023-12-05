package brc20

type Activity struct {
	TxId    string `json:"tx_id"`
	TxIndex int    `json:"tx_index"`
	Event   string `json:"event"`

	BlockHeight int64  `json:"block_height"`
	BlockTime   int64  `json:"block_time"`
	BlockHash   string `json:"block_hash"`

	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	Tick        string `json:"tick"`
	Amount      string `json:"amount"`

	InscriptionNumber int64  `json:"inscription_number"`
	InscriptionId     string `json:"inscription_id"`
}

type DataActivities struct {
	Count  int        `json:"count"`
	Limit  string     `json:"limit"`
	Offset string     `json:"offset"`
	List   []Activity `json:"list"`
}

type ResponseActivities struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    DataActivities `json:"data"`
}
