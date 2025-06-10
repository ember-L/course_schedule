package repository

import (
	"fmt"
	"log"

	"course_schedule/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Database 数据库连接
type Database struct {
	DB *gorm.DB
}

// NewDatabase 创建数据库连接
func NewDatabase(username, password, host, port, dbname, charset string) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, dbname, charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return nil, err
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	
	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(100)

	return &Database{DB: db}, nil
}

// AutoMigrate 自动迁移数据库表结构
func (d *Database) AutoMigrate() error {
	// 迁移基本模型
	err := d.DB.AutoMigrate(
		&model.Teacher{},
		&model.Course{},
		&model.Classroom{},
		&model.TimeSlot{},
		&model.Section{},
		&model.Student{},
		&model.Enrollment{},
		&model.AdminAuth{},
	)
	if err != nil {
		return err
	}

	// 检查是否需要创建默认管理员账号
	var count int64
	d.DB.Model(&model.AdminAuth{}).Count(&count)
	if count == 0 {
		// 创建默认管理员账号
		hashedPassword, err := HashPassword("admin123")
		if err != nil {
			return err
		}

		admin := model.AdminAuth{
			Username: "admin",
			Password: hashedPassword,
			Name:     "系统管理员",
		}

		if err := d.DB.Create(&admin).Error; err != nil {
			return err
		}

		log.Println("Created default admin account: username=admin, password=admin123")
	}

	// 创建默认时间段
	var timeSlotCount int64
	d.DB.Model(&model.TimeSlot{}).Count(&timeSlotCount)
	if timeSlotCount == 0 {
		// 创建默认时间段
		timeSlots := []model.TimeSlot{
			{DayOfWeek: 1, StartTime: "08:00", EndTime: "09:40"}, // 周一第1-2节
			{DayOfWeek: 1, StartTime: "10:00", EndTime: "11:40"}, // 周一第3-4节
			{DayOfWeek: 1, StartTime: "14:00", EndTime: "15:40"}, // 周一第5-6节
			{DayOfWeek: 1, StartTime: "16:00", EndTime: "17:40"}, // 周一第7-8节
			{DayOfWeek: 1, StartTime: "19:00", EndTime: "20:40"}, // 周一第9-10节
			{DayOfWeek: 2, StartTime: "08:00", EndTime: "09:40"}, // 周二第1-2节
			{DayOfWeek: 2, StartTime: "10:00", EndTime: "11:40"}, // 周二第3-4节
			{DayOfWeek: 2, StartTime: "14:00", EndTime: "15:40"}, // 周二第5-6节
			{DayOfWeek: 2, StartTime: "16:00", EndTime: "17:40"}, // 周二第7-8节
			{DayOfWeek: 2, StartTime: "19:00", EndTime: "20:40"}, // 周二第9-10节
			{DayOfWeek: 3, StartTime: "08:00", EndTime: "09:40"}, // 周三第1-2节
			{DayOfWeek: 3, StartTime: "10:00", EndTime: "11:40"}, // 周三第3-4节
			{DayOfWeek: 3, StartTime: "14:00", EndTime: "15:40"}, // 周三第5-6节
			{DayOfWeek: 3, StartTime: "16:00", EndTime: "17:40"}, // 周三第7-8节
			{DayOfWeek: 3, StartTime: "19:00", EndTime: "20:40"}, // 周三第9-10节
			{DayOfWeek: 4, StartTime: "08:00", EndTime: "09:40"}, // 周四第1-2节
			{DayOfWeek: 4, StartTime: "10:00", EndTime: "11:40"}, // 周四第3-4节
			{DayOfWeek: 4, StartTime: "14:00", EndTime: "15:40"}, // 周四第5-6节
			{DayOfWeek: 4, StartTime: "16:00", EndTime: "17:40"}, // 周四第7-8节
			{DayOfWeek: 4, StartTime: "19:00", EndTime: "20:40"}, // 周四第9-10节
			{DayOfWeek: 5, StartTime: "08:00", EndTime: "09:40"}, // 周五第1-2节
			{DayOfWeek: 5, StartTime: "10:00", EndTime: "11:40"}, // 周五第3-4节
			{DayOfWeek: 5, StartTime: "14:00", EndTime: "15:40"}, // 周五第5-6节
			{DayOfWeek: 5, StartTime: "16:00", EndTime: "17:40"}, // 周五第7-8节
			{DayOfWeek: 5, StartTime: "19:00", EndTime: "20:40"}, // 周五第9-10节
		}

		for _, ts := range timeSlots {
			if err := d.DB.Create(&ts).Error; err != nil {
				return err
			}
		}

		log.Println("Created default time slots")
	}

	return nil
}