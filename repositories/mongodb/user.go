package mongodb

import (
	"context"
	"strings"
	"time"

	"github.com/masterraf21/dss-backend/configs"
	"github.com/masterraf21/dss-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionName = "user"
	identifier     = "id_user"
)

type userRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewUserRepo will create new userrepo
func NewUserRepo(instance *mongo.Database, ctr models.CounterRepository) models.UserRepository {
	return &userRepo{
		Instance:    instance,
		CounterRepo: ctr,
	}
}

func (r *userRepo) Store(user *models.User) (id uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err = r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	user.ID = id

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return
	}

	// _id := result.InsertedID
	// uid = _id.(primitive.ObjectID)

	return
}

func (r *userRepo) GetByID(id uint32) (res *models.User, err error) {
	collection := r.Instance.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{identifier: id}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *userRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collection := r.Instance.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{identifier: id},
		bson.M{"$set": bson.M{key: value}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) BulkStore(users []*models.User) (res []uint32, err error) {
	var id uint32
	var input []interface{}

	res = make([]uint32, 0)

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	collection := r.Instance.Collection(collectionName)

	for i := 0; i < len(users); i++ {
		id, err = r.CounterRepo.Get(collectionName, identifier)
		if err != nil {
			return
		}

		users[i].ID = id
	}

	for _, user := range users {
		input = append(input, user)
	}

	_, err = collection.InsertMany(ctx, input)
	if err != nil {
		return
	}

	return
}

func (r *userRepo) GetAll() (res []models.User, err error) {
	collection := r.Instance.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.User, 0)
			err = nil
			return
		}

		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}
