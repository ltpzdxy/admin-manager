package Pet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PetAnimalsRouter struct {}

// InitPetAnimalsRouter 初始化 宠物信息表 路由信息
func (s *PetAnimalsRouter) InitPetAnimalsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	petAnimalsRouter := Router.Group("petAnimals").Use(middleware.OperationRecord())
	petAnimalsRouterWithoutRecord := Router.Group("petAnimals")
	petAnimalsRouterWithoutAuth := PublicRouter.Group("petAnimals")
	{
		petAnimalsRouter.POST("createPetAnimals", petAnimalsApi.CreatePetAnimals)   // 新建宠物信息表
		petAnimalsRouter.DELETE("deletePetAnimals", petAnimalsApi.DeletePetAnimals) // 删除宠物信息表
		petAnimalsRouter.DELETE("deletePetAnimalsByIds", petAnimalsApi.DeletePetAnimalsByIds) // 批量删除宠物信息表
		petAnimalsRouter.PUT("updatePetAnimals", petAnimalsApi.UpdatePetAnimals)    // 更新宠物信息表
	}
	{
		petAnimalsRouterWithoutRecord.GET("findPetAnimals", petAnimalsApi.FindPetAnimals)        // 根据ID获取宠物信息表
		petAnimalsRouterWithoutRecord.GET("getPetAnimalsList", petAnimalsApi.GetPetAnimalsList)  // 获取宠物信息表列表
	}
	{
	    petAnimalsRouterWithoutAuth.GET("getPetAnimalsPublic", petAnimalsApi.GetPetAnimalsPublic)  // 宠物信息表开放接口
	}
}
