package main

import (
	"fmt"
	"log"

	"course_schedule/internal/handler"
	"course_schedule/internal/repository"
	"course_schedule/internal/service"
	"course_schedule/pkg/config"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// 自定义验证器
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// 加载配置
	configPath, err := config.FindConfigFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to find config file: %v", err)
	}

	cfg, err := config.LoadConfigWithEnv(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库
	db, err := repository.NewDatabase(
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.Charset,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移数据库表结构
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 初始化仓库
	courseRepo := repository.NewCourseRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)
	classroomRepo := repository.NewClassroomRepository(db)
	timeSlotRepo := repository.NewTimeSlotRepository(db)
	sectionRepo := repository.NewSectionRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)
	authRepo := repository.NewAuthRepository(db)

	// 初始化服务
	courseService := service.NewCourseService(courseRepo)
	teacherService := service.NewTeacherService(teacherRepo)
	classroomService := service.NewClassroomService(classroomRepo)
	timeSlotService := service.NewTimeSlotService(timeSlotRepo)
	sectionService := service.NewSectionService(sectionRepo)
	studentService := service.NewStudentService(studentRepo)
	enrollmentService := service.NewEnrollmentService(enrollmentRepo, sectionRepo)
	authService := service.NewAuthService(authRepo, cfg.Auth.JWTSecret, cfg.Auth.JWTDuration)
	scheduleService := service.NewScheduleService(sectionService)

	// 初始化处理器
	courseHandler := handler.NewCourseHandler(courseService, authService)
	teacherHandler := handler.NewTeacherHandler(teacherService, authService)
	classroomHandler := handler.NewClassroomHandler(classroomService, authService)
	timeSlotHandler := handler.NewTimeSlotHandler(timeSlotService)
	sectionHandler := handler.NewSectionHandler(sectionService)
	studentHandler := handler.NewStudentHandler(studentService)
	enrollmentHandler := handler.NewEnrollmentHandler(enrollmentService)
	authHandler := handler.NewAuthHandler(authService)
	scheduleHandler := handler.NewScheduleHandler(scheduleService, authService)

	// 创建Echo实例
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// 设置中间件
	handler.SetupMiddlewares(e)

	// 注册路由
	courseHandler.Register(e)
	teacherHandler.Register(e)
	classroomHandler.Register(e)
	timeSlotHandler.Register(e)
	sectionHandler.Register(e)
	studentHandler.Register(e)
	enrollmentHandler.Register(e)
	authHandler.Register(e)
	scheduleHandler.Register(e)

	// 启动服务器
	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", serverAddr)
	if err := e.Start(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
