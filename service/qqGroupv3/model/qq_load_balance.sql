/*
 Navicat Premium Data Transfer

 Source Server         : 内网新运维测试后台_10.10.88.229
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 10.10.88.229:3306
 Source Schema         : qq_group

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 28/04/2022 18:22:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for qq_load_balance
-- ----------------------------
DROP TABLE IF EXISTS `qq_load_balance`;
CREATE TABLE `qq_load_balance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `qq_api` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT 'qqapi地址',
  `is_master` smallint(255) NULL DEFAULT 0 COMMENT '是否为主q 0：未激活 1： 激活',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
