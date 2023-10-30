package resource

import model "jora/app/models"


func UserIndexResourceCollection(items []model.User) (result []interface{}) {

	for _, item := range items {
		result = append(result, UserIndexResource(item))
	}

	return
}

func UserIndexResource(item model.User) map[string]interface{}{
	return map[string]interface{}{
		"id": item.ID,
		"register_number": item.RegisterNumber,
		"first_name": item.FirstName,
		"last_name": item.LastName,
		"avatar": item.Avatar,


		"created_at": item.CreatedAt,
		"deleted_at": item.DeletedAt,
	}
}