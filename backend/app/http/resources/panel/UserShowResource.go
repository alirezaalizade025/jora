package resource

import model "jora/app/models"



func UserShowResource(item model.User) map[string]interface{}{

	team := item.GetTeam()

	return map[string]interface{}{
		"id": item.ID,
		"register_number": item.RegisterNumber,
		"first_name": item.FirstName,
		"last_name": item.LastName,
		"avatar": item.Avatar,
		"team": map[string]interface{}{
			"id": team["id"],
			"title": team["title"],
		},


		"created_at": item.CreatedAt,
		"deleted_at": item.DeletedAt,
	}
}