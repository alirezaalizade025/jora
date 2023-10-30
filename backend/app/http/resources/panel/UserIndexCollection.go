package resource

import model "jora/app/models"


func UserIndexCollection(items []model.User, pagination map[string]int) (result map[string]interface{}) {

	var data []interface{}
	if data = UserIndexResourceCollection(items); len(data) == 0 {
		data = []interface{}{}
	}

	result = map[string]interface{}{
		"items": data,
		"pagination": pagination,
	}

	return
}