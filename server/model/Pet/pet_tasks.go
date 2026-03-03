
// 自动生成模板PetTasks
package Pet
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 任务信息表 结构体  PetTasks
type PetTasks struct {
    global.GVA_MODEL
  Shop_id  *int64 `json:"shop_id" form:"shop_id" gorm:"column:shop_id;"`  //店铺名
  Customer_id  *int64 `json:"customer_id" form:"customer_id" gorm:"comment:关联客户;column:customer_id;"`  //关联客户
  Pet_id  *int64 `json:"pet_id" form:"pet_id" gorm:"comment:关联宠物;column:pet_id;"`  //关联宠物
  Title  *string `json:"title" form:"title" gorm:"comment:任务标题;column:title;"`  //任务标题
  Type  *int64 `json:"type" form:"type" gorm:"comment:任务类型;column:type;"`  //任务类型
  DueAt  *time.Time `json:"dueAt" form:"dueAt" gorm:"comment:到期时间;column:due_at;"`  //到期时间
  Priority  *int64 `json:"priority" form:"priority" gorm:"comment:优先级;column:priority;"`  //优先级
  Avatar  string `json:"avatar" form:"avatar" gorm:"comment:照片;column:avatar;"`  //照片
  Status  *int64 `json:"status" form:"status" gorm:"comment:状态;column:status;"`  //状态
}


// TableName 任务信息表 PetTasks自定义表名 pet_tasks
func (PetTasks) TableName() string {
    return "pet_tasks"
}





