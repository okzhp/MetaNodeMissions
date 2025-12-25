package dto

type QueryAuctionRequest struct {
	SellerAddress string `form:"sellerAddress"`
	NftAddress    string `form:"nftAddress"`
	TokenId       string `form:"tokenId"`
	BeginTime     string `form:"beginTime"`
	EndTime       string `form:"endTime"`
}

type BidRequest struct {
	BuyerAddress string `form:"buyerAddress"`
	MinBid       string `form:"minBid"`
	MaxBid       string `form:"maxBid"`
}
