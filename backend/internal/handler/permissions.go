package handler

// APIPermissions 定义API权限配置
type APIPermissions struct {
	Public       []string // 公开API，无需认证
	Admin        []string // 仅管理员可访问
	Teacher      []string // 仅教师可访问
	Student      []string // 仅学生可访问
	AdminTeacher []string // 管理员和教师可访问
	AllAuth      []string // 所有认证用户可访问
}

// GetAPIPermissions 获取API权限配置
func GetAPIPermissions() *APIPermissions {
	return &APIPermissions{
		Public: []string{
			"GET /courses",
			"GET /courses/:id",
			"GET /classrooms",
			"GET /classrooms/:id",
			"GET /time-slots",
			"GET /time-slots/:id",
			"POST /admin/login",
			"POST /teacher/login",
			"POST /student/login",
			"POST /logout",
		},
		Admin: []string{
			"POST /courses",
			"PUT /courses/:id",
			"DELETE /courses/:id",
			"GET /teachers",
			"GET /teachers/:id",
			"POST /teachers",
			"PUT /teachers/:id",
			"DELETE /teachers/:id",
			"POST /classrooms",
			"PUT /classrooms/:id",
			"DELETE /classrooms/:id",
			"POST /time-slots",
			"PUT /time-slots/:id",
			"DELETE /time-slots/:id",
			"GET /sections",
			"GET /sections/:id",
			"POST /sections",
			"PUT /sections/:id",
			"DELETE /sections/:id",
			"GET /students",
			"GET /students/:id",
			"POST /students",
			"PUT /students/:id",
			"DELETE /students/:id",
			"GET /enrollments",
			"GET /enrollments/:id",
			"POST /enrollments",
			"PUT /enrollments/:id",
			"DELETE /enrollments/:id",
			"GET /admin/profile/:id",
		},
		Teacher: []string{
			"GET /teachers/profile",
			"GET /courses/teacher/:teacherId",
			"GET /sections/teacher/:teacherId",
		},
		Student: []string{
			"GET /students/profile",
			"GET /enrollments/student/:studentId",
			"POST /enrollments/enroll",
			"DELETE /enrollments/drop",
		},
		AdminTeacher: []string{
			"GET /sections",
			"GET /sections/:id",
			"POST /sections",
			"PUT /sections/:id",
			"DELETE /sections/:id",
		},
		AllAuth: []string{
			"POST /change-password",
			"GET /schedule",
			"GET /schedule/today",
		},
	}
}

// CheckPermission 检查用户是否有权限访问指定API
func CheckPermission(userType, method, path string) bool {
	permissions := GetAPIPermissions()
	apiKey := method + " " + path

	// 检查公开API
	for _, api := range permissions.Public {
		if api == apiKey {
			return true
		}
	}

	// 检查管理员权限
	if userType == "admin" {
		for _, api := range permissions.Admin {
			if api == apiKey {
				return true
			}
		}
		for _, api := range permissions.AdminTeacher {
			if api == apiKey {
				return true
			}
		}
		for _, api := range permissions.AllAuth {
			if api == apiKey {
				return true
			}
		}
	}

	// 检查教师权限
	if userType == "teacher" {
		for _, api := range permissions.Teacher {
			if api == apiKey {
				return true
			}
		}
		for _, api := range permissions.AdminTeacher {
			if api == apiKey {
				return true
			}
		}
		for _, api := range permissions.AllAuth {
			if api == apiKey {
				return true
			}
		}
	}

	// 检查学生权限
	if userType == "student" {
		for _, api := range permissions.Student {
			if api == apiKey {
				return true
			}
		}
		for _, api := range permissions.AllAuth {
			if api == apiKey {
				return true
			}
		}
	}

	return false
}
