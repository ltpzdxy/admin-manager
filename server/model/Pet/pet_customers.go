
// 自动生成模板PetCustomers
package Pet
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 客户信息表 结构体  PetCustomers
type PetCustomers struct {
    global.GVA_MODEL
  Shop_id  *int64 `json:"shop_id" form:"shop_id" gorm:"column:shop_id;"`  //店铺名
  Name  *string `json:"name" form:"name" gorm:"comment:姓名;column:name;" binding:"required"`  //姓名
  Phone  *int64 `json:"phone" form:"phone" gorm:"comment:手机号;column:phone;"`  //手机号
  Address  *string `json:"address" form:"address" gorm:"comment:地址;column:address;"`  //地址
  Channel  *string `json:"channel" form:"channel" gorm:"comment:渠道;column:channel;"`  //渠道
  Tags  *string `json:"tags" form:"tags" gorm:"comment:标签;column:tags;"`  //标签
  Balance  *float64 `json:"balance" form:"balance" gorm:"comment:账户余额;column:balance;"`  //账户余额
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;"`  //备注
}


// TableName 客户信息表 PetCustomers自定义表名 pet_customers
func (PetCustomers) TableName() string {
    return "pet_customers"
}





