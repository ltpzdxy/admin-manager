
// 自动生成模板PetAnimals
package Pet
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 宠物信息表 结构体  PetAnimals
type PetAnimals struct {
    global.GVA_MODEL
  Shop_id  *int64 `json:"shop_id" form:"shop_id" gorm:"column:shop_id;"`  //店铺名
  Customer_id  *int64 `json:"customer_id" form:"customer_id" gorm:"comment:所属客户;column:customer_id;"`  //所属客户
  Nickname  *string `json:"nickname" form:"nickname" gorm:"comment:昵称;column:nickname;"`  //昵称
  Breed  *string `json:"breed" form:"breed" gorm:"comment:品种;column:breed;"`  //品种
  Is_neutered  *bool `json:"is_neutered" form:"is_neutered" gorm:"comment:是否绝育;column:is_neutered;"`  //是否绝育
  Weight  *float64 `json:"weight" form:"weight" gorm:"comment:体重;column:weight;"`  //体重
  Allergy  *string `json:"allergy" form:"allergy" gorm:"comment:过敏/禁忌;column:allergy;"`  //过敏/禁忌
  Avatar  string `json:"avatar" form:"avatar" gorm:"comment:照片;column:avatar;"`  //照片
}


// TableName 宠物信息表 PetAnimals自定义表名 pet_animals
func (PetAnimals) TableName() string {
    return "pet_animals"
}





