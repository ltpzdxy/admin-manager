
package Pet

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Pet"
    PetReq "github.com/flipped-aurora/gin-vue-admin/server/model/Pet/request"
)

type PetCustomersService struct {}
// CreatePetCustomers 创建客户信息表记录
// Author [yourname](https://github.com/yourname)
func (petCustomersService *PetCustomersService) CreatePetCustomers(ctx context.Context, petCustomers *Pet.PetCustomers) (err error) {
	err = global.GVA_DB.Create(petCustomers).Error
	return err
}

// DeletePetCustomers 删除客户信息表记录
// Author [yourname](https://github.com/yourname)
func (petCustomersService *PetCustomersService)DeletePetCustomers(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&Pet.PetCustomers{},"id = ?",ID).Error
	return err
}

// DeletePetCustomersByIds 批量删除客户信息表记录
// Author [yourname](https://github.com/yourname)
func (petCustomersService *PetCustomersService)DeletePetCustomersByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Pet.PetCustomers{},"id in ?",IDs).Error
	return err
}

// UpdatePetCustomers 更新客户信息表记录
// Author [yourname](https://github.com/yourname)
func (petCustomersService *PetCustomersService)UpdatePetCustomers(ctx context.Context, petCustomers Pet.PetCustomers) (err error) {
	err = global.GVA_DB.Model(&Pet.PetCustomers{}).Where("id = ?",petCustomers.ID).Updates(&petCustomers).Error
	return err
}

// GetPetCustomers 根据ID获取客户信息表记录
// Author [yourname](https://github.com/yourname)
func (petCustomersService *PetCustomersService)GetPetCustomers(ctx context.Context, ID string) (petCustomers Pet.PetCustomers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&petCustomers).Error
	return
}
// GetPetCustomersInfoList 分页获取客户信息表记录
// Author [yourname](https://github.com/yourname)
func (petCustomersService *PetCustomersService)GetPetCustomersInfoList(ctx context.Context, info PetReq.PetCustomersSearch) (list []Pet.PetCustomers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Pet.PetCustomers{})
    var petCustomerss []Pet.PetCustomers
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
	if info.StartPhone != nil && info.EndPhone != nil {
		db = db.Where("phone BETWEEN ? AND ? ", *info.StartPhone, *info.EndPhone)
	}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&petCustomerss).Error
	return  petCustomerss, total, err
}
func (petCustomersService *PetCustomersService)GetPetCustomersPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
