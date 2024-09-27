package repository

import (
	"context"

	"github.com/naol86/addis-software-starter/project/backend/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	db         *mongo.Database
	collection string
}

// CreateUser implements domain.UserRepository.
func (u *UserRepository) CreateUser(c context.Context, user domain.UserSignupRequest) (domain.User, error) {
	collection := u.db.Collection(u.collection)
	res, err := collection.InsertOne(c, user)
	if err != nil {
		return domain.User{}, err
	}
	var new_user domain.User
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: res.InsertedID}}).Decode(&new_user)
	if err != nil {
		return domain.User{}, err
	}
	return new_user, nil
}

// FindByEmail implements domain.UserRepository.
func (u *UserRepository) FindByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	collection := u.db.Collection(u.collection)
	err := collection.FindOne(c, bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// FindByID implements domain.UserRepository.
func (u *UserRepository) FindByID(c context.Context, id string) (domain.User, error) {
	var user domain.User
	collection := u.db.Collection(u.collection)
	err := collection.FindOne(c, bson.D{{Key: "_id", Value: id}}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &UserRepository{
		db:         db,
		collection: collection,
	}
}
