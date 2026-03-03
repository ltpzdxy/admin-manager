
package Pet

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Pet"
    PetReq "github.com/flipped-aurora/gin-vue-admin/server/model/Pet/request"
)

type PetShopService struct {}
// CreatePetShop 创建宠物店信息记录
// Author [yourname](https://github.com/yourname)
func (petshopService *PetShopService) CreatePetShop(ctx context.Context, petshop *Pet.PetShop) (err error) {
	err = global.GVA_DB.Create(petshop).Error
	return err
}

// DeletePetShop 删除宠物店信息记录
// Author [yourname](https://github.com/yourname)
func (petshopService *PetShopService)DeletePetShop(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&Pet.PetShop{},"id = ?",ID).Error
	return err
}

// DeletePetShopByIds 批量删除宠物店信息记录
// Author [yourname](https://github.com/yourname)
func (petshopService *PetShopService)DeletePetShopByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Pet.PetShop{},"id in ?",IDs).Error
	return err
}

// UpdatePetShop 更新宠物店信息记录
// Author [yourname](https://github.com/yourname)
func (petshopService *PetShopService)UpdatePetShop(ctx context.Context, petshop Pet.PetShop) (err error) {
	err = global.GVA_DB.Model(&Pet.PetShop{}).Where("id = ?",petshop.ID).Updates(&petshop).Error
	return err
}

// GetPetShop 根据ID获取宠物店信息记录
// Author [yourname](https://github.com/yourname)
func (petshopService *PetShopService)GetPetShop(ctx context.Context, ID string) (petshop Pet.PetShop, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&petshop).Error
	return
}
// GetPetShopInfoList 分页获取宠物店信息记录
// Author [yourname](https://github.com/yourname)
func (petshopService *PetShopService)GetPetShopInfoList(ctx context.Context, info PetReq.PetShopSearch) (list []Pet.PetShop, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Pet.PetShop{})
    var petshops []Pet.PetShop
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.ShopName != nil && *info.ShopName != "" {
        db = db.Where("ShopName LIKE ?", "%"+ *info.ShopName+"%")
    }
    if info.BossName != nil && *info.BossName != "" {
        db = db.Where("BossName LIKE ?", "%"+ *info.BossName+"%")
    }
    if info.Phone != nil {
        db = db.Where("Phone = ?", *info.Phone)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&petshops).Error
	return  petshops, total, err
}
func (petshopService *PetShopService)GetPetShopPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
