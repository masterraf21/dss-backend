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

type menuRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

//  NewMenuRepo will create new menu repository object
func NewMenuRepo(inst *mongo.Database, ctr models.CounterRepository) models.MenuRepository {
	return &menuRepo{
		Instance:    inst,
		CounterRepo: ctr,
	}
}

func (r *menuRepo) Store(menu *models.Menu) (id uint32, err error) {
	collectionName := "menu"
	identifier := "id_menu"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err = r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	menu.ID = id

	_, err = collection.InsertOne(ctx, menu)
	if err != nil {
		return
	}

	return
}

func (r *menuRepo) BulkStore(menus []*models.Menu) (res []uint32, err error) {
	collectionName := "menu"
	identifier := "id_menu"

	var id uint32
	var input []interface{}

	res = make([]uint32, 0)

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	collection := r.Instance.Collection(collectionName)

	for i := 0; i < len(menus); i++ {
		id, err = r.CounterRepo.Get(collectionName, identifier)
		if err != nil {
			return
		}

		menus[i].ID = id
		res = append(res, id)
	}

	for _, menu := range menus {
		input = append(input, menu)
	}

	_, err = collection.InsertMany(ctx, input)
	if err != nil {
		return
	}

	return
}

func (r *menuRepo) GetAll() (res []models.Menu, err error) {
	collectionName := "menu"

	collection := r.Instance.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Menu, 0)
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

func (r *menuRepo) GetByID(id uint32) (res *models.Menu, err error) {
	collectionName := "menu"
	identifier := "id_menu"

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

func (r *menuRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collectionName := "menu"
	identifier := "id_menu"

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
