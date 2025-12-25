package controller

import (
  "auctionBackEnd/dto"
  "auctionBackEnd/model"
  "auctionBackEnd/utils"
  "log"
  "time"

  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
)

type AuctionController struct {
  Db *gorm.DB
}

func (ctl *AuctionController) QueryAuction(c *gin.Context) {
  var req dto.QueryAuctionRequest

  err := c.ShouldBind(&req)
  if err != nil {
    log.Fatalln("参数绑定有误")
  }

  var auctions []model.Auction
  tx := ctl.Db.Model(&model.Auction{})

  if req.SellerAddress != "" {
    tx = tx.Where("seller_address = ?", req.SellerAddress)
  }

  if req.NftAddress != "" {
    tx = tx.Where("nft_address = ?", req.NftAddress)
  }

  if req.TokenId != "" {
    tx = tx.Where("token_id = ?", req.TokenId)
  }

  layout := "2006-01-02 15:04:05"
  if req.BeginTime != "" && req.EndTime != "" {
    beginTime, err := time.ParseInLocation(layout, req.BeginTime, time.Local)
    if err != nil {
      log.Fatalln("开始时间格式有误")
    }
    endTime, err := time.ParseInLocation(layout, req.EndTime, time.Local)
    if err != nil {
      log.Fatalln("结束时间格式有误")
    }
    tx = tx.Where("end_time between ? and ?", beginTime, endTime)
  }

  err = tx.Find(&auctions).Error
  if err != nil {
    log.Fatalln("query auctions fail")
  }

  utils.Success(c, auctions)
}

func (ctl *AuctionController) QueryBid(c *gin.Context) {
  var req dto.BidRequest
  err := c.ShouldBind(&req)
  if err != nil {
    log.Fatalln("参数绑定有误")
  }

  tx := ctl.Db
  var bids []model.Bid
  if req.BuyerAddress != "" {
    tx = tx.Where("buyer = ?", req.BuyerAddress)
  }

  if req.MinBid != "" {
    tx = tx.Where("bid >= ?", req.MinBid)
  }

  if req.MaxBid != "" {
    tx = tx.Where("bid <= ?", req.MaxBid)
  }

  err = tx.Find(&bids).Error
  if err != nil {
    log.Fatalln("query bids fail")
  }

  utils.Success(c, bids)

}
