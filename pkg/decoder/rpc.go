package decoder

import (
	"log"
	"time"

	"github.com/goccy/go-json"

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
		now := time.Now()
		var subscribeData pumpfun.SubscribeData
		err := json.Unmarshal(data, &subscribeData)
		if err != nil {
			log.Println(err)
			log.Println(string(data))
		}

		log.Printf("decode block: %v - time: %v\n", subscribeData.Params.Result.Context.Slot, time.Since(now))

		go NewBlock(subscribeData)
		go NewToken(subscribeData)
		// log.Printf("block: %d totalTransactions: %d\n", subscribeData.Params.Result.Value.Block.BlockHeight, len(subscribeData.Params.Result.Value.Block.Transactions))
	case Block:
		// Add block decoder
		log.Println(string(data))
	}
}
