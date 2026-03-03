
package Pet

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Pet"
    PetReq "github.com/flipped-aurora/gin-vue-admin/server/model/Pet/request"
)

type PetTasksService struct {}
// CreatePetTasks 创建任务信息表记录
// Author [yourname](https://github.com/yourname)
func (petTasksService *PetTasksService) CreatePetTasks(ctx context.Context, petTasks *Pet.PetTasks) (err error) {
	err = global.GVA_DB.Create(petTasks).Error
	return err
}

// DeletePetTasks 删除任务信息表记录
// Author [yourname](https://github.com/yourname)
func (petTasksService *PetTasksService)DeletePetTasks(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&Pet.PetTasks{},"id = ?",ID).Error
	return err
}

// DeletePetTasksByIds 批量删除任务信息表记录
// Author [yourname](https://github.com/yourname)
func (petTasksService *PetTasksService)DeletePetTasksByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Pet.PetTasks{},"id in ?",IDs).Error
	return err
}

// UpdatePetTasks 更新任务信息表记录
// Author [yourname](https://github.com/yourname)
func (petTasksService *PetTasksService)UpdatePetTasks(ctx context.Context, petTasks Pet.PetTasks) (err error) {
	err = global.GVA_DB.Model(&Pet.PetTasks{}).Where("id = ?",petTasks.ID).Updates(&petTasks).Error
	return err
}

// GetPetTasks 根据ID获取任务信息表记录
// Author [yourname](https://github.com/yourname)
func (petTasksService *PetTasksService)GetPetTasks(ctx context.Context, ID string) (petTasks Pet.PetTasks, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&petTasks).Error
	return
}
// GetPetTasksInfoList 分页获取任务信息表记录
// Author [yourname](https://github.com/yourname)
func (petTasksService *PetTasksService)GetPetTasksInfoList(ctx context.Context, info PetReq.PetTasksSearch) (list []Pet.PetTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Pet.PetTasks{})
    var petTaskss []Pet.PetTasks
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
			if len(info.DueAtRange) == 2 {
				db = db.Where("due_at BETWEEN ? AND ? ", info.DueAtRange[0], info.DueAtRange[1])
			}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&petTaskss).Error
	return  petTaskss, total, err
}
func (petTasksService *PetTasksService)GetPetTasksPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
