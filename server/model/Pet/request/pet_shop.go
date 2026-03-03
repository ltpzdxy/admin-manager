
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PetShopSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      ShopName  *string `json:"ShopName" form:"ShopName"` 
      BossName  *string `json:"BossName" form:"BossName"` 
      Phone  *int `json:"Phone" form:"Phone"` 
    request.PageInfo
}
