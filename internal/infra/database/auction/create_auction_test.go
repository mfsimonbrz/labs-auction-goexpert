package auction

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateAuction(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("AutoCloseAuction", func(t *mtest.T) {
		t.AddMockResponses(mtest.CreateSuccessResponse())

		repo := &AuctionRepository{
			Collection: t.Coll,
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		t.Setenv("AUCTION_INTERVAL", "500ms")

		entity := &auction_entity.Auction{
			Id:          uuid.New().String(),
			Status:      auction_entity.Active,
			Description: "Automóvel Fiat Mobi Vermelho",
			Category:    "Automóveis",
			ProductName: "Fiat Mobi 2024",
			Timestamp:   time.Now(),
		}
		err := repo.CreateAuction(ctx, entity)

		if assert.Nil(t, err, "Expected to create the auction") {
			if ev := t.GetSucceededEvent(); assert.NotNil(t, ev) && assert.Equal(t, "insert", ev.CommandName) {
				assert.Nil(t, t.GetSucceededEvent(), "No more events expected")

				t.AddMockResponses(mtest.CreateSuccessResponse())

				time.Sleep(time.Second)

				if ev = t.GetSucceededEvent(); assert.NotNil(t, ev, "Update event expected for complete auction") {
					assert.Equal(t, "update", ev.CommandName, "Update event expected for complete auction")
				}
			}
		}
	})
}
