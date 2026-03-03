package Pet

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Pet"
    PetReq "github.com/flipped-aurora/gin-vue-admin/server/model/Pet/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PetCustomersApi struct {}



// CreatePetCustomers 创建客户信息表
// @Tags PetCustomers
// @Summary 创建客户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body Pet.PetCustomers true "创建客户信息表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /petCustomers/createPetCustomers [post]
func (petCustomersApi *PetCustomersApi) CreatePetCustomers(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var petCustomers Pet.PetCustomers
	err := c.ShouldBindJSON(&petCustomers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = petCustomersService.CreatePetCustomers(ctx,&petCustomers)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePetCustomers 删除客户信息表
// @Tags PetCustomers
// @Summary 删除客户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body Pet.PetCustomers true "删除客户信息表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /petCustomers/deletePetCustomers [delete]
func (petCustomersApi *PetCustomersApi) DeletePetCustomers(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := petCustomersService.DeletePetCustomers(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePetCustomersByIds 批量删除客户信息表
// @Tags PetCustomers
// @Summary 批量删除客户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /petCustomers/deletePetCustomersByIds [delete]
func (petCustomersApi *PetCustomersApi) DeletePetCustomersByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := petCustomersService.DeletePetCustomersByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePetCustomers 更新客户信息表
// @Tags PetCustomers
// @Summary 更新客户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body Pet.PetCustomers true "更新客户信息表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /petCustomers/updatePetCustomers [put]
func (petCustomersApi *PetCustomersApi) UpdatePetCustomers(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var petCustomers Pet.PetCustomers
	err := c.ShouldBindJSON(&petCustomers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = petCustomersService.UpdatePetCustomers(ctx,petCustomers)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPetCustomers 用id查询客户信息表
// @Tags PetCustomers
// @Summary 用id查询客户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询客户信息表"
// @Success 200 {object} response.Response{data=Pet.PetCustomers,msg=string} "查询成功"
// @Router /petCustomers/findPetCustomers [get]
func (petCustomersApi *PetCustomersApi) FindPetCustomers(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	repetCustomers, err := petCustomersService.GetPetCustomers(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(repetCustomers, c)
}
// GetPetCustomersList 分页获取客户信息表列表
// @Tags PetCustomers
// @Summary 分页获取客户信息表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query PetReq.PetCustomersSearch true "分页获取客户信息表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /petCustomers/getPetCustomersList [get]
func (petCustomersApi *PetCustomersApi) GetPetCustomersList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo PetReq.PetCustomersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := petCustomersService.GetPetCustomersInfoList(ctx,pageInfo)
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

// GetPetCustomersPublic 不需要鉴权的客户信息表接口
// @Tags PetCustomers
// @Summary 不需要鉴权的客户信息表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /petCustomers/getPetCustomersPublic [get]
func (petCustomersApi *PetCustomersApi) GetPetCustomersPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    petCustomersService.GetPetCustomersPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的客户信息表接口信息",
    }, "获取成功", c)
}
