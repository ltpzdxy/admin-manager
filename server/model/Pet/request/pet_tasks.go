
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PetTasksSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      DueAtRange  []time.Time  `json:"dueAtRange" form:"dueAtRange[]"`
    request.PageInfo
}
