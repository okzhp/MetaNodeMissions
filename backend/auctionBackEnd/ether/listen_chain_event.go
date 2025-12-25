package ether

import (
	"auctionBackEnd/abi"
	"auctionBackEnd/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

const SepoliaWssRpc = "wss://eth-sepolia.g.alchemy.com/v2/QyTJFlHI_6sViFRorM5P0"

// 拍卖合约地址 0x979166bFF7A7013F059730719C3A7F3A305130ec
const AuctionAddress = "0x9e341fE62F798B4d1E7AC36A68a912116627E197"

// 要监听合约的事件 topic hash 签名哈希
var (
	SigCreateAuction = crypto.Keccak256Hash([]byte("CreateAuction(address,address,uint256,uint256)"))

	SigBidUSD = crypto.Keccak256Hash([]byte("BidUSD(address,uint256)"))

	SigAuctionEnd = crypto.Keccak256Hash([]byte("AuctionEnd(address,uint256)"))
)

func ListeningChainEvent(db *gorm.DB) {
	client, err := ethclient.Dial(SepoliaWssRpc)
	if err != nil {
		log.Fatalln(err)
	}

	//blockHash := common.HexToHash("0x158b2a83e8a60509a24f09d334f2c1134f028f2e08243b878e5863a162c034c6")
	query := ethereum.FilterQuery{
		//BlockHash: &blockHash,
		Addresses: []common.Address{common.HexToAddress(AuctionAddress)},
	}

	logsChan := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logsChan)
	if err != nil {
		log.Fatalln("监听logs失败:", err)
	}

	//logs, err := client.FilterLogs(context.Background(), query)
	//if err != nil {
	//	log.Fatalln("查询logs失败:", err)
	//}

	filterer, err := abi.NewAuctionContractFilterer(common.HexToAddress(AuctionAddress), client)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case err = <-sub.Err():
			fmt.Println("监听订阅事件err:", err)
		case vLog := <-logsChan:
			switch vLog.Topics[0] {
			case SigCreateAuction:
				processEventCreateAuction(db, filterer, vLog)
			case SigBidUSD:
				processEventBidUSD(db, filterer, vLog)
			case SigAuctionEnd:
				processEventAuctionEnd(db, filterer, vLog)
			default:
				fmt.Println("接收到未定义事件:", vLog.Topics[0].Hex())
			}
		}
	}
}

// 处理创建合约事件
func processEventCreateAuction(db *gorm.DB, filterer *abi.AuctionContractFilterer, vLog types.Log) {
	fmt.Println("监听到CreateAuction!")
	event, err := filterer.ParseCreateAuction(vLog)
	if err != nil {
		log.Fatalln(err)
	}

	//保存入库
	auction := model.Auction{
		SellerAddress:   event.Seller.Hex(),
		NftAddress:      event.Nft.Hex(),
		TokenId:         event.TokenId.String(),
		DurationMinutes: event.DurationMinutes.Uint64(),
		EndTime:         time.Unix(int64(vLog.BlockTimestamp), 0).Add(time.Duration(event.DurationMinutes.Uint64()) * time.Minute),
	}
	err = db.Create(&auction).Error
	if err != nil {
		log.Fatalln("CreateAuction处理失败:", err)
	}

	fmt.Println("CreateAuction处理完毕!")
}

// 处理出价合约事件
func processEventBidUSD(db *gorm.DB, filterer *abi.AuctionContractFilterer, vLog types.Log) {
	fmt.Println("监听到BidUSD!")
	event, err := filterer.ParseBidUSD(vLog)
	if err != nil {
		log.Fatalln(err)
	}

	bid := model.Bid{
		Buyer: event.Buyer.Hex(),
		Bid:   event.Bid.String(),
	}
	err = db.Create(&bid).Error
	if err != nil {
		log.Fatalln("BidUSD处理失败:", err)
	}
	fmt.Println("BidUSD处理完毕!")
}

// 处理敲定合约事件
func processEventAuctionEnd(db *gorm.DB, filterer *abi.AuctionContractFilterer, vLog types.Log) {
	fmt.Println("监听到AuctionEnd!")
	event, err := filterer.ParseAuctionEnd(vLog)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Model(&model.Bid{}).
		Where("buyer = ? and bid = ?", event.Buyer.Hex(), event.Bid.String()).
		Updates(map[string]interface{}{
			"deal":   1,
			"dealAt": time.Now(),
		}).Error
	if err != nil {
		log.Fatalln("AuctionEnd处理失败:", err)
	}
	fmt.Println("AuctionEnd处理完毕!")
}
