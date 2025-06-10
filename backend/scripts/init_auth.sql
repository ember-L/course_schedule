-- 为教师表添加密码字段
ALTER TABLE teachers ADD COLUMN password VARCHAR(100) NOT NULL DEFAULT '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy';
-- 默认密码为 123456

-- 为学生表添加密码字段
ALTER TABLE students ADD COLUMN password VARCHAR(100) NOT NULL DEFAULT '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy';
-- 默认密码为 123456

-- 创建管理员表
CREATE TABLE IF NOT EXISTS admin_auths (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(100) NOT NULL,
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP NULL,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL
);

-- 插入默认管理员账号
INSERT INTO admin_auths (username, password, name, created_at, updated_at)
VALUES ('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '系统管理员', NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();
-- 默认密码为 123456