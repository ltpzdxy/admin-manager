
// 自动生成模板PetShop
package Pet
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 宠物店信息 结构体  PetShop
type PetShop struct {
    global.GVA_MODEL
  ShopName  *string `json:"ShopName" form:"ShopName" gorm:"comment:店铺名称;column:ShopName;" binding:"required"`  //店铺名称
  BossName  *string `json:"BossName" form:"BossName" gorm:"comment:负责人;column:BossName;" binding:"required"`  //负责人
  Phone  *int64 `json:"Phone" form:"Phone" gorm:"comment:手机号;column:Phone;" binding:"required"`  //手机号
}


// TableName 宠物店信息 PetShop自定义表名 pet_shops
func (PetShop) TableName() string {
    return "pet_shops"
}





