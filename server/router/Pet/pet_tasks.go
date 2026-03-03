package Pet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PetTasksRouter struct {}

// InitPetTasksRouter 初始化 任务信息表 路由信息
func (s *PetTasksRouter) InitPetTasksRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	petTasksRouter := Router.Group("petTasks").Use(middleware.OperationRecord())
	petTasksRouterWithoutRecord := Router.Group("petTasks")
	petTasksRouterWithoutAuth := PublicRouter.Group("petTasks")
	{
		petTasksRouter.POST("createPetTasks", petTasksApi.CreatePetTasks)   // 新建任务信息表
		petTasksRouter.DELETE("deletePetTasks", petTasksApi.DeletePetTasks) // 删除任务信息表
		petTasksRouter.DELETE("deletePetTasksByIds", petTasksApi.DeletePetTasksByIds) // 批量删除任务信息表
		petTasksRouter.PUT("updatePetTasks", petTasksApi.UpdatePetTasks)    // 更新任务信息表
	}
	{
		petTasksRouterWithoutRecord.GET("findPetTasks", petTasksApi.FindPetTasks)        // 根据ID获取任务信息表
		petTasksRouterWithoutRecord.GET("getPetTasksList", petTasksApi.GetPetTasksList)  // 获取任务信息表列表
	}
	{
	    petTasksRouterWithoutAuth.GET("getPetTasksPublic", petTasksApi.GetPetTasksPublic)  // 任务信息表开放接口
	}
}
