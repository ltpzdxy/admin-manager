package Pet

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Pet"
    PetReq "github.com/flipped-aurora/gin-vue-admin/server/model/Pet/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PetAnimalsApi struct {}



// CreatePetAnimals 创建宠物信息表
// @Tags PetAnimals
// @Summary 创建宠物信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body Pet.PetAnimals true "创建宠物信息表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /petAnimals/createPetAnimals [post]
func (petAnimalsApi *PetAnimalsApi) CreatePetAnimals(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var petAnimals Pet.PetAnimals
	err := c.ShouldBindJSON(&petAnimals)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = petAnimalsService.CreatePetAnimals(ctx,&petAnimals)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePetAnimals 删除宠物信息表
// @Tags PetAnimals
// @Summary 删除宠物信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body Pet.PetAnimals true "删除宠物信息表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /petAnimals/deletePetAnimals [delete]
func (petAnimalsApi *PetAnimalsApi) DeletePetAnimals(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := petAnimalsService.DeletePetAnimals(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePetAnimalsByIds 批量删除宠物信息表
// @Tags PetAnimals
// @Summary 批量删除宠物信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /petAnimals/deletePetAnimalsByIds [delete]
func (petAnimalsApi *PetAnimalsApi) DeletePetAnimalsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := petAnimalsService.DeletePetAnimalsByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePetAnimals 更新宠物信息表
// @Tags PetAnimals
// @Summary 更新宠物信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body Pet.PetAnimals true "更新宠物信息表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /petAnimals/updatePetAnimals [put]
func (petAnimalsApi *PetAnimalsApi) UpdatePetAnimals(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var petAnimals Pet.PetAnimals
	err := c.ShouldBindJSON(&petAnimals)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = petAnimalsService.UpdatePetAnimals(ctx,petAnimals)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPetAnimals 用id查询宠物信息表
// @Tags PetAnimals
// @Summary 用id查询宠物信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询宠物信息表"
// @Success 200 {object} response.Response{data=Pet.PetAnimals,msg=string} "查询成功"
// @Router /petAnimals/findPetAnimals [get]
func (petAnimalsApi *PetAnimalsApi) FindPetAnimals(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	repetAnimals, err := petAnimalsService.GetPetAnimals(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(repetAnimals, c)
}
// GetPetAnimalsList 分页获取宠物信息表列表
// @Tags PetAnimals
// @Summary 分页获取宠物信息表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query PetReq.PetAnimalsSearch true "分页获取宠物信息表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /petAnimals/getPetAnimalsList [get]
func (petAnimalsApi *PetAnimalsApi) GetPetAnimalsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo PetReq.PetAnimalsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := petAnimalsService.GetPetAnimalsInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetPetAnimalsPublic 不需要鉴权的宠物信息表接口
// @Tags PetAnimals
// @Summary 不需要鉴权的宠物信息表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /petAnimals/getPetAnimalsPublic [get]
func (petAnimalsApi *PetAnimalsApi) GetPetAnimalsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    petAnimalsService.GetPetAnimalsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的宠物信息表接口信息",
    }, "获取成功", c)
}
