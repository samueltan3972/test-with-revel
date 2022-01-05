package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"test-with-revel/app/repository"

	"github.com/revel/revel"
)

type DatabaseOperationController struct {
	*revel.Controller
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (c DatabaseOperationController) DatabaseOperation() revel.Result {

	// // Create
	// rand.Seed(time.Now().UnixNano())
	// randNum := rand.Intn(1000-1) + 1
	// dt := time.Now()
	// md5hash := GetMD5Hash(dt.String() + fmt.Sprint(randNum))
	// repository.GetDummyRepository().Create(md5hash, "Test", "Test2")

	// // Update
	// repository.GetDummyRepository().Update(md5hash, "Test3")

	// // Delete
	// repository.GetDummyRepository().Delete(md5hash)

	// Return fruits
	fruits := repository.GetFruitRepository().GetFruits()

	if fruits == nil {
		return c.RenderText("Null")
	}

	return c.RenderJSON(fruits)
}
