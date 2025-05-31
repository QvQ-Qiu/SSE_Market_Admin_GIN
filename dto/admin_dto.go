// token结构体
package dto

import "sse_market_admin/model"

type AdminDto struct {
	Account string `json:"account"`
}

func ToAdminDto(admin model.Admin) AdminDto {
	return AdminDto{
		Account: admin.Account,
	}
}
