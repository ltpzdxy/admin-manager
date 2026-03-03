package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Pet"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(Pet.PetShop{}, Pet.PetCustomers{}, Pet.PetAnimals{}, Pet.PetTasks{})
	if err != nil {
		return err
	}
	return nil
}
