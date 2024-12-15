package connection

type Subscribe struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type SubiscribeParams struct {
	MentionsAccountOrProgram       string `json:"mentionsAccountOrProgram,omitempty"`
	ProgramSubscribe               string `json:"programSubscribe,omitempty"`
	AccountSubscribe               string `json:"accountSubscribe,omitempty"`
	Commitment                     string `json:"commitment,omitempty"`
	Encoding                       string `json:"encoding,omitempty"`
	TransactionDetails             string `json:"transactionDetails,omitempty"`
	ShowRewards                    *bool  `json:"showRewards,omitempty"`
	MaxSupportedTransactionVersion *int   `json:"maxSupportedTransactionVersion,omitempty"`
}

var SUBSCRIBE_PUMPFUN = Subscribe{
	Jsonrpc: "2.0",
	Id:      1,
	Method:  "blockSubscribe",
	Params: []interface{}{
		SubiscribeParams{
			MentionsAccountOrProgram: "6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P",
		},
		SubiscribeParams{
			Commitment:                     "confirmed",
			Encoding:                       "jsonParsed",
			TransactionDetails:             "full",
			ShowRewards:                    &[]bool{false}[0], // false
			MaxSupportedTransactionVersion: &[]int{0}[0],      // 0
		},
	},
}

var SUBSCRIBE_ALL = Subscribe{
	Jsonrpc: "2.0",
	Id:      1,
	Method:  "blockSubscribe",
	Params: []interface{}{
		"all",
		SubiscribeParams{
			Commitment:                     "confirmed",
			Encoding:                       "jsonParsed",
			TransactionDetails:             "full",
			ShowRewards:                    &[]bool{false}[0], // false
			MaxSupportedTransactionVersion: &[]int{0}[0],      // 0
		},
	},
}
