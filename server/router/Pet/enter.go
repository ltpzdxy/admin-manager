package Pet

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	PetShopRouter
	PetCustomersRouter
	PetAnimalsRouter
	PetTasksRouter
}

var (
	petshopApi      = api.ApiGroupApp.PetApiGroup.PetShopApi
	petCustomersApi = api.ApiGroupApp.PetApiGroup.PetCustomersApi
	petAnimalsApi   = api.ApiGroupApp.PetApiGroup.PetAnimalsApi
	petTasksApi     = api.ApiGroupApp.PetApiGroup.PetTasksApi
)
