package repository

import (
	"test-with-revel/app"
	"test-with-revel/app/models"

	"github.com/revel/revel"
)

type DummyRepository interface {
	Create(string, string, string)
	Update(string, string)
	Delete(string)
}

type DBDummyRepository struct {
	dummy []models.Dummy
}

// Create
func (r DBDummyRepository) Create(col1 string, col2 string, col3 string) {
	d1 := &models.Dummy{col1, col2, col3}
	err := app.Dbm.Insert(d1)

	if err != nil {
		revel.AppLog.Debug("DB Error: ", err)
	}
}

// Update ...
func (r DBDummyRepository) Update(col1 string, col3 string) {
	operation, err := app.Dbm.Exec("UPDATE dummy SET dummy_col3=? WHERE dummy_col1=?", col3, col1)
	if err != nil {
		revel.AppLog.Debug("DB Error: ", err)
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		revel.AppLog.Debug("DB Error: updated 0 records")
	}
}

//Delete ...
func (r DBDummyRepository) Delete(col1 string) {

	operation, err := app.Dbm.Exec("DELETE FROM dummy WHERE dummy_col1=?", col1)
	if err != nil {
		revel.AppLog.Debug("DB Error: ", err)
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		revel.AppLog.Debug("DB Error: no record was deleted")
	}
}

var dummyRepository = new(DBDummyRepository)

func GetDummyRepository() (r DummyRepository) {
	return dummyRepository
}
