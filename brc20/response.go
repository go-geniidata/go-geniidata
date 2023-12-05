package brc20

// Tick 出现在tick list和tick info接口中，结构完全一致。
type Tick struct {
	Tick              string `json:"tick"`
	InscriptionId     string `json:"inscription_id"`
	InscriptionNumber int64  `json:"inscription_number"`
	Max               string `json:"max"`
	Limit             string `json:"limit"`
	Decimals          int    `json:"decimals"`
	Deployer          string `json:"deployer"`
	DeployTime        int64  `json:"deploy_time"`

	Minted       string `json:"minted"`
	MintProgress string `json:"mint_progress"`
	Transactions int64  `json:"transactions"`
	Holders      int64  `json:"holders"`
}

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

type DataTickList struct {
	Count  int    `json:"count"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	List   []Tick `json:"list"`
}

type ResponseTickList struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    DataTickList `json:"data"`
}

type ResponseTickInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Tick   `json:"data"`
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
