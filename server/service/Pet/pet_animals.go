
package Pet

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Pet"
    PetReq "github.com/flipped-aurora/gin-vue-admin/server/model/Pet/request"
)

type PetAnimalsService struct {}
// CreatePetAnimals 创建宠物信息表记录
// Author [yourname](https://github.com/yourname)
func (petAnimalsService *PetAnimalsService) CreatePetAnimals(ctx context.Context, petAnimals *Pet.PetAnimals) (err error) {
	err = global.GVA_DB.Create(petAnimals).Error
	return err
}

// DeletePetAnimals 删除宠物信息表记录
// Author [yourname](https://github.com/yourname)
func (petAnimalsService *PetAnimalsService)DeletePetAnimals(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&Pet.PetAnimals{},"id = ?",ID).Error
	return err
}

// DeletePetAnimalsByIds 批量删除宠物信息表记录
// Author [yourname](https://github.com/yourname)
func (petAnimalsService *PetAnimalsService)DeletePetAnimalsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Pet.PetAnimals{},"id in ?",IDs).Error
	return err
}

// UpdatePetAnimals 更新宠物信息表记录
// Author [yourname](https://github.com/yourname)
func (petAnimalsService *PetAnimalsService)UpdatePetAnimals(ctx context.Context, petAnimals Pet.PetAnimals) (err error) {
	err = global.GVA_DB.Model(&Pet.PetAnimals{}).Where("id = ?",petAnimals.ID).Updates(&petAnimals).Error
	return err
}

// GetPetAnimals 根据ID获取宠物信息表记录
// Author [yourname](https://github.com/yourname)
func (petAnimalsService *PetAnimalsService)GetPetAnimals(ctx context.Context, ID string) (petAnimals Pet.PetAnimals, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&petAnimals).Error
	return
}
// GetPetAnimalsInfoList 分页获取宠物信息表记录
// Author [yourname](https://github.com/yourname)
func (petAnimalsService *PetAnimalsService)GetPetAnimalsInfoList(ctx context.Context, info PetReq.PetAnimalsSearch) (list []Pet.PetAnimals, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Pet.PetAnimals{})
    var petAnimalss []Pet.PetAnimals
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Nickname != nil && *info.Nickname != "" {
        db = db.Where("nickname LIKE ?", "%"+ *info.Nickname+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&petAnimalss).Error
	return  petAnimalss, total, err
}
func (petAnimalsService *PetAnimalsService)GetPetAnimalsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
