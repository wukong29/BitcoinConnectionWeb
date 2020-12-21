package entity

type RPCResult struct {
	Result interface{} `json:"result"`
	Error string `json:"error"`
	Id int64 `json:"id"`
}

type BlockCount struct {
	Result int `json:"result"`
	Error string `json:"error"`
	Id int64 `json:"id"`
}

type Besthash struct {
	Result string `json:"result"`
	Error string `json:"error"`
	Id int64 `json:"id"`
}

type BlockHash struct {
	Result string `json:"result"`
	Error string `json:"error"`
	Id int64 `json:"id"`
}
