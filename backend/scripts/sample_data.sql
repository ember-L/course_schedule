USE course_schedule;

-- 插入示例教师数据
INSERT INTO teacher (name, title, email, created_at, updated_at) VALUES
('张三', '教授', 'zhangsan@example.com', NOW(), NOW()),
('李四', '副教授', 'lisi@example.com', NOW(), NOW()),
('王五', '讲师', 'wangwu@example.com', NOW(), NOW()),
('赵六', '助教', 'zhaoliu@example.com', NOW(), NOW()),
('钱七', '教授', 'qianqi@example.com', NOW(), NOW());

-- 插入示例课程数据
INSERT INTO course (code, name, credits, description, created_at, updated_at) VALUES
('CS101', '计算机导论', 3.0, '计算机科学的基础入门课程', NOW(), NOW()),
('CS201', '数据结构', 4.0, '介绍基本数据结构和算法', NOW(), NOW()),
('CS301', '数据库系统', 3.5, '数据库设计与实现', NOW(), NOW()),
('CS401', '操作系统', 4.0, '操作系统原理与设计', NOW(), NOW()),
('CS501', '软件工程', 3.0, '软件开发方法与实践', NOW(), NOW());

-- 插入教师课程关联
INSERT INTO teacher_courses (teacher_id, course_id) VALUES
(1, 1), -- 张三教授计算机导论
(1, 2), -- 张三教授数据结构
(2, 3), -- 李四教授数据库系统
(3, 4), -- 王五教授操作系统
(4, 5), -- 赵六教授软件工程
(5, 1); -- 钱七教授计算机导论

-- 插入示例教室数据
INSERT INTO classroom (building, room, capacity, type, created_at, updated_at) VALUES
('信息楼', '101', 60, '普通教室', NOW(), NOW()),
('信息楼', '102', 80, '多媒体教室', NOW(), NOW()),
('信息楼', '201', 40, '实验室', NOW(), NOW()),
('理科楼', '301', 100, '阶梯教室', NOW(), NOW()),
('理科楼', '302', 50, '普通教室', NOW(), NOW());

-- 插入示例学生数据
INSERT INTO student (student_id, name, grade, major, created_at, updated_at) VALUES
('2023001', '学生1', '2023', '计算机科学', NOW(), NOW()),
('2023002', '学生2', '2023', '软件工程', NOW(), NOW()),
('2023003', '学生3', '2023', '人工智能', NOW(), NOW()),
('2022001', '学生4', '2022', '计算机科学', NOW(), NOW()),
('2022002', '学生5', '2022', '软件工程', NOW(), NOW());

-- 插入示例课程安排数据
INSERT INTO section (course_id, teacher_id, classroom_id, time_slot_id, semester, max_students, created_at, updated_at) VALUES
(1, 1, 1, 1, '2023-2024-1', 50, NOW(), NOW()),  -- 计算机导论，张三，信息楼101，周一1-2节
(2, 1, 2, 6, '2023-2024-1', 60, NOW(), NOW()),  -- 数据结构，张三，信息楼102，周二1-2节
(3, 2, 3, 11, '2023-2024-1', 40, NOW(), NOW()), -- 数据库系统，李四，信息楼201，周三1-2节
(4, 3, 4, 16, '2023-2024-1', 80, NOW(), NOW()), -- 操作系统，王五，理科楼301，周四1-2节
(5, 4, 5, 21, '2023-2024-1', 40, NOW(), NOW()); -- 软件工程，赵六，理科楼302，周五1-2节

-- 插入示例选课记录
INSERT INTO enrollment (student_id, section_id, enroll_time, grade, created_at, updated_at) VALUES
(1, 1, NOW(), 'A', NOW(), NOW()),  -- 学生1选修计算机导论
(1, 2, NOW(), 'B', NOW(), NOW()),  -- 学生1选修数据结构
(2, 1, NOW(), 'A', NOW(), NOW()),  -- 学生2选修计算机导论
(2, 3, NOW(), 'B+', NOW(), NOW()), -- 学生2选修数据库系统
(3, 4, NOW(), 'A-', NOW(), NOW()), -- 学生3选修操作系统
(4, 5, NOW(), 'B+', NOW(), NOW()), -- 学生4选修软件工程
(5, 1, NOW(), 'A', NOW(), NOW());  -- 学生5选修计算机导论