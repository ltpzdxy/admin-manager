import service from '@/utils/request'
// @Tags PetCustomers
// @Summary 创建客户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetCustomers true "创建客户信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /petCustomers/createPetCustomers [post]
export const createPetCustomers = (data) => {
  return service({
    url: '/petCustomers/createPetCustomers',
    method: 'post',
    data
  })
}

// @Tags PetCustomers
// @Summary 删除客户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetCustomers true "删除客户信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /petCustomers/deletePetCustomers [delete]
export const deletePetCustomers = (params) => {
  return service({
    url: '/petCustomers/deletePetCustomers',
    method: 'delete',
    params
  })
}

// @Tags PetCustomers
// @Summary 批量删除客户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除客户信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /petCustomers/deletePetCustomers [delete]
export const deletePetCustomersByIds = (params) => {
  return service({
    url: '/petCustomers/deletePetCustomersByIds',
    method: 'delete',
    params
  })
}

// @Tags PetCustomers
// @Summary 更新客户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetCustomers true "更新客户信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /petCustomers/updatePetCustomers [put]
export const updatePetCustomers = (data) => {
  return service({
    url: '/petCustomers/updatePetCustomers',
    method: 'put',
    data
  })
}

// @Tags PetCustomers
// @Summary 用id查询客户信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PetCustomers true "用id查询客户信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /petCustomers/findPetCustomers [get]
export const findPetCustomers = (params) => {
  return service({
    url: '/petCustomers/findPetCustomers',
    method: 'get',
    params
  })
}

// @Tags PetCustomers
// @Summary 分页获取客户信息表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取客户信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /petCustomers/getPetCustomersList [get]
export const getPetCustomersList = (params) => {
  return service({
    url: '/petCustomers/getPetCustomersList',
    method: 'get',
    params
  })
}

// @Tags PetCustomers
// @Summary 不需要鉴权的客户信息表接口
// @Accept application/json
// @Produce application/json
// @Param data query PetReq.PetCustomersSearch true "分页获取客户信息表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /petCustomers/getPetCustomersPublic [get]
export const getPetCustomersPublic = () => {
  return service({
    url: '/petCustomers/getPetCustomersPublic',
    method: 'get',
  })
}
