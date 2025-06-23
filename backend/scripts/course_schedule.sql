/*
 Navicat Premium Data Transfer

 Source Server         : better
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : 127.0.0.1:3306
 Source Schema         : course_schedule

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 24/06/2025 00:10:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_auth
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth`;
CREATE TABLE `admin_auth`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_admin_auth_username`(`username`) USING BTREE,
  INDEX `idx_admin_auth_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_auth
-- ----------------------------
INSERT INTO `admin_auth` VALUES (1, '2025-06-11 00:34:03.153', '2025-06-11 00:34:03.153', NULL, 'admin', '$2a$10$PHgi3Chxn0wCITDFBXOHWO9ipIjRvkS1cewy8dTeel.ahCJQJ4HGG', '系统管理员');

-- ----------------------------
-- Table structure for announcement
-- ----------------------------
DROP TABLE IF EXISTS `announcement`;
CREATE TABLE `announcement`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of announcement
-- ----------------------------
INSERT INTO `announcement` VALUES (1, '端午节放假公告', '端午节放假5天', '2025-06-24 00:10:04.177', '2025-06-24 00:10:04.177');

-- ----------------------------
-- Table structure for classroom
-- ----------------------------
DROP TABLE IF EXISTS `classroom`;
CREATE TABLE `classroom`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `building` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `room` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `capacity` bigint NOT NULL,
  `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_classroom_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_classroom_building`(`building`) USING BTREE,
  INDEX `idx_classroom_type`(`type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of classroom
-- ----------------------------
INSERT INTO `classroom` VALUES (1, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '教学楼A', '101', 60, '普通教室');
INSERT INTO `classroom` VALUES (2, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '教学楼A', '102', 60, '普通教室');
INSERT INTO `classroom` VALUES (3, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '教学楼A', '201', 80, '多媒体教室');
INSERT INTO `classroom` VALUES (4, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '教学楼A', '202', 80, '多媒体教室');
INSERT INTO `classroom` VALUES (5, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '教学楼B', '101', 100, '阶梯教室');
INSERT INTO `classroom` VALUES (6, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '教学楼B', '102', 100, '阶梯教室');
INSERT INTO `classroom` VALUES (7, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '教学楼B', '201', 40, '讨论室');
INSERT INTO `classroom` VALUES (8, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '实验楼C', '101', 50, '计算机实验室');
INSERT INTO `classroom` VALUES (9, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '实验楼C', '102', 50, '物理实验室');
INSERT INTO `classroom` VALUES (10, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '实验楼C', '201', 30, '语音室');

-- ----------------------------
-- Table structure for course
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `code` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `credits` float NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_course_code`(`code`) USING BTREE,
  INDEX `idx_course_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_course_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of course
-- ----------------------------
INSERT INTO `course` VALUES (1, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 'MATH101', '高等数学I', 4, '微积分基础、极限、导数、积分等内容');
INSERT INTO `course` VALUES (2, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 'MATH102', '高等数学II', 4, '多元微积分、级数、微分方程等内容');
INSERT INTO `course` VALUES (3, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 'CS101', '计算机导论', 3, '计算机基础知识、发展历史、基本原理');
INSERT INTO `course` VALUES (4, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 'CS201', '数据结构', 4, '基本数据结构、算法分析、排序和搜索');
INSERT INTO `course` VALUES (5, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 'CS301', '数据库原理', 3.5, '数据库设计、SQL语言、事务处理');
INSERT INTO `course` VALUES (6, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 'PHY101', '大学物理I', 4, '力学、热学基础知识');
INSERT INTO `course` VALUES (7, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 'PHY102', '大学物理II', 4, '电磁学、光学基础知识');
INSERT INTO `course` VALUES (8, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 'ENG101', '大学英语I', 3, '英语听说读写基础训练');
INSERT INTO `course` VALUES (9, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 'ENG102', '大学英语II', 3, '英语听说读写进阶训练');
INSERT INTO `course` VALUES (10, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 'MGT101', '管理学原理', 3, '管理学基本理论和方法');
INSERT INTO `course` VALUES (12, '2025-06-22 12:03:31.968', '2025-06-22 12:03:31.968', NULL, 'CS108', '计算机网络基础', 4, '');

-- ----------------------------
-- Table structure for enrollment
-- ----------------------------
DROP TABLE IF EXISTS `enrollment`;
CREATE TABLE `enrollment`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `student_id` bigint UNSIGNED NOT NULL,
  `section_id` bigint UNSIGNED NOT NULL,
  `enroll_time` datetime(3) NOT NULL,
  `grade` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_enrollment_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_enrollment_student_id`(`student_id`) USING BTREE,
  INDEX `idx_enrollment_section_id`(`section_id`) USING BTREE,
  CONSTRAINT `fk_section_enrollments` FOREIGN KEY (`section_id`) REFERENCES `section` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_student_enrollments` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of enrollment
-- ----------------------------
INSERT INTO `enrollment` VALUES (1, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 1, 1, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (2, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 1, 3, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (3, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', '2025-06-20 23:20:06.820', 1, 6, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (4, '2025-06-15 15:50:33.000', '2025-06-21 10:10:56.618', NULL, 2, 1, '2025-06-15 15:50:33.000', '12');
INSERT INTO `enrollment` VALUES (5, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 2, 3, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (6, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 2, 8, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (7, '2025-06-15 15:50:33.000', '2025-06-22 12:49:01.576', NULL, 3, 2, '2025-06-15 15:50:33.000', '20');
INSERT INTO `enrollment` VALUES (8, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 3, 5, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (9, '2025-06-15 15:50:33.000', '2025-06-21 17:01:44.837', NULL, 4, 4, '2025-06-15 15:50:33.000', '89');
INSERT INTO `enrollment` VALUES (10, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 4, 7, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (11, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 5, 4, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (12, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 5, 10, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (13, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 6, 5, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (14, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 6, 9, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (15, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 7, 6, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (16, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 8, 7, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (17, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 9, 8, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (18, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 10, 9, '2025-06-15 15:50:33.000', NULL);
INSERT INTO `enrollment` VALUES (19, '2025-06-22 11:54:43.044', '2025-06-22 11:54:43.044', NULL, 3, 3, '2025-06-22 11:54:43.044', '');

-- ----------------------------
-- Table structure for section
-- ----------------------------
DROP TABLE IF EXISTS `section`;
CREATE TABLE `section`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `course_id` bigint UNSIGNED NOT NULL,
  `teacher_id` bigint UNSIGNED NOT NULL,
  `classroom_id` bigint UNSIGNED NOT NULL,
  `time_slot_id` bigint UNSIGNED NOT NULL,
  `semester` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `max_students` bigint NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_section_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_section_course_id`(`course_id`) USING BTREE,
  INDEX `idx_section_teacher_id`(`teacher_id`) USING BTREE,
  INDEX `idx_section_classroom_id`(`classroom_id`) USING BTREE,
  INDEX `idx_section_time_slot_id`(`time_slot_id`) USING BTREE,
  INDEX `idx_section_semester`(`semester`) USING BTREE,
  CONSTRAINT `fk_classroom_sections` FOREIGN KEY (`classroom_id`) REFERENCES `classroom` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_course_sections` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_teacher_sections` FOREIGN KEY (`teacher_id`) REFERENCES `teacher` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_time_slot_sections` FOREIGN KEY (`time_slot_id`) REFERENCES `time_slot` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of section
-- ----------------------------
INSERT INTO `section` VALUES (1, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 1, 1, 1, 22, '2023-2024-1', 50);
INSERT INTO `section` VALUES (2, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 2, 1, 1, 21, '2023-2024-1', 50);
INSERT INTO `section` VALUES (3, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 3, 2, 3, 2, '2023-2024-1', 60);
INSERT INTO `section` VALUES (4, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 4, 3, 8, 3, '2023-2024-1', 40);
INSERT INTO `section` VALUES (5, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 5, 4, 3, 6, '2023-2024-1', 60);
INSERT INTO `section` VALUES (6, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 6, 5, 9, 9, '2023-2024-1', 45);
INSERT INTO `section` VALUES (7, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 7, 5, 9, 13, '2023-2024-1', 45);
INSERT INTO `section` VALUES (8, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 8, 6, 10, 4, '2023-2024-1', 30);
INSERT INTO `section` VALUES (9, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 9, 6, 10, 8, '2023-2024-1', 30);
INSERT INTO `section` VALUES (10, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 10, 7, 5, 11, '2023-2024-1', 80);
INSERT INTO `section` VALUES (11, '2025-06-23 22:21:58.174', '2025-06-23 22:21:58.174', NULL, 2, 3, 1, 1, '2023-2024-1', 33);
INSERT INTO `section` VALUES (12, '2025-06-23 23:40:27.811', '2025-06-23 23:40:27.811', NULL, 5, 4, 2, 1, '2023-2024-1', 23);

-- ----------------------------
-- Table structure for student
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `student_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `grade` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `major` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '$2a$10$PHgi3Chxn0wCITDFBXOHWO9ipIjRvkS1cewy8dTeel.ahCJQJ4HGG',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_student_student_id`(`student_id`) USING BTREE,
  INDEX `idx_student_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_student_name`(`name`) USING BTREE,
  INDEX `idx_student_grade`(`grade`) USING BTREE,
  INDEX `idx_student_major`(`major`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of student
-- ----------------------------
INSERT INTO `student` VALUES (1, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '2023001', '张三', '2024级', '计算机科学', '$2a$10$PHgi3Chxn0wCITDFBXOHWO9ipIjRvkS1cewy8dTeel.ahCJQJ4HGG');
INSERT INTO `student` VALUES (2, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '2023002', '李四', '2024级', '软件工程', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `student` VALUES (3, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '2023003', '王五', '2024级', '数据科学', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `student` VALUES (4, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '2022001', '赵六', '2024级', '计算机科学', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `student` VALUES (5, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '2022002', '钱七', '2024级', '软件工程', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `student` VALUES (6, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '2022003', '孙八', '2024级', '人工智能', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `student` VALUES (7, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '2021001', '周九', '2024级', '计算机科学', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `student` VALUES (8, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '2021002', '吴十', '2024级', '软件工程', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `student` VALUES (9, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '2021003', '郑十一', '2024级', '网络工程', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `student` VALUES (10, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '2020001', '王十二', '2024级', '计算机科学', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `student` VALUES (11, '2025-06-23 00:35:02.818', '2025-06-23 00:35:02.818', '2025-06-23 01:26:06.039', '20240001', 'ember', '2024级', '计算机科学与技术', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `student` VALUES (13, '2025-06-23 01:27:43.183', '2025-06-23 01:27:43.183', NULL, '20240002', 'ember1', '2021级', '软件工程', '$2a$10$PHgi3Chxn0wCITDFBXOHWO9ipIjRvkS1cewy8dTeel.ahCJQJ4HGG');

-- ----------------------------
-- Table structure for system_config
-- ----------------------------
DROP TABLE IF EXISTS `system_config`;
CREATE TABLE `system_config`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `semester` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `week` bigint NULL DEFAULT NULL,
  `updated_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `start_date` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of system_config
-- ----------------------------
INSERT INTO `system_config` VALUES (1, '2024-2025-1', 3, '', '2025-03-01');

-- ----------------------------
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '$2a$10$PHgi3Chxn0wCITDFBXOHWO9ipIjRvkS1cewy8dTeel.ahCJQJ4HGG',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_teacher_email`(`email`) USING BTREE,
  INDEX `idx_teacher_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_teacher_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of teacher
-- ----------------------------
INSERT INTO `teacher` VALUES (1, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '张教授', '教授', 'zhang@university.edu', '$2a$10$PHgi3Chxn0wCITDFBXOHWO9ipIjRvkS1cewy8dTeel.ahCJQJ4HGG');
INSERT INTO `teacher` VALUES (2, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '李副教授', '副教授', 'li@university.edu', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `teacher` VALUES (3, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '王讲师', '讲师', 'wang@university.edu', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `teacher` VALUES (4, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '刘教授', '教授', 'liu@university.edu', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `teacher` VALUES (5, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '陈副教授', '副教授', 'chen@university.edu', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `teacher` VALUES (6, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '杨讲师', '讲师', 'yang@university.edu', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `teacher` VALUES (7, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '赵教授', '教授', 'zhao@university.edu', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `teacher` VALUES (8, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '孙副教授', '副教授', 'sun@university.edu', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `teacher` VALUES (9, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '周讲师', '讲师', 'zhou@university.edu', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
INSERT INTO `teacher` VALUES (10, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, '吴教授', '教授', 'wu@university.edu', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');

-- ----------------------------
-- Table structure for teacher_courses
-- ----------------------------
DROP TABLE IF EXISTS `teacher_courses`;
CREATE TABLE `teacher_courses`  (
  `course_id` bigint UNSIGNED NOT NULL,
  `teacher_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`course_id`, `teacher_id`) USING BTREE,
  INDEX `fk_teacher_courses_teacher`(`teacher_id`) USING BTREE,
  CONSTRAINT `fk_teacher_courses_course` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_teacher_courses_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `teacher` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of teacher_courses
-- ----------------------------
INSERT INTO `teacher_courses` VALUES (1, 1);
INSERT INTO `teacher_courses` VALUES (2, 2);
INSERT INTO `teacher_courses` VALUES (2, 3);
INSERT INTO `teacher_courses` VALUES (3, 4);
INSERT INTO `teacher_courses` VALUES (1, 5);

-- ----------------------------
-- Table structure for time_slot
-- ----------------------------
DROP TABLE IF EXISTS `time_slot`;
CREATE TABLE `time_slot`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `day_of_week` bigint NOT NULL,
  `start_time` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `end_time` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_time_slot_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_time_slot_day_of_week`(`day_of_week`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of time_slot
-- ----------------------------
INSERT INTO `time_slot` VALUES (1, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 1, '08:00', '09:40');
INSERT INTO `time_slot` VALUES (2, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 1, '10:00', '11:40');
INSERT INTO `time_slot` VALUES (3, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 1, '14:00', '15:40');
INSERT INTO `time_slot` VALUES (4, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 1, '16:00', '17:40');
INSERT INTO `time_slot` VALUES (5, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 2, '08:00', '09:40');
INSERT INTO `time_slot` VALUES (6, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 2, '10:00', '11:40');
INSERT INTO `time_slot` VALUES (7, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 2, '14:00', '15:40');
INSERT INTO `time_slot` VALUES (8, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 2, '16:00', '17:40');
INSERT INTO `time_slot` VALUES (9, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 3, '08:00', '09:40');
INSERT INTO `time_slot` VALUES (10, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 3, '10:00', '11:40');
INSERT INTO `time_slot` VALUES (11, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 3, '14:00', '15:40');
INSERT INTO `time_slot` VALUES (12, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 3, '16:00', '17:40');
INSERT INTO `time_slot` VALUES (13, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 4, '08:00', '09:40');
INSERT INTO `time_slot` VALUES (14, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 4, '10:00', '11:40');
INSERT INTO `time_slot` VALUES (15, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 4, '14:00', '15:40');
INSERT INTO `time_slot` VALUES (16, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 4, '16:00', '17:40');
INSERT INTO `time_slot` VALUES (17, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 5, '08:00', '09:40');
INSERT INTO `time_slot` VALUES (18, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 5, '10:00', '11:40');
INSERT INTO `time_slot` VALUES (19, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 5, '14:00', '15:40');
INSERT INTO `time_slot` VALUES (20, '2025-06-15 15:50:33.000', '2025-06-15 15:50:33.000', NULL, 5, '16:00', '17:40');
INSERT INTO `time_slot` VALUES (21, NULL, NULL, NULL, 6, '08:00', '09:40');
INSERT INTO `time_slot` VALUES (22, NULL, NULL, NULL, 6, '10:00', '11:40');
INSERT INTO `time_slot` VALUES (23, NULL, NULL, NULL, 6, '14:00', '15:40');
INSERT INTO `time_slot` VALUES (24, NULL, NULL, NULL, 6, '16:00', '17:40');

SET FOREIGN_KEY_CHECKS = 1;
