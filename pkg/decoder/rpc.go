package decoder

import (
	"encoding/json"
	"log"

	"github.com/bruxaodev/sol-listner/pkg/decoder/pumpfun"
)

type Decoder int

const (
	Pumpfun Decoder = iota
	Block
)

func BlockDecoder(data []byte, decoder Decoder) {
	switch decoder {
	case Pumpfun:
		var subscribeData pumpfun.SubscribeData
		err := json.Unmarshal(data, &subscribeData)
		if err != nil {
			log.Println(err)
			log.Println(string(data))
		}
		go NewToken(subscribeData)
		// log.Printf("block: %d totalTransactions: %d\n", subscribeData.Params.Result.Value.Block.BlockHeight, len(subscribeData.Params.Result.Value.Block.Transactions))
	case Block:
		// Add block decoder
		log.Println(string(data))
	}
}
