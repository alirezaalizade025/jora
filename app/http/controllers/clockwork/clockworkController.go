package clockwork

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"

	"jora/app/models/attendance"
	db "jora/database/postgres"
)

func ClockIn(c *gin.Context) {

	// get user id from context
	id, exists := c.Get("userId")
	if !exists {
		id = 0
	}

	now := time.Now()

	// save clock in time
	attendanceModel := &attendance.Attendance{
		UserID:  uint(id.(float64)),
		ClockIn: &now,
		Type:    "working",
	}
	attendanceModel.SetTypeName()

	db.DB.Model(&attendance.Attendance{}).Create(attendanceModel)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    attendanceModel,
	})
}

func ClockOut(c *gin.Context) {

	// get user id from context
	userId, exists := c.Get("userId")
	if !exists {
		userId = 0
	}

	now := time.Now()

	// save clock in time
	attendanceModel := &attendance.Attendance{
		UserID:  uint(userId.(float64)),
		ClockOut: &now,
		Type:    "working",
	}
	attendanceModel.SetTypeName()


	// update (if has clock in today) or create
	query := db.DB.Model(&attendanceModel)
	query.Clauses(clause.Returning{})
	query.Where("user_id = ?", userId)
	query.Where("DATE(clock_in) = ?", time.Now().Format("2006-01-02"))
	query.Where("clock_out IS NULL")
	query.Order("clock_in desc")
	row := query.Updates(attendanceModel)

	if row.RowsAffected == 0 {
		db.DB.Model(&attendanceModel).Save(attendanceModel)
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    attendanceModel,
	})
}
