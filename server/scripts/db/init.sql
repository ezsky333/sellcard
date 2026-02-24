-- Database initialization SQL for account-related tables
CREATE DATABASE IF NOT EXISTS `sellcard` CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;
USE `sellcard`;

-- users table
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(64) NOT NULL,
  `password_hash` VARCHAR(255) NOT NULL,
  `role` VARCHAR(32) NOT NULL DEFAULT 'user',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- optional roles table
CREATE TABLE IF NOT EXISTS `roles` (
  `name` VARCHAR(32) NOT NULL PRIMARY KEY,
  `description` VARCHAR(255) NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- sample insert (password should be stored as bcrypt hash by the application)
INSERT INTO users (username, password_hash, role) VALUES ('admin', '$2a$10$UVUrkYo0zrhYCU8EZ.gQiufbbr3z3Jm5UDI9t4IP2Hy87x61aV0AK', 'admin');
