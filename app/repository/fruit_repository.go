package repository

import (
	"test-with-revel/app"
	"test-with-revel/app/models"

	"github.com/revel/revel"
)

type FruitRepository interface {
	GetFruits() []models.Fruit
}

type DBFruitRepository struct {
	fruits []models.Fruit
}

func (r DBFruitRepository) GetFruits() (fruits []models.Fruit) {
	_, err := app.Dbm.Select(&fruits, "SELECT * FROM fruit")

	if err != nil {
		revel.AppLog.Debug("DB Error: ", err)
	}

	return fruits
}

var fruitRepository = new(DBFruitRepository)

func GetFruitRepository() (r FruitRepository) {
	return fruitRepository
}
