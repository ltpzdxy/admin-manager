package Pet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PetShopRouter struct {}

// InitPetShopRouter 初始化 宠物店信息 路由信息
func (s *PetShopRouter) InitPetShopRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	petshopRouter := Router.Group("petshop").Use(middleware.OperationRecord())
	petshopRouterWithoutRecord := Router.Group("petshop")
	petshopRouterWithoutAuth := PublicRouter.Group("petshop")
	{
		petshopRouter.POST("createPetShop", petshopApi.CreatePetShop)   // 新建宠物店信息
		petshopRouter.DELETE("deletePetShop", petshopApi.DeletePetShop) // 删除宠物店信息
		petshopRouter.DELETE("deletePetShopByIds", petshopApi.DeletePetShopByIds) // 批量删除宠物店信息
		petshopRouter.PUT("updatePetShop", petshopApi.UpdatePetShop)    // 更新宠物店信息
	}
	{
		petshopRouterWithoutRecord.GET("findPetShop", petshopApi.FindPetShop)        // 根据ID获取宠物店信息
		petshopRouterWithoutRecord.GET("getPetShopList", petshopApi.GetPetShopList)  // 获取宠物店信息列表
	}
	{
	    petshopRouterWithoutAuth.GET("getPetShopPublic", petshopApi.GetPetShopPublic)  // 宠物店信息开放接口
	}
}
