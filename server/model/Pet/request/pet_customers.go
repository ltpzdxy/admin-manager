
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PetCustomersSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      StartPhone  *int  `json:"startPhone" form:"startPhone"`
EndPhone  *int  `json:"endPhone" form:"endPhone"`
    request.PageInfo
}
