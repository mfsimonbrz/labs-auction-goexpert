package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"fullcycle-auction_go/internal/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}
type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	auctionInterval := utils.GetAuctionInterval()
	go func() {
		bgCtx := context.Background()
		for {
			auctionEndTime := auctionEntity.Timestamp.Add(auctionInterval).Truncate(time.Second)
			if time.Now().Truncate(time.Second).Equal(auctionEndTime) {
				logger.Info(fmt.Sprintf("auction %s ended", auctionEntity.Id))
				filter := bson.M{"_id": auctionEntity.Id}
				update := bson.M{
					"$set": bson.M{
						"status": auction_entity.Completed,
					},
				}
				_, err := ar.Collection.UpdateOne(bgCtx, filter, update)
				if err != nil {
					logger.Error("Error trying to update auction", err)
				}
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()
	return nil
}
