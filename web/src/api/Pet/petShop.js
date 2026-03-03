import service from '@/utils/request'
// @Tags PetShop
// @Summary 创建宠物店信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetShop true "创建宠物店信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /petshop/createPetShop [post]
export const createPetShop = (data) => {
  return service({
    url: '/petshop/createPetShop',
    method: 'post',
    data
  })
}

// @Tags PetShop
// @Summary 删除宠物店信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetShop true "删除宠物店信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /petshop/deletePetShop [delete]
export const deletePetShop = (params) => {
  return service({
    url: '/petshop/deletePetShop',
    method: 'delete',
    params
  })
}

// @Tags PetShop
// @Summary 批量删除宠物店信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除宠物店信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /petshop/deletePetShop [delete]
export const deletePetShopByIds = (params) => {
  return service({
    url: '/petshop/deletePetShopByIds',
    method: 'delete',
    params
  })
}

// @Tags PetShop
// @Summary 更新宠物店信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetShop true "更新宠物店信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /petshop/updatePetShop [put]
export const updatePetShop = (data) => {
  return service({
    url: '/petshop/updatePetShop',
    method: 'put',
    data
  })
}

// @Tags PetShop
// @Summary 用id查询宠物店信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PetShop true "用id查询宠物店信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /petshop/findPetShop [get]
export const findPetShop = (params) => {
  return service({
    url: '/petshop/findPetShop',
    method: 'get',
    params
  })
}

// @Tags PetShop
// @Summary 分页获取宠物店信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取宠物店信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /petshop/getPetShopList [get]
export const getPetShopList = (params) => {
  return service({
    url: '/petshop/getPetShopList',
    method: 'get',
    params
  })
}

// @Tags PetShop
// @Summary 不需要鉴权的宠物店信息接口
// @Accept application/json
// @Produce application/json
// @Param data query PetReq.PetShopSearch true "分页获取宠物店信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /petshop/getPetShopPublic [get]
export const getPetShopPublic = () => {
  return service({
    url: '/petshop/getPetShopPublic',
    method: 'get',
  })
}
