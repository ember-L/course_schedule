-- 创建数据库
CREATE DATABASE IF NOT EXISTS course_schedule CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE course_schedule;

-- 创建教师表
CREATE TABLE IF NOT EXISTS teacher (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  title VARCHAR(50),
  email VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(100) NOT NULL DEFAULT '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
  created_at TIMESTAMP NULL,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL,
  INDEX idx_teacher_name (name),
  INDEX idx_teacher_deleted_at (deleted_at)
);

-- 创建课程表
CREATE TABLE IF NOT EXISTS course (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  code VARCHAR(20) NOT NULL UNIQUE,
  name VARCHAR(100) NOT NULL,
  credits FLOAT NOT NULL,
  description TEXT,
  created_at TIMESTAMP NULL,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL,
  INDEX idx_course_name (name),
  INDEX idx_course_deleted_at (deleted_at)
);

-- 创建教师课程关联表
CREATE TABLE IF NOT EXISTS teacher_courses (
  teacher_id INT UNSIGNED NOT NULL,
  course_id INT UNSIGNED NOT NULL,
  PRIMARY KEY (teacher_id, course_id),
  CONSTRAINT fk_teacher_courses_teacher FOREIGN KEY (teacher_id) REFERENCES teacher (id) ON DELETE CASCADE,
  CONSTRAINT fk_teacher_courses_course FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE
);

-- 创建教室表
CREATE TABLE IF NOT EXISTS classroom (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  building VARCHAR(50) NOT NULL,
  room VARCHAR(20) NOT NULL,
  capacity INT NOT NULL,
  type VARCHAR(50),
  created_at TIMESTAMP NULL,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL,
  INDEX idx_classroom_building (building),
  INDEX idx_classroom_type (type),
  INDEX idx_classroom_deleted_at (deleted_at)
);

-- 创建时间段表
CREATE TABLE IF NOT EXISTS time_slot (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  day_of_week INT NOT NULL,
  start_time VARCHAR(5) NOT NULL,
  end_time VARCHAR(5) NOT NULL,
  created_at TIMESTAMP NULL,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL,
  INDEX idx_time_slot_day_of_week (day_of_week),
  INDEX idx_time_slot_deleted_at (deleted_at)
);

-- 创建学生表
CREATE TABLE IF NOT EXISTS student (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  student_id VARCHAR(20) NOT NULL UNIQUE,
  name VARCHAR(50) NOT NULL,
  grade VARCHAR(10) NOT NULL,
  major VARCHAR(50) NOT NULL,
  password VARCHAR(100) NOT NULL DEFAULT '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
  created_at TIMESTAMP NULL,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL,
  INDEX idx_student_name (name),
  INDEX idx_student_grade (grade),
  INDEX idx_student_major (major),
  INDEX idx_student_deleted_at (deleted_at)
);

-- 创建课程安排表
CREATE TABLE IF NOT EXISTS section (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  course_id INT UNSIGNED NOT NULL,
  teacher_id INT UNSIGNED NOT NULL,
  classroom_id INT UNSIGNED NOT NULL,
  time_slot_id INT UNSIGNED NOT NULL,
  semester VARCHAR(20) NOT NULL,
  max_students INT NOT NULL,
  created_at TIMESTAMP NULL,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL,
  INDEX idx_section_course_id (course_id),
  INDEX idx_section_teacher_id (teacher_id),
  INDEX idx_section_classroom_id (classroom_id),
  INDEX idx_section_time_slot_id (time_slot_id),
  INDEX idx_section_semester (semester),
  INDEX idx_section_deleted_at (deleted_at),
  CONSTRAINT fk_section_course FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE,
  CONSTRAINT fk_section_teacher FOREIGN KEY (teacher_id) REFERENCES teacher (id) ON DELETE CASCADE,
  CONSTRAINT fk_section_classroom FOREIGN KEY (classroom_id) REFERENCES classroom (id) ON DELETE CASCADE,
  CONSTRAINT fk_section_time_slot FOREIGN KEY (time_slot_id) REFERENCES time_slot (id) ON DELETE CASCADE
);

-- 创建选课记录表
CREATE TABLE IF NOT EXISTS enrollment (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  student_id INT UNSIGNED NOT NULL,
  section_id INT UNSIGNED NOT NULL,
  enroll_time TIMESTAMP NOT NULL,
  grade VARCHAR(2),
  created_at TIMESTAMP NULL,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL,
  INDEX idx_enrollment_student_id (student_id),
  INDEX idx_enrollment_section_id (section_id),
  INDEX idx_enrollment_deleted_at (deleted_at),
  CONSTRAINT fk_enrollment_student FOREIGN KEY (student_id) REFERENCES student (id) ON DELETE CASCADE,
  CONSTRAINT fk_enrollment_section FOREIGN KEY (section_id) REFERENCES section (id) ON DELETE CASCADE
);

-- 创建管理员表
CREATE TABLE IF NOT EXISTS admin_auth (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(100) NOT NULL,
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP NULL,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL,
  INDEX idx_admin_auth_deleted_at (deleted_at)
);

-- 插入默认管理员账号
INSERT INTO admin_auth (username, password, name, created_at, updated_at)
VALUES ('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '系统管理员', NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();
-- 默认密码为 admin123

-- 插入默认时间段
INSERT INTO time_slot (day_of_week, start_time, end_time, created_at, updated_at) VALUES
(1, '08:00', '09:40', NOW(), NOW()), -- 周一第1-2节
(1, '10:00', '11:40', NOW(), NOW()), -- 周一第3-4节
(1, '14:00', '15:40', NOW(), NOW()), -- 周一第5-6节
(1, '16:00', '17:40', NOW(), NOW()), -- 周一第7-8节
(1, '19:00', '20:40', NOW(), NOW()), -- 周一第9-10节
(2, '08:00', '09:40', NOW(), NOW()), -- 周二第1-2节
(2, '10:00', '11:40', NOW(), NOW()), -- 周二第3-4节
(2, '14:00', '15:40', NOW(), NOW()), -- 周二第5-6节
(2, '16:00', '17:40', NOW(), NOW()), -- 周二第7-8节
(2, '19:00', '20:40', NOW(), NOW()), -- 周二第9-10节
(3, '08:00', '09:40', NOW(), NOW()), -- 周三第1-2节
(3, '10:00', '11:40', NOW(), NOW()), -- 周三第3-4节
(3, '14:00', '15:40', NOW(), NOW()), -- 周三第5-6节
(3, '16:00', '17:40', NOW(), NOW()), -- 周三第7-8节
(3, '19:00', '20:40', NOW(), NOW()), -- 周三第9-10节
(4, '08:00', '09:40', NOW(), NOW()), -- 周四第1-2节
(4, '10:00', '11:40', NOW(), NOW()), -- 周四第3-4节
(4, '14:00', '15:40', NOW(), NOW()), -- 周四第5-6节
(4, '16:00', '17:40', NOW(), NOW()), -- 周四第7-8节
(4, '19:00', '20:40', NOW(), NOW()), -- 周四第9-10节
(5, '08:00', '09:40', NOW(), NOW()), -- 周五第1-2节
(5, '10:00', '11:40', NOW(), NOW()), -- 周五第3-4节
(5, '14:00', '15:40', NOW(), NOW()), -- 周五第5-6节
(5, '16:00', '17:40', NOW(), NOW()), -- 周五第7-8节
(5, '19:00', '20:40', NOW(), NOW())  -- 周五第9-10节
ON DUPLICATE KEY UPDATE updated_at = NOW();