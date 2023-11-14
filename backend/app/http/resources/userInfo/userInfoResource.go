package resource

import model "jora/app/models"


func UserInfoResource(item model.User) map[string]interface{}{
	return map[string]interface{}{
		"id": item.ID,
		"register_number": item.RegisterNumber,
		"first_name": item.FirstName,
		"last_name": item.LastName,
		"avatar": item.Avatar,
		"roles": item.Roles,


		"created_at": item.CreatedAt,
		"deleted_at": item.DeletedAt,
	}
}