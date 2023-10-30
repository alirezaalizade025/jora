package usersController

import (
	"fmt"
	request "jora/app/http/requests"
	resource "jora/app/http/resources/panel"
	model "jora/app/models"
	"jora/database/postgres"
	"jora/utility"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	ptime "github.com/yaa110/go-persian-calendar"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	req := &request.PanelUsersIndexRequest{}
	if !request.Validation(c, req) {
		return
	}

	authId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	query := companyQuery(authId)

	if req.User != "" {
		query.Where("id = ? OR concat('first_name', ' ', 'last_name') = ?", req.User, req.User)
	}

	if req.TeamId != nil {
		query.Where("team_id = ?", req.TeamId)
	}

	var users []model.User
	query.Order("created_at DESC").Scopes(postgres.Paginate(c.Request)).Find(&users)

	req.Page = utility.If(req.Page == 0, 1, req.Page)
	req.PerPage = utility.If(req.PerPage == 0, 20, req.PerPage)

	var total int64
	query.Count(&total)
	pagination := map[string]int{
		"total":    int(total),
		"per_page": req.PerPage,
		"page":     req.Page,
		"count":    len(users),
	}

	c.JSON(http.StatusOK, resource.UserIndexCollection(users, pagination))
}

func Create(c *gin.Context) {

	req := &request.PanelUsersCreateRequest{}
	if !request.Validation(c, req) {
		return
	}

	authId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}


	var count int64
	companyQuery(authId).Where("CONCAT(first_name, ' ', last_name) = ?", req.FirstName + " " + req.LastName).Count(&count)
	if count > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "user exists!"})
		return
	}

	if req.RegisterNumber != 0 {
		companyQuery(authId).Where("register_number = ?", req.RegisterNumber).Count(&count)

		if count > 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "register_number exists!"})
			return
		}
	} else {

		req.RegisterNumber = generateRegisterNumber(companyQuery(authId))
	}

	if req.TeamId != 0 {
		postgres.DB.Model(model.Team{}).Where("company_id = ?", authId).Where("team_id = ?", req.TeamId).Count(&count)
		if count == 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "team not exists!"})
			return
		}
	} 

	newUser := model.User{
		CompanyID: uint(authId.(float64)),
		RegisterNumber: req.RegisterNumber,
		FirstName: req.FirstName,
		LastName: req.LastName,
		TeamID: req.TeamId,
		
	}

	postgres.DB.Save(&newUser)

	c.JSON(http.StatusAccepted, resource.UserShowResource(newUser))
}

func companyQuery(authId interface{}) *gorm.DB {
	query := postgres.DB.Model(model.User{}).Where("company_id = ?", authId)
	return query
}

func generateRegisterNumber(query *gorm.DB) (newRegisterNumber uint) {

	var u model.User
	err := query.Order("register_number DESC").First(&u)


	if err == nil {
		newRegisterNumber = u.RegisterNumber + 1
	} else {
		pt := ptime.Now()

		var count int64
		var registerNumber string

		count = 1
		for count != 0 {
			registerNumber = fmt.Sprintf("%s%s", pt.Format("yyMM"), "1111")
			print()
			query.Where("register_number = ?", registerNumber).Count(&count)
		}

		number, _ := strconv.ParseUint(registerNumber, 10, 0)
		newRegisterNumber = uint(number)
	}

	return
}
