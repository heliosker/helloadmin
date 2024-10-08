/*
 Navicat Premium Data Transfer

 Source Server         : Tiger-MySQL
 Source Server Type    : MySQL
 Source Server Version : 80034
 Source Host           : 47.103.204.136:3306
 Source Schema         : helloadmin

 Target Server Type    : MySQL
 Target Server Version : 80034
 File Encoding         : 65001

 Date: 08/10/2024 17:12:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `handle` varchar(128) DEFAULT NULL COMMENT 'handle',
  `title` varchar(128) DEFAULT NULL COMMENT '标题',
  `path` varchar(128) DEFAULT NULL COMMENT '地址',
  `method` varchar(16) DEFAULT NULL COMMENT '请求类型',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='接口';

-- ----------------------------
-- Records of api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for department
-- ----------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `parent_id` bigint DEFAULT '0' COMMENT '上级部门ID',
  `sort` bigint DEFAULT '0' COMMENT '排序值，值越大越靠前',
  `leader` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门负责人',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建于',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新于',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of department
-- ----------------------------
BEGIN;
INSERT INTO `department` (`id`, `name`, `parent_id`, `sort`, `leader`, `created_at`, `updated_at`) VALUES (1, '研发部', 0, 33, 'labore', '2024-02-08 18:30:41.677', '2024-03-18 18:27:32.106');
COMMIT;

-- ----------------------------
-- Table structure for login_record
-- ----------------------------
DROP TABLE IF EXISTS `login_record`;
CREATE TABLE `login_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `browser` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `email` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录邮箱',
  `ip` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '客户端IP',
  `os` varchar(60) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作系统',
  `platform` varchar(60) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '平台',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建于',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新于',
  `error_message` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=83 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of login_record
-- ----------------------------
BEGIN;
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (1, 'Apifox', 'admin@helloadmin.com', '::1', '', '', '2024-02-02 18:21:55.927', '2024-02-02 18:21:55.927', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (2, 'Apifox', 'admin@helloadmin.com', '::1', '', '', '2024-02-02 18:22:13.216', '2024-02-02 18:22:13.216', 'The password is incorrect');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (3, 'Apifox', 'admin@helloadmin.com', '::1', '', '', '2024-02-02 18:22:22.181', '2024-02-02 18:22:22.181', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (4, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-02 18:57:41.354', '2024-02-02 18:57:41.354', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (5, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-02 18:59:34.726', '2024-02-02 18:59:34.726', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (6, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 09:55:32.998', '2024-02-04 09:55:32.998', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (7, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 09:56:54.128', '2024-02-04 09:56:54.128', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (8, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 09:57:14.284', '2024-02-04 09:57:14.284', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (9, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 09:57:28.374', '2024-02-04 09:57:28.374', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (10, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 09:57:54.805', '2024-02-04 09:57:54.805', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (11, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 10:07:36.918', '2024-02-04 10:07:36.918', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (12, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 10:12:49.244', '2024-02-04 10:12:49.244', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (13, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 10:29:17.315', '2024-02-04 10:29:17.315', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (14, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:00:07.797', '2024-02-04 11:00:07.797', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (15, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:00:17.586', '2024-02-04 11:00:17.586', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (16, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:00:56.581', '2024-02-04 11:00:56.581', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (17, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:01:42.197', '2024-02-04 11:01:42.197', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (18, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:03:03.830', '2024-02-04 11:03:03.830', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (19, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:04:05.228', '2024-02-04 11:04:05.228', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (20, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:04:09.519', '2024-02-04 11:04:09.519', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (21, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:04:33.334', '2024-02-04 11:04:33.334', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (22, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:04:47.616', '2024-02-04 11:04:47.616', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (23, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:07:47.713', '2024-02-04 11:07:47.713', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (24, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:12:55.015', '2024-02-04 11:12:55.015', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (25, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-04 11:13:45.548', '2024-02-04 11:13:45.548', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (26, 'Apifox', 'admin@helloadmin.com', '::1', '', '', '2024-02-04 17:37:46.611', '2024-02-04 17:37:46.611', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (27, 'Apifox', 'admin@helloadmin.com', '::1', '', '', '2024-02-04 17:38:16.130', '2024-02-04 17:38:16.130', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (28, 'Apifox', 'admin@helloadmin.com', '::1', '', '', '2024-02-04 19:34:38.833', '2024-02-04 19:34:38.833', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (29, 'Apifox', 'admin@helloadmin.com', '101.230.253.82', '', '', '2024-02-05 14:35:58.455', '2024-02-05 14:35:58.455', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (30, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-18 09:49:18.412', '2024-02-18 09:49:18.412', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (31, 'Apifox', 'admin@helloadmin.com', '::1', '', '', '2024-02-20 11:11:15.584', '2024-02-20 11:11:15.584', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (32, 'Apifox', 'admin@helloadmin.com', '::1', '', '', '2024-02-20 11:11:50.605', '2024-02-20 11:11:50.605', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (33, 'Apifox', 'admin@helloadmin.com', '101.230.253.82', '', '', '2024-02-20 11:28:45.839', '2024-02-20 11:28:45.839', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (34, 'Chrome', 'test@qq.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-02-23 15:28:05.098', '2024-02-23 15:28:05.098', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (35, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-01 17:33:18.844', '2024-03-01 17:33:18.844', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (36, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-04 11:09:08.737', '2024-03-04 11:09:08.737', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (37, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-05 10:38:52.729', '2024-03-05 10:38:52.729', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (38, 'Chrome', 'admin@helloadmin.com', '101.230.253.82', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-05 10:43:54.123', '2024-03-05 10:43:54.123', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (39, 'Chrome', 'admin@helloadmin.com', '172.27.0.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-18 14:20:10.783', '2024-03-18 14:20:10.783', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (40, 'Chrome', 'admin@example.com', '172.27.0.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-18 16:05:59.095', '2024-03-18 16:05:59.095', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (41, 'Chrome', 'admin@example.com', '172.27.0.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-18 16:06:09.369', '2024-03-18 16:06:09.369', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (42, 'Chrome', 'admin@helloadmin.com', '172.27.0.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-18 16:07:16.750', '2024-03-18 16:07:16.750', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (43, 'Chrome', 'admin@helloadmin.com', '172.27.0.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-18 18:08:48.980', '2024-03-18 18:08:48.980', 'The password is incorrect');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (44, 'Chrome', 'admin@helloadmin.com', '172.27.0.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-18 18:09:00.234', '2024-03-18 18:09:00.234', 'The password is incorrect');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (45, 'Chrome', 'admin@helloadmin.com', '172.27.0.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-18 18:09:13.286', '2024-03-18 18:09:13.286', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (46, 'Chrome', 'admin@helloadmin.com', '172.27.0.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-18 18:24:56.945', '2024-03-18 18:24:56.945', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (47, 'Chrome', 'admin@123.com', '192.168.16.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-18 20:25:03.066', '2024-03-18 20:25:03.066', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (48, 'Chrome', 'admin@123.com', '192.168.16.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-03-18 20:25:20.019', '2024-03-18 20:25:20.019', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (49, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-03-22 02:00:49.572', '2024-03-22 02:00:49.572', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (50, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-03-22 02:00:58.611', '2024-03-22 02:00:58.611', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (51, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-03-22 02:01:03.544', '2024-03-22 02:01:03.544', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (52, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-03-28 09:21:04.442', '2024-03-28 09:21:04.442', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (53, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-03-28 09:21:14.503', '2024-03-28 09:21:14.503', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (54, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-03-28 09:22:05.887', '2024-03-28 09:22:05.887', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (55, 'Edge', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-04-07 14:14:09.307', '2024-04-07 14:14:09.307', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (56, 'Chrome', 'admin@helloadmin.cn', '192.168.32.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-04-07 16:47:07.053', '2024-04-07 16:47:07.053', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (57, 'Chrome', 'admin@123.com', '192.168.32.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-04-07 16:47:26.831', '2024-04-07 16:47:26.831', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (58, 'Chrome', 'admin@123.com', '192.168.32.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-04-07 16:47:39.901', '2024-04-07 16:47:39.901', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (59, 'Chrome', 'admin@helloadmin.com', '192.168.32.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-04-07 16:48:58.136', '2024-04-07 16:48:58.136', '');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (60, 'Chrome', 'admin@helloadmin.com', '192.168.32.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-04-07 16:49:14.714', '2024-04-07 16:49:14.714', 'The password is incorrect');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (61, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-04-18 06:59:00.233', '2024-04-18 06:59:00.233', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (62, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-04-18 06:59:23.421', '2024-04-18 06:59:23.421', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (63, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-04-23 16:27:22.890', '2024-04-23 16:27:22.890', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (64, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-04-23 16:28:35.674', '2024-04-23 16:28:35.674', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (65, 'Chrome', 'admin@123.com', '192.168.32.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-04-26 14:26:40.257', '2024-04-26 14:26:40.257', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (66, 'Chrome', 'admin@123.com', '192.168.32.1', 'Intel Mac OS X 10_15_7', 'Macintosh', '2024-04-26 14:26:55.520', '2024-04-26 14:26:55.520', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (67, 'Edge', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-11 00:09:50.187', '2024-05-11 00:09:50.187', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (68, 'Edge', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-11 00:09:56.694', '2024-05-11 00:09:56.694', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (69, 'Edge', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-11 00:10:03.314', '2024-05-11 00:10:03.314', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (70, 'Edge', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-11 15:50:20.946', '2024-05-11 15:50:20.946', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (71, 'Chrome', 'admin@123.com', '192.168.32.1', 'Android 10', 'Linux', '2024-05-12 00:27:02.508', '2024-05-12 00:27:02.508', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (72, 'Chrome', 'admin@123.com', '192.168.32.1', 'Android 10', 'Linux', '2024-05-12 00:27:32.775', '2024-05-12 00:27:32.775', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (73, 'Edge', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-12 01:37:45.291', '2024-05-12 01:37:45.291', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (74, 'Edge', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-12 01:39:01.291', '2024-05-12 01:39:01.291', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (75, 'Edge', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-12 01:39:05.086', '2024-05-12 01:39:05.086', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (76, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-16 09:05:51.987', '2024-05-16 09:05:51.987', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (77, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-16 09:05:58.405', '2024-05-16 09:05:58.405', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (78, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-16 09:06:12.814', '2024-05-16 09:06:12.814', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (79, 'Edge', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-05-26 23:28:01.291', '2024-05-26 23:28:01.291', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (80, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-06-09 21:44:16.727', '2024-06-09 21:44:16.727', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (81, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-06-09 21:44:40.807', '2024-06-09 21:44:40.807', 'The user does not exist');
INSERT INTO `login_record` (`id`, `browser`, `email`, `ip`, `os`, `platform`, `created_at`, `updated_at`, `error_message`) VALUES (82, 'Chrome', 'admin@123.com', '192.168.32.1', 'Windows 10', 'Windows', '2024-06-09 21:45:21.319', '2024-06-09 21:45:21.319', 'The user does not exist');
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单标题',
  `icon` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `path` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单路径',
  `type` char(1) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单类型 目录D 菜单M 按钮B',
  `parent_id` bigint DEFAULT '0' COMMENT '上级菜单ID',
  `component` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `sort` bigint DEFAULT '0' COMMENT '排序值，值越大越靠前',
  `visible` char(1) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '是否可见，Y可见 N不可见',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建于',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新于',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (5, 'Workplace', '仪表盘', 'bx-analyse', '/workplace', 'M', 0, 'Home', 1, 'Y', '2024-03-04 18:34:49.145', '2024-03-04 18:34:49.145');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (6, 'DepartmentUser', '部门与员工管理', 'user', '/departmentUser', 'M', 0, 'departmentUser', 2, 'Y', '2024-03-04 18:39:12.328', '2024-03-04 18:39:12.328');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (7, 'Auth', '权限管理', 'auth', 'auth', 'D', 0, 'RouteView', 3, 'Y', '2024-03-04 18:47:30.706', '2024-03-04 18:47:30.706');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (8, 'RoleList', '角色管理', 'role', 'role-list', 'M', 7, 'role', 1, 'Y', '2024-03-05 10:10:57.614', '2024-03-05 10:10:57.614');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (9, 'MenuList', '菜单列表', 'menu', 'menu-list', 'M', 7, 'menu', 2, 'Y', '2024-03-05 10:11:50.850', '2024-03-05 10:11:50.850');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (10, 'Log', '日志管理', 'log1', 'log', 'D', 0, 'RouteView', 4, 'Y', '2024-03-05 10:12:43.515', '2024-03-05 10:12:43.515');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (11, 'LoginLog', '登录日志', 'log', 'login-log', 'M', 10, 'loginLog', 1, 'Y', '2024-03-05 10:13:26.552', '2024-03-05 10:13:26.552');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (12, 'Edit User', '编辑', '', 'user:edit', 'B', 6, '', 2, 'Y', '2024-03-06 14:02:49.827', '2024-03-06 14:02:49.827');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (13, 'Add User', '新增员工', '', 'user:add', 'B', 6, '', 1, 'Y', '2024-03-06 14:19:02.237', '2024-03-06 14:19:02.237');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (14, 'Delete User', '删除员工', '', 'user:delete', 'B', 6, '', 3, 'Y', '2024-03-06 14:20:35.011', '2024-03-06 14:20:35.011');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (15, 'Add Department', '新增部门', '', 'department:add', 'B', 6, '', 4, 'Y', '2024-03-06 14:21:54.302', '2024-03-06 14:21:54.302');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (16, 'Edit Department', '编辑部门', '', 'department:edit', 'B', 6, '', 5, 'Y', '2024-03-06 14:22:39.341', '2024-03-06 14:22:39.341');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (17, 'Delete Department', '删除部门', '', 'department:delete', 'B', 6, '', 6, 'Y', '2024-03-06 14:23:24.066', '2024-03-06 14:23:24.066');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (18, 'Add Role', '新增角色', '', 'role:add', 'B', 8, '', 1, 'Y', '2024-03-06 14:44:49.834', '2024-03-06 14:44:49.834');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (19, 'Edit Role', '编辑角色', '', 'role:edit', 'B', 8, '', 2, 'Y', '2024-03-06 14:46:21.039', '2024-03-06 14:46:21.039');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (20, 'Delete Role', '删除角色', '', 'role:delete', 'B', 8, '', 3, 'Y', '2024-03-06 14:46:56.260', '2024-03-06 14:46:56.260');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (21, 'Auth Save', '权限保存', '', 'auth:save', 'B', 8, '', 4, 'Y', '2024-03-06 14:48:03.435', '2024-03-06 14:48:03.435');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (22, 'Add Menu', '新增菜单', '', 'menu:add', 'B', 9, '', 1, 'Y', '2024-03-06 14:48:44.358', '2024-03-06 14:48:44.358');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (23, 'Edit Menu', '编辑菜单', '', 'menu:edit', 'B', 9, '', 2, 'Y', '2024-03-06 14:49:15.557', '2024-03-06 14:49:15.557');
INSERT INTO `menu` (`id`, `name`, `title`, `icon`, `path`, `type`, `parent_id`, `component`, `sort`, `visible`, `created_at`, `updated_at`) VALUES (24, 'Delete Menu', '删除菜单', '', 'menu:delete', 'B', 9, '', 3, 'Y', '2024-03-06 14:49:45.357', '2024-03-06 14:49:45.357');
COMMIT;

-- ----------------------------
-- Table structure for operation_record
-- ----------------------------
DROP TABLE IF EXISTS `operation_record`;
CREATE TABLE `operation_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账号唯一ID',
  `operation` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作类型',
  `path` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作路径',
  `method` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求方式',
  `ip` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '客户端IP',
  `http_status` bigint NOT NULL DEFAULT '0' COMMENT 'HTTP状态码',
  `payload` text COLLATE utf8mb4_general_ci COMMENT '请求参数',
  `response` text COLLATE utf8mb4_general_ci COMMENT '响应结果',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建于',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新于',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of operation_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `describe` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色描述',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建于',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新于',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`, `describe`, `created_at`, `updated_at`) VALUES (1, '超级管理员', 'exercitation id adipisicing enim', '2024-02-20 11:04:59.163', '2024-03-07 11:42:42.515');
COMMIT;

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `model_id` bigint unsigned NOT NULL,
  `menu_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`model_id`,`menu_id`),
  KEY `fk_role_menu_menus` (`menu_id`),
  CONSTRAINT `fk_role_menu_menus` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`),
  CONSTRAINT `fk_role_menu_model` FOREIGN KEY (`model_id`) REFERENCES `role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
BEGIN;
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 5);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 6);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 7);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 8);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 9);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 10);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 11);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 12);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 13);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 14);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 15);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 16);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 17);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 18);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 19);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 20);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 21);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 22);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 23);
INSERT INTO `role_menu` (`model_id`, `menu_id`) VALUES (1, 24);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账号唯一ID',
  `nickname` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '昵称',
  `password` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `email` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '盐字段',
  `role_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `dept_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '部门ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建于',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新于',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除于',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `fk_user_role` (`role_id`),
  KEY `fk_user_department` (`dept_id`),
  CONSTRAINT `fk_user_department` FOREIGN KEY (`dept_id`) REFERENCES `department` (`id`),
  CONSTRAINT `fk_user_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `user_id`, `nickname`, `password`, `email`, `salt`, `role_id`, `dept_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'B2v7HdVlrm', '高静', '$2a$10$LdA79sXzj8nYUEC7ASWZJOhlxIjSHrmR0gcjqdpljEDZMTdf10jzC', 'w.sbfylgb@qq.com', 'GoS76zvddPBOBzbW', 1, 1, '2024-02-20 11:05:09.615', '2024-02-20 11:05:09.615', NULL);
INSERT INTO `user` (`id`, `user_id`, `nickname`, `password`, `email`, `salt`, `role_id`, `dept_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'B2vjjUToZe', 'admin', '$2a$10$fwkBcF.Z70z7kl6uXdFwdeKlfm03DtOzwY3hj5mnsUi0Yl4Ym7Whe', 'admin@helloadmin.com', 'kcTvOIxxy2YKS6bX', 1, 1, '2024-02-20 11:11:43.153', '2024-02-20 11:11:43.153', NULL);
INSERT INTO `user` (`id`, `user_id`, `nickname`, `password`, `email`, `salt`, `role_id`, `dept_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'B3bq9Wz7XQ', 'test', '$2a$10$6k4I4p0/HkSEFfaB0Td2COdwre06zn5qqbp9mMrikmvbPRNTTLm3u', '123@qq.com', 'wBzdZvXMXv5Djzpp', 1, 1, '2024-02-21 11:44:54.630', '2024-02-21 11:44:54.630', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
