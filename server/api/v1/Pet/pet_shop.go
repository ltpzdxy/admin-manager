package Pet

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Pet"
    PetReq "github.com/flipped-aurora/gin-vue-admin/server/model/Pet/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PetShopApi struct {}



// CreatePetShop 创建宠物店信息
// @Tags PetShop
// @Summary 创建宠物店信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body Pet.PetShop true "创建宠物店信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /petshop/createPetShop [post]
func (petshopApi *PetShopApi) CreatePetShop(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var petshop Pet.PetShop
	err := c.ShouldBindJSON(&petshop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = petshopService.CreatePetShop(ctx,&petshop)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePetShop 删除宠物店信息
// @Tags PetShop
// @Summary 删除宠物店信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body Pet.PetShop true "删除宠物店信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /petshop/deletePetShop [delete]
func (petshopApi *PetShopApi) DeletePetShop(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := petshopService.DeletePetShop(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePetShopByIds 批量删除宠物店信息
// @Tags PetShop
// @Summary 批量删除宠物店信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /petshop/deletePetShopByIds [delete]
func (petshopApi *PetShopApi) DeletePetShopByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := petshopService.DeletePetShopByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePetShop 更新宠物店信息
// @Tags PetShop
// @Summary 更新宠物店信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body Pet.PetShop true "更新宠物店信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /petshop/updatePetShop [put]
func (petshopApi *PetShopApi) UpdatePetShop(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var petshop Pet.PetShop
	err := c.ShouldBindJSON(&petshop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = petshopService.UpdatePetShop(ctx,petshop)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPetShop 用id查询宠物店信息
// @Tags PetShop
// @Summary 用id查询宠物店信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询宠物店信息"
// @Success 200 {object} response.Response{data=Pet.PetShop,msg=string} "查询成功"
// @Router /petshop/findPetShop [get]
func (petshopApi *PetShopApi) FindPetShop(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	repetshop, err := petshopService.GetPetShop(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(repetshop, c)
}
// GetPetShopList 分页获取宠物店信息列表
// @Tags PetShop
// @Summary 分页获取宠物店信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query PetReq.PetShopSearch true "分页获取宠物店信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /petshop/getPetShopList [get]
func (petshopApi *PetShopApi) GetPetShopList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo PetReq.PetShopSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := petshopService.GetPetShopInfoList(ctx,pageInfo)
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

// GetPetShopPublic 不需要鉴权的宠物店信息接口
// @Tags PetShop
// @Summary 不需要鉴权的宠物店信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /petshop/getPetShopPublic [get]
func (petshopApi *PetShopApi) GetPetShopPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    petshopService.GetPetShopPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的宠物店信息接口信息",
    }, "获取成功", c)
}
