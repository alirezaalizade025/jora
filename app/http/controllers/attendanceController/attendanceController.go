package attendanceController

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	request "jora/app/http/requests"
	"jora/app/models/attendance"
	db "jora/database/postgres"
	"jora/utility"

	"gitee.com/golang-module/carbon/v2"
)

func Start(c *gin.Context) {

	// get user id from context
	id, exists := c.Get("userId")
	if !exists {
		id = 0
	}

	now := time.Now()

	// save clock in time
	attendanceModel := &attendance.Attendance{
		UserID:  uint(id.(float64)),
		StartAt: &now,
		Type:    "working",
	}
	attendanceModel.SetType()

	db.DB.Model(&attendance.Attendance{}).Create(attendanceModel)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    attendanceModel,
	})
}

func Finish(c *gin.Context) {

	// get user id from context
	userId, exists := c.Get("userId")
	if !exists {
		userId = 0
	}

	now := time.Now()

	// save clock in time
	attendanceModel := &attendance.Attendance{
		UserID:   uint(userId.(float64)),
		FinishAt: &now,
		Type:     "working",
	}
	attendanceModel.SetType()

	// update (if has clock in today) or create
	query := db.DB.Model(&attendanceModel)
	query.Clauses(clause.Returning{})
	query.Where("user_id = ?", userId)
	query.Where("DATE(start_at) = ?", time.Now().Format("2006-01-02"))
	query.Where("finish_at IS NULL")
	query.Order("start_at desc")
	row := query.Updates(attendanceModel)

	if row.RowsAffected == 0 {
		db.DB.Model(&attendanceModel).Save(attendanceModel)
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    attendanceModel,
	})
}

func Leave(c *gin.Context) {

	// get user id from context
	userId, exists := c.Get("userId")
	if !exists {
		userId = 0
	}

	req := &request.LeaveRequest{}

	if !request.Validation(c, req) {
		return
	}

	// convert clock string to time.Time
	StartAt := carbon.Parse(req.StartAt).ToStdTime()
	FinishAt := carbon.Parse(req.FinishAt).ToStdTime()

	// validate clock in and clock out
	if FinishAt.Before(StartAt) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "finish_at must be grater then start_at",
		})
		return
	}

	// save clock in time
	attendanceModel := &attendance.Attendance{
		UserID:   uint(userId.(float64)),
		FinishAt: &FinishAt,
		StartAt:  &StartAt,
		Type:     "hourly_leave",
		Note:     &req.Note,
	}
	attendanceModel.SetType()

	var count int64
	// check for duplicate request
	query := db.DB.Model(&attendanceModel)
	query.Where("user_id = ?", userId)
	query.Where("start_at = ?", StartAt)
	query.Where("finish_at = ?", FinishAt)
	query.Where("type = ?", attendanceModel.TypeInt)
	query.Count(&count)

	if count > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "duplicate request",
		})
		return
	}

	db.DB.Model(&attendanceModel).Save(attendanceModel)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    attendanceModel,
	})
}

func HourlyLeave(c *gin.Context) {

	// get user id from context
	userId, exists := c.Get("userId")
	if !exists {
		userId = 0
	}

	req := &request.HourlyLeaveRequest{}

	if !request.Validation(c, req) {
		return
	}

	// convert clock string to time.Time
	StartAt := carbon.Parse(req.StartAt).EndOfDay().ToStdTime()
	FinishAt := carbon.Parse(req.FinishAt).StartOfDay().ToStdTime()

	// validate clock in and clock out
	if FinishAt.Before(StartAt) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "finish_at must be grater then start_at",
		})
		return
	}

	// save clock in time
	attendanceModel := &attendance.Attendance{
		UserID:   uint(userId.(float64)),
		FinishAt: &FinishAt,
		StartAt:  &StartAt,
		Note:     &req.Note,
	}
	attendanceModel.SetType()

	var count int64
	// check for duplicate request
	query := db.DB.Model(&attendanceModel)
	query.Where("user_id = ?", userId)
	query.Where("start_at = ?", StartAt)
	query.Where("finish_at = ?", FinishAt)
	query.Where("type = ?", attendanceModel.TypeInt)
	query.Count(&count)

	if count > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "duplicate request",
		})
		return
	}

	db.DB.Model(&attendanceModel).Save(attendanceModel)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    attendanceModel,
	})
}

func BusinessTrip(c *gin.Context) {

	// get user id from context
	userId, exists := c.Get("userId")
	if !exists {
		userId = 0
	}

	req := &request.BusinessTripRequest{}

	if !request.Validation(c, req) {
		return
	}

	// convert clock string to time.Time
	StartAt := carbon.Parse(req.StartAt).ToStdTime()
	FinishAt := carbon.Parse(req.FinishAt).ToStdTime()

	// validate clock in and clock out
	if FinishAt.Before(StartAt) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "finish_at must be grater then start_at",
		})
		return
	}

	// save clock in time
	attendanceModel := &attendance.Attendance{
		UserID:   uint(userId.(float64)),
		FinishAt: &FinishAt,
		StartAt:  &StartAt,
		Note:     &req.Note,
		Type:     "business_trip",
	}
	attendanceModel.SetType()

	var count int64
	// check for duplicate request
	query := db.DB.Model(&attendanceModel)
	query.Where("user_id = ?", userId)
	query.Where("DATE(start_at) = ?", StartAt)
	query.Where("DATE(finish_at) = ?", FinishAt)
	query.Where("type = ?", attendanceModel.TypeInt)
	query.Count(&count)

	if count > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "duplicate request",
		})
		return
	}

	db.DB.Model(&attendanceModel).Save(attendanceModel)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    attendanceModel,
	})
}

func RemoteWork(c *gin.Context) {

	// get user id from context
	userId, exists := c.Get("userId")
	if !exists {
		userId = 0
	}

	req := &request.RemoteWorkRequest{}

	if !request.Validation(c, req) {
		return
	}

	// convert clock string to time.Time
	StartAt := carbon.Parse(req.StartAt).ToStdTime()
	FinishAt := carbon.Parse(req.FinishAt).ToStdTime()

	// validate clock in and clock out
	if FinishAt.Before(StartAt) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "finish_at must be grater then start_at",
		})
		return
	}

	// save clock in time
	attendanceModel := &attendance.Attendance{
		UserID:   uint(userId.(float64)),
		FinishAt: &FinishAt,
		StartAt:  &StartAt,
		Note:     &req.Note,
		Type:     "remote_work",
	}
	attendanceModel.SetType()

	var count int64
	// check for duplicate request
	query := db.DB.Model(&attendanceModel)
	query.Where("user_id = ?", userId)
	query.Where("DATE(start_at) = ?", StartAt)
	query.Where("DATE(finish_at) = ?", FinishAt)
	query.Where("type = ?", attendanceModel.TypeInt)
	query.Count(&count)

	if count > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "duplicate request",
		})
		return
	}

	db.DB.Model(&attendanceModel).Save(attendanceModel)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    attendanceModel,
	})
}

func Update(c *gin.Context) {

	// get user id from context
	userId, exists := c.Get("userId")
	if !exists {
		userId = 0
	}

	// read id from path
	attendanceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid id",
		})
		return
	}
	// validation
	req := &request.AttendanceUpdateRequest{}
	if !request.Validation(c, req) {
		return
	}

	// find or fail attendance
	attendanceModel := &attendance.Attendance{}
	result := db.DB.Where("user_id", userId).Find(&attendanceModel, attendanceId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	// convert clock string to time.Time
	if req.StartAt != "" {
		var StartAt time.Time
		if utility.InArray(attendanceModel.Type, []string{"sick_leave", "annual_leave", "vacation_leave"}) {
			StartAt = carbon.Parse(req.StartAt).StartOfDay().ToStdTime()
		} else {
			StartAt = carbon.Parse(req.StartAt).ToStdTime()
		}

		attendanceModel.StartAt = &StartAt
	} else {
		attendanceModel.StartAt = nil
	}

	if req.FinishAt != "" {
		var FinishAt time.Time
		if utility.InArray(attendanceModel.Type, []string{"sick_leave", "annual_leave", "vacation_leave"}) {
			FinishAt = carbon.Parse(req.FinishAt).EndOfDay().ToStdTime()
		} else {
			FinishAt = carbon.Parse(req.FinishAt).ToStdTime()
		}
		
		attendanceModel.FinishAt = &FinishAt
	} else {
		attendanceModel.FinishAt = nil
	}

	attendanceModel.Note = req.Note
	attendanceModel.SetType()

	db.DB.Model(&attendanceModel).Save(attendanceModel)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    attendanceModel,
	})

}
