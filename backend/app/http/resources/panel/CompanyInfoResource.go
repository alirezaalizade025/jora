package resource

import (
	model "jora/app/models"
	"jora/database/postgres"
)


func CompanyInfoResource(item model.Company) map[string]interface{}{

    // Get the count of orders for the user
	var usersCount int64
    postgres.DB.Table("users").Where("company_id = ?", item.ID).Count(&usersCount)

	return map[string]interface{}{
		"id": item.ID,

		"title": item.Title,
		"phone": item.Phone,
		
		"users_count": usersCount,


		"created_at": item.CreatedAt,
		"deleted_at": item.DeletedAt,
	}
}