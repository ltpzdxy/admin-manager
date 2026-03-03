import service from '@/utils/request'
// @Tags PetTasks
// @Summary 创建任务信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetTasks true "创建任务信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /petTasks/createPetTasks [post]
export const createPetTasks = (data) => {
  return service({
    url: '/petTasks/createPetTasks',
    method: 'post',
    data
  })
}

// @Tags PetTasks
// @Summary 删除任务信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetTasks true "删除任务信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /petTasks/deletePetTasks [delete]
export const deletePetTasks = (params) => {
  return service({
    url: '/petTasks/deletePetTasks',
    method: 'delete',
    params
  })
}

// @Tags PetTasks
// @Summary 批量删除任务信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除任务信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /petTasks/deletePetTasks [delete]
export const deletePetTasksByIds = (params) => {
  return service({
    url: '/petTasks/deletePetTasksByIds',
    method: 'delete',
    params
  })
}

// @Tags PetTasks
// @Summary 更新任务信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PetTasks true "更新任务信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /petTasks/updatePetTasks [put]
export const updatePetTasks = (data) => {
  return service({
    url: '/petTasks/updatePetTasks',
    method: 'put',
    data
  })
}

// @Tags PetTasks
// @Summary 用id查询任务信息表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PetTasks true "用id查询任务信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /petTasks/findPetTasks [get]
export const findPetTasks = (params) => {
  return service({
    url: '/petTasks/findPetTasks',
    method: 'get',
    params
  })
}

// @Tags PetTasks
// @Summary 分页获取任务信息表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取任务信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /petTasks/getPetTasksList [get]
export const getPetTasksList = (params) => {
  return service({
    url: '/petTasks/getPetTasksList',
    method: 'get',
    params
  })
}

// @Tags PetTasks
// @Summary 不需要鉴权的任务信息表接口
// @Accept application/json
// @Produce application/json
// @Param data query PetReq.PetTasksSearch true "分页获取任务信息表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /petTasks/getPetTasksPublic [get]
export const getPetTasksPublic = () => {
  return service({
    url: '/petTasks/getPetTasksPublic',
    method: 'get',
  })
}
