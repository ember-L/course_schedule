-- 修复teacher_courses表的主键问题
USE course_schedule;

-- 删除现有表（如果存在）
DROP TABLE IF EXISTS teacher_courses;

-- 重新创建teacher_courses表，正确定义复合主键
CREATE TABLE teacher_courses (
  teacher_id INT UNSIGNED NOT NULL,
  course_id INT UNSIGNED NOT NULL,
  PRIMARY KEY (teacher_id, course_id),
  CONSTRAINT fk_tc_teacher FOREIGN KEY (teacher_id) REFERENCES teacher (id) ON DELETE CASCADE,
  CONSTRAINT fk_tc_course FOREIGN KEY (course_id) REFERENCES course (id) ON DELETE CASCADE
);