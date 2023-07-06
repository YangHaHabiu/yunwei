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

 Date: 28/04/2022 18:22:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for qq_message_history
-- ----------------------------
DROP TABLE IF EXISTS `qq_message_history`;
CREATE TABLE `qq_message_history`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `message_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '消息id',
  `user_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户id',
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '消息内容',
  `create_time` int(11) NULL DEFAULT NULL COMMENT '创建时间',
  `executed` int(11) NULL DEFAULT 0 COMMENT '执行',
  `group_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '群组类型',
  `group_id` int(11) NULL DEFAULT NULL COMMENT '群id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
