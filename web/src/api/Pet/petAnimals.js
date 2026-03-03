import service from '@/utils/request'
// @Tags PetAnimals
// @Summary 创建宠物信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetAnimals true "创建宠物信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /petAnimals/createPetAnimals [post]
export const createPetAnimals = (data) => {
  return service({
    url: '/petAnimals/createPetAnimals',
    method: 'post',
    data
  })
}

// @Tags PetAnimals
// @Summary 删除宠物信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetAnimals true "删除宠物信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /petAnimals/deletePetAnimals [delete]
export const deletePetAnimals = (params) => {
  return service({
    url: '/petAnimals/deletePetAnimals',
    method: 'delete',
    params
  })
}

// @Tags PetAnimals
// @Summary 批量删除宠物信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除宠物信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /petAnimals/deletePetAnimals [delete]
export const deletePetAnimalsByIds = (params) => {
  return service({
    url: '/petAnimals/deletePetAnimalsByIds',
    method: 'delete',
    params
  })
}

// @Tags PetAnimals
// @Summary 更新宠物信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetAnimals true "更新宠物信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /petAnimals/updatePetAnimals [put]
export const updatePetAnimals = (data) => {
  return service({
    url: '/petAnimals/updatePetAnimals',
    method: 'put',
    data
  })
}

// @Tags PetAnimals
// @Summary 用id查询宠物信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PetAnimals true "用id查询宠物信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /petAnimals/findPetAnimals [get]
export const findPetAnimals = (params) => {
  return service({
    url: '/petAnimals/findPetAnimals',
    method: 'get',
    params
  })
}

// @Tags PetAnimals
// @Summary 分页获取宠物信息表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取宠物信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /petAnimals/getPetAnimalsList [get]
export const getPetAnimalsList = (params) => {
  return service({
    url: '/petAnimals/getPetAnimalsList',
    method: 'get',
    params
  })
}

// @Tags PetAnimals
// @Summary 不需要鉴权的宠物信息表接口
// @Accept application/json
// @Produce application/json
// @Param data query PetReq.PetAnimalsSearch true "分页获取宠物信息表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /petAnimals/getPetAnimalsPublic [get]
export const getPetAnimalsPublic = () => {
  return service({
    url: '/petAnimals/getPetAnimalsPublic',
    method: 'get',
  })
}
