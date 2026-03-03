package Pet

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	PetShopApi
	PetCustomersApi
	PetAnimalsApi
	PetTasksApi
}

var (
	petshopService      = service.ServiceGroupApp.PetServiceGroup.PetShopService
	petCustomersService = service.ServiceGroupApp.PetServiceGroup.PetCustomersService
	petAnimalsService   = service.ServiceGroupApp.PetServiceGroup.PetAnimalsService
	petTasksService     = service.ServiceGroupApp.PetServiceGroup.PetTasksService
)
