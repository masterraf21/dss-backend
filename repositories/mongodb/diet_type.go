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

type dietTypeRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

func NewDietTypeRepository(ins *mongo.Database, ctr models.CounterRepository) models.DietTypeRepository {
	return &dietTypeRepo{
		Instance:    ins,
		CounterRepo: ctr,
	}
}

func (r *dietTypeRepo) Store(dietType *models.DietType) (id uint32, err error) {
	collectionName := "diet_type"
	identifier := "id_diet_type"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err = r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	dietType.ID = id

	_, err = collection.InsertOne(ctx, dietType)
	if err != nil {
		return
	}

	return
}

func (r *dietTypeRepo) BulkStore(dietTypes []*models.DietType) (res []uint32, err error) {
	collectionName := "diet_type"
	identifier := "id_diet_type"

	var id uint32
	var input []interface{}

	res = make([]uint32, 0)

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	collection := r.Instance.Collection(collectionName)

	for i := 0; i < len(dietTypes); i++ {
		id, err = r.CounterRepo.Get(collectionName, identifier)
		if err != nil {
			return
		}

		dietTypes[i].ID = id
		res = append(res, id)
	}

	for _, menu := range dietTypes {
		input = append(input, menu)
	}

	_, err = collection.InsertMany(ctx, input)
	if err != nil {
		return
	}

	return
}

func (r *dietTypeRepo) GetAll() (res []models.DietType, err error) {
	collectionName := "diet_type"

	collection := r.Instance.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.DietType, 0)
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

func (r *dietTypeRepo) GetByID(id uint32) (res *models.DietType, err error) {
	collectionName := "diet_type"
	identifier := "id_diet_type"

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

func (r *dietTypeRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collectionName := "diet_type"
	identifier := "id_diet_type"

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
