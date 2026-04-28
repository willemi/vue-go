# MySQL 操作指南

## 执行逻辑

- SHOW TABLES; 是 MySQL 内部命令，不是 shell 命令。你需要先进入 MySQL 再执行：

```bash
mysql -u root -p
```

输入密码进入 MySQL shell 后，再执行：

```bash
USE fullstack_db;
SHOW TABLES;
```

或者一行命令完成（不进入交互式界面）：

```bash
mysql -u root -p -e "USE fullstack_db; SHOW TABLES;"
```

## 连接数据库

```bash
mysql -u root -p
mysql -u root -p -h 127.0.0.1 -P 3306
mysql -u root -p fullstack_db
```

## 数据库操作

```sql
-- 查看所有数据库
SHOW DATABASES;

-- 创建数据库
CREATE DATABASE fullstack_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE fullstack_db;

-- 删除数据库
DROP DATABASE fullstack_db;

-- 查看当前数据库
SELECT DATABASE();
```

## 表操作

```sql
-- 查看所有表
SHOW TABLES;

-- 查看表结构
DESC users;
DESCRIBE users;

-- 查看建表语句
SHOW CREATE TABLE users;

-- 创建表
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- 删除表
DROP TABLE users;

-- 清空表（自增ID重置）
TRUNCATE TABLE users;
```

## 增删改查（CRUD）

```sql
-- 插入数据
INSERT INTO users (username, password, role) VALUES ('admin', 'hashed_pwd', 'admin');
INSERT INTO users (username, password, role) VALUES
    ('user1', 'pwd1', 'user'),
    ('user2', 'pwd2', 'user');

-- 查询数据
SELECT * FROM users;
例子：mysql -u root -p352608ww -e "SELECT * FROM fullstack_db.users;"
SELECT id, username, role FROM users WHERE role = 'admin';
SELECT * FROM users WHERE deleted_at IS NULL;

-- 更新数据
UPDATE users SET role = 'admin' WHERE username = 'test';

-- 删除数据（物理删除）
DELETE FROM users WHERE id = 1;

例子：
mysql -u root -p352608ww -e "DELETE FROM fullstack_db.users WHERE username='admin';"

-- 软删除（配合 GORM 使用）
UPDATE users SET deleted_at = NOW() WHERE id = 1;
```

## 条件查询

```sql
-- WHERE 条件
SELECT * FROM users WHERE role = 'admin' AND deleted_at IS NULL;

-- LIKE 模糊匹配
SELECT * FROM users WHERE username LIKE '%admin%';

-- IN 查询
SELECT * FROM users WHERE id IN (1, 2, 3);

-- LIMIT 和 OFFSET
SELECT * FROM users LIMIT 10 OFFSET 20;

-- 排序
SELECT * FROM users ORDER BY created_at DESC;

-- 统计数量
SELECT COUNT(*) FROM users WHERE deleted_at IS NULL;
```

## 索引操作

```sql
-- 查看表的索引
SHOW INDEX FROM users;

-- 创建索引
CREATE INDEX idx_username ON users(username);
CREATE UNIQUE INDEX idx_email ON users(email);

-- 删除索引
DROP INDEX idx_username ON users;
```

## 用户与权限

```sql
-- 创建用户
CREATE USER 'testuser'@'localhost' IDENTIFIED BY 'password123';

-- 授权
GRANT ALL PRIVILEGES ON fullstack_db.* TO 'testuser'@'localhost';
GRANT SELECT, INSERT ON fullstack_db.* TO 'testuser'@'localhost';

-- 刷新权限
FLUSH PRIVILEGES;

-- 查看用户权限
SHOW GRANTS FOR 'testuser'@'localhost';

-- 删除用户
DROP USER 'testuser'@'localhost';
```

## 常用运维命令

```sql
-- 查看所有进程
SHOW PROCESSLIST;

-- 查看变量配置
SHOW VARIABLES LIKE 'max_connections';

-- 查看表状态
SHOW TABLE STATUS FROM fullstack_db;

-- 分析查询（查看查询计划）
EXPLAIN SELECT * FROM users WHERE username = 'admin';
EXPLAIN ANALYZE SELECT * FROM users WHERE role = 'admin';
```

## 导入导出

```bash
# 导出整个数据库
mysqldump -u root -p fullstack_db > fullstack_db.sql

# 导出指定表
mysqldump -u root -p fullstack_db users > users_table.sql

# 导出所有数据库
mysqldump -u root -p --all-databases > all_db.sql

# 导入
mysql -u root -p fullstack_db < fullstack_db.sql

# 导出数据（纯数据，无表结构）
mysqldump -u root -p --no-create-info fullstack_db users > users_data.sql
```

## 事务操作

```sql
-- 开启事务
START TRANSACTION;

-- 提交事务
COMMIT;

-- 回滚事务
ROLLBACK;

-- 设置保存点
SAVEPOINT savepoint_name;

-- 回滚到保存点
ROLLBACK TO SAVEPOINT savepoint_name;
```

## 聚合函数与分组

```sql
-- 统计
SELECT COUNT(*) FROM users;
SELECT COUNT(*) FROM users GROUP BY role;

-- 分组统计
SELECT role, COUNT(*) as count FROM users WHERE deleted_at IS NULL GROUP BY role;

-- HAVING 筛选分组结果
SELECT role, COUNT(*) as count FROM users GROUP BY role HAVING count > 1;
```

## 修改表结构

```sql
-- 添加列
ALTER TABLE users ADD COLUMN phone VARCHAR(20);

-- 修改列
ALTER TABLE users MODIFY COLUMN phone VARCHAR(30);

-- 删除列
ALTER TABLE users DROP COLUMN phone;

-- 重命名表
RENAME TABLE users TO user_table;
```
