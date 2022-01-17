package controllers

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	enums "job-com-backend/enums"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Oportunity struct {
	IdOportunity            int `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Title                   string
	TypeOportunity          int
	Quantity                int
	ExpertiseLevel          int
	IsRemote                bool
	SalaryMin               float32
	SalaryMax               float32
	Description             string
	RequerimentsDescription string
	BenefitsDescription     string
	ContactPhone            string
	ContactEmail            string
	Address                 string
	ExpirationDate          time.Time
	RegisterDate            time.Time
	HourMin                 string
	HourMax                 string
	City                    string
	IdUserFk                int
	User                    UserWithoutPassword `gorm:"ForeignKey:IdUserFk"`
}

func GetOportunities(c *gin.Context) {
	var oportunity []Oportunity
	db.Preload("User").Find(&oportunity)
	c.JSON(200, gin.H{
		"data": oportunity,
	})
}

func FindOportunitiesByUser(c *gin.Context) {
	var oportunity []Oportunity
	userId := c.Param("userId")
	db.Where("id_user_fk = @IdUserFk", sql.Named("IdUserFk", userId)).Preload("User").Find(&oportunity)
	c.JSON(200, gin.H{
		"data": oportunity,
	})
}

func filterOportunityBySalaryRange(query *gorm.DB, salaryRangeType int) *gorm.DB {
	minSalary := 0
	maxSalary := 0
	INFINITY_CONSTANT := "-1"
	switch salaryRangeType {
	case enums.TYPE_FIRST_SALARY_RANGE:
		minSalary, maxSalary = int(enums.FIRST_SALARY_RANGE[0]), int(enums.FIRST_SALARY_RANGE[1])
	case enums.TYPE_SECOND_SALARY_RANGE:
		minSalary, maxSalary = int(enums.SECOND_SALARY_RANGE[0]), int(enums.SECOND_SALARY_RANGE[1])
	case enums.TYPE_THIRD_SALARY_RANGE:
		minSalary, maxSalary = int(enums.THIRD_SALARY_RANGE[0]), int(enums.THIRD_SALARY_RANGE[1])
	case enums.TYPE_FOURTH_SALARY_RANGE:
		minSalary, maxSalary = int(enums.FOURTH_SALARY_RANGE[0]), int(enums.FOURTH_SALARY_RANGE[1])
	case enums.TYPE_LAST_SALARY_RANGE:
		minSalary, maxSalary = int(enums.LAST_SALARY_RANGE[0]), int(enums.LAST_SALARY_RANGE[1])
	}
	return query.Where(
		"(salary_min > @MinSalary OR salary_max > @MinSalary) AND "+
			"(salary_min < @MaxSalary OR salary_max < @MaxSalary OR @MaxSalary = "+INFINITY_CONSTANT+")",
		sql.Named("MinSalary", minSalary), sql.Named("MaxSalary", maxSalary))
}

func filterOportunityByRegisterRange(query *gorm.DB, registerRange int) *gorm.DB {
	now := time.Now()
	subtractDay := time.Hour * -24
	then := time.Now()
	switch registerRange {
	case enums.TYPE_DAY:
		then = now.Add(enums.DAY * subtractDay)
	case enums.TYPE_WEEK:
		then = now.Add(enums.WEEK * subtractDay)
	case enums.TYPE_MONTH:
		then = now.Add(enums.MONTH * subtractDay)
	case enums.TYPE_THREE_MONTHS:
		then = now.Add(enums.THREE_MONTHS * subtractDay)
	}
	return query.Where("register_date > @DateRange", sql.Named("DateRange", then))
}

func GetOportunitiesPaginated(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	textSearch := c.Query("text_search")
	levelFilter := c.Query("level_filter")
	isRemote, errorRemote := strconv.ParseBool(c.Query("is_remote"))
	vacancyType, errorVacancyType := strconv.Atoi(c.Query("vacancy_type"))
	salaryRangeType, errorSalaryRangeType := strconv.Atoi(c.Query("salary_range_type"))
	registerRange, errorRegisterRange := strconv.Atoi(c.Query("register_range"))

	var oportunity []Oportunity
	var query *gorm.DB = db

	if errorRemote == nil && isRemote {
		query = query.Where("is_remote")
	}
	if len(textSearch) > 0 {
		textSearch = "%" + textSearch + "%"
		query = query.Where("description like @TextSearch OR title like @TextSearch", sql.Named("TextSearch", textSearch))
	}
	if errorVacancyType == nil {
		query = query.Where("type_oportunity = @VacancyType", sql.Named("VacancyType", vacancyType))
	}
	if errorSalaryRangeType == nil {
		query = filterOportunityBySalaryRange(query, salaryRangeType)
	}
	if errorRegisterRange == nil {
		query = filterOportunityByRegisterRange(query, registerRange)
	}
	if len(levelFilter) > 0 {
		query = query.Where("expertise_level IN (?)", strings.Split(levelFilter, ","))
	}
	var pagination Pagination = Pagination{Limit: pageSize, Page: page}
	query.Scopes(paginate(oportunity, &pagination, query)).Preload("User").Find(&oportunity)
	pagination.Rows = oportunity
	c.JSON(200, gin.H{
		"data": pagination,
	})
}

func GetOportunity(c *gin.Context) {
	id := c.Param("id")
	var oportunity Oportunity
	db.Where("id_oportunity = @IdOportunity", sql.Named("IdOportunity", id)).Preload("User").Find(&oportunity)
	c.JSON(200, gin.H{
		"data": oportunity,
	})
}

func PostOportunity(c *gin.Context) {
	var oportunity Oportunity
	c.BindJSON(&oportunity)
	oportunity.RegisterDate = time.Now()
	result := db.Create(&oportunity)
	if result.Error == nil {
		c.JSON(200, gin.H{
			"data": oportunity,
		})
	}
}

func PutOportunity(c *gin.Context) {
	var oportunity Oportunity
	c.BindJSON(&oportunity)
	db.Save(&oportunity)
	c.JSON(200, gin.H{
		"data": oportunity,
	})
}

func DeleteOportunity(c *gin.Context) {
	id := c.Param("id")
	var oportunity Oportunity
	db.Where("id_oportunity = @IdOportunity", sql.Named("IdOportunity", id)).Delete(&oportunity)
	c.JSON(200, gin.H{
		"data": oportunity,
	})
}
