package user

import (
	"github.com/Shravan-Chaudhary/revamp-server/internal/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"
type UserRepository interface {
	GetUserById(string) (*types.User, error)
}

type MongoUserRepository struct {
	client *mongo.Client
	coll *mongo.Collection
}

func NewMongoUserRepository(client *mongo.Client, cfg *types.Config) *MongoUserRepository {
	return &MongoUserRepository{
		client: client,
		coll: client.Database(cfg.DATABASE_NAME).Collection(userColl),
	}
}

func (m *MongoUserRepository) GetUserById(id string) (*types.User, error) {
	return &types.User{
		ID: "1",
		FIRSTNAME: "Shravan",
		LASTNAME: "Chaudhary",
	}, nil
}
