package Pet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PetCustomersRouter struct {}

// InitPetCustomersRouter 初始化 客户信息表 路由信息
func (s *PetCustomersRouter) InitPetCustomersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	petCustomersRouter := Router.Group("petCustomers").Use(middleware.OperationRecord())
	petCustomersRouterWithoutRecord := Router.Group("petCustomers")
	petCustomersRouterWithoutAuth := PublicRouter.Group("petCustomers")
	{
		petCustomersRouter.POST("createPetCustomers", petCustomersApi.CreatePetCustomers)   // 新建客户信息表
		petCustomersRouter.DELETE("deletePetCustomers", petCustomersApi.DeletePetCustomers) // 删除客户信息表
		petCustomersRouter.DELETE("deletePetCustomersByIds", petCustomersApi.DeletePetCustomersByIds) // 批量删除客户信息表
		petCustomersRouter.PUT("updatePetCustomers", petCustomersApi.UpdatePetCustomers)    // 更新客户信息表
	}
	{
		petCustomersRouterWithoutRecord.GET("findPetCustomers", petCustomersApi.FindPetCustomers)        // 根据ID获取客户信息表
		petCustomersRouterWithoutRecord.GET("getPetCustomersList", petCustomersApi.GetPetCustomersList)  // 获取客户信息表列表
	}
	{
	    petCustomersRouterWithoutAuth.GET("getPetCustomersPublic", petCustomersApi.GetPetCustomersPublic)  // 客户信息表开放接口
	}
}
