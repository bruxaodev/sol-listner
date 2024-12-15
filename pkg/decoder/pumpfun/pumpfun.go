package pumpfun

type SubscribeData struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscriptions int `json:"subscriptions"`
		Result        struct {
			Context struct {
				Slot int `json:"slot"`
			} `json:"context"`
			Value struct {
				Slot  int   `json:"slot"`
				Block Block `json:"block"`
			} `json:"value"`
		} `json:"result"`
		// Err
	} `json:"params"`
}

type Block struct {
	PreviousBlockhash string        `json:"previousBlockhash"`
	BlockHash         string        `json:"blockhash"`
	ParentSlot        int           `json:"parentSlot"`
	BlockTime         int           `json:"blockTime"`
	BlockHeight       int           `json:"blockHeight"`
	Transactions      []Transaction `json:"transactions"`
}

type Transaction struct {
	Transaction struct {
		Signatures []string `json:"signatures"`
		Message    struct {
			AccountKeys         []AccountKeys `json:"accountKeys"`
			RecentBlockhash     string        `json:"recentBlockhash"`
			Instructions        []Instruction `json:"instructions"`
			AddressTableLookups []interface{} `json:"addressTableLookups,omitempty"`
		} `json:"message"`
	} `json:"transaction"`
	Meta    struct{}    `json:"meta"`
	Version interface{} `json:"version,omitempty"`
}

type AccountKeys struct {
	Pubkey   string `json:"pubkey"`
	Writable bool   `json:"writable"`
	Signer   bool   `json:"signer"`
	Source   string `json:"source"`
}

type Instruction struct {
	ProgramId   string      `json:"programId,omitempty"`
	Program     string      `json:"program,omitempty"`
	Accounts    []string    `json:"accounts,omitempty"`
	Data        string      `json:"data,omitempty"`
	StackHeight int         `json:"stackHeight,omitempty"`
	Parsed      interface{} `json:"parsed,omitempty"`
}

type ParsedInstruction struct {
	Info *Info   `json:"info,omitempty"`
	Type *string `json:"type,omitempty"`
}

type Info struct {
	Account        string   `json:"account,omitempty"`
	Mint           string   `json:"mint,omitempty"`
	Source         string   `json:"source,omitempty"`
	SystemProgram  string   `json:"systemProgram,omitempty"`
	TokenProgram   string   `json:"tokenProgram,omitempty"`
	Wallet         string   `json:"wallet,omitempty"`
	ExtensionTypes []string `json:"extensionTypes,omitempty"`
	Lamports       int      `json:"lamports,omitempty"`
	NewAccount     string   `json:"newAccount,omitempty"`
	Owner          string   `json:"owner,omitempty"`
	Space          int      `json:"space,omitempty"`
}

type Meta struct {
	Err struct {
		InstructionError []interface{} `json:"InstructionError"`
	} `json:"err"`
	Status struct {
		Err struct {
			InstructionError []interface{} `json:"InstructionError"`
		} `json:"Err"`
	} `json:"status"`
	Fee                  int               `json:"fee"`
	PreBalances          []int             `json:"preBalances"`
	PostBalances         []int             `json:"postBalances"`
	InnerInstructions    InnerInstructions `json:"innerInstructions"`
	LogMessages          []string          `json:"logMessages"`
	PreTokenBalances     []TokenBalances   `json:"preTokenBalances"`
	PostTokenBalances    []TokenBalances   `json:"postTokenBalances"`
	Rewards              []struct{}        `json:"rewards,omitempty"`
	ComputeUnitsConsumed int               `json:"computeUnitsConsumed"`
}

type InnerInstructions struct {
	Index        int           `json:"index"`
	Instructions []Instruction `json:"instructions"`
}

type TokenBalances struct {
	AccountIndex  int    `json:"accountIndex"`
	Mint          string `json:"mint"`
	UiTokenAmount struct {
		UiAmount       float64 `json:"uiAmount"`
		Decimals       int     `json:"decimals"`
		Amount         string  `json:"amount"`
		UiAmountString string  `json:"uiAmountString"`
	} `json:"uiTokenAmount"`
	Owner     string `json:"owner"`
	ProgramId string `json:"programId"`
}
