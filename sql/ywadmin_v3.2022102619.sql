-- MySQL dump 10.13  Distrib 8.0.29-21, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: ywadmin_v3
-- ------------------------------------------------------
-- Server version	8.0.29-21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*!50717 SELECT COUNT(*) INTO @rocksdb_has_p_s_session_variables FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'performance_schema' AND TABLE_NAME = 'session_variables' */;
/*!50717 SET @rocksdb_get_is_supported = IF (@rocksdb_has_p_s_session_variables, 'SELECT COUNT(*) INTO @rocksdb_is_supported FROM performance_schema.session_variables WHERE VARIABLE_NAME=\'rocksdb_bulk_load\'', 'SELECT 0') */;
/*!50717 PREPARE s FROM @rocksdb_get_is_supported */;
/*!50717 EXECUTE s */;
/*!50717 DEALLOCATE PREPARE s */;
/*!50717 SET @rocksdb_enable_bulk_load = IF (@rocksdb_is_supported, 'SET SESSION rocksdb_bulk_load = 1', 'SET @rocksdb_dummy_bulk_load = 0') */;
/*!50717 PREPARE s FROM @rocksdb_enable_bulk_load */;
/*!50717 EXECUTE s */;
/*!50717 DEALLOCATE PREPARE s */;

--
-- Table structure for table `asset`
--

DROP TABLE IF EXISTS `asset`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `asset` (
  `asset_id` smallint unsigned NOT NULL AUTO_INCREMENT COMMENT '资产ID',
  `outer_ip` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '公网IP',
  `inner_ip` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '内网IP',
  `accelerate_domain` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '加速域名',
  `host_role_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '服务器用途ID(game|cross|center...)',
  `provider_id` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '云商ID',
  `hardware_info` json NOT NULL COMMENT '服务器硬件信息',
  `ssh_port` int unsigned NOT NULL DEFAULT '12580' COMMENT 'SSH端口',
  `init_type` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '初始化状态：2:未初始化;1:已初始化',
  `clean_type` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '清理状态：2:未清理;1:已清理',
  `recycle_type` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '回收状态：2:使用中;1:已回收',
  `init_login_info` json NOT NULL COMMENT '初始登录信息(用户、密码/key路径等,只在未初始化时需要此)',
  `change_status_remark` json NOT NULL COMMENT '状态变更备注信息(如初始化、清理、归还等信息),记录该操作时间,操作人,操作备注',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '备注信息',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `del_flag` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '删除状态：0:未删除;1:已删除',
  PRIMARY KEY (`asset_id`) USING BTREE,
  KEY `asset_index_outer_ip` (`outer_ip`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='资产信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `asset`
--

LOCK TABLES `asset` WRITE;
/*!40000 ALTER TABLE `asset` DISABLE KEYS */;
/*!40000 ALTER TABLE `asset` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `company`
--

DROP TABLE IF EXISTS `company`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `company` (
  `company_id` tinyint unsigned NOT NULL AUTO_INCREMENT COMMENT '公司ID',
  `company_cn` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '公司中文名',
  `company_en` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '公司英文名',
  `supply_company_status` tinyint DEFAULT '1' COMMENT '出机方状态：1：正常 2：下线',
  `del_flag` tinyint unsigned DEFAULT '0' COMMENT '删除状态：0:未删除(数据使用中);1:已删除(回收)',
  PRIMARY KEY (`company_id`) USING BTREE,
  KEY `company_index_company_en` (`company_en`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='公司表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `company`
--

LOCK TABLES `company` WRITE;
/*!40000 ALTER TABLE `company` DISABLE KEYS */;
/*!40000 ALTER TABLE `company` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `config_file`
--

DROP TABLE IF EXISTS `config_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `config_file` (
  `id` int NOT NULL AUTO_INCREMENT,
  `project_id` int NOT NULL DEFAULT '0',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '',
  `dest_path` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '',
  `file_mod_time` int DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `project_id_and_name_unique` (`project_id`,`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='配置文件信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config_file`
--

LOCK TABLES `config_file` WRITE;
/*!40000 ALTER TABLE `config_file` DISABLE KEYS */;
/*!40000 ALTER TABLE `config_file` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `config_mng_log`
--

DROP TABLE IF EXISTS `config_mng_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `config_mng_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `asset_id` int DEFAULT '0',
  `config_file_id` int DEFAULT '0',
  `config_time` int DEFAULT '0',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `asset_id_config_file_id` (`asset_id`,`config_file_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='配置管理日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config_mng_log`
--

LOCK TABLES `config_mng_log` WRITE;
/*!40000 ALTER TABLE `config_mng_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `config_mng_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `feature_server_info`
--

DROP TABLE IF EXISTS `feature_server_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `feature_server_info` (
  `feature_server_id` smallint NOT NULL AUTO_INCREMENT COMMENT '功能服信息ID',
  `project_id` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '项目ID',
  `feature_server_info` json NOT NULL COMMENT '功能服相关信息',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '备注信息',
  `del_flag` tinyint unsigned DEFAULT '0' COMMENT '删除状态：0:未删除(数据使用中);1:已删除(回收)',
  PRIMARY KEY (`feature_server_id`) USING BTREE,
  KEY `fsi_index_project_id` (`project_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='功能服';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `feature_server_info`
--

LOCK TABLES `feature_server_info` WRITE;
/*!40000 ALTER TABLE `feature_server_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `feature_server_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `game_server`
--

DROP TABLE IF EXISTS `game_server`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `game_server` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `center_id` int unsigned DEFAULT '0' COMMENT '中央后台中序号',
  `project_id` tinyint unsigned NOT NULL COMMENT '项目ID',
  `platform_id` smallint unsigned NOT NULL COMMENT '平台ID',
  `server_id` smallint unsigned NOT NULL COMMENT '服ID',
  `server_alias` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '服别名',
  `open_time` int NOT NULL COMMENT '开服时间',
  `asset_id` smallint unsigned NOT NULL COMMENT '资产ID',
  `server_status` tinyint unsigned NOT NULL COMMENT '服状态：0:停服;1:收费服;2:合服;3:被合服;4:测试服;5:跨服',
  `combine_remark` varchar(300) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '合服信息',
  `operate_info` json NOT NULL COMMENT '重要操作信息',
  `del_flag` tinyint unsigned DEFAULT '0' COMMENT '删除状态：0:未删除(数据使用中);1:已删除(回收)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='游戏服信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `game_server`
--

LOCK TABLES `game_server` WRITE;
/*!40000 ALTER TABLE `game_server` DISABLE KEYS */;
/*!40000 ALTER TABLE `game_server` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `game_server_ext`
--

DROP TABLE IF EXISTS `game_server_ext`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `game_server_ext` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `server_alias` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '服别名',
  `node_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '目录名',
  `game_port` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '游戏端口',
  `console_port` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '调试端口',
  `http_port` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '游戏服web端口',
  `router` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '路由地址和端口',
  `center` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '中心服地址和端口',
  `mysql_host` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据库IP',
  `mysql_port` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '3306' COMMENT '数据库端口',
  `mysql_user` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据库用户',
  `mysql_password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据库密码',
  `mysql_database` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '数据库名称',
  `mergedate` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '合服时间',
  `gameserver_title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服名',
  `wss_port` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转发端口',
  `wan_ip` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '服外网IP',
  `wan_host` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '转发域名',
  `open_status` int NOT NULL DEFAULT '3' COMMENT '服务器开启状态(0:新服;1:老服;2:维护状态;3:未开服)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='游戏服信息扩展表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `game_server_ext`
--

LOCK TABLES `game_server_ext` WRITE;
/*!40000 ALTER TABLE `game_server_ext` DISABLE KEYS */;
/*!40000 ALTER TABLE `game_server_ext` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `hot_log_history`
--

DROP TABLE IF EXISTS `hot_log_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `hot_log_history` (
  `id` int NOT NULL AUTO_INCREMENT,
  `hot_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '热更标题',
  `project_id` int DEFAULT '0' COMMENT '游戏名',
  `oper_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '操作类型',
  `oper_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '操作日志',
  `oper_status` smallint DEFAULT '0' COMMENT '操作状态 0 成功|1 失败',
  `create_by` int DEFAULT '0' COMMENT '操作人',
  `create_time` int DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hot_log_history`
--

LOCK TABLES `hot_log_history` WRITE;
/*!40000 ALTER TABLE `hot_log_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `hot_log_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `key_manage`
--

DROP TABLE IF EXISTS `key_manage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `key_manage` (
  `key_id` int NOT NULL AUTO_INCREMENT COMMENT '密钥id',
  `key_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密钥名称',
  `key_path` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密钥路径',
  `key_pass` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '密钥密文',
  `key_type` smallint NOT NULL DEFAULT '1' COMMENT '密钥类型：1:临时key 2:个人key',
  `public_key` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '公钥',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '备注',
  `del_flag` tinyint DEFAULT '0' COMMENT '删除状态：0:未删除(数据使用中);1:已删除(回收)',
  PRIMARY KEY (`key_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='密钥管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `key_manage`
--

LOCK TABLES `key_manage` WRITE;
/*!40000 ALTER TABLE `key_manage` DISABLE KEYS */;
/*!40000 ALTER TABLE `key_manage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `label`
--

DROP TABLE IF EXISTS `label`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `label` (
  `label_id` bigint NOT NULL AUTO_INCREMENT,
  `label_type` tinyint NOT NULL COMMENT '标签类型 1,集群',
  `label_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `label_values` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '值',
  `label_remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '备注',
  `create_by` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `del_flag` tinyint NOT NULL DEFAULT '0' COMMENT '删除标识 0：未删除 1：删除',
  PRIMARY KEY (`label_id`) USING BTREE,
  UNIQUE KEY `label_name_values` (`label_name`,`label_values`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='标签信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `label`
--

LOCK TABLES `label` WRITE;
/*!40000 ALTER TABLE `label` DISABLE KEYS */;
/*!40000 ALTER TABLE `label` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `label_global`
--

DROP TABLE IF EXISTS `label_global`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `label_global` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `label_id` bigint NOT NULL COMMENT '标签id',
  `resource_en` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '资源英文名称',
  `binding_id` bigint NOT NULL COMMENT '绑定对象值',
  `project_id` bigint NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_index` (`label_id`,`resource_en`,`binding_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='标签组全局关联';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `label_global`
--

LOCK TABLES `label_global` WRITE;
/*!40000 ALTER TABLE `label_global` DISABLE KEYS */;
/*!40000 ALTER TABLE `label_global` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `maintain_plan`
--

DROP TABLE IF EXISTS `maintain_plan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `maintain_plan` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `project_id` tinyint unsigned NOT NULL COMMENT '项目ID',
  `maintain_type` tinyint unsigned DEFAULT NULL COMMENT '维护类型：1:例行维护;2:临时维护;3:关服计划;4:迁移计划',
  `start_time` int NOT NULL COMMENT '维护开始时间',
  `end_time` int NOT NULL COMMENT '维护结束时间',
  `maintain_range` text CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '维护范围',
  `title` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '维护内容',
  `create_by` int DEFAULT '0' COMMENT '创建人',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` int DEFAULT '0' COMMENT '修改人',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `maintain_operator` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '操作人',
  `del_flag` tinyint DEFAULT '0' COMMENT '删除状态：0:未删除(数据使用中);1:已删除(回收)',
  `cluster_id` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '集群ids',
  `task_id` int DEFAULT '-1' COMMENT '任务id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='维护计划';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `maintain_plan`
--

LOCK TABLES `maintain_plan` WRITE;
/*!40000 ALTER TABLE `maintain_plan` DISABLE KEYS */;
/*!40000 ALTER TABLE `maintain_plan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `merge_plan`
--

DROP TABLE IF EXISTS `merge_plan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `merge_plan` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `project_id` tinyint unsigned NOT NULL COMMENT '项目ID',
  `platform_id` smallint unsigned NOT NULL COMMENT '平台ID',
  `server_id` smallint unsigned NOT NULL COMMENT '服ID',
  `input_range` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '输入合服范围',
  `combine_range` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '实际合服范围',
  `start_time` int NOT NULL COMMENT '合服开始时间',
  `end_time` int NOT NULL COMMENT '合服结束时间',
  `merge_status` tinyint DEFAULT '-1' COMMENT '合服状态：-1:未合服;1:已合服',
  `merge_operator` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '合服人',
  `del_flag` tinyint unsigned DEFAULT '0' COMMENT '删除状态：0:未删除(数据使用中);1:已删除(回收)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='合服计划';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `merge_plan`
--

LOCK TABLES `merge_plan` WRITE;
/*!40000 ALTER TABLE `merge_plan` DISABLE KEYS */;
/*!40000 ALTER TABLE `merge_plan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `open_plan`
--

DROP TABLE IF EXISTS `open_plan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `open_plan` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `project_id` tinyint unsigned NOT NULL COMMENT '项目ID',
  `platform_id` smallint unsigned NOT NULL COMMENT '平台ID',
  `server_id` smallint unsigned NOT NULL COMMENT '服ID',
  `gameserver_title` text CHARACTER SET utf8mb3 COLLATE utf8_general_ci COMMENT '服务器名称',
  `open_time` bigint NOT NULL COMMENT '开服时间',
  `install_status` tinyint DEFAULT '-1' COMMENT '安装状态：-1:未安装;1:已安装',
  `install_operator` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '安装人',
  `initdb_status` tinyint DEFAULT '-1' COMMENT '清档状态：-1:未清档;1:已清档',
  `initdb_operator` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '清档人',
  `remark` text CHARACTER SET utf8mb3 COLLATE utf8_general_ci COMMENT '备注信息',
  `del_flag` tinyint unsigned DEFAULT '0' COMMENT '删除状态：0:未删除(数据使用中);1:已删除(回收)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='开服计划';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `open_plan`
--

LOCK TABLES `open_plan` WRITE;
/*!40000 ALTER TABLE `open_plan` DISABLE KEYS */;
/*!40000 ALTER TABLE `open_plan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `platform`
--

DROP TABLE IF EXISTS `platform`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `platform` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `project_id` tinyint unsigned NOT NULL COMMENT '项目ID',
  `platform_id` smallint unsigned NOT NULL COMMENT '平台ID',
  `platform_en` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '平台英文名',
  `platform_cn` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '平台中文名',
  `domain_format` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '单服域名格式',
  `remark` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '备注信息',
  `del_flag` tinyint unsigned DEFAULT '0' COMMENT '删除状态：0:未删除(数据使用中);1:已删除(回收)',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `platform_unique_index` (`platform_id`,`del_flag`,`project_id`) USING BTREE COMMENT '唯一索引[项目ID+平台ID+删除状态]',
  KEY `platform_index_platform_en` (`platform_en`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='平台管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `platform`
--

LOCK TABLES `platform` WRITE;
/*!40000 ALTER TABLE `platform` DISABLE KEYS */;
/*!40000 ALTER TABLE `platform` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project`
--

DROP TABLE IF EXISTS `project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `project` (
  `project_id` tinyint unsigned NOT NULL AUTO_INCREMENT COMMENT '项目ID',
  `project_cn` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '项目中文名',
  `project_en` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '项目英文名',
  `project_team` int DEFAULT NULL COMMENT '项目组',
  `project_type` tinyint unsigned DEFAULT '1' COMMENT '项目类型(1:自研;2:发行)',
  `group_qq` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '群qq号',
  `group_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '消息群类型：group，discuss',
  `group_dev_qq` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '通知开发qq组，多个用逗号分割',
  `del_flag` tinyint unsigned DEFAULT '0' COMMENT '删除状态：0:未删除(数据使用中);1:已删除(回收)',
  PRIMARY KEY (`project_id`) USING BTREE,
  KEY `project_index_project_en` (`project_en`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='项目表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project`
--

LOCK TABLES `project` WRITE;
/*!40000 ALTER TABLE `project` DISABLE KEYS */;
/*!40000 ALTER TABLE `project` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project_relationship`
--

DROP TABLE IF EXISTS `project_relationship`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `project_relationship` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `company_id` tinyint unsigned NOT NULL COMMENT '公司ID',
  `project_id` tinyint unsigned NOT NULL COMMENT '项目ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='公司项目关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project_relationship`
--

LOCK TABLES `project_relationship` WRITE;
/*!40000 ALTER TABLE `project_relationship` DISABLE KEYS */;
/*!40000 ALTER TABLE `project_relationship` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `server_affiliation`
--

DROP TABLE IF EXISTS `server_affiliation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `server_affiliation` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pr_id` smallint unsigned NOT NULL COMMENT '公司项目关系ID',
  `company_id` tinyint unsigned NOT NULL COMMENT '公司ID',
  `asset_id` smallint unsigned NOT NULL COMMENT '资产ID',
  `del_flag` tinyint unsigned DEFAULT '2' COMMENT '删除状态：2:未删除(数据使用中);1:已删除(回收)',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `asset_id` (`asset_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='服务器归属表(出机方及使用方)';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `server_affiliation`
--

LOCK TABLES `server_affiliation` WRITE;
/*!40000 ALTER TABLE `server_affiliation` DISABLE KEYS */;
/*!40000 ALTER TABLE `server_affiliation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stat_server_game_info`
--

DROP TABLE IF EXISTS `stat_server_game_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `stat_server_game_info` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `project_id` int NOT NULL,
  `project_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `counts` bigint NOT NULL,
  `detail` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` int NOT NULL,
  `count_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务器server游戏服game',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stat_server_game_info`
--

LOCK TABLES `stat_server_game_info` WRITE;
/*!40000 ALTER TABLE `stat_server_game_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `stat_server_game_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dept`
--

DROP TABLE IF EXISTS `sys_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dept` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '机构名称',
  `parent_id` bigint DEFAULT NULL COMMENT '上级机构ID，一级机构为0',
  `order_num` int DEFAULT NULL COMMENT '排序',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `last_update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '更新人',
  `last_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `del_flag` tinyint DEFAULT '0' COMMENT '是否删除  1：已删除  0：正常',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='部门管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dept`
--

LOCK TABLES `sys_dept` WRITE;
/*!40000 ALTER TABLE `sys_dept` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict`
--

DROP TABLE IF EXISTS `sys_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dict` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `pid` bigint DEFAULT NULL COMMENT 'pid',
  `value` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '数据值',
  `label` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '标签名',
  `types` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '类型（字典的key）',
  `description` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '描述',
  `sort` decimal(10,0) NOT NULL COMMENT '排序（升序）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=117 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='字典表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict`
--

LOCK TABLES `sys_dict` WRITE;
/*!40000 ALTER TABLE `sys_dict` DISABLE KEYS */;
INSERT INTO `sys_dict` VALUES (1,7,'2','停用','sys_normal_disable','用户停用状态',1),(2,7,'1','启用','sys_normal_disable','用户启用状态',2),(3,8,'1','显示','sys_show_hide','用户显示菜单',1),(4,8,'2','隐藏','sys_show_hide','用户隐藏菜单',2),(5,9,'1','自研','project_status','项目类型',1),(6,9,'2','发行','project_status','项目类型',2),(7,-1,'0','用户状态','sys_normal_disable','用户状态',1),(8,-1,'0','菜单状态','sys_show_hide','菜单显示',2),(9,-1,'0','项目类型','project_status','项目类型显示',3),(23,-1,'0','群组类型','project_group_type','项目组群组类型',1),(24,23,'group','群组','project_group_type','群组',1),(25,23,'discuss','讨论组','project_group_type','讨论组',2),(30,-1,'0','云商','cloud_provider_type','云商',1),(31,30,'1','qcloud','cloud_provider_type','腾讯云',1),(32,30,'2','aliyun','cloud_provider_type','阿里云',2),(33,-1,'0','服务器用途','host_role_type','服务器用途',1),(34,33,'1','game','host_role_type','游戏服',1),(35,33,'2','cross','host_role_type','跨服',2),(36,33,'3','center','host_role_type','中央服',3),(37,-1,'0','初始化状态','init_type','初始化状态',2),(38,-1,'0','清理状态','clean_type','清理状态',1),(39,-1,'0','删除状态','del_flag_type','删除状态',1),(40,39,'2','未删除','del_flag_type','数据使用中',1),(41,39,'1','已删除','del_flag_type','回收',1),(42,38,'2','未清理','clean_type','未清理',1),(43,38,'1','已清理','clean_type','已清理',1),(44,37,'2','未初始化','init_type','未初始化',1),(45,37,'1','已初始化','init_type','已初始化',1),(46,-1,'0','标签类型','label_type','标签类型',3),(47,46,'1','集群','label_type','集群标签',1),(48,46,'2','功能服','label_type','功能服信息标签',2),(49,46,'3','装服','label_type','装服标签',3),(50,46,'4','其他','label_type','其他标签',4),(51,33,'4','login','host_role_type','登录服',4),(52,33,'5','pay','host_role_type','充值服',5),(53,30,'3','hwcloud','cloud_provider_type','华为云',3),(54,-1,'0','项目状态','project_show_status','项目状态',1),(55,54,'-1','在线','project_show_status','在线',1),(56,54,'1','下线','project_show_status','下线',1),(57,-1,'0','回收状态','recycle_type','回收状态',1),(58,57,'2','未回收','recycle_type','未回收',1),(59,57,'1','已回收','recycle_type','已回收',2),(60,-1,'0','安装状态','install_status','安装状态',1),(61,-1,'0','清档状态','initdb_status','清档状态',1),(62,61,'-1','未清档','initdb_status','未清档',1),(63,61,'1','已清档','initdb_status','已清档',2),(64,60,'-1','未安装','install_status','未安装',1),(65,60,'1','已安装','install_status','已安装',2),(66,-1,'0','游戏服状态','game_server_status','游戏服状态',1),(67,66,'0','停服','game_server_status','停服',1),(68,66,'1','收费服','game_server_status','收费服',1),(69,66,'2','合服','game_server_status','合服',1),(70,66,'3','被合服','game_server_status','被合服',1),(71,66,'4','测试服','game_server_status','测试服',1),(72,66,'5','跨服','game_server_status','跨服',1),(73,-1,'0','维护类型','maintain_type','维护类型',1),(74,73,'1','例行维护','maintain_type','例行维护',1),(75,73,'3','停服','maintain_type','停服',3),(76,73,'4','其他','maintain_type','其他',4),(77,-1,'0','合服状态','merge_status','是否已合服',1),(78,77,'-1','否','merge_status','未合服',2),(79,77,'1','是','merge_status','已合服',1),(80,73,'2','临时维护','maintain_type','临时维护',2),(81,-1,'0','功能服类型','feature_server_type','功能服类型',1),(82,81,'1','center','feature_server_type','中央服',1),(84,81,'2','login','feature_server_type','登录服',2),(85,81,'3','pay','feature_server_type','充值服',3),(86,81,'4','source','feature_server_type','游戏源',4),(87,81,'5','cdn','feature_server_type','CDN源',5),(88,81,'6','cross','feature_server_type','跨服',6),(89,33,'6','test','host_role_type','测试服',6),(90,33,'7','backup','host_role_type','备份服',7),(91,81,'7','db','feature_server_type','数据库',7),(92,81,'8','gate','feature_server_type','网关',8),(93,33,'8','web','host_role_type','网站',8),(94,33,'9','db','host_role_type','数据库',9),(95,33,'10','cdn','host_role_type','CDN源',10),(96,33,'11','gate','host_role_type','网关',11),(97,-1,'0','任务类型','task_types','任务类型',1),(98,97,'1','临时维护','task_types','临时维护',1),(99,97,'2','日常维护','task_types','日常维护',2),(100,-1,'0','任务状态','task_status','任务状态',0),(101,100,'3','执行成功','task_status','执行成功',4),(102,100,'2','执行失败','task_status','执行失败',3),(103,100,'1','执行中','task_status','执行中',2),(104,100,'-1','未开始','task_status','未开始',1),(105,100,'4','取消任务','task_status','取消任务',5),(109,30,'4','gcloud','cloud_provider_type','谷歌云',4),(110,30,'5','capitalonline','cloud_provider_type','首都在线',5),(111,30,'6','aws','cloud_provider_type','亚马逊云',6),(112,30,'7','ucloud','cloud_provider_type','优刻得',7),(113,30,'8','chinac','cloud_provider_type','华云',8),(114,30,'9','jdcloud','cloud_provider_type','京东云',9),(115,81,'9','backup','feature_server_type','备份服',9),(116,100,'7','未开始(已删除)','task_status','未开始的任务',7);
/*!40000 ALTER TABLE `sys_dict` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_log`
--

DROP TABLE IF EXISTS `sys_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '用户名',
  `operation` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '用户操作',
  `method` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '请求方法',
  `params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '请求参数',
  `time` float NOT NULL COMMENT '执行时长(秒)',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'IP地址',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='系统操作日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_log`
--

LOCK TABLES `sys_log` WRITE;
/*!40000 ALTER TABLE `sys_log` DISABLE KEYS */;
INSERT INTO `sys_log` VALUES (1,'admin','/yunying/openPlan/list','POST','{\"current\":1,\"pageSize\":10}',0.00368079,'183.60.191.113','2022-10-26 11:47:39'),(2,'admin','/yunying/maintainPlan/list','POST','{\"current\":1,\"pageSize\":10}',0.00231353,'183.60.191.113','2022-10-26 11:47:40'),(3,'admin','/yunying/mergePlan/list','POST','{\"current\":1,\"pageSize\":10}',0.00299518,'183.60.191.113','2022-10-26 11:47:42');
/*!40000 ALTER TABLE `sys_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_login_log`
--

DROP TABLE IF EXISTS `sys_login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_login_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名',
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IP地址',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='系统登录日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_login_log`
--

LOCK TABLES `sys_login_log` WRITE;
/*!40000 ALTER TABLE `sys_login_log` DISABLE KEYS */;
INSERT INTO `sys_login_log` VALUES (1,'dengxiguang','login','183.60.191.113','2022-10-26 11:44:36'),(2,'dengxiguang','logout','183.60.191.113','2022-10-26 11:45:36'),(3,'dengxiguang','online','183.60.191.113','2022-10-26 11:45:59'),(4,'admin','login','183.60.191.113','2022-10-26 11:46:59'),(5,'admin','logout','183.60.191.113','2022-10-26 11:47:58');
/*!40000 ALTER TABLE `sys_login_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menu`
--

DROP TABLE IF EXISTS `sys_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '菜单名称',
  `parent_id` bigint DEFAULT '0' COMMENT '父菜单ID，一级菜单为0',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
  `perms` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)',
  `tp` int DEFAULT '0' COMMENT '类型   0：目录   1：菜单   2：按钮',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '菜单图标',
  `order_num` int DEFAULT '0' COMMENT '排序',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `last_update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '更新人',
  `last_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `vue_path` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'vue系统的path',
  `vue_component` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'vue的页面',
  `vue_icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'vue的图标',
  `vue_redirect` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'vue的路由重定向',
  `del_flag` tinyint DEFAULT '0' COMMENT '是否删除  -1：已删除  0：正常',
  `table_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '表',
  `keep_alive` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'false' COMMENT '是否缓存，默认否',
  `is_show` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'true' COMMENT '是否显示，默认是',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `vue_path` (`vue_path`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='菜单管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu`
--

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;
INSERT INTO `sys_menu` VALUES (1,'欢迎',0,'/welcome','',0,'',1,'admin','2021-02-26 06:45:04','admin','2021-02-26 14:45:04','','','','',1,'','false','true'),(2,'系统管理',0,'/system','',0,'',8,'admin','2021-02-26 06:45:04','admin','2022-08-30 14:45:06','/sys','Layout','setting','/sys/userList',0,'','false','true'),(3,'用户管理',2,'/system/user/list','',1,'',1,'admin','2021-02-26 06:45:04','admin','2022-09-05 17:32:54','userList','system/user/index','users','',0,'sys_user','false','true'),(4,'角色管理',2,'/system/role/list','',1,'',2,'admin','2021-02-26 06:45:04','admin','2021-02-26 14:45:04','roleList','system/role/index','users','',0,'sys_role','false','true'),(6,'部门管理',12,'/system/dept/list','',1,'',2,'admin','2021-02-26 06:45:04','admin','2021-02-26 14:45:04','deptList','system/dept/index','tree','',0,'sys_dept','false','true'),(7,'字典管理',2,'/system/dict/list','',1,'',5,'admin','2021-02-26 06:45:05','admin','2022-09-05 17:32:58','dictList','system/dict/index','users','',0,'sys_dict','false','true'),(8,'日志管理',0,'/log','',0,'',97,'admin','2021-02-26 06:45:05','admin','2022-08-30 14:45:15','/log','Layout','guide','/log/loginLogList',0,'','false','true'),(9,'登录日志',8,'/log/loginLog/list','',1,'',3,'admin','2021-02-26 06:45:05','admin','2021-02-26 14:45:05','loginLogList','log/loginlog/index','logininfor','',0,'sys_login_log','false','true'),(10,'操作日志',8,'/log/sysLog/list','',1,'',3,'admin','2021-02-26 06:45:05','admin','2021-02-26 14:45:05','sysLogList','log/syslog/index','form','',0,'sys_log','false','true'),(12,'项目中心',0,'/prcenter','',0,'',3,'admin','2021-02-26 06:45:05','admin','2022-08-30 14:44:42','/prcenter','Layout','guide','/prcenter/companyList',0,'','false','true'),(13,'公司管理',12,'/prcenter/company/list','',1,'',1,'admin','2022-04-08 07:35:03','admin','2022-04-08 15:39:29','companyList','prcenter/company/index','fa-building','',0,'company','false','true'),(14,'项目管理',12,'/prcenter/project/list','',1,'',3,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','projectList','prcenter/project/index','code','',0,'project','false','true'),(17,'策略管理',2,'/system/stgroup/list','',1,'',4,'admin','2021-02-26 06:45:05','admin','2021-02-26 14:45:05','stgroupList','system/stgroup/index','swagger','',0,'sys_stgroup','false','true'),(18,'标签列表',22,'/prcenter/label/list','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','labelList','prcenter/label/index','fa-tags','',0,'label','false','true'),(20,'用户组管理',2,'/system/ugroup/list','',1,'',1,'admin','2022-04-08 09:05:51','admin','2022-04-08 17:06:38','ugroupList','system/ugroup/index','post','',0,'sys_ugroup','false','true'),(22,'标签管理',0,'/label','',0,'',2,'admin','2022-04-08 07:35:38','admin','2022-08-31 15:09:00','/label','Layout','guide','/label/labelList',0,'','false','true'),(25,'资产管理',0,'/asset','',0,'',4,'admin','2021-02-26 06:45:05','admin','2022-08-30 14:44:46','/asset','Layout','guide','/asset/keyList',0,'','false','true'),(26,'资产信息',25,'/asset/asset/list','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','assetList','/asset/asset/index','fa-tags','',0,'asset','false','true'),(29,'运维管理',0,'/yunwei','',0,'',5,'admin','2021-02-26 06:45:05','admin','2022-08-30 14:44:49','/yunwei','Layout','guide','/yunwei/yunweiList',0,'','false','true'),(31,'游戏服信息',29,'/yunwei/gameServer/list','',1,'',3,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','gameServerList','/yunwei/gameServer/index','fa-tags','',0,'game','false','true'),(32,'集群管理',29,'/yunwei/cluster/list','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','clusterList','/yunwei/cluster/index','fa-tags','',0,'cluster','false','true'),(33,'平台管理',29,'/yunwei/platform/list','',1,'',4,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','platformList','/yunwei/platform/index','fa-tags','',0,'platform','false','true'),(34,'服务器信息',29,'/yunwei/hosts/list','',1,'',1,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','hostsList','/yunwei/hosts/index','fa-tags','',0,'hosts','false','true'),(35,'任务管理',0,'/taskMng','',0,'',6,'admin','2021-02-26 06:45:05','admin','2022-08-30 14:44:51','/taskMng','Layout','guide','/taskMng/taskMngList',0,'','false','true'),(36,'操作队列',35,'/taskMng/taskQueue/list','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-09-28 14:37:24','taskQueueList','/taskMng/taskQueue/index','fa-tags','',0,'game','false','true'),(37,'操作日志',35,'/taskMng/taskLog/list','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-09-28 14:37:26','taskLogList','/taskMng/taskLog/index','fa-tags','',0,'game','false','true'),(43,'功能服信息',29,'/yunwei/featureServer/list','',1,'',2,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','featureList','/yunwei/feature/index','fa-tags','',0,'feature_server_info','false','true'),(44,'运营管理',0,'/yunying','',0,'',7,'admin','2021-02-26 06:45:05','admin','2022-08-30 14:44:54','/yunying','Layout','guide','/yunying/yunyingList',0,'','false','true'),(47,'维护计划',44,'/yunying/maintainPlan/list','',1,'',2,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','maintainPlanList','/yunying/maintainPlan/index','fa-tags','',0,'game','false','true'),(48,'开服计划',44,'/yunying/openPlan/list','',1,'',1,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','openPlanList','/yunying/openPlan/index','fa-tags','',0,'game','false','true'),(49,'合服计划',44,'/yunying/mergePlan/list','',1,'',3,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','mergePlanList','/yunying/mergePlan/index','fa-tags','',0,'game','false','true'),(50,'添加资产',25,'/asset/assetAdd/list','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','assetAdd','/asset/assetAdd/index','fa-tags','',0,'asset','true','false'),(52,'平台详情',29,'/yunwei/platform/data','',1,'',4,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','platformDesc/:Id(\\w+)','/yunwei/platform/data','fa-tags','',0,'platform','false','false'),(53,'集群详情',29,'/yunwei/cluster/data','',1,'',4,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','clusterDesc/:Id(\\w+)','/yunwei/cluster/data','fa-tags','',0,'cluster','false','false'),(54,'帮助中心',0,'/help','',0,'',98,'admin','2021-02-26 06:45:05','admin','2022-08-30 14:45:04','/help','Layout','guide','/help/helpList',0,'','false','true'),(55,'格式化工具',54,'/help/format/list','',1,'',1,'admin','2022-04-08 07:35:38','admin','2022-08-02 11:41:05','formatList','/help/format/index','fa-tags','',0,'','false','true'),(56,'小助手更新模板',54,'/help/operate/list','',1,'',2,'admin','2022-07-12 02:11:36','admin','2022-09-30 10:04:20','operateList','/help/operate/index','fa-tags','',0,'','false','true'),(64,'运维',100,'yunwei','',2,'',1,'admin','2022-04-08 07:35:38','admin','2022-10-10 12:30:20','yunwei','yunwei','fa-tags','',0,'','false','false'),(65,'启动服务端',64,'start_game','',3,'',1,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:16:03','start_game','start_game','fa-tags','',0,'','false','false'),(66,'宕机重启',64,'crash_restart','',3,'',3,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:19:23','crash_restart','stop_game','fa-tags','',0,'','false','false'),(68,'检查游戏进程',64,'check_process','',3,'',5,'admin','2022-07-12 06:50:17','admin','2022-07-15 17:19:27','check_process','restart_game','fa-tags','',0,'','false','false'),(69,'清档和修改开服时间',64,'initdb_settime','',3,'',6,'admin','2022-07-12 06:50:17','admin','2022-07-15 17:19:28','initdb_settime','crash_restart','fa-tags','',0,'','false','false'),(70,'关停平台和清理回收',64,'stop_recycle','',3,'',7,'admin','2022-07-12 06:50:17','admin','2022-07-15 17:19:29','stop_recycle','check_process','fa-tags','',0,'','false','false'),(71,'获取服务器信息',64,'get_server_conf','',3,'',8,'admin','2022-07-12 06:50:17','admin','2022-07-15 17:19:31','get_server_conf','initdb_settime','fa-tags','',0,'','false','false'),(72,'后端',100,'backend','',2,'',2,'admin','2022-04-08 07:35:38','admin','2022-10-10 12:30:20','backend','backend','fa-tags','',0,'','false','false'),(73,'关闭服务端',64,'stop_game','',3,'',2,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:19:22','stop_game','stop_recycle','fa-tags','',0,'','false','false'),(74,'重启服务端',64,'restart_game','',3,'',4,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:19:24','restart_game','get_server_conf','fa-tags','',0,'','false','false'),(75,'前端',100,'client','',2,'',3,'admin','2022-04-08 07:35:38','admin','2022-10-10 12:30:20','client','client','fa-tags','',0,'','false','false'),(77,'热更配置',72,'server_hot','',3,'',1,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:16:03','server_hot','server_hot','fa-tags','',0,'','false','false'),(78,'闪断更新[含开关服]',72,'server_program','',3,'',3,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:19:13','server_program','server_cmd','fa-tags','',0,'','false','false'),(79,'日常更新',72,'daily_update','',3,'',5,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:19:15','daily_update','server_program','fa-tags','',0,'','false','false'),(80,'PHP',100,'php','',2,'',4,'admin','2022-04-08 07:35:38','admin','2022-10-10 12:30:20','php','php','fa-tags','',0,'','false','false'),(81,'执行服务端命令',72,'server_cmd','',3,'',2,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:19:12','server_cmd','server_update','fa-tags','',0,'','false','false'),(82,'更新服务端[含开关服]',72,'server_update','',3,'',4,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:19:14','server_update','daily_update','fa-tags','',0,'','false','false'),(83,'更新前端',75,'client_hot','',3,'',1,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:19:04','client_hot','client_hot','fa-tags','',0,'','false','false'),(84,'执行PHP命令',80,'php_cmd','',3,'',2,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:18:53','php_cmd','update_php','fa-tags','',0,'','false','false'),(86,'数据库',100,'sql','',2,'',5,'admin','2022-04-08 07:35:38','admin','2022-10-10 12:30:23','sql','sql','fa-tags','',0,'','false','false'),(87,'更新PHP',80,'update_php','',3,'',1,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:16:03','update_php','php_cmd','fa-tags','',0,'','false','false'),(88,'添加操作',35,'/taskMng/addTaskQueue','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-09-28 14:37:27','addTaskQueue/:Type(\\w+)/:Id(\\w+)','/taskMng/addTaskQueue/index','fa-tags','',0,'game','true','false'),(89,'同步数据库结构',86,'sync_db','',3,'',1,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:18:56','sync_db','sync_db','fa-tags','',0,'','false','false'),(90,'备份数据库',86,'backup_db','',3,'',2,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:18:57','backup_db','backup_db','fa-tags','',0,'','false','false'),(91,'执行SQL文件',86,'exec_sql_file','',3,'',3,'admin','2022-04-08 07:35:38','admin','2022-07-15 17:18:59','exec_sql_file','exec_sql_file','fa-tags','',0,'','false','false'),(93,'WEB终端',25,'/asset/web/list','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-07-25 16:56:24','webList','/asset/web/index','fa-tags','',0,'','false','true'),(94,'查看操作日志',35,'/taskMng/showTaskLog','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-09-28 14:37:29','showTaskLog/:taskId(\\w+)','/taskMng/showTaskLog/index','fa-tags','',0,'game','false','false'),(95,'终端日志',8,'/log/terminalLog/list','',1,'',3,'admin','2021-02-26 06:45:05','admin','2021-02-26 14:45:05','terminalLogList','log/terminallog/index','form','',0,'','false','true'),(96,'容器管理',0,'/dockerMng','',0,'',99,'admin','2021-02-26 06:45:05','admin','2022-08-30 14:44:13','/dockerMng','Layout','guide','/dcokerMng/dockerMngList',0,'','false','true'),(97,'容器列表',96,'/dockerMng/docker/list','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-08-25 10:44:16','dockerMngList','/dockerMng/dockerMng/index','fa-tags','',0,'game','false','true'),(98,'配置管理',29,'/yunwei/configFileDelivery/list','',1,'',6,'admin','2022-04-08 07:35:38','admin','2022-04-08 15:39:31','configFileDeliveryList','/yunwei/configFileDelivery/index','fa-tags','',0,'cluster','false','true'),(99,'装服日志',35,'/taskMng/installLog/list','',1,'',5,'admin','2022-04-08 07:35:38','admin','2022-10-10 12:29:55','installLogList','/taskMng/installLog/index','fa-tags','',0,'','false','true'),(100,'操作类型',35,'/taskMng/taskQueue/types','',1,'',3,'admin','2022-04-08 07:35:38','admin','2022-10-10 12:29:50','taskQueueTypes','/taskMng/taskQueue/Types','fa-tags','',0,'','false','false'),(101,'执行SQL语句',86,'export_sql','',3,'',3,'admin','2022-04-08 07:35:38','admin','2022-10-18 20:43:41','export_sql','export_sql','fa-tags','',0,'','false','false');
/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role`
--

DROP TABLE IF EXISTS `sys_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '角色名称',
  `remark` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `last_update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '更新人',
  `last_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `del_flag` tinyint DEFAULT '0' COMMENT '是否删除  1：已删除  0：正常',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='角色管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role`
--

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_menu`
--

DROP TABLE IF EXISTS `sys_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `role_id` bigint DEFAULT NULL COMMENT '角色ID',
  `menu_id` bigint DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='角色菜单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_menu`
--

LOCK TABLES `sys_role_menu` WRITE;
/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_stgroup`
--

DROP TABLE IF EXISTS `sys_stgroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_stgroup` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `st_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '策略名称',
  `st_json` json DEFAULT NULL COMMENT '策略json',
  `st_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
  `create_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `last_update_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `del_flag` smallint DEFAULT '0' COMMENT '是否删除  1：已删除  0：正常',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='策略组信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_stgroup`
--

LOCK TABLES `sys_stgroup` WRITE;
/*!40000 ALTER TABLE `sys_stgroup` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_stgroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_stgroup_ugroup`
--

DROP TABLE IF EXISTS `sys_stgroup_ugroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_stgroup_ugroup` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `stgroup_id` bigint NOT NULL COMMENT '决策组id',
  `ugroup_id` bigint NOT NULL COMMENT '用户组id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='策略组与用户组关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_stgroup_ugroup`
--

LOCK TABLES `sys_stgroup_ugroup` WRITE;
/*!40000 ALTER TABLE `sys_stgroup_ugroup` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_stgroup_ugroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_stgroup_user`
--

DROP TABLE IF EXISTS `sys_stgroup_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_stgroup_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `stgroup_id` bigint DEFAULT NULL COMMENT '策略组id',
  `user_id` bigint DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_stgroup_user`
--

LOCK TABLES `sys_stgroup_user` WRITE;
/*!40000 ALTER TABLE `sys_stgroup_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_stgroup_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_strategy`
--

DROP TABLE IF EXISTS `sys_strategy`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_strategy` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `st_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '名称',
  `st_sort` smallint DEFAULT '0' COMMENT '排序',
  `st_pid` smallint DEFAULT NULL COMMENT '上级pid',
  `st_level` smallint DEFAULT NULL COMMENT '授权值',
  `st_urls` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'urls',
  `st_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `st_is_auth` tinyint DEFAULT '1' COMMENT '是否鉴权 是：1 否 ：2',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=234 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='策略明细';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_strategy`
--

LOCK TABLES `sys_strategy` WRITE;
/*!40000 ALTER TABLE `sys_strategy` DISABLE KEYS */;
INSERT INTO `sys_strategy` VALUES (1,'admin',0,-1,1,'/admin','系统模块',1),(2,'user',0,1,2,'/admin/user','用户服务',1),(3,'updateUserStatus',0,2,3,'/admin/user/updateUserStatus','更改用户状态',1),(4,'add',0,2,3,'/admin/user/add','用户新增',1),(5,'currentUser',0,2,3,'/admin/user/currentUser','用户信息',2),(6,'delete',0,2,3,'/admin/user/delete','用户删除',1),(7,'edit',0,2,3,'/admin/user/edit','根据id获取对应用户数据',2),(8,'getRouters',0,2,3,'/admin/user/getRouters','获取个人路由信息',2),(9,'getUserAssignmentPolicy',0,2,3,'/admin/user/getUserAssignmentPolicy','获取用户已选策略',2),(10,'list',0,2,3,'/admin/user/list','用户列表',1),(11,'login',0,2,3,'/admin/user/login','用户登录',2),(12,'logout',0,2,3,'/admin/user/logout','用户退出',2),(13,'reSetPassword',0,2,3,'/admin/user/reSetPassword','重置用户密码',1),(14,'selectAllData',0,2,3,'/admin/user/selectAllData','用户相关所有数据',2),(15,'update',0,2,3,'/admin/user/update','用户更新',1),(16,'updatePersonalData',0,2,3,'/admin/user/updatePersonalData','修改个人用户相关信息',2),(17,'updatePersonalPasswordData',0,2,3,'/admin/user/updatePersonalPasswordData','修改个人密码信息',2),(18,'userAssignmentPolicy',0,2,3,'/admin/user/userAssignmentPolicy','分配用户策略',1),(19,'ugroup',0,1,2,'/admin/ugroup','用户组服务',1),(20,'add',0,19,3,'/admin/ugroup/add','用户组新增',1),(21,'delete',0,19,3,'/admin/ugroup/delete','用户组删除',1),(22,'getUgroupAssignmentPolicy',0,19,3,'/admin/ugroup/getUgroupAssignmentPolicy','获取用户组已选策略',2),(23,'list',0,19,3,'/admin/ugroup/list','用户组列表',1),(24,'ugroupAssignmentPolicy',0,19,3,'/admin/ugroup/ugroupAssignmentPolicy','分配用户组策略',1),(25,'update',0,19,3,'/admin/ugroup/update','用户组更新',1),(26,'stgroup',0,1,2,'/admin/stgroup','策略服务',1),(27,'add',0,26,3,'/admin/stgroup/add','策略组新增',1),(28,'delete',0,26,3,'/admin/stgroup/delete','策略组删除',1),(29,'getDistributionConfig',0,26,3,'/admin/stgroup/getDistributionConfig','获取策略相关信息根据id',2),(30,'list',0,26,3,'/admin/stgroup/list','策略组列表',1),(31,'policyAssociatedUsers',0,26,3,'/admin/stgroup/policyAssociatedUsers','策略关联用户或组提交',1),(32,'update',0,26,3,'/admin/stgroup/update','策略组更新',1),(33,'role',0,1,2,'/admin/role','角色服务',1),(34,'add',0,33,3,'/admin/role/add','角色新增',1),(35,'delete',0,33,3,'/admin/role/delete','角色删除',1),(36,'list',0,33,3,'/admin/role/list','角色列表',1),(37,'queryMenuByRoleId',0,33,3,'/admin/role/queryMenuByRoleId','根据角色id查询菜单',2),(38,'update',0,33,3,'/admin/role/update','角色更新',1),(39,'updateRoleMenu',0,33,3,'/admin/role/updateRoleMenu','更新角色菜单',1),(40,'project',0,1,2,'/admin/project','项目服务',1),(41,'add',0,40,3,'/admin/project/add','项目新增',1),(42,'delete',0,40,3,'/admin/project/delete','项目删除',1),(43,'list',0,40,3,'/admin/project/list','项目列表',2),(44,'update',0,40,3,'/admin/project/update','项目更新',1),(45,'menu',0,1,2,'/admin/menu','菜单服务',2),(46,'label',0,1,2,'/admin/label','标签服务',1),(47,'dept',0,1,2,'/admin/dept','部门服务',1),(48,'dict',0,1,2,'/admin/dict','字典服务',1),(49,'company',0,1,2,'/admin/company','公司服务',1),(50,'add',0,49,3,'/admin/company/add','公司新增',1),(51,'delete',0,49,3,'/admin/company/delete','公司删除',1),(52,'list',0,49,3,'/admin/company/list','公司列表',2),(53,'update',0,49,3,'/admin/company/update','公司更新',1),(54,'add',0,47,3,'/admin/dept/add','部门新增',1),(55,'delete',0,47,3,'/admin/dept/delete','部门删除',1),(56,'list',0,47,3,'/admin/dept/list','部门列表',2),(57,'treeselect',0,47,3,'/admin/dept/treeselect','部门生成树',2),(58,'update',0,47,3,'/admin/dept/update','部门更新',1),(59,'add',0,48,3,'/admin/dict/add','字典新增',1),(60,'delete',0,48,3,'/admin/dict/delete','字典删除',1),(61,'list',0,48,3,'/admin/dict/list','字典列表',2),(62,'type',0,48,3,'/admin/dict/type','根据字典类型获取',2),(63,'update',0,48,3,'/admin/dict/update','字典更新',1),(64,'add',0,46,3,'/admin/label/add','标签新增',1),(65,'delete',0,46,3,'/admin/label/delete','标签删除',1),(66,'list',0,46,3,'/admin/label/list','标签列表',2),(67,'update',0,46,3,'/admin/label/update','标签更新',1),(68,'add',0,45,3,'/admin/menu/add','菜单新增',2),(69,'delete',0,45,3,'/admin/menu/delete','菜单删除',2),(70,'list',0,45,3,'/admin/menu/list','菜单列表',2),(71,'update',0,45,3,'/admin/menu/update','菜单更新',2),(72,'log',0,1,2,'/admin/log','日志服务',1),(73,'loginlog',0,72,3,'/admin/log/loginlog','登录日志列表',1),(74,'syslog',0,72,3,'/admin/log/syslog','系统日志列表',1),(75,'edit',0,26,3,'/admin/stgroup/edit','根据ID策略获取',1),(76,'resource',0,1,2,'/admin/resource','资源服务',1),(77,'add',0,76,3,'/admin/resource/add','资源新增',1),(78,'batchDelete',0,76,3,'/admin/resource/batchDelete','资源删除',1),(79,'list',0,76,3,'/admin/resource/list','获取所有资源类型',2),(81,'rsourceObjectValueList',0,76,3,'/admin/resource/rsourceObjectValueList','根据标签条件查询所需资源',2),(82,'yunwei',0,-1,1,'/yunwei','运维模块',1),(83,'asset',0,82,2,'/yunwei/asset','资产服务',1),(84,'add',0,83,3,'/yunwei/asset/add','资产新增',1),(85,'delete',0,83,3,'/yunwei/asset/delete','资产删除',1),(86,'update',0,83,3,'/yunwei/asset/update','资产修改',1),(87,'list',0,83,3,'/yunwei/asset/list','资产列表',1),(120,'keyManage',0,82,2,'/yunwei/keyManage','密钥管理',2),(121,'add',0,120,3,'/yunwei/keyManage/add','密钥管理增加',2),(122,'delete',0,120,3,'/yunwei/keyManage/delete','密钥管理删除',2),(123,'get',0,120,3,'/yunwei/keyManage/get','密钥管理获取',2),(124,'list',0,120,3,'/yunwei/keyManage/list','密钥管理列表',2),(125,'update',0,120,3,'/yunwei/keyManage/update','密钥管理更新',2),(126,'platform',0,82,2,'/yunwei/platform','平台管理',1),(127,'add',0,126,3,'/yunwei/platform/add','平台管理增加',1),(128,'delete',0,126,3,'/yunwei/platform/delete','平台管理删除',1),(129,'get',0,126,3,'/yunwei/platform/get','平台管理获取',2),(130,'list',0,126,3,'/yunwei/platform/list','平台管理列表',1),(131,'update',0,126,3,'/yunwei/platform/update','平台管理更新',1),(132,'featureServer',0,82,2,'/yunwei/featureServer','功能服',1),(133,'add',0,132,3,'/yunwei/featureServer/add','功能服增加',1),(134,'delete',0,132,3,'/yunwei/featureServer/delete','功能服删除',1),(135,'get',0,132,3,'/yunwei/featureServer/get','功能服获取',2),(136,'list',0,132,3,'/yunwei/featureServer/list','功能服列表',1),(137,'update',0,132,3,'/yunwei/featureServer/update','功能服更新',1),(138,'maintainPlan',0,200,2,'/yunying/maintainPlan','维护计划',1),(139,'add',0,138,3,'/yunying/maintainPlan/add','维护计划增加',1),(140,'delete',0,138,3,'/yunying/maintainPlan/delete','维护计划删除',1),(141,'get',0,138,3,'/yunying/maintainPlan/get','维护计划获取',2),(142,'list',0,138,3,'/yunying/maintainPlan/list','维护计划列表',1),(143,'update',0,138,3,'/yunying/maintainPlan/update','维护计划更新',1),(144,'mergePlan',0,200,2,'/yunying/mergePlan','合服计划',1),(145,'add',0,144,3,'/yunying/mergePlan/add','合服计划增加',1),(146,'delete',0,144,3,'/yunying/mergePlan/delete','合服计划删除',1),(147,'get',0,144,3,'/yunying/mergePlan/get','合服计划获取',2),(148,'list',0,144,3,'/yunying/mergePlan/list','合服计划列表',1),(149,'update',0,144,3,'/yunying/mergePlan/update','合服计划更新',1),(150,'openPlan',0,200,2,'/yunying/openPlan','开服计划',1),(151,'add',0,150,3,'/yunying/openPlan/add','开服计划增加',1),(152,'delete',0,150,3,'/yunying/openPlan/delete','开服计划删除',1),(153,'get',0,150,3,'/yunying/openPlan/get','开服计划获取',2),(154,'list',0,150,3,'/yunying/openPlan/list','开服计划列表',1),(155,'update',0,150,3,'/yunying/openPlan/update','开服计划更新',1),(156,'taskQueue',0,201,2,'/taskMng/taskQueue','任务管理',1),(157,'add',0,156,3,'/taskMng/taskQueue/add','任务管理增加',1),(158,'delete',0,156,3,'/taskMng/taskQueue/delete','任务管理删除',1),(159,'get',0,156,3,'/taskMng/taskQueue/get','任务管理获取',2),(160,'list',0,156,3,'/taskMng/taskQueue/list','任务管理列表',1),(161,'update',0,156,3,'/taskMng/taskQueue/update','任务管理更新',1),(162,'file',0,1,2,'/admin/file','文件服务',2),(163,'download',0,162,3,'/admin/file/download','下载文件',2),(164,'upload',0,162,3,'/admin/file/upload','上传文件',2),(165,'search',0,1,2,'/admin/search','搜索服务',2),(166,'article',0,165,3,'/admin/search/article','全局搜索',2),(167,'strategy',0,1,2,'/admin/strategy','策略服务',2),(168,'list',0,167,3,'/admin/strategy/list','策略列表',2),(169,'taskLog',0,201,2,'/taskMng/taskLog','任务日志',1),(170,'list',0,169,3,'/taskMng/taskLog/list','任务日志列表',1),(171,'get',0,169,3,'/taskMng/taskLog/get','任务日志详情',2),(172,'cluster',0,82,2,'/yunwei/cluster','集群服务',1),(173,'gameServer',0,82,2,'/yunwei/gameServer','游戏服服务',1),(174,'hosts',0,82,2,'/yunwei/hosts','服务器服务',1),(175,'list',0,172,3,'/yunwei/cluster/list','集群列表',1),(176,'list',0,173,3,'/yunwei/gameServer/list','游戏服列表',1),(177,'list',0,174,3,'/yunwei/hosts/list','服务器列表',1),(178,'getTaskOperationList',0,156,3,'/taskMng/taskQueue/getTaskOperationList','获取当前用户操作列表',2),(179,'getWebSocketAddr',0,156,3,'/taskMng/taskQueue/getWebSocketAddr','获取websokcet地址',2),(180,'start',0,156,3,'/taskMng/taskQueue/start','任务管理开始',1),(181,'stop',0,156,3,'/taskMng/taskQueue/stop','任务管理停止',1),(182,'assetInfoData',0,83,3,'/yunwei/asset/assetInfoData','项目公司视图',2),(183,'ownerProjectData',0,83,3,'/yunwei/asset/ownerProjectData','个人项目数据',2),(184,'recycle',0,83,3,'/yunwei/asset/recycle','资产回收',1),(185,'detail',0,126,3,'/yunwei/platform/detail','平台详情',1),(186,'projects',0,126,3,'/yunwei/platform/project','根据项目ID取平台',2),(187,'mergePlanRangeList',0,144,3,'/yunying/mergePlan/mergePlanRangeList','获取合服范围',2),(188,'getMaintanListByPri',0,138,3,'/yunying/maintainPlan/getMaintanListByPri','根据项目获取维护计划',2),(189,'maintainGetClusterInfo',0,138,3,'/yunying/maintainPlan/maintainGetClusterInfo','根据项目获取集群信息',2),(190,'maintainPlanRangeList',0,138,3,'/yunying/maintainPlan/maintainPlanRangeList','维护计划范围获取',2),(191,'terminallog',0,72,3,'/admin/log/terminallog','终端日志列表',1),(192,'download',0,72,3,'/admin/log/download','终端日志下载',1),(193,'taskLogHistroy',0,201,2,'/taskMng/taskLogHistroy','任务历史日志',1),(194,'list',0,193,3,'/taskMng/taskLogHistroy/list','任务历史日志列表',1),(195,'detail',0,193,3,'/taskMng/taskLogHistroy/detail','任务历史日志详情',1),(196,'getWebSshTree',0,83,3,'/yunwei/asset/getWebSshTree','webssh树形',2),(197,'assetFileDownload',0,83,3,'/yunwei/asset/assetFileDownload','文件管理器下载',1),(198,'assetFileUpload',0,83,3,'/yunwei/asset/assetFileUpload','文件管理器上传',1),(199,'assetFile',0,83,3,'/yunwei/asset/assetFile','文件管理器列表',1),(200,'yunying',0,-1,1,'/yunying','运营模块',1),(201,'taskMng',0,-1,1,'/taskMng','任务模块',1),(202,'getRoleAssignmentUser',0,33,3,'/admin/role/getRoleAssignmentUser','获取角色分配用户',2),(203,'roleAssignmentUser',0,33,3,'/admin/role/roleAssignmentUser','更新角色分配用户',1),(204,'getUgroupAssignmentUser',0,19,3,'/admin/ugroup/getUgroupAssignmentUser','获取用户组分配用户',2),(205,'ugroupAssignmentUser',0,19,3,'/admin/ugroup/ugroupAssignmentUser','更新用户组分配用户',1),(206,'assetBatchDistribute',0,83,3,'/yunwei/asset/assetBatchDistribute','资产批量操作',1),(207,'get',0,83,3,'/yunwei/asset/get','资产获取',2),(208,'userBatchEditItems',0,2,3,'/admin/user/userBatchEditItems','批量分配（删除）用户项目',1),(215,'configFileDelivery',0,82,2,'/yunwei/configFileDelivery','配置文件管理',1),(216,'add',0,215,3,'/yunwei/configFileDelivery/add','配置文件下发',1),(217,'configFileDeliveryGet',0,215,3,'/yunwei/configFileDelivery/configFileDeliveryGet','配置文件下发获取服务器列表',1),(218,'configFileDeliveryGetLog',0,215,3,'/yunwei/configFileDelivery/configFileDeliveryGetLog','配置文件下发获取状态日志',1),(219,'list',0,215,3,'/yunwei/configFileDelivery/list','配置文件下发列表',1),(220,'configFileDeliveryGetFileContent',0,215,3,'/yunwei/configFileDelivery/configFileDeliveryGetFileContent','配置文件下发获取模版或生成文件内容',1),(221,'configFileDeliveryUpdateTemplate',0,215,3,'/yunwei/configFileDelivery/configFileDeliveryUpdateTemplate','配置文件下发修改模版文件内容',1),(222,'detail',0,172,3,'/yunwei/cluster/detail','集群详情',1),(223,'updateSupplyCompany',0,49,3,'/admin/company/updateSupplyCompany','出机方状态更新',1),(225,'taskGetFormatJson',0,156,3,'/taskMng/taskQueue/taskGetFormatJson','格式化任务',2),(226,'refreshProfileList',0,215,3,'/yunwei/configFileDelivery/refreshProfileList','一键刷新配置文件',1),(227,'mergeCheckServerRange',0,138,3,'/yunying/mergePlan/mergeCheckServerRange','检查合服范围',2),(228,'getInstallLogList',0,156,3,'/taskMng/installLogList/getInstallLogList','获取装服日志列表',2),(229,'help',0,82,2,'/yunwei/help','帮助中心',2),(230,'list',0,229,3,'/yunwei/help/list','小助手更新模板',2),(231,'getGameTrendChart',0,229,3,'/yunwei/dashboard/getGameTrendChart','游戏服趋势图',2),(232,'getServerTrendChart',0,229,3,'/yunwei/dashboard/getServerTrendChart','服务器趋势图',2),(233,'getSumOfCurrentInfor',0,229,3,'/yunwei/dashboard/getSumOfCurrentInfor','当前信息汇总',2);
/*!40000 ALTER TABLE `sys_strategy` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_ugroup`
--

DROP TABLE IF EXISTS `sys_ugroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_ugroup` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '组id',
  `ug_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '组名',
  `ug_json` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '组备注名',
  `create_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `last_update_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
  `last_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `del_flag` smallint DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户组信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_ugroup`
--

LOCK TABLES `sys_ugroup` WRITE;
/*!40000 ALTER TABLE `sys_ugroup` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_ugroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户名',
  `nick_name` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '头像',
  `password` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '密码',
  `salt` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '加密盐',
  `email` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '邮箱',
  `mobile` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '手机号',
  `status` tinyint DEFAULT '1' COMMENT '状态  1：正常 2：禁用 ',
  `dept_id` bigint DEFAULT NULL COMMENT '机构ID',
  `create_by` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `last_update_by` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '更新人',
  `last_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `del_flag` tinyint DEFAULT '0' COMMENT '是否删除  1：已删除  0：正常',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='用户信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user`
--

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;
INSERT INTO `sys_user` VALUES (1,'admin','超级管理员','','cfb97aaee68e51613a714eb287c0e2ea','abcd12','admin@qq.com','12345678903',1,3,'admin','2018-08-14 03:11:11','admin','2022-10-20 15:27:23',0),(3,'zhangqingxian','张清贤','','94d2ae2072c90023fbf6abcbe97394aa','sV0AB8','1062829477@qq.com','13535427611',1,3,'admin','2022-09-05 07:40:45','zhangqingxian','2022-10-21 08:52:39',0),(5,'jiayuanhao','贾源皓','','35a1e476242738fe647876fe2abab58f','dDWRLe','459105919@qq.com','18620261315',1,3,'zhangqingxian','2022-09-05 08:27:57','zhangqingxian','2022-10-24 09:26:28',0),(6,'dengxiguang','邓锡光','','fed19e33ef25dcf534e1c3c6ec3e676a','KbpTr1','963853870@qq.com','1213123',1,3,'zhangqingxian','2022-09-05 08:29:26','zhangqingxian','2022-10-26 19:45:19',0);
/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_project`
--

DROP TABLE IF EXISTS `sys_user_project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_project` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户id',
  `project_id` bigint NOT NULL COMMENT '项目id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_project`
--

LOCK TABLES `sys_user_project` WRITE;
/*!40000 ALTER TABLE `sys_user_project` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_user_project` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_role`
--

DROP TABLE IF EXISTS `sys_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `role_id` bigint NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_role`
--

LOCK TABLES `sys_user_role` WRITE;
/*!40000 ALTER TABLE `sys_user_role` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_user_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_ugroup`
--

DROP TABLE IF EXISTS `sys_user_ugroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_ugroup` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `ugroup_id` bigint NOT NULL COMMENT '用户组id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户和用户组关联';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_ugroup`
--

LOCK TABLES `sys_user_ugroup` WRITE;
/*!40000 ALTER TABLE `sys_user_ugroup` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_user_ugroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `task_log_histroy`
--

DROP TABLE IF EXISTS `task_log_histroy`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `task_log_histroy` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '日志id',
  `tasks_id` int DEFAULT '0' COMMENT '任务id',
  `tasks_time` int DEFAULT '0' COMMENT '时间戳',
  `tasks_logs` longtext CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '任务日志',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `task_id` (`tasks_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='任务历史记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task_log_histroy`
--

LOCK TABLES `task_log_histroy` WRITE;
/*!40000 ALTER TABLE `task_log_histroy` DISABLE KEYS */;
/*!40000 ALTER TABLE `task_log_histroy` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tasks`
--

DROP TABLE IF EXISTS `tasks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tasks` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务id',
  `project_id` int NOT NULL DEFAULT '0' COMMENT '项目id',
  `cluster_id` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '集群ids',
  `task_type` smallint NOT NULL DEFAULT '0' COMMENT '任务类型 1：临时维护 2：日常维护',
  `level` smallint NOT NULL DEFAULT '0' COMMENT '任务层级',
  `maintain_id` int NOT NULL DEFAULT '0' COMMENT '关联日常维护计划表',
  `name` longtext CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '任务名称',
  `types` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '任务操作类型',
  `pid` int NOT NULL DEFAULT '0' COMMENT '任务pid',
  `task_start_time` int DEFAULT NULL COMMENT '任务开始时间',
  `task_end_time` int DEFAULT '0' COMMENT '任务结束时间',
  `task_exec_time` int DEFAULT '0' COMMENT '任务执行时间',
  `cmd` longtext CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '任务命令',
  `content` longtext CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '任务执行日志',
  `task_status` smallint DEFAULT '-1' COMMENT '任务状态 -1：未开始 1：执行中 2：执行失败 3：执行成功 4：取消任务 5：失败删除 6：删除',
  `task_step` smallint DEFAULT '0' COMMENT '任务步骤',
  `outer_ip` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '外网源IP',
  `create_by` int DEFAULT '0' COMMENT '创建者',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` int DEFAULT '0' COMMENT '修改者',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `export_file_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '导出文件名',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT '' COMMENT '任务备注',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='任务管理表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tasks`
--

LOCK TABLES `tasks` WRITE;
/*!40000 ALTER TABLE `tasks` DISABLE KEYS */;
/*!40000 ALTER TABLE `tasks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tasks_tid_pid`
--

DROP TABLE IF EXISTS `tasks_tid_pid`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tasks_tid_pid` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tid` int DEFAULT '0' COMMENT '任务id',
  `pid` int DEFAULT '0' COMMENT '进程pid',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='任务id和进程id对应表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tasks_tid_pid`
--

LOCK TABLES `tasks_tid_pid` WRITE;
/*!40000 ALTER TABLE `tasks_tid_pid` DISABLE KEYS */;
/*!40000 ALTER TABLE `tasks_tid_pid` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `view_asset_config_file`
--

DROP TABLE IF EXISTS `view_asset_config_file`;
/*!50001 DROP VIEW IF EXISTS `view_asset_config_file`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_asset_config_file` AS SELECT 
 1 AS `view_asset_id`,
 1 AS `view_outer_ip`,
 1 AS `view_inner_ip`,
 1 AS `view_en_host_role`,
 1 AS `view_provider_name_cn`,
 1 AS `view_host_sort`,
 1 AS `view_asset_describe`,
 1 AS `view_ssh_port`,
 1 AS `view_user_project_id`,
 1 AS `view_user_project_cn`,
 1 AS `view_user_project_en`,
 1 AS `view_cluster_name`,
 1 AS `view_config_file_id`,
 1 AS `view_config_file_name`,
 1 AS `view_config_dest_path`,
 1 AS `view_file_mod_time`,
 1 AS `view_config_time`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_assets`
--

DROP TABLE IF EXISTS `view_assets`;
/*!50001 DROP VIEW IF EXISTS `view_assets`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_assets` AS SELECT 
 1 AS `view_asset_id`,
 1 AS `view_outer_ip`,
 1 AS `view_inner_ip`,
 1 AS `view_accelerate_domain`,
 1 AS `view_host_role_id`,
 1 AS `view_host_role_cn`,
 1 AS `view_en_host_role`,
 1 AS `view_provider_id`,
 1 AS `view_provider_name_en`,
 1 AS `view_provider_name_cn`,
 1 AS `view_hardware_info`,
 1 AS `view_ssh_port`,
 1 AS `view_init_type`,
 1 AS `view_init_type_cn`,
 1 AS `view_clean_type`,
 1 AS `view_clean_type_cn`,
 1 AS `view_recycle_type`,
 1 AS `view_recycle_type_cn`,
 1 AS `view_init_login_info`,
 1 AS `view_change_status_remark`,
 1 AS `view_remark`,
 1 AS `view_asset_create_time`,
 1 AS `view_asset_update_time`,
 1 AS `view_asset_del_flag`,
 1 AS `view_pr_id`,
 1 AS `view_asset_ownership_company_id`,
 1 AS `view_asset_ownership_company_cn`,
 1 AS `view_asset_ownership_company_en`,
 1 AS `view_asset_ownership_company_deleted`,
 1 AS `view_server_affiliation_deleted`,
 1 AS `view_user_company_id`,
 1 AS `view_user_company_cn`,
 1 AS `view_user_company_en`,
 1 AS `view_user_company_deleted`,
 1 AS `view_user_project_id`,
 1 AS `view_user_project_cn`,
 1 AS `view_user_project_en`,
 1 AS `view_user_project_deleted`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_company_project`
--

DROP TABLE IF EXISTS `view_company_project`;
/*!50001 DROP VIEW IF EXISTS `view_company_project`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_company_project` AS SELECT 
 1 AS `view_company_id`,
 1 AS `view_company_cn`,
 1 AS `view_company_en`,
 1 AS `view_company_del_flag`,
 1 AS `view_pr_id`,
 1 AS `view_project_id`,
 1 AS `view_project_cn`,
 1 AS `view_project_en`,
 1 AS `view_dept_id`,
 1 AS `view_dept_name`,
 1 AS `view_project_type`,
 1 AS `view_project_type_cn`,
 1 AS `view_group_qq`,
 1 AS `view_group_type_cn`,
 1 AS `view_group_type_en`,
 1 AS `view_group_dev_qq`,
 1 AS `view_project_del_flag`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_feature_label`
--

DROP TABLE IF EXISTS `view_feature_label`;
/*!50001 DROP VIEW IF EXISTS `view_feature_label`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_feature_label` AS SELECT 
 1 AS `view_feature_server_id`,
 1 AS `view_project_id`,
 1 AS `view_project_en`,
 1 AS `view_project_cn`,
 1 AS `view_feature_server_info`,
 1 AS `view_feature_server_remark`,
 1 AS `view_label_id`,
 1 AS `view_label_type`,
 1 AS `view_label_name`,
 1 AS `view_label_values`,
 1 AS `view_label_remark`,
 1 AS `view_project_del_flag`,
 1 AS `view_feature_server_del_flag`,
 1 AS `view_label_del_flag`,
 1 AS `view_table_name`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_ip_list`
--

DROP TABLE IF EXISTS `view_ip_list`;
/*!50001 DROP VIEW IF EXISTS `view_ip_list`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_ip_list` AS SELECT 
 1 AS `view_gatfile`,
 1 AS `view_placeholder_1`,
 1 AS `view_server_id`,
 1 AS `view_platform_en`,
 1 AS `view_cluster_name`,
 1 AS `view_server_alias`,
 1 AS `view_outer_ip`,
 1 AS `view_ssh_port`,
 1 AS `view_inner_ip`,
 1 AS `view_placeholder_2`,
 1 AS `view_open_time`,
 1 AS `view_provider_name_en`,
 1 AS `view_new_platform_en`,
 1 AS `view_project_en`,
 1 AS `view_unix_timestamp_opentime`,
 1 AS `view_server_status`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_label_info`
--

DROP TABLE IF EXISTS `view_label_info`;
/*!50001 DROP VIEW IF EXISTS `view_label_info`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_label_info` AS SELECT 
 1 AS `view_label_id`,
 1 AS `view_label_type`,
 1 AS `view_label_name`,
 1 AS `view_label_values`,
 1 AS `view_label_remark`,
 1 AS `view_stop_status`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_platform_feature`
--

DROP TABLE IF EXISTS `view_platform_feature`;
/*!50001 DROP VIEW IF EXISTS `view_platform_feature`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_platform_feature` AS SELECT 
 1 AS `view_project_id`,
 1 AS `view_project_en`,
 1 AS `view_project_cn`,
 1 AS `view_platform_autoid`,
 1 AS `view_platform_id`,
 1 AS `view_platform_en`,
 1 AS `view_platform_cn`,
 1 AS `view_domain_format`,
 1 AS `view_platform_remark`,
 1 AS `view_cluster_name`,
 1 AS `view_cluster_platform_group`,
 1 AS `view_cluster_label_id`,
 1 AS `view_cluster_label_name`,
 1 AS `view_cluster_label_values`,
 1 AS `view_labels`,
 1 AS `view_cluster_feature_info`,
 1 AS `view_platform_feature_info`,
 1 AS `view_feature_info`,
 1 AS `view_center_cluster_info`,
 1 AS `view_center_platform_info`,
 1 AS `view_login_cluster_info`,
 1 AS `view_login_platform_info`,
 1 AS `view_pay_cluster_info`,
 1 AS `view_pay_platform_info`,
 1 AS `view_source_cluster_info`,
 1 AS `view_source_platform_info`,
 1 AS `view_cdn_cluster_info`,
 1 AS `view_cdn_platform_info`,
 1 AS `view_cross_cluster_info`,
 1 AS `view_cross_platform_info`,
 1 AS `view_db_cluster_info`,
 1 AS `view_db_platform_info`,
 1 AS `view_gate_cluster_info`,
 1 AS `view_gate_platform_info`,
 1 AS `view_backup_cluster_info`,
 1 AS `view_backup_platform_info`,
 1 AS `view_project_del_flag`,
 1 AS `view_platform_del_flag`,
 1 AS `view_feature_server_del_flag`,
 1 AS `view_label_del_flag`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_platform_label`
--

DROP TABLE IF EXISTS `view_platform_label`;
/*!50001 DROP VIEW IF EXISTS `view_platform_label`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_platform_label` AS SELECT 
 1 AS `view_platform_autoid`,
 1 AS `view_project_id`,
 1 AS `view_project_en`,
 1 AS `view_project_cn`,
 1 AS `view_platform_id`,
 1 AS `view_platform_en`,
 1 AS `view_platform_cn`,
 1 AS `view_domain_format`,
 1 AS `view_platform_remark`,
 1 AS `view_platform_group_1`,
 1 AS `view_platform_group`,
 1 AS `view_args_platform_group`,
 1 AS `view_label_id`,
 1 AS `view_label_type`,
 1 AS `view_label_name`,
 1 AS `view_label_values`,
 1 AS `view_label_remark`,
 1 AS `view_project_del_flag`,
 1 AS `view_platform_del_flag`,
 1 AS `view_label_del_flag`,
 1 AS `view_table_name`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_project_cluster_info`
--

DROP TABLE IF EXISTS `view_project_cluster_info`;
/*!50001 DROP VIEW IF EXISTS `view_project_cluster_info`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_project_cluster_info` AS SELECT 
 1 AS `view_project_id`,
 1 AS `view_project_en`,
 1 AS `view_project_cn`,
 1 AS `view_feature_table_name`,
 1 AS `view_label_id`,
 1 AS `view_label_type`,
 1 AS `view_label_name`,
 1 AS `view_label_values`,
 1 AS `view_label_remark`,
 1 AS `view_cluster_name`,
 1 AS `view_cluster_platform_group`,
 1 AS `view_center_cluster_info`,
 1 AS `view_login_cluster_info`,
 1 AS `view_pay_cluster_info`,
 1 AS `view_source_cluster_info`,
 1 AS `view_cdn_cluster_info`,
 1 AS `view_cross_cluster_info`,
 1 AS `view_db_cluster_info`,
 1 AS `view_gate_cluster_info`,
 1 AS `view_backup_cluster_info`,
 1 AS `view_project_del_flag`,
 1 AS `view_feature_server_del_flag_abandon`,
 1 AS `view_label_del_flag`,
 1 AS `view_cluster_feature_info`,
 1 AS `view_feature_server_del_flag`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_search_label`
--

DROP TABLE IF EXISTS `view_search_label`;
/*!50001 DROP VIEW IF EXISTS `view_search_label`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_search_label` AS SELECT 
 1 AS `view_resource_cn_name`,
 1 AS `view_resource_en_name`,
 1 AS `view_resource_remark`,
 1 AS `view_project_id`,
 1 AS `view_primary_key`,
 1 AS `view_primary_key_value`,
 1 AS `view_resource_type`,
 1 AS `view_resource_value`,
 1 AS `view_data_content`,
 1 AS `view_data_url`,
 1 AS `view_json_id`,
 1 AS `view_table_name`,
 1 AS `view_show_cluster`,
 1 AS `view_show_feature`,
 1 AS `view_show_install`,
 1 AS `view_show_other`,
 1 AS `view_system_show`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_sh_platform_info`
--

DROP TABLE IF EXISTS `view_sh_platform_info`;
/*!50001 DROP VIEW IF EXISTS `view_sh_platform_info`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_sh_platform_info` AS SELECT 
 1 AS `view_project_id`,
 1 AS `view_project_en`,
 1 AS `view_project_cn`,
 1 AS `view_platform_en`,
 1 AS `view_platform_info`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_sh_split_ip_and_region`
--

DROP TABLE IF EXISTS `view_sh_split_ip_and_region`;
/*!50001 DROP VIEW IF EXISTS `view_sh_split_ip_and_region`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_sh_split_ip_and_region` AS SELECT 
 1 AS `view_project_id`,
 1 AS `view_project_en`,
 1 AS `view_project_cn`,
 1 AS `view_single_ip_pool`,
 1 AS `view_cross_ip_pool`,
 1 AS `view_split_single_region`,
 1 AS `view_split_cross_region`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_task_logs_relationship_with_id`
--

DROP TABLE IF EXISTS `view_task_logs_relationship_with_id`;
/*!50001 DROP VIEW IF EXISTS `view_task_logs_relationship_with_id`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_task_logs_relationship_with_id` AS SELECT 
 1 AS `v_log_inner`,
 1 AS `v_pid`,
 1 AS `v_id`,
 1 AS `v_step`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_user_asset`
--

DROP TABLE IF EXISTS `view_user_asset`;
/*!50001 DROP VIEW IF EXISTS `view_user_asset`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_user_asset` AS SELECT 
 1 AS `view_asset_id`,
 1 AS `view_user_id`,
 1 AS `view_cluster_name`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `view_user_strategy`
--

DROP TABLE IF EXISTS `view_user_strategy`;
/*!50001 DROP VIEW IF EXISTS `view_user_strategy`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_user_strategy` AS SELECT 
 1 AS `view_sys_user_id`,
 1 AS `view_sys_user_name`,
 1 AS `view_stgroup_st_json`*/;
SET character_set_client = @saved_cs_client;

--
-- Final view structure for view `view_asset_config_file`
--

/*!50001 DROP VIEW IF EXISTS `view_asset_config_file`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_asset_config_file` (`view_asset_id`,`view_outer_ip`,`view_inner_ip`,`view_en_host_role`,`view_provider_name_cn`,`view_host_sort`,`view_asset_describe`,`view_ssh_port`,`view_user_project_id`,`view_user_project_cn`,`view_user_project_en`,`view_cluster_name`,`view_config_file_id`,`view_config_file_name`,`view_config_dest_path`,`view_file_mod_time`,`view_config_time`) AS select `asset_file`.`view_asset_id` AS `view_asset_id`,`asset_file`.`view_outer_ip` AS `view_outer_ip`,`asset_file`.`view_inner_ip` AS `view_inner_ip`,`asset_file`.`view_en_host_role` AS `view_en_host_role`,`asset_file`.`view_provider_name_cn` AS `view_provider_name_cn`,(case `asset_file`.`view_en_host_role` when 'game' then '游戏服' when 'cross' then '跨服' else '其他' end) AS `host_sort`,concat(`asset_file`.`view_outer_ip`,'_',`asset_file`.`view_inner_ip`,'_',`asset_file`.`view_provider_name_cn`,'_',`asset_file`.`view_en_host_role`) AS `asset_describe`,`asset_file`.`view_ssh_port` AS `view_ssh_port`,`asset_file`.`view_user_project_id` AS `view_user_project_id`,`asset_file`.`view_user_project_cn` AS `view_user_project_cn`,`asset_file`.`view_user_project_en` AS `view_user_project_en`,`asset_file`.`view_cluster_name` AS `view_cluster_name`,`asset_file`.`config_file_id` AS `config_file_id`,`asset_file`.`name` AS `config_file_name`,`asset_file`.`dest_path` AS `config_dest_path`,`asset_file`.`file_mod_time` AS `file_mod_time`,`asset_file`.`config_time` AS `config_time` from (select `asset_config_file`.`view_asset_id` AS `view_asset_id`,`asset_config_file`.`view_outer_ip` AS `view_outer_ip`,`asset_config_file`.`view_inner_ip` AS `view_inner_ip`,`asset_config_file`.`view_accelerate_domain` AS `view_accelerate_domain`,`asset_config_file`.`view_host_role_id` AS `view_host_role_id`,`asset_config_file`.`view_host_role_cn` AS `view_host_role_cn`,`asset_config_file`.`view_en_host_role` AS `view_en_host_role`,`asset_config_file`.`view_provider_id` AS `view_provider_id`,`asset_config_file`.`view_provider_name_en` AS `view_provider_name_en`,`asset_config_file`.`view_provider_name_cn` AS `view_provider_name_cn`,`asset_config_file`.`view_hardware_info` AS `view_hardware_info`,`asset_config_file`.`view_ssh_port` AS `view_ssh_port`,`asset_config_file`.`view_init_type` AS `view_init_type`,`asset_config_file`.`view_init_type_cn` AS `view_init_type_cn`,`asset_config_file`.`view_clean_type` AS `view_clean_type`,`asset_config_file`.`view_clean_type_cn` AS `view_clean_type_cn`,`asset_config_file`.`view_recycle_type` AS `view_recycle_type`,`asset_config_file`.`view_recycle_type_cn` AS `view_recycle_type_cn`,`asset_config_file`.`view_init_login_info` AS `view_init_login_info`,`asset_config_file`.`view_change_status_remark` AS `view_change_status_remark`,`asset_config_file`.`view_remark` AS `view_remark`,`asset_config_file`.`view_asset_create_time` AS `view_asset_create_time`,`asset_config_file`.`view_asset_update_time` AS `view_asset_update_time`,`asset_config_file`.`view_asset_del_flag` AS `view_asset_del_flag`,`asset_config_file`.`view_pr_id` AS `view_pr_id`,`asset_config_file`.`view_asset_ownership_company_id` AS `view_asset_ownership_company_id`,`asset_config_file`.`view_asset_ownership_company_cn` AS `view_asset_ownership_company_cn`,`asset_config_file`.`view_asset_ownership_company_en` AS `view_asset_ownership_company_en`,`asset_config_file`.`view_asset_ownership_company_deleted` AS `view_asset_ownership_company_deleted`,`asset_config_file`.`view_server_affiliation_deleted` AS `view_server_affiliation_deleted`,`asset_config_file`.`view_user_company_id` AS `view_user_company_id`,`asset_config_file`.`view_user_company_cn` AS `view_user_company_cn`,`asset_config_file`.`view_user_company_en` AS `view_user_company_en`,`asset_config_file`.`view_user_company_deleted` AS `view_user_company_deleted`,`asset_config_file`.`view_user_project_id` AS `view_user_project_id`,`asset_config_file`.`view_user_project_cn` AS `view_user_project_cn`,`asset_config_file`.`view_user_project_en` AS `view_user_project_en`,`asset_config_file`.`view_user_project_deleted` AS `view_user_project_deleted`,`asset_config_file`.`view_cluster_name` AS `view_cluster_name`,`asset_config_file`.`id` AS `id`,`asset_config_file`.`project_id` AS `project_id`,`asset_config_file`.`name` AS `name`,`asset_config_file`.`dest_path` AS `dest_path`,`asset_config_file`.`file_mod_time` AS `file_mod_time`,`asset_config_file`.`id` AS `config_file_id`,ifnull(`config_mng_log`.`config_time`,946656000) AS `config_time` from ((select `view_assets`.`view_asset_id` AS `view_asset_id`,`view_assets`.`view_outer_ip` AS `view_outer_ip`,`view_assets`.`view_inner_ip` AS `view_inner_ip`,`view_assets`.`view_accelerate_domain` AS `view_accelerate_domain`,`view_assets`.`view_host_role_id` AS `view_host_role_id`,`view_assets`.`view_host_role_cn` AS `view_host_role_cn`,`view_assets`.`view_en_host_role` AS `view_en_host_role`,`view_assets`.`view_provider_id` AS `view_provider_id`,`view_assets`.`view_provider_name_en` AS `view_provider_name_en`,`view_assets`.`view_provider_name_cn` AS `view_provider_name_cn`,`view_assets`.`view_hardware_info` AS `view_hardware_info`,`view_assets`.`view_ssh_port` AS `view_ssh_port`,`view_assets`.`view_init_type` AS `view_init_type`,`view_assets`.`view_init_type_cn` AS `view_init_type_cn`,`view_assets`.`view_clean_type` AS `view_clean_type`,`view_assets`.`view_clean_type_cn` AS `view_clean_type_cn`,`view_assets`.`view_recycle_type` AS `view_recycle_type`,`view_assets`.`view_recycle_type_cn` AS `view_recycle_type_cn`,`view_assets`.`view_init_login_info` AS `view_init_login_info`,`view_assets`.`view_change_status_remark` AS `view_change_status_remark`,`view_assets`.`view_remark` AS `view_remark`,`view_assets`.`view_asset_create_time` AS `view_asset_create_time`,`view_assets`.`view_asset_update_time` AS `view_asset_update_time`,`view_assets`.`view_asset_del_flag` AS `view_asset_del_flag`,`view_assets`.`view_pr_id` AS `view_pr_id`,`view_assets`.`view_asset_ownership_company_id` AS `view_asset_ownership_company_id`,`view_assets`.`view_asset_ownership_company_cn` AS `view_asset_ownership_company_cn`,`view_assets`.`view_asset_ownership_company_en` AS `view_asset_ownership_company_en`,`view_assets`.`view_asset_ownership_company_deleted` AS `view_asset_ownership_company_deleted`,`view_assets`.`view_server_affiliation_deleted` AS `view_server_affiliation_deleted`,`view_assets`.`view_user_company_id` AS `view_user_company_id`,`view_assets`.`view_user_company_cn` AS `view_user_company_cn`,`view_assets`.`view_user_company_en` AS `view_user_company_en`,`view_assets`.`view_user_company_deleted` AS `view_user_company_deleted`,`view_assets`.`view_user_project_id` AS `view_user_project_id`,`view_assets`.`view_user_project_cn` AS `view_user_project_cn`,`view_assets`.`view_user_project_en` AS `view_user_project_en`,`view_assets`.`view_user_project_deleted` AS `view_user_project_deleted`,`view_assets`.`view_cluster_name` AS `view_cluster_name`,`config_file`.`id` AS `id`,`config_file`.`project_id` AS `project_id`,`config_file`.`name` AS `name`,`config_file`.`dest_path` AS `dest_path`,`config_file`.`file_mod_time` AS `file_mod_time` from ((select `view_assets`.`view_asset_id` AS `view_asset_id`,`view_assets`.`view_outer_ip` AS `view_outer_ip`,`view_assets`.`view_inner_ip` AS `view_inner_ip`,`view_assets`.`view_accelerate_domain` AS `view_accelerate_domain`,`view_assets`.`view_host_role_id` AS `view_host_role_id`,`view_assets`.`view_host_role_cn` AS `view_host_role_cn`,`view_assets`.`view_en_host_role` AS `view_en_host_role`,`view_assets`.`view_provider_id` AS `view_provider_id`,`view_assets`.`view_provider_name_en` AS `view_provider_name_en`,`view_assets`.`view_provider_name_cn` AS `view_provider_name_cn`,`view_assets`.`view_hardware_info` AS `view_hardware_info`,`view_assets`.`view_ssh_port` AS `view_ssh_port`,`view_assets`.`view_init_type` AS `view_init_type`,`view_assets`.`view_init_type_cn` AS `view_init_type_cn`,`view_assets`.`view_clean_type` AS `view_clean_type`,`view_assets`.`view_clean_type_cn` AS `view_clean_type_cn`,`view_assets`.`view_recycle_type` AS `view_recycle_type`,`view_assets`.`view_recycle_type_cn` AS `view_recycle_type_cn`,`view_assets`.`view_init_login_info` AS `view_init_login_info`,`view_assets`.`view_change_status_remark` AS `view_change_status_remark`,`view_assets`.`view_remark` AS `view_remark`,`view_assets`.`view_asset_create_time` AS `view_asset_create_time`,`view_assets`.`view_asset_update_time` AS `view_asset_update_time`,`view_assets`.`view_asset_del_flag` AS `view_asset_del_flag`,`view_assets`.`view_pr_id` AS `view_pr_id`,`view_assets`.`view_asset_ownership_company_id` AS `view_asset_ownership_company_id`,`view_assets`.`view_asset_ownership_company_cn` AS `view_asset_ownership_company_cn`,`view_assets`.`view_asset_ownership_company_en` AS `view_asset_ownership_company_en`,`view_assets`.`view_asset_ownership_company_deleted` AS `view_asset_ownership_company_deleted`,`view_assets`.`view_server_affiliation_deleted` AS `view_server_affiliation_deleted`,`view_assets`.`view_user_company_id` AS `view_user_company_id`,`view_assets`.`view_user_company_cn` AS `view_user_company_cn`,`view_assets`.`view_user_company_en` AS `view_user_company_en`,`view_assets`.`view_user_company_deleted` AS `view_user_company_deleted`,`view_assets`.`view_user_project_id` AS `view_user_project_id`,`view_assets`.`view_user_project_cn` AS `view_user_project_cn`,`view_assets`.`view_user_project_en` AS `view_user_project_en`,`view_assets`.`view_user_project_deleted` AS `view_user_project_deleted`,ifnull(`view_user_asset`.`view_cluster_name`,'') AS `view_cluster_name` from (`view_assets` left join `view_user_asset` on((`view_assets`.`view_asset_id` = `view_user_asset`.`view_asset_id`))) where ((`view_assets`.`view_recycle_type` = 2) and (`view_assets`.`view_asset_del_flag` = 0) and (`view_assets`.`view_user_project_deleted` = -(1)))) `view_assets` join `config_file`) where (`view_assets`.`view_user_project_id` = `config_file`.`project_id`)) `asset_config_file` left join `config_mng_log` on(((`asset_config_file`.`view_asset_id` = `config_mng_log`.`asset_id`) and (`asset_config_file`.`id` = `config_mng_log`.`config_file_id`)))) having (`asset_config_file`.`file_mod_time` > `config_time`)) `asset_file` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_assets`
--

/*!50001 DROP VIEW IF EXISTS `view_assets`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_assets` (`view_asset_id`,`view_outer_ip`,`view_inner_ip`,`view_accelerate_domain`,`view_host_role_id`,`view_host_role_cn`,`view_en_host_role`,`view_provider_id`,`view_provider_name_en`,`view_provider_name_cn`,`view_hardware_info`,`view_ssh_port`,`view_init_type`,`view_init_type_cn`,`view_clean_type`,`view_clean_type_cn`,`view_recycle_type`,`view_recycle_type_cn`,`view_init_login_info`,`view_change_status_remark`,`view_remark`,`view_asset_create_time`,`view_asset_update_time`,`view_asset_del_flag`,`view_pr_id`,`view_asset_ownership_company_id`,`view_asset_ownership_company_cn`,`view_asset_ownership_company_en`,`view_asset_ownership_company_deleted`,`view_server_affiliation_deleted`,`view_user_company_id`,`view_user_company_cn`,`view_user_company_en`,`view_user_company_deleted`,`view_user_project_id`,`view_user_project_cn`,`view_user_project_en`,`view_user_project_deleted`) AS select `a`.`asset_id` AS `asset_id`,`a`.`outer_ip` AS `outer_ip`,`a`.`inner_ip` AS `inner_ip`,`a`.`accelerate_domain` AS `accelerate_domain`,`a`.`host_role_id` AS `host_role_id`,left(`a`.`cn_host_role`,(char_length(`a`.`cn_host_role`) - 1)) AS `cn_host_role`,left(`a`.`en_host_role`,(char_length(`a`.`en_host_role`) - 1)) AS `en_host_role`,`a`.`provider_id` AS `provider_id`,`sl1`.`label` AS `provider_name_en`,`sl1`.`description` AS `provider_name_cn`,`a`.`hardware_info` AS `hardware_info`,`a`.`ssh_port` AS `ssh_port`,`a`.`init_type` AS `init_type`,`sl2`.`label` AS `init_type_cn`,`a`.`clean_type` AS `clean_type`,`sl3`.`label` AS `clean_type_cn`,`a`.`recycle_type` AS `recycle_type`,`sl4`.`label` AS `recycle_type_cn`,`a`.`init_login_info` AS `init_login_info`,`a`.`change_status_remark` AS `change_status_remark`,`a`.`remark` AS `remark`,`a`.`create_time` AS `create_time`,`a`.`update_time` AS `update_time`,`a`.`del_flag` AS `asset_del_flag`,`ownership_table`.`pr_id` AS `pr_id`,`ownership_table`.`asset_ownership_company_id` AS `asset_ownership_company_id`,`ownership_table`.`asset_ownership_company_cn` AS `asset_ownership_company_cn`,`ownership_table`.`asset_ownership_company_en` AS `asset_ownership_company_en`,`ownership_table`.`asset_ownership_company_deleted` AS `asset_ownership_company_deleted`,`ownership_table`.`server_affiliation_deleted` AS `server_affiliation_deleted`,`ownership_table`.`user_company_id` AS `user_company_id`,`ownership_table`.`user_company_cn` AS `user_company_cn`,`ownership_table`.`user_company_en` AS `user_company_en`,`ownership_table`.`user_company_deleted` AS `user_company_deleted`,`ownership_table`.`user_project_id` AS `user_project_id`,`ownership_table`.`user_project_cn` AS `user_project_cn`,`ownership_table`.`user_project_en` AS `user_project_en`,`ownership_table`.`user_project_deleted` AS `user_project_deleted` from ((((((select `a`.`asset_id` AS `asset_id`,`a`.`outer_ip` AS `outer_ip`,`a`.`inner_ip` AS `inner_ip`,`a`.`accelerate_domain` AS `accelerate_domain`,`a`.`host_role_id` AS `host_role_id`,concat(if((`ar`.`cn_host_role_game` is null),'',concat(`ar`.`cn_host_role_game`,',')),if((`ar`.`cn_host_role_cross` is null),'',concat(`ar`.`cn_host_role_cross`,',')),if((`ar`.`cn_host_role_center` is null),'',concat(`ar`.`cn_host_role_center`,',')),if((`ar`.`cn_host_role_login` is null),'',concat(`ar`.`cn_host_role_login`,',')),if((`ar`.`cn_host_role_pay` is null),'',concat(`ar`.`cn_host_role_pay`,',')),if((`ar`.`cn_host_role_test` is null),'',concat(`ar`.`cn_host_role_test`,',')),if((`ar`.`cn_host_role_backup` is null),'',concat(`ar`.`cn_host_role_backup`,',')),if((`ar`.`cn_host_role_web` is null),'',concat(`ar`.`cn_host_role_web`,',')),if((`ar`.`cn_host_role_db` is null),'',concat(`ar`.`cn_host_role_db`,',')),if((`ar`.`cn_host_role_cdn` is null),'',concat(`ar`.`cn_host_role_cdn`,',')),if((`ar`.`cn_host_role_gate` is null),'',concat(`ar`.`cn_host_role_gate`,','))) AS `cn_host_role`,concat(if((`ar`.`en_host_role_game` is null),'',concat(`ar`.`en_host_role_game`,',')),if((`ar`.`en_host_role_cross` is null),'',concat(`ar`.`en_host_role_cross`,',')),if((`ar`.`en_host_role_center` is null),'',concat(`ar`.`en_host_role_center`,',')),if((`ar`.`en_host_role_login` is null),'',concat(`ar`.`en_host_role_login`,',')),if((`ar`.`en_host_role_pay` is null),'',concat(`ar`.`en_host_role_pay`,',')),if((`ar`.`en_host_role_test` is null),'',concat(`ar`.`en_host_role_test`,',')),if((`ar`.`en_host_role_backup` is null),'',concat(`ar`.`en_host_role_backup`,',')),if((`ar`.`en_host_role_web` is null),'',concat(`ar`.`en_host_role_web`,',')),if((`ar`.`en_host_role_db` is null),'',concat(`ar`.`en_host_role_db`,',')),if((`ar`.`en_host_role_cdn` is null),'',concat(`ar`.`en_host_role_cdn`,',')),if((`ar`.`en_host_role_gate` is null),'',concat(`ar`.`en_host_role_gate`,','))) AS `en_host_role`,`a`.`provider_id` AS `provider_id`,`a`.`hardware_info` AS `hardware_info`,`a`.`ssh_port` AS `ssh_port`,`a`.`init_type` AS `init_type`,`a`.`clean_type` AS `clean_type`,`a`.`recycle_type` AS `recycle_type`,`a`.`init_login_info` AS `init_login_info`,`a`.`change_status_remark` AS `change_status_remark`,`a`.`remark` AS `remark`,`a`.`create_time` AS `create_time`,`a`.`update_time` AS `update_time`,`a`.`del_flag` AS `del_flag` from (`asset` `a` join (select `asset`.`asset_id` AS `asset_id`,`asset`.`outer_ip` AS `outer_ip`,(case when (find_in_set(1,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 1))) else NULL end) AS `cn_host_role_game`,(case when (find_in_set(1,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 1))) else NULL end) AS `en_host_role_game`,(case when (find_in_set(2,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 2))) else NULL end) AS `cn_host_role_cross`,(case when (find_in_set(2,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 2))) else NULL end) AS `en_host_role_cross`,(case when (find_in_set(3,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 3))) else NULL end) AS `cn_host_role_center`,(case when (find_in_set(3,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 3))) else NULL end) AS `en_host_role_center`,(case when (find_in_set(4,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 4))) else NULL end) AS `cn_host_role_login`,(case when (find_in_set(4,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 4))) else NULL end) AS `en_host_role_login`,(case when (find_in_set(5,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 5))) else NULL end) AS `cn_host_role_pay`,(case when (find_in_set(5,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 5))) else NULL end) AS `en_host_role_pay`,(case when (find_in_set(6,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 6))) else NULL end) AS `cn_host_role_test`,(case when (find_in_set(6,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 6))) else NULL end) AS `en_host_role_test`,(case when (find_in_set(7,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 7))) else NULL end) AS `cn_host_role_backup`,(case when (find_in_set(7,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 7))) else NULL end) AS `en_host_role_backup`,(case when (find_in_set(8,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 8))) else NULL end) AS `cn_host_role_web`,(case when (find_in_set(8,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 8))) else NULL end) AS `en_host_role_web`,(case when (find_in_set(9,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 9))) else NULL end) AS `cn_host_role_db`,(case when (find_in_set(9,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 9))) else NULL end) AS `en_host_role_db`,(case when (find_in_set(10,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 10))) else NULL end) AS `cn_host_role_cdn`,(case when (find_in_set(10,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 10))) else NULL end) AS `en_host_role_cdn`,(case when (find_in_set(11,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`description` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 11))) else NULL end) AS `cn_host_role_gate`,(case when (find_in_set(11,`asset`.`host_role_id`) >= 1) then (select `sys_dict`.`label` from `sys_dict` where ((`sys_dict`.`types` = 'host_role_type') and (`sys_dict`.`pid` <> -(1)) and (`sys_dict`.`value` = 11))) else NULL end) AS `en_host_role_gate` from `asset`) `ar`) where (`a`.`asset_id` = `ar`.`asset_id`)) `a` left join (select `sa`.`id` AS `id`,`sa`.`pr_id` AS `pr_id`,`sa`.`asset_id` AS `asset_id`,`sa`.`company_id` AS `asset_ownership_company_id`,`c`.`company_cn` AS `asset_ownership_company_cn`,`c`.`company_en` AS `asset_ownership_company_en`,`c`.`supply_company_status` AS `asset_ownership_company_deleted`,`sa`.`del_flag` AS `server_affiliation_deleted`,`vcp`.`view_company_id` AS `user_company_id`,`vcp`.`view_company_cn` AS `user_company_cn`,`vcp`.`view_company_en` AS `user_company_en`,`vcp`.`view_company_del_flag` AS `user_company_deleted`,`vcp`.`view_project_id` AS `user_project_id`,`vcp`.`view_project_cn` AS `user_project_cn`,`vcp`.`view_project_en` AS `user_project_en`,`vcp`.`view_project_del_flag` AS `user_project_deleted` from ((`server_affiliation` `sa` left join `view_company_project` `vcp` on((`sa`.`pr_id` = `vcp`.`view_pr_id`))) left join `company` `c` on((`sa`.`company_id` = `c`.`company_id`)))) `ownership_table` on((`a`.`asset_id` = `ownership_table`.`asset_id`))) left join (select cast(`sys_dict`.`value` as signed) AS `value`,`sys_dict`.`label` AS `label`,`sys_dict`.`description` AS `description` from `sys_dict` where ((`sys_dict`.`types` = 'cloud_provider_type') and (`sys_dict`.`pid` <> -(1)))) `sl1` on((`a`.`provider_id` = `sl1`.`value`))) left join (select cast(`sys_dict`.`value` as signed) AS `value`,`sys_dict`.`label` AS `label`,`sys_dict`.`description` AS `description` from `sys_dict` where ((`sys_dict`.`types` = 'init_type') and (`sys_dict`.`pid` <> -(1)))) `sl2` on((`a`.`init_type` = `sl2`.`value`))) left join (select cast(`sys_dict`.`value` as signed) AS `value`,`sys_dict`.`label` AS `label`,`sys_dict`.`description` AS `description` from `sys_dict` where ((`sys_dict`.`types` = 'clean_type') and (`sys_dict`.`pid` <> -(1)))) `sl3` on((`a`.`clean_type` = `sl3`.`value`))) left join (select cast(`sys_dict`.`value` as signed) AS `value`,`sys_dict`.`label` AS `label`,`sys_dict`.`description` AS `description` from `sys_dict` where ((`sys_dict`.`types` = 'recycle_type') and (`sys_dict`.`pid` <> -(1)))) `sl4` on((`a`.`recycle_type` = `sl4`.`value`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_company_project`
--

/*!50001 DROP VIEW IF EXISTS `view_company_project`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_company_project` (`view_company_id`,`view_company_cn`,`view_company_en`,`view_company_del_flag`,`view_pr_id`,`view_project_id`,`view_project_cn`,`view_project_en`,`view_dept_id`,`view_dept_name`,`view_project_type`,`view_project_type_cn`,`view_group_qq`,`view_group_type_cn`,`view_group_type_en`,`view_group_dev_qq`,`view_project_del_flag`) AS select `c`.`company_id` AS `company_id`,`c`.`company_cn` AS `company_cn`,`c`.`company_en` AS `company_en`,`c`.`del_flag` AS `company_del_flag`,`pr`.`id` AS `pr_id`,`p`.`project_id` AS `project_id`,`p`.`project_cn` AS `project_cn`,`p`.`project_en` AS `project_en`,`p`.`project_team` AS `project_team`,`sd`.`name` AS `name`,`sl1`.`value` AS `project_type`,`sl1`.`label` AS `project_type_cn`,`p`.`group_qq` AS `group_qq`,`sl2`.`label` AS `group_type_cn`,`p`.`group_type` AS `group_type_en`,`p`.`group_dev_qq` AS `group_dev_qq`,if((`p`.`del_flag` = 0),-(1),`p`.`del_flag`) AS `project_del_flag` from (((((`company` `c` left join `project_relationship` `pr` on((`c`.`company_id` = `pr`.`company_id`))) left join `project` `p` on((`pr`.`project_id` = `p`.`project_id`))) left join `sys_dept` `sd` on((`p`.`project_team` = `sd`.`id`))) left join (select cast(`sys_dict`.`value` as signed) AS `value`,`sys_dict`.`label` AS `label`,`sys_dict`.`description` AS `description` from `sys_dict` where ((`sys_dict`.`types` = 'project_status') and (`sys_dict`.`pid` <> -(1)))) `sl1` on((`p`.`project_type` = `sl1`.`value`))) left join (select `sys_dict`.`value` AS `value`,`sys_dict`.`label` AS `label`,`sys_dict`.`description` AS `description` from `sys_dict` where ((`sys_dict`.`types` = 'project_group_type') and (`sys_dict`.`pid` <> -(1)))) `sl2` on((`p`.`group_type` = `sl2`.`value`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_feature_label`
--

/*!50001 DROP VIEW IF EXISTS `view_feature_label`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_feature_label` (`view_feature_server_id`,`view_project_id`,`view_project_en`,`view_project_cn`,`view_feature_server_info`,`view_feature_server_remark`,`view_label_id`,`view_label_type`,`view_label_name`,`view_label_values`,`view_label_remark`,`view_project_del_flag`,`view_feature_server_del_flag`,`view_label_del_flag`,`view_table_name`) AS select `fsi`.`feature_server_id` AS `feature_server_id`,`fsi`.`project_id` AS `project_id`,`fsi`.`project_en` AS `project_en`,`fsi`.`project_cn` AS `project_cn`,`fsi`.`feature_server_info` AS `feature_server_info`,`fsi`.`remark` AS `feature_server_remark`,`l`.`label_id` AS `label_id`,`l`.`label_type` AS `label_type`,`l`.`label_name` AS `label_name`,`l`.`label_values` AS `label_values`,`l`.`label_remark` AS `label_remark`,`fsi`.`project_del_flag` AS `project_del_flag`,`fsi`.`feature_server_del_flag` AS `feature_server_del_flag`,`l`.`del_flag` AS `label_del_flag`,`fsi`.`table_name` AS `table_name` from ((select `fsi`.`feature_server_id` AS `feature_server_id`,`p`.`project_id` AS `project_id`,`p`.`project_en` AS `project_en`,`p`.`project_cn` AS `project_cn`,`fsi`.`feature_server_info` AS `feature_server_info`,`fsi`.`remark` AS `remark`,`p`.`del_flag` AS `project_del_flag`,`fsi`.`del_flag` AS `feature_server_del_flag`,'feature_server_info' AS `table_name` from (`feature_server_info` `fsi` join `project` `p`) where ((`fsi`.`project_id` = `p`.`project_id`) and (`p`.`del_flag` = 0))) `fsi` left join (select `l`.`label_id` AS `label_id`,`l`.`label_type` AS `label_type`,`l`.`label_name` AS `label_name`,`l`.`label_values` AS `label_values`,`l`.`label_remark` AS `label_remark`,`l`.`del_flag` AS `del_flag`,`lg`.`resource_en` AS `resource_en`,`lg`.`binding_id` AS `binding_id`,`lg`.`project_id` AS `project_id` from (`label` `l` left join `label_global` `lg` on((`l`.`label_id` = `lg`.`label_id`))) where ((`l`.`del_flag` = 0) and (`lg`.`resource_en` = 'feature_server_info')) order by `l`.`label_type`,`l`.`label_id`) `l` on(((convert(`fsi`.`table_name` using utf8mb3) = `l`.`resource_en`) and (`fsi`.`feature_server_id` = `l`.`binding_id`) and (`fsi`.`project_id` = `l`.`project_id`)))) order by `fsi`.`project_id`,`fsi`.`feature_server_id` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_ip_list`
--

/*!50001 DROP VIEW IF EXISTS `view_ip_list`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_ip_list` (`view_gatfile`,`view_placeholder_1`,`view_server_id`,`view_platform_en`,`view_cluster_name`,`view_server_alias`,`view_outer_ip`,`view_ssh_port`,`view_inner_ip`,`view_placeholder_2`,`view_open_time`,`view_provider_name_en`,`view_new_platform_en`,`view_project_en`,`view_unix_timestamp_opentime`,`view_server_status`) AS select `A`.`gatfile` AS `gatfile`,`A`.`time_zone` AS `time_zone`,`A`.`server_id` AS `server_id`,`A`.`view_platform_en` AS `view_platform_en`,`A`.`view_cluster_name` AS `view_cluster_name`,`A`.`server_alias` AS `server_alias`,`A`.`view_outer_ip` AS `view_outer_ip`,`A`.`view_ssh_port` AS `view_ssh_port`,`A`.`view_inner_ip` AS `view_inner_ip`,`A`.`mongodb_port` AS `mongodb_port`,`A`.`open_time` AS `open_time`,`A`.`view_provider_name_en` AS `view_provider_name_en`,`A`.`new_platform_en` AS `new_platform_en`,`A`.`view_project_en` AS `view_project_en`,`A`.`unix_timestamp_opentime` AS `unix_timestamp_opentime`,`A`.`server_status` AS `server_status` from (select 'gatfile' AS `gatfile`,'NULL' AS `time_zone`,`game_server`.`server_id` AS `server_id`,`view_platform_feature`.`view_platform_en` AS `view_platform_en`,(case when regexp_like(`view_platform_feature`.`view_platform_en`,'cross') then concat(`view_platform_feature`.`view_cluster_name`,'CR') else `view_platform_feature`.`view_cluster_name` end) AS `view_cluster_name`,`game_server`.`server_alias` AS `server_alias`,`view_assets`.`view_outer_ip` AS `view_outer_ip`,`view_assets`.`view_ssh_port` AS `view_ssh_port`,`view_assets`.`view_inner_ip` AS `view_inner_ip`,'NULL' AS `mongodb_port`,date_format(from_unixtime(`game_server`.`open_time`),'%Y-%m-%d_%H:%i:%s') AS `open_time`,`view_assets`.`view_provider_name_en` AS `view_provider_name_en`,(case when (regexp_like(`game_server`.`server_alias`,'jswar_jszzd_s0[0-1][0-9]') or regexp_like(`game_server`.`server_alias`,'jswar_Jszzdcross_s10[1-9]a') or (`game_server`.`server_alias` in ('jswar_Jszzdmidcross_s2001a','jswar_Jszzdmidcross_s2002a','jswar_Jszzdcross_s110a','jswar_jszzd_s020a','jswar_Jszzdbigcross_s3001a'))) then concat(left(`view_platform_feature`.`view_platform_en`,5),'1',right(`view_platform_feature`.`view_platform_en`,(length(`view_platform_feature`.`view_platform_en`) - 5))) else `view_platform_feature`.`view_platform_en` end) AS `new_platform_en`,`view_platform_feature`.`view_project_en` AS `view_project_en`,`game_server`.`open_time` AS `unix_timestamp_opentime`,`game_server`.`server_status` AS `server_status` from ((`game_server` join `view_assets`) join `view_platform_feature`) where ((`game_server`.`project_id` = `view_assets`.`view_user_project_id`) and (`game_server`.`asset_id` = `view_assets`.`view_asset_id`) and (`game_server`.`project_id` = `view_platform_feature`.`view_project_id`) and (`game_server`.`platform_id` = `view_platform_feature`.`view_platform_id`) and (`game_server`.`server_status` in (1,2,5)) and (`view_assets`.`view_recycle_type` = 2) and (`view_assets`.`view_asset_del_flag` = 0))) `A` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_label_info`
--

/*!50001 DROP VIEW IF EXISTS `view_label_info`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_label_info` (`view_label_id`,`view_label_type`,`view_label_name`,`view_label_values`,`view_label_remark`,`view_stop_status`) AS select `A`.`label_id` AS `label_id`,`A`.`label_type` AS `label_type`,`A`.`label_name` AS `label_name`,`A`.`label_values` AS `label_values`,`A`.`label_remark` AS `label_remark`,(case `A`.`stop_status` when '1' then '1' else '2' end) AS `stop_status` from (select `label_table`.`label_id` AS `label_id`,`label_table`.`label_type` AS `label_type`,`label_table`.`label_name` AS `label_name`,`label_table`.`label_values` AS `label_values`,`label_table`.`label_remark` AS `label_remark`,group_concat(distinct `data_status`.`stop_status` separator ',') AS `stop_status` from ((select `l`.`label_id` AS `label_id`,`l`.`label_type` AS `label_type`,`l`.`label_name` AS `label_name`,`l`.`label_values` AS `label_values`,`l`.`label_remark` AS `label_remark`,`lg`.`resource_en` AS `resource_en`,`lg`.`binding_id` AS `binding_id` from (`label` `l` left join `label_global` `lg` on((`l`.`label_id` = `lg`.`label_id`))) where (`l`.`del_flag` = 0) order by `l`.`label_type`,`l`.`label_id`) `label_table` left join (select `view_search_label`.`view_resource_en_name` AS `view_resource_en_name`,`view_search_label`.`view_primary_key_value` AS `view_primary_key_value`,(case json_unquote(json_extract(`view_search_label`.`view_json_id`,'$.view_recycle_type')) when '2' then '0' when '1' then '1' end) AS `stop_status` from `view_search_label` where (`view_search_label`.`view_system_show` in (1,3))) `data_status` on(((`label_table`.`resource_en` = convert(`data_status`.`view_resource_en_name` using utf8mb3)) and (`label_table`.`binding_id` = `data_status`.`view_primary_key_value`)))) group by `label_table`.`label_id` order by `label_table`.`label_type`,`label_table`.`label_id`) `A` order by (case `A`.`stop_status` when '1' then '1' else '2' end),`A`.`label_type`,`A`.`label_id` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_platform_feature`
--

/*!50001 DROP VIEW IF EXISTS `view_platform_feature`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_platform_feature` (`view_project_id`,`view_project_en`,`view_project_cn`,`view_platform_autoid`,`view_platform_id`,`view_platform_en`,`view_platform_cn`,`view_domain_format`,`view_platform_remark`,`view_cluster_name`,`view_cluster_platform_group`,`view_cluster_label_id`,`view_cluster_label_name`,`view_cluster_label_values`,`view_labels`,`view_cluster_feature_info`,`view_platform_feature_info`,`view_feature_info`,`view_center_cluster_info`,`view_center_platform_info`,`view_login_cluster_info`,`view_login_platform_info`,`view_pay_cluster_info`,`view_pay_platform_info`,`view_source_cluster_info`,`view_source_platform_info`,`view_cdn_cluster_info`,`view_cdn_platform_info`,`view_cross_cluster_info`,`view_cross_platform_info`,`view_db_cluster_info`,`view_db_platform_info`,`view_gate_cluster_info`,`view_gate_platform_info`,`view_backup_cluster_info`,`view_backup_platform_info`,`view_project_del_flag`,`view_platform_del_flag`,`view_feature_server_del_flag`,`view_label_del_flag`) AS select `D`.`project_id` AS `project_id`,`D`.`project_en` AS `project_en`,`D`.`project_cn` AS `project_cn`,`D`.`platform_autoid` AS `platform_autoid`,`D`.`platform_id` AS `platform_id`,`D`.`platform_en` AS `platform_en`,`D`.`platform_cn` AS `platform_cn`,`D`.`domain_format` AS `domain_format`,`D`.`platform_remark` AS `platform_remark`,`D`.`cluster_name` AS `cluster_name`,`D`.`cluster_platform_group` AS `cluster_platform_group`,`D`.`cluster_label_id` AS `cluster_label_id`,`D`.`cluster_label_name` AS `cluster_label_name`,`D`.`cluster_label_values` AS `cluster_label_values`,`D`.`labels` AS `labels`,`D`.`cluster_feature_info` AS `cluster_feature_info`,json_object('center_feature_info',json_merge_patch(`D`.`center_cluster_info`,`D`.`center_platform_info`),'login_feature_info',json_merge_patch(`D`.`login_cluster_info`,`D`.`login_platform_info`),'pay_feature_info',json_merge_patch(`D`.`pay_cluster_info`,`D`.`pay_platform_info`),'source_feature_info',json_merge_patch(`D`.`source_cluster_info`,`D`.`source_platform_info`),'cdn_feature_info',json_merge_patch(`D`.`cdn_cluster_info`,`D`.`cdn_platform_info`),'cross_feature_info',json_merge_patch(`D`.`cross_cluster_info`,`D`.`cross_platform_info`),'db_feature_info',json_merge_patch(`D`.`db_cluster_info`,`D`.`db_platform_info`),'gate_feature_info',json_merge_patch(`D`.`gate_cluster_info`,`D`.`gate_platform_info`),'backup_feature_info',json_merge_patch(`D`.`backup_cluster_info`,`D`.`backup_platform_info`)) AS `platform_feature_info`,concat('{"center":[',`D`.`center_cluster_info`,',',`D`.`center_platform_info`,'],"login":[',`D`.`login_cluster_info`,',',`D`.`login_platform_info`,'],"pay":[',`D`.`pay_cluster_info`,',',`D`.`pay_platform_info`,'],"source":[',`D`.`source_cluster_info`,',',`D`.`source_platform_info`,'],"cdn":[',`D`.`cdn_cluster_info`,',',`D`.`cdn_platform_info`,'],"cross":[',`D`.`cross_cluster_info`,',',`D`.`cross_platform_info`,'],"db":[',`D`.`db_cluster_info`,',',`D`.`db_platform_info`,'],"gate":[',`D`.`gate_cluster_info`,',',`D`.`gate_platform_info`,'],"backup":[',`D`.`backup_cluster_info`,',',`D`.`backup_platform_info`,']}') AS `feature_info`,`D`.`center_cluster_info` AS `center_cluster_info`,`D`.`center_platform_info` AS `center_platform_info`,`D`.`login_cluster_info` AS `login_cluster_info`,`D`.`login_platform_info` AS `login_platform_info`,`D`.`pay_cluster_info` AS `pay_cluster_info`,`D`.`pay_platform_info` AS `pay_platform_info`,`D`.`source_cluster_info` AS `source_cluster_info`,`D`.`source_platform_info` AS `source_platform_info`,`D`.`cdn_cluster_info` AS `cdn_cluster_info`,`D`.`cdn_platform_info` AS `cdn_platform_info`,`D`.`cross_cluster_info` AS `cross_cluster_info`,`D`.`cross_platform_info` AS `cross_platform_info`,`D`.`backup_cluster_info` AS `backup_cluster_info`,`D`.`backup_platform_info` AS `backup_platform_info`,`D`.`db_cluster_info` AS `db_cluster_info`,`D`.`db_platform_info` AS `db_platform_info`,`D`.`gate_cluster_info` AS `gate_cluster_info`,`D`.`gate_platform_info` AS `gate_platform_info`,`D`.`project_del_flag` AS `project_del_flag`,`D`.`platform_del_flag` AS `platform_del_flag`,`D`.`feature_server_del_flag` AS `feature_server_del_flag`,`D`.`label_del_flag` AS `label_del_flag` from (select `C`.`project_id` AS `project_id`,`C`.`project_en` AS `project_en`,`C`.`project_cn` AS `project_cn`,`C`.`platform_autoid` AS `platform_autoid`,`C`.`platform_id` AS `platform_id`,`C`.`platform_en` AS `platform_en`,`C`.`platform_cn` AS `platform_cn`,`C`.`domain_format` AS `domain_format`,`C`.`platform_remark` AS `platform_remark`,group_concat(distinct `C`.`cluster_name` separator ',') AS `cluster_name`,group_concat(distinct `C`.`cluster_platform_group` separator ',') AS `cluster_platform_group`,group_concat(distinct `C`.`cluster_label_id` separator ',') AS `cluster_label_id`,group_concat(distinct `C`.`cluster_label_name` separator ',') AS `cluster_label_name`,group_concat(distinct `C`.`cluster_label_values` separator ',') AS `cluster_label_values`,group_concat(distinct concat(`C`.`label_name`,'(',`C`.`label_values`,')') order by `C`.`label_type` ASC,`C`.`label_id` ASC separator ',') AS `labels`,if((group_concat(distinct `C`.`center_cluster_info` separator ',') is null),json_object(),group_concat(distinct `C`.`center_cluster_info` separator ',')) AS `center_cluster_info`,if((group_concat(distinct `C`.`login_cluster_info` separator ',') is null),json_object(),group_concat(distinct `C`.`login_cluster_info` separator ',')) AS `login_cluster_info`,if((group_concat(distinct `C`.`pay_cluster_info` separator ',') is null),json_object(),group_concat(distinct `C`.`pay_cluster_info` separator ',')) AS `pay_cluster_info`,if((group_concat(distinct `C`.`source_cluster_info` separator ',') is null),json_object(),group_concat(distinct `C`.`source_cluster_info` separator ',')) AS `source_cluster_info`,if((group_concat(distinct `C`.`cdn_cluster_info` separator ',') is null),json_object(),group_concat(distinct `C`.`cdn_cluster_info` separator ',')) AS `cdn_cluster_info`,if((group_concat(distinct `C`.`cross_cluster_info` separator ',') is null),json_object(),group_concat(distinct `C`.`cross_cluster_info` separator ',')) AS `cross_cluster_info`,if((group_concat(distinct `C`.`db_cluster_info` separator ',') is null),json_object(),group_concat(distinct `C`.`db_cluster_info` separator ',')) AS `db_cluster_info`,if((group_concat(distinct `C`.`gate_cluster_info` separator ',') is null),json_object(),group_concat(distinct `C`.`gate_cluster_info` separator ',')) AS `gate_cluster_info`,if((group_concat(distinct `C`.`backup_cluster_info` separator ',') is null),json_object(),group_concat(distinct `C`.`backup_cluster_info` separator ',')) AS `backup_cluster_info`,if((group_concat(distinct `C`.`center_platform_info` separator ',') is null),json_object(),group_concat(distinct `C`.`center_platform_info` separator ',')) AS `center_platform_info`,if((group_concat(distinct `C`.`login_platform_info` separator ',') is null),json_object(),group_concat(distinct `C`.`login_platform_info` separator ',')) AS `login_platform_info`,if((group_concat(distinct `C`.`pay_platform_info` separator ',') is null),json_object(),group_concat(distinct `C`.`pay_platform_info` separator ',')) AS `pay_platform_info`,if((group_concat(distinct `C`.`source_platform_info` separator ',') is null),json_object(),group_concat(distinct `C`.`source_platform_info` separator ',')) AS `source_platform_info`,if((group_concat(distinct `C`.`cdn_platform_info` separator ',') is null),json_object(),group_concat(distinct `C`.`cdn_platform_info` separator ',')) AS `cdn_platform_info`,if((group_concat(distinct `C`.`cross_platform_info` separator ',') is null),json_object(),group_concat(distinct `C`.`cross_platform_info` separator ',')) AS `cross_platform_info`,if((group_concat(distinct `C`.`db_platform_info` separator ',') is null),json_object(),group_concat(distinct `C`.`db_platform_info` separator ',')) AS `db_platform_info`,if((group_concat(distinct `C`.`gate_platform_info` separator ',') is null),json_object(),group_concat(distinct `C`.`gate_platform_info` separator ',')) AS `gate_platform_info`,if((group_concat(distinct `C`.`backup_platform_info` separator ',') is null),json_object(),group_concat(distinct `C`.`backup_platform_info` separator ',')) AS `backup_platform_info`,if((group_concat(distinct `C`.`cluster_feature_info` separator ',') is null),json_object(),group_concat(distinct `C`.`cluster_feature_info` separator ',')) AS `cluster_feature_info`,`C`.`project_del_flag` AS `project_del_flag`,`C`.`platform_del_flag` AS `platform_del_flag`,if((group_concat(distinct `C`.`feature_server_del_flag` separator ',') is null),'0',group_concat(distinct `C`.`feature_server_del_flag` separator ',')) AS `feature_server_del_flag`,`C`.`label_del_flag` AS `label_del_flag` from (select `A`.`project_id` AS `project_id`,`A`.`project_en` AS `project_en`,`A`.`project_cn` AS `project_cn`,`A`.`tb_platform` AS `platform_table_name`,`A`.`platform_autoid` AS `platform_autoid`,`A`.`platform_id` AS `platform_id`,`A`.`platform_en` AS `platform_en`,`A`.`platform_cn` AS `platform_cn`,`A`.`domain_format` AS `domain_format`,`A`.`platform_remark` AS `platform_remark`,`A`.`tb_feature` AS `feature_table_name`,`A`.`feature_server_id` AS `feature_server_id`,`A`.`feature_server_info` AS `feature_server_info`,`A`.`feature_server_remark` AS `feature_server_remark`,`A`.`label_id` AS `label_id`,`A`.`label_type` AS `label_type`,`A`.`label_name` AS `label_name`,`A`.`label_values` AS `label_values`,`A`.`label_remark` AS `label_remark`,(case when (`A`.`label_type` = 1) then upper(replace(`A`.`label_values`,'cluster_','')) end) AS `cluster_name`,(case when (`A`.`label_type` = 1) then if((`A`.`label_name` is null),'__',concat(`A`.`label_name`,'__',`A`.`label_values`)) end) AS `cluster_platform_group`,(case when (`A`.`label_type` = 1) then `A`.`label_id` end) AS `cluster_label_id`,(case when (`A`.`label_type` = 1) then `A`.`label_name` end) AS `cluster_label_name`,(case when (`A`.`label_type` = 1) then `A`.`label_values` end) AS `cluster_label_values`,`B`.`view_center_cluster_info` AS `center_cluster_info`,`B`.`view_login_cluster_info` AS `login_cluster_info`,`B`.`view_pay_cluster_info` AS `pay_cluster_info`,`B`.`view_source_cluster_info` AS `source_cluster_info`,`B`.`view_cdn_cluster_info` AS `cdn_cluster_info`,`B`.`view_cross_cluster_info` AS `cross_cluster_info`,`B`.`view_db_cluster_info` AS `db_cluster_info`,`B`.`view_gate_cluster_info` AS `gate_cluster_info`,`B`.`view_backup_cluster_info` AS `backup_cluster_info`,(case when ((json_unquote(json_extract(`A`.`feature_server_info`,'$.type')) = 'center') and (`A`.`label_type` = 2)) then `A`.`feature_server_info` end) AS `center_platform_info`,(case when ((json_unquote(json_extract(`A`.`feature_server_info`,'$.type')) = 'login') and (`A`.`label_type` = 2)) then `A`.`feature_server_info` end) AS `login_platform_info`,(case when ((json_unquote(json_extract(`A`.`feature_server_info`,'$.type')) = 'pay') and (`A`.`label_type` = 2)) then `A`.`feature_server_info` end) AS `pay_platform_info`,(case when ((json_unquote(json_extract(`A`.`feature_server_info`,'$.type')) = 'source') and (`A`.`label_type` = 2)) then `A`.`feature_server_info` end) AS `source_platform_info`,(case when ((json_unquote(json_extract(`A`.`feature_server_info`,'$.type')) = 'cdn') and (`A`.`label_type` = 2)) then `A`.`feature_server_info` end) AS `cdn_platform_info`,(case when ((json_unquote(json_extract(`A`.`feature_server_info`,'$.type')) = 'cross') and (`A`.`label_type` = 2)) then `A`.`feature_server_info` end) AS `cross_platform_info`,(case when ((json_unquote(json_extract(`A`.`feature_server_info`,'$.type')) = 'db') and (`A`.`label_type` = 2)) then `A`.`feature_server_info` end) AS `db_platform_info`,(case when ((json_unquote(json_extract(`A`.`feature_server_info`,'$.type')) = 'gate') and (`A`.`label_type` = 2)) then `A`.`feature_server_info` end) AS `gate_platform_info`,(case when ((json_unquote(json_extract(`A`.`feature_server_info`,'$.type')) = 'backup') and (`A`.`label_type` = 2)) then `A`.`feature_server_info` end) AS `backup_platform_info`,`B`.`view_cluster_feature_info` AS `cluster_feature_info`,`A`.`project_del_flag` AS `project_del_flag`,`A`.`platform_del_flag` AS `platform_del_flag`,`A`.`feature_server_del_flag` AS `feature_server_del_flag`,`A`.`label_del_flag` AS `label_del_flag` from ((select `view_platform_label`.`view_project_id` AS `project_id`,`view_platform_label`.`view_project_en` AS `project_en`,`view_platform_label`.`view_project_cn` AS `project_cn`,`view_platform_label`.`view_table_name` AS `tb_platform`,`view_platform_label`.`view_platform_autoid` AS `platform_autoid`,`view_platform_label`.`view_platform_id` AS `platform_id`,`view_platform_label`.`view_platform_en` AS `platform_en`,`view_platform_label`.`view_platform_cn` AS `platform_cn`,`view_platform_label`.`view_domain_format` AS `domain_format`,`view_platform_label`.`view_platform_remark` AS `platform_remark`,`view_feature_label`.`view_table_name` AS `tb_feature`,`view_feature_label`.`view_feature_server_id` AS `feature_server_id`,`view_feature_label`.`view_feature_server_info` AS `feature_server_info`,`view_feature_label`.`view_feature_server_remark` AS `feature_server_remark`,`view_platform_label`.`view_label_id` AS `label_id`,`view_platform_label`.`view_label_type` AS `label_type`,`view_platform_label`.`view_label_name` AS `label_name`,`view_platform_label`.`view_label_values` AS `label_values`,`view_platform_label`.`view_label_remark` AS `label_remark`,`view_platform_label`.`view_project_del_flag` AS `project_del_flag`,`view_platform_label`.`view_platform_del_flag` AS `platform_del_flag`,`view_feature_label`.`view_feature_server_del_flag` AS `feature_server_del_flag`,`view_platform_label`.`view_label_del_flag` AS `label_del_flag` from (`view_platform_label` left join `view_feature_label` on(((`view_platform_label`.`view_project_id` = `view_feature_label`.`view_project_id`) and (`view_platform_label`.`view_label_id` = `view_feature_label`.`view_label_id`))))) `A` left join `view_project_cluster_info` `B` on(((`A`.`project_id` = `B`.`view_project_id`) and (`A`.`label_id` = `B`.`view_label_id`))))) `C` group by `C`.`platform_autoid`) `D` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_platform_label`
--

/*!50001 DROP VIEW IF EXISTS `view_platform_label`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_platform_label` (`view_platform_autoid`,`view_project_id`,`view_project_en`,`view_project_cn`,`view_platform_id`,`view_platform_en`,`view_platform_cn`,`view_domain_format`,`view_platform_remark`,`view_platform_group_1`,`view_platform_group`,`view_args_platform_group`,`view_label_id`,`view_label_type`,`view_label_name`,`view_label_values`,`view_label_remark`,`view_project_del_flag`,`view_platform_del_flag`,`view_label_del_flag`,`view_table_name`) AS select `p`.`id` AS `id`,`p`.`project_id` AS `project_id`,`p`.`project_en` AS `project_en`,`p`.`project_cn` AS `project_cn`,`p`.`platform_id` AS `platform_id`,`p`.`platform_en` AS `platform_en`,`p`.`platform_cn` AS `platform_cn`,`p`.`domain_format` AS `domain_format`,`p`.`remark` AS `platform_remark`,if((`l`.`label_values` is null),'',`l`.`label_values`) AS `platform_group_1`,if((`l`.`label_name` is null),'',concat(`l`.`label_name`,'(',`l`.`label_values`,')')) AS `platform_group`,if((`l`.`label_name` is null),'__',concat(`l`.`label_name`,'__',`l`.`label_values`)) AS `args_platform_group`,`l`.`label_id` AS `label_id`,`l`.`label_type` AS `label_type`,`l`.`label_name` AS `label_name`,`l`.`label_values` AS `label_values`,`l`.`label_remark` AS `label_remark`,`p`.`project_del_flag` AS `project_del_flag`,`p`.`platform_del_flag` AS `platform_del_flag`,`l`.`del_flag` AS `label_del_flag`,`p`.`table_name` AS `table_name` from ((select `platform`.`id` AS `id`,`project`.`project_id` AS `project_id`,`project`.`project_en` AS `project_en`,`project`.`project_cn` AS `project_cn`,`platform`.`platform_id` AS `platform_id`,`platform`.`platform_en` AS `platform_en`,`platform`.`platform_cn` AS `platform_cn`,`platform`.`domain_format` AS `domain_format`,`platform`.`remark` AS `remark`,`project`.`del_flag` AS `project_del_flag`,`platform`.`del_flag` AS `platform_del_flag`,'platform' AS `table_name` from (`platform` join `project`) where ((`platform`.`project_id` = `project`.`project_id`) and (`project`.`del_flag` = 0))) `p` left join (select `l`.`label_id` AS `label_id`,`l`.`label_type` AS `label_type`,`l`.`label_name` AS `label_name`,`l`.`label_values` AS `label_values`,`l`.`label_remark` AS `label_remark`,`l`.`del_flag` AS `del_flag`,`lg`.`resource_en` AS `resource_en`,`lg`.`binding_id` AS `binding_id`,`lg`.`project_id` AS `project_id` from (`label` `l` left join `label_global` `lg` on((`l`.`label_id` = `lg`.`label_id`))) where ((`l`.`del_flag` = 0) and (`lg`.`resource_en` = 'platform')) order by `l`.`label_type`,`l`.`label_id`) `l` on(((convert(`p`.`table_name` using utf8mb3) = `l`.`resource_en`) and (`p`.`id` = `l`.`binding_id`) and (`p`.`project_id` = `l`.`project_id`)))) order by `p`.`project_id`,`p`.`platform_id` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_project_cluster_info`
--

/*!50001 DROP VIEW IF EXISTS `view_project_cluster_info`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_project_cluster_info` (`view_project_id`,`view_project_en`,`view_project_cn`,`view_feature_table_name`,`view_label_id`,`view_label_type`,`view_label_name`,`view_label_values`,`view_label_remark`,`view_cluster_name`,`view_cluster_platform_group`,`view_center_cluster_info`,`view_login_cluster_info`,`view_pay_cluster_info`,`view_source_cluster_info`,`view_cdn_cluster_info`,`view_cross_cluster_info`,`view_db_cluster_info`,`view_gate_cluster_info`,`view_backup_cluster_info`,`view_project_del_flag`,`view_feature_server_del_flag_abandon`,`view_label_del_flag`,`view_cluster_feature_info`,`view_feature_server_del_flag`) AS select `B`.`view_project_id` AS `view_project_id`,`B`.`view_project_en` AS `view_project_en`,`B`.`view_project_cn` AS `view_project_cn`,`B`.`view_feature_table_name` AS `view_feature_table_name`,`B`.`view_label_id` AS `view_label_id`,`B`.`view_label_type` AS `view_label_type`,`B`.`view_label_name` AS `view_label_name`,`B`.`view_label_values` AS `view_label_values`,`B`.`view_label_remark` AS `view_label_remark`,`B`.`cluster_name` AS `cluster_name`,`B`.`cluster_platform_group` AS `cluster_platform_group`,`B`.`center_cluster_info` AS `center_cluster_info`,`B`.`login_cluster_info` AS `login_cluster_info`,`B`.`pay_cluster_info` AS `pay_cluster_info`,`B`.`source_cluster_info` AS `source_cluster_info`,`B`.`cdn_cluster_info` AS `cdn_cluster_info`,`B`.`cross_cluster_info` AS `cross_cluster_info`,`B`.`db_cluster_info` AS `db_cluster_info`,`B`.`gate_cluster_info` AS `gate_cluster_info`,`B`.`backup_cluster_info` AS `backup_cluster_info`,`B`.`project_del_flag` AS `project_del_flag`,`B`.`feature_server_del_flag` AS `feature_server_del_flag`,`B`.`label_del_flag` AS `label_del_flag`,replace(replace(replace(json_object('center_feature_info',`B`.`center_cluster_info`,'login_feature_info',`B`.`login_cluster_info`,'pay_feature_info',`B`.`pay_cluster_info`,'source_feature_info',`B`.`source_cluster_info`,'cdn_feature_info',`B`.`cdn_cluster_info`,'cross_feature_info',`B`.`cross_cluster_info`,'db_feature_info',`B`.`db_cluster_info`,'gate_feature_info',`B`.`gate_cluster_info`,'backup_feature_info',`B`.`backup_cluster_info`),'\\',''),'"{','{'),'}"','}') AS `cluster_feature_info`,(case `B`.`feature_server_del_flag` when '1' then '1' else '2' end) AS `feature_server_del_flag` from (select `A`.`view_project_id` AS `view_project_id`,`A`.`view_project_en` AS `view_project_en`,`A`.`view_project_cn` AS `view_project_cn`,`A`.`view_feature_table_name` AS `view_feature_table_name`,`A`.`view_label_id` AS `view_label_id`,`A`.`view_label_type` AS `view_label_type`,`A`.`view_label_name` AS `view_label_name`,`A`.`view_label_values` AS `view_label_values`,`A`.`view_label_remark` AS `view_label_remark`,`A`.`cluster_name` AS `cluster_name`,`A`.`cluster_platform_group` AS `cluster_platform_group`,if((group_concat(`A`.`center_cluster_info` separator ',') is null),json_object(),group_concat(`A`.`center_cluster_info` separator ',')) AS `center_cluster_info`,if((group_concat(`A`.`login_cluster_info` separator ',') is null),json_object(),group_concat(`A`.`login_cluster_info` separator ',')) AS `login_cluster_info`,if((group_concat(`A`.`pay_cluster_info` separator ',') is null),json_object(),group_concat(`A`.`pay_cluster_info` separator ',')) AS `pay_cluster_info`,if((group_concat(`A`.`source_cluster_info` separator ',') is null),json_object(),group_concat(`A`.`source_cluster_info` separator ',')) AS `source_cluster_info`,if((group_concat(`A`.`cdn_cluster_info` separator ',') is null),json_object(),group_concat(`A`.`cdn_cluster_info` separator ',')) AS `cdn_cluster_info`,if((group_concat(`A`.`cross_cluster_info` separator ',') is null),json_object(),group_concat(`A`.`cross_cluster_info` separator ',')) AS `cross_cluster_info`,if((group_concat(`A`.`db_cluster_info` separator ',') is null),json_object(),group_concat(`A`.`db_cluster_info` separator ',')) AS `db_cluster_info`,if((group_concat(`A`.`gate_cluster_info` separator ',') is null),json_object(),group_concat(`A`.`gate_cluster_info` separator ',')) AS `gate_cluster_info`,if((group_concat(`A`.`backup_cluster_info` separator ',') is null),json_object(),group_concat(`A`.`backup_cluster_info` separator ',')) AS `backup_cluster_info`,`A`.`project_del_flag` AS `project_del_flag`,group_concat(distinct `A`.`feature_server_del_flag` separator ',') AS `feature_server_del_flag`,`A`.`label_del_flag` AS `label_del_flag` from (select `view_feature_label`.`view_project_id` AS `view_project_id`,`view_feature_label`.`view_project_en` AS `view_project_en`,`view_feature_label`.`view_project_cn` AS `view_project_cn`,`view_feature_label`.`view_table_name` AS `view_feature_table_name`,`view_feature_label`.`view_feature_server_id` AS `view_feature_server_id`,`view_feature_label`.`view_feature_server_info` AS `view_feature_server_info`,`view_feature_label`.`view_feature_server_remark` AS `view_feature_server_remark`,`view_feature_label`.`view_label_id` AS `view_label_id`,`view_feature_label`.`view_label_type` AS `view_label_type`,`view_feature_label`.`view_label_name` AS `view_label_name`,`view_feature_label`.`view_label_values` AS `view_label_values`,`view_feature_label`.`view_label_remark` AS `view_label_remark`,upper(replace(`view_feature_label`.`view_label_values`,'cluster_','')) AS `cluster_name`,if((`view_feature_label`.`view_label_name` is null),'__',concat(`view_feature_label`.`view_label_name`,'__',`view_feature_label`.`view_label_values`)) AS `cluster_platform_group`,(case when (json_unquote(json_extract(`view_feature_label`.`view_feature_server_info`,'$.type')) = 'center') then `view_feature_label`.`view_feature_server_info` end) AS `center_cluster_info`,(case when (json_unquote(json_extract(`view_feature_label`.`view_feature_server_info`,'$.type')) = 'login') then `view_feature_label`.`view_feature_server_info` end) AS `login_cluster_info`,(case when (json_unquote(json_extract(`view_feature_label`.`view_feature_server_info`,'$.type')) = 'pay') then `view_feature_label`.`view_feature_server_info` end) AS `pay_cluster_info`,(case when (json_unquote(json_extract(`view_feature_label`.`view_feature_server_info`,'$.type')) = 'source') then `view_feature_label`.`view_feature_server_info` end) AS `source_cluster_info`,(case when (json_unquote(json_extract(`view_feature_label`.`view_feature_server_info`,'$.type')) = 'cdn') then `view_feature_label`.`view_feature_server_info` end) AS `cdn_cluster_info`,(case when (json_unquote(json_extract(`view_feature_label`.`view_feature_server_info`,'$.type')) = 'cross') then `view_feature_label`.`view_feature_server_info` end) AS `cross_cluster_info`,(case when (json_unquote(json_extract(`view_feature_label`.`view_feature_server_info`,'$.type')) = 'db') then `view_feature_label`.`view_feature_server_info` end) AS `db_cluster_info`,(case when (json_unquote(json_extract(`view_feature_label`.`view_feature_server_info`,'$.type')) = 'gate') then `view_feature_label`.`view_feature_server_info` end) AS `gate_cluster_info`,(case when (json_unquote(json_extract(`view_feature_label`.`view_feature_server_info`,'$.type')) = 'backup') then `view_feature_label`.`view_feature_server_info` end) AS `backup_cluster_info`,`view_feature_label`.`view_project_del_flag` AS `project_del_flag`,`view_feature_label`.`view_feature_server_del_flag` AS `feature_server_del_flag`,`view_feature_label`.`view_label_del_flag` AS `label_del_flag` from `view_feature_label` where (`view_feature_label`.`view_label_type` = 1)) `A` group by `A`.`view_project_id`,`A`.`cluster_name`) `B` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_search_label`
--

/*!50001 DROP VIEW IF EXISTS `view_search_label`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_search_label` (`view_resource_cn_name`,`view_resource_en_name`,`view_resource_remark`,`view_project_id`,`view_primary_key`,`view_primary_key_value`,`view_resource_type`,`view_resource_value`,`view_data_content`,`view_data_url`,`view_json_id`,`view_table_name`,`view_show_cluster`,`view_show_feature`,`view_show_install`,`view_show_other`,`view_system_show`) AS select `table_data`.`resource_cn_name` AS `resource_cn_name`,`table_data`.`resource_en_name` AS `resource_en_name`,`table_data`.`resource_remark` AS `resource_remark`,`table_data`.`project_id` AS `project_id`,`table_data`.`primary_key` AS `primary_key`,`table_data`.`primary_key_value` AS `primary_key_value`,`table_data`.`resource_type` AS `resource_type`,`table_data`.`resource_value` AS `resource_value`,concat(`table_data`.`resource_en_name`,':',`table_data`.`resource_value`) AS `data_content`,ifnull(`table_url`.`url`,'') AS `url`,concat('{"',`table_data`.`primary_key`,'":',`table_data`.`primary_key_value`,`table_data`.`data_status`,'}') AS `json_id`,`table_data`.`resource_en_name` AS `resource_en_name`,`table_data`.`show_cluster` AS `show_cluster`,`table_data`.`show_feature` AS `show_feature`,`table_data`.`show_install` AS `show_install`,`table_data`.`show_other` AS `show_other`,`table_data`.`system_show` AS `system_show` from ((select '标签' AS `resource_cn_name`,'label' AS `resource_en_name`,'标签信息' AS `resource_remark`,0 AS `project_id`,'label_id' AS `primary_key`,`label`.`label_id` AS `primary_key_value`,'' AS `resource_type`,'' AS `data_status`,concat(`label`.`label_name`,'-',`label`.`label_values`) AS `resource_value`,0 AS `show_cluster`,0 AS `show_feature`,0 AS `show_install`,0 AS `show_other`,2 AS `system_show` from `label` where (`label`.`del_flag` = 0) union all select '用户' AS `resource_cn_name`,'sys_user' AS `resource_en_name`,'用户信息' AS `resource_remark`,-(1) AS `project_id`,'id' AS `primary_key`,`sys_user`.`id` AS `primary_key_value`,'' AS `resource_type`,(case `sys_user`.`del_flag` when '1' then concat(',"view_recycle_type":1') when '0' then concat(',"view_recycle_type":2') end) AS `data_status`,concat(`sys_user`.`nick_name`,'-',`sys_user`.`name`) AS `resource_value`,0 AS `show_cluster`,0 AS `show_feature`,0 AS `show_install`,1 AS `show_other`,3 AS `system_show` from `sys_user` where (`sys_user`.`del_flag` = 0) union all select '角色' AS `resource_cn_name`,'sys_role' AS `resource_en_name`,'角色信息' AS `resource_remark`,0 AS `project_id`,'id' AS `primary_key`,`sys_role`.`id` AS `primary_key_value`,'' AS `resource_type`,'' AS `data_status`,concat(`sys_role`.`remark`,'-',`sys_role`.`name`) AS `resource_value`,0 AS `show_cluster`,0 AS `show_feature`,0 AS `show_install`,0 AS `show_other`,2 AS `system_show` from `sys_role` where (`sys_role`.`del_flag` = 0) union all select '字典' AS `resource_cn_name`,'sys_dict' AS `resource_en_name`,'字典信息' AS `resource_remark`,0 AS `project_id`,'id' AS `primary_key`,`sys_dict`.`id` AS `primary_key_value`,'' AS `resource_type`,'' AS `data_status`,concat(`sys_dict`.`description`,'-',`sys_dict`.`types`) AS `resource_value`,0 AS `show_cluster`,0 AS `show_feature`,0 AS `show_install`,0 AS `show_other`,2 AS `system_show` from `sys_dict` where (`sys_dict`.`pid` = -(1)) union all select '项目' AS `resource_cn_name`,'project' AS `resource_en_name`,'项目信息' AS `resource_remark`,`project`.`project_id` AS `project_id`,'project_id' AS `primary_key`,`project`.`project_id` AS `primary_key_value`,'' AS `resource_type`,'' AS `data_status`,concat(`project`.`project_id`,'-',`project`.`project_cn`,'-',`project`.`project_en`) AS `resource_value`,0 AS `show_cluster`,0 AS `show_feature`,0 AS `show_install`,0 AS `show_other`,2 AS `system_show` from `project` where (`project`.`del_flag` = 0) union all select '服务器' AS `resource_cn_name`,'asset' AS `resource_en_name`,'服务器信息' AS `resource_remark`,`view_assets`.`view_user_project_id` AS `project_id`,'asset_id' AS `primary_key`,`view_assets`.`view_asset_id` AS `primary_key_value`,`view_assets`.`view_host_role_cn` AS `resource_type`,concat(',"view_recycle_type":',`view_assets`.`view_recycle_type`) AS `data_status`,concat(if((`view_assets`.`view_user_project_en` is null),'',concat(`view_assets`.`view_user_project_en`,'_')),if((`view_assets`.`view_outer_ip` is null),'',concat(`view_assets`.`view_outer_ip`,'_')),if((`view_assets`.`view_inner_ip` is null),'',concat(`view_assets`.`view_inner_ip`,'_')),if((`view_assets`.`view_provider_name_en` is null),'',concat(`view_assets`.`view_provider_name_en`,'_')),if((`view_assets`.`view_en_host_role` is null),'',`view_assets`.`view_en_host_role`)) AS `resource_value`,1 AS `show_cluster`,-(1) AS `show_feature`,1 AS `show_install`,1 AS `show_other`,3 AS `system_show` from `view_assets` where ((`view_assets`.`view_user_project_deleted` = -(1)) and (`view_assets`.`view_asset_del_flag` = 0)) union all select '平台' AS `resource_cn_name`,'platform' AS `resource_en_name`,'平台信息' AS `resource_remark`,`pr`.`project_id` AS `project_id`,'id' AS `primary_key`,`pf`.`id` AS `primary_key_value`,'' AS `resource_type`,(case `pf`.`del_flag` when '1' then concat(',"view_recycle_type":1') when '0' then concat(',"view_recycle_type":2') end) AS `data_status`,concat(`pr`.`project_en`,'_',`pf`.`platform_id`,'_',`pf`.`platform_cn`,'_',`pf`.`platform_en`) AS `resource_value`,1 AS `show_cluster`,1 AS `show_feature`,1 AS `show_install`,1 AS `show_other`,3 AS `system_show` from (`platform` `pf` join `project` `pr`) where ((`pf`.`project_id` = `pr`.`project_id`) and (`pr`.`del_flag` = 0)) union all select '功能服' AS `resource_cn_name`,'feature_server_info' AS `resource_en_name`,'功能服信息' AS `resource_remark`,`pr`.`project_id` AS `project_id`,'feature_server_id' AS `primary_key`,`fsi`.`feature_server_id` AS `primary_key_value`,json_unquote(json_extract(`fsi`.`feature_server_info`,'$.type')) AS `resource_type`,(case `fsi`.`del_flag` when '1' then concat(',"view_recycle_type":1') when '0' then concat(',"view_recycle_type":2') end) AS `data_status`,json_object('project',`pr`.`project_en`,'feature_remark',`fsi`.`remark`,'type',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.type')),'domain',if((json_unquote(json_extract(`fsi`.`feature_server_info`,'$.domain')) is null),'',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.domain'))),'inner_ip',if((json_unquote(json_extract(`fsi`.`feature_server_info`,'$.ip[0].inner_ip')) is null),'',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.ip[0].inner_ip'))),'outer_ip',if((json_unquote(json_extract(`fsi`.`feature_server_info`,'$.ip[0].outer_ip')) is null),'',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.ip[0].outer_ip'))),'db_host',if((json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.db_host')) is null),'',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.db_host'))),'db_port',if((json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.db_port')) is null),'',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.db_port'))),'db_name',if((json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.db_name')) is null),'',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.db_name'))),'single_db_log',if((json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.single_db_log')) is null),'',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.single_db_log'))),'cross_db_name',if((json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.cross_db_name')) is null),'',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.cross_db_name'))),'cross_db_log',if((json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.cross_db_log')) is null),'',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.db_info.cross_db_log'))),'base_dir',if((json_unquote(json_extract(`fsi`.`feature_server_info`,'$.base_dir')) is null),'',json_unquote(json_extract(`fsi`.`feature_server_info`,'$.base_dir')))) AS `resource_value`,1 AS `show_cluster`,1 AS `show_feature`,-(1) AS `show_install`,1 AS `show_other`,3 AS `system_show` from (`feature_server_info` `fsi` join `project` `pr`) where ((`fsi`.`project_id` = `pr`.`project_id`) and (`pr`.`del_flag` = 0))) `table_data` left join (select `sm1`.`table_name` AS `table_name`,concat(`sm2`.`vue_path`,'/',`sm1`.`vue_path`) AS `url` from (`sys_menu` `sm1` join `sys_menu` `sm2`) where ((`sm1`.`parent_id` = `sm2`.`id`) and (`sm1`.`tp` = 1) and (`sm1`.`del_flag` = 0) and (`sm2`.`del_flag` = 0) and (`sm1`.`vue_path` like '%List'))) `table_url` on((`table_url`.`table_name` = `table_data`.`resource_en_name`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_sh_platform_info`
--

/*!50001 DROP VIEW IF EXISTS `view_sh_platform_info`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_sh_platform_info` (`view_project_id`,`view_project_en`,`view_project_cn`,`view_platform_en`,`view_platform_info`) AS select `A`.`view_project_id` AS `view_project_id`,`A`.`view_project_en` AS `view_project_en`,`A`.`view_project_cn` AS `view_project_cn`,`A`.`view_platform_en` AS `view_platform_en`,concat(`A`.`view_platform_en`,')export MMO_SRC_DB="',convert(`A`.`source_db_game` using utf8mb3),'";export LOG_SRC_DB="',convert(`A`.`source_db_log` using utf8mb3),'";;') AS `platform_feature_info` from (select `view_platform_feature`.`view_project_id` AS `view_project_id`,`view_platform_feature`.`view_project_en` AS `view_project_en`,`view_platform_feature`.`view_project_cn` AS `view_project_cn`,`view_platform_feature`.`view_platform_en` AS `view_platform_en`,(case when regexp_like(`view_platform_feature`.`view_platform_en`,'cross') then json_unquote(json_extract(`view_platform_feature`.`view_platform_feature_info`,'$.source_feature_info.db_info.cross_db_name')) else json_unquote(json_extract(`view_platform_feature`.`view_platform_feature_info`,'$.source_feature_info.db_info.db_name')) end) AS `source_db_game`,(case when regexp_like(`view_platform_feature`.`view_platform_en`,'cross') then json_unquote(json_extract(`view_platform_feature`.`view_platform_feature_info`,'$.source_feature_info.db_info.cross_db_log')) else json_unquote(json_extract(`view_platform_feature`.`view_platform_feature_info`,'$.source_feature_info.db_info.single_db_log')) end) AS `source_db_log` from `view_platform_feature`) `A` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_sh_split_ip_and_region`
--

/*!50001 DROP VIEW IF EXISTS `view_sh_split_ip_and_region`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_sh_split_ip_and_region` (`view_project_id`,`view_project_en`,`view_project_cn`,`view_single_ip_pool`,`view_cross_ip_pool`,`view_split_single_region`,`view_split_cross_region`) AS select `A`.`view_project_id` AS `view_project_id`,`A`.`view_project_en` AS `view_project_en`,`A`.`view_project_cn` AS `view_project_cn`,concat(`A`.`view_cluster_name`,'_IP_POOL=\'',convert(`A`.`source_inner_ip` using utf8mb3),':',convert(`A`.`source_outer_ip` using utf8mb3),':',convert(`A`.`source_ssh_port` using utf8mb3),'|',convert(`A`.`center_inner_ip` using utf8mb3),':',convert(`A`.`center_outer_ip` using utf8mb3),':',convert(`A`.`center_ssh_port` using utf8mb3),':',convert(`A`.`center_db_name` using utf8mb3),':',convert(`A`.`center_domain` using utf8mb3),'|',convert(`A`.`source_single_db_game` using utf8mb3),'|',convert(`A`.`source_single_db_log` using utf8mb3),'\'') AS `single_ip_pool`,concat(`A`.`view_cluster_name`,'CR_IP_POOL=\'',convert(`A`.`source_inner_ip` using utf8mb3),':',convert(`A`.`source_outer_ip` using utf8mb3),':',convert(`A`.`source_ssh_port` using utf8mb3),'|',convert(`A`.`center_inner_ip` using utf8mb3),':',convert(`A`.`center_outer_ip` using utf8mb3),':',convert(`A`.`center_ssh_port` using utf8mb3),':',convert(`A`.`center_db_name` using utf8mb3),':',convert(`A`.`center_domain` using utf8mb3),'|',convert(`A`.`source_cross_db_game` using utf8mb3),'|',convert(`A`.`source_cross_db_log` using utf8mb3),'\'') AS `cross_ip_pool`,concat(`A`.`view_cluster_name`,')split_ip_pool ${',`A`.`view_cluster_name`,'_IP_POOL};;') AS `split_single_region`,concat(`A`.`view_cluster_name`,'CR)split_ip_pool ${',`A`.`view_cluster_name`,'CR_IP_POOL};;') AS `split_cross_region` from (select `view_project_cluster_info`.`view_project_id` AS `view_project_id`,`view_project_cluster_info`.`view_project_en` AS `view_project_en`,`view_project_cluster_info`.`view_project_cn` AS `view_project_cn`,`view_project_cluster_info`.`view_cluster_name` AS `view_cluster_name`,json_unquote(json_extract(`view_project_cluster_info`.`view_source_cluster_info`,'$.ip[0].inner_ip')) AS `source_inner_ip`,json_unquote(json_extract(`view_project_cluster_info`.`view_source_cluster_info`,'$.ip[0].outer_ip')) AS `source_outer_ip`,json_unquote(json_extract(`view_project_cluster_info`.`view_source_cluster_info`,'$.ip[0].ssh_port')) AS `source_ssh_port`,json_unquote(json_extract(`view_project_cluster_info`.`view_center_cluster_info`,'$.ip[0].inner_ip')) AS `center_inner_ip`,json_unquote(json_extract(`view_project_cluster_info`.`view_center_cluster_info`,'$.ip[0].outer_ip')) AS `center_outer_ip`,json_unquote(json_extract(`view_project_cluster_info`.`view_center_cluster_info`,'$.ip[0].ssh_port')) AS `center_ssh_port`,json_unquote(json_extract(`view_project_cluster_info`.`view_center_cluster_info`,'$.db_info.db_name')) AS `center_db_name`,json_unquote(json_extract(`view_project_cluster_info`.`view_center_cluster_info`,'$.domain')) AS `center_domain`,json_unquote(json_extract(`view_project_cluster_info`.`view_source_cluster_info`,'$.db_info.db_name')) AS `source_single_db_game`,json_unquote(json_extract(`view_project_cluster_info`.`view_source_cluster_info`,'$.db_info.single_db_log')) AS `source_single_db_log`,json_unquote(json_extract(`view_project_cluster_info`.`view_source_cluster_info`,'$.db_info.cross_db_name')) AS `source_cross_db_game`,json_unquote(json_extract(`view_project_cluster_info`.`view_source_cluster_info`,'$.db_info.cross_db_log')) AS `source_cross_db_log` from `view_project_cluster_info`) `A` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_task_logs_relationship_with_id`
--

/*!50001 DROP VIEW IF EXISTS `view_task_logs_relationship_with_id`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_task_logs_relationship_with_id` (`v_log_inner`,`v_pid`,`v_id`,`v_step`) AS select (case `B`.`task_status` when 5 then concat('{"label":"[出错删除]',substring_index(`tasks`.`name`,'(',1),'",',convert(`B`.`log_inner` using utf8mb3),'}') when 2 then concat('{"label":"[ERROR]',substring_index(`tasks`.`name`,'(',1),'",','"is_open":1,',convert(`B`.`log_inner` using utf8mb3),'}') else concat('{"label":"',substring_index(`tasks`.`name`,'(',1),'",','"is_open":1,',convert(`B`.`log_inner` using utf8mb3),'}') end) AS `log_inner`,`tasks`.`pid` AS `pid`,`tasks`.`id` AS `id`,`B`.`step` AS `step` from (`tasks` join (select (case `A`.`task_status` when 5 then if((`A`.`log_inner` = '0'),'0',concat('"children":[',group_concat(`A`.`log_inner` order by `A`.`tasks_time` DESC separator ','),']')) else if((`A`.`log_inner` = '0'),concat('"children":[{"label":"实时日志","value":-1',',"tasks_id":',`A`.`tasks_id`,'}]'),concat('"children":[{"label":"实时日志","value":-1',',"tasks_id":',`A`.`tasks_id`,'},',group_concat(`A`.`log_inner` order by `A`.`tasks_time` DESC separator ','),']')) end) AS `log_inner`,`A`.`tasks_id` AS `tasks_id`,`A`.`task_status` AS `task_status`,`A`.`step` AS `step` from (select (case `new_task_log_histroy`.`task_status` when 5 then concat('{"label":"[出错删除]',date_format(from_unixtime(`new_task_log_histroy`.`tasks_time`),'%Y%m%d%H%i'),'","value":',`new_task_log_histroy`.`id`,',"tasks_id":',`new_task_log_histroy`.`tasks_id`,'}') else if((`new_task_log_histroy`.`id` = '0'),'0',concat('{"label":',date_format(from_unixtime(`new_task_log_histroy`.`tasks_time`),'%Y%m%d%H%i'),',"value":',`new_task_log_histroy`.`id`,',"tasks_id":',`new_task_log_histroy`.`tasks_id`,'}')) end) AS `log_inner`,`new_task_log_histroy`.`tasks_id` AS `tasks_id`,`new_task_log_histroy`.`tasks_time` AS `tasks_time`,`new_task_log_histroy`.`task_status` AS `task_status`,`new_task_log_histroy`.`step` AS `step` from (select ifnull(`task_log_history`.`id`,'0') AS `id`,`tb_tid`.`id` AS `tasks_id`,`tb_tid`.`name` AS `name`,`tb_tid`.`types` AS `types`,`tb_tid`.`pid` AS `pid`,`tb_tid`.`task_status` AS `task_status`,`tb_tid`.`step` AS `step`,`task_log_history`.`tasks_time` AS `tasks_time` from ((select `tasks`.`id` AS `id`,`tasks`.`level` AS `level`,`tasks`.`name` AS `name`,`tasks`.`types` AS `types`,`tasks`.`pid` AS `pid`,`tasks`.`task_status` AS `task_status`,concat(rpad(`tasks`.`level`,2,0),rpad(`tb_cid`.`task_step`,2,0),rpad(`tasks`.`task_step`,2,0)) AS `step` from (`tasks` join (select `tasks`.`id` AS `id`,`tasks`.`task_step` AS `task_step` from `tasks`) `tb_cid`) where ((`tasks`.`pid` = `tb_cid`.`id`) and (`tasks`.`task_status` in (-(1),2,3,5)) and (`tasks`.`level` = 3))) `tb_tid` left join (select `task_log_histroy`.`id` AS `id`,`task_log_histroy`.`tasks_id` AS `tasks_id`,`task_log_histroy`.`tasks_time` AS `tasks_time` from `task_log_histroy`) `task_log_history` on((`tb_tid`.`id` = `task_log_history`.`tasks_id`))) order by `tb_tid`.`step`,`tasks_id` desc) `new_task_log_histroy`) `A` group by `A`.`tasks_id`) `B`) where (`tasks`.`id` = `B`.`tasks_id`) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_user_asset`
--

/*!50001 DROP VIEW IF EXISTS `view_user_asset`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_user_asset` (`view_asset_id`,`view_user_id`,`view_cluster_name`) AS select `A`.`binding_id` AS `binding_id`,ifnull(group_concat(`A`.`label_user_id` order by `A`.`label_user_id` ASC separator ','),'') AS `label_user_id`,group_concat(distinct `A`.`cluster_name` separator ',') AS `cluster_name` from (select `label_asset`.`label_id` AS `label_id`,`label_asset`.`label_type` AS `label_type`,`label_asset`.`label_name` AS `label_name`,`label_asset`.`label_values` AS `label_values`,`label_asset`.`label_remark` AS `label_remark`,`label_asset`.`del_flag` AS `del_flag`,`label_asset`.`resource_en` AS `resource_en`,`label_asset`.`binding_id` AS `binding_id`,`label_asset`.`project_id` AS `project_id`,`label_user`.`id` AS `label_user_id`,ifnull(`label_user`.`name`,'') AS `label_user_name`,ifnull(`label_user`.`nick_name`,'') AS `label_user_nick_name`,(case `label_asset`.`label_type` when 1 then upper(replace(`label_asset`.`label_values`,'cluster_','')) end) AS `cluster_name` from ((select `l`.`label_id` AS `label_id`,`l`.`label_type` AS `label_type`,`l`.`label_name` AS `label_name`,`l`.`label_values` AS `label_values`,`l`.`label_remark` AS `label_remark`,`l`.`del_flag` AS `del_flag`,`lg`.`resource_en` AS `resource_en`,`lg`.`binding_id` AS `binding_id`,`lg`.`project_id` AS `project_id` from (`label` `l` left join `label_global` `lg` on((`l`.`label_id` = `lg`.`label_id`))) where ((`l`.`del_flag` = 0) and (`lg`.`resource_en` = 'asset')) order by `l`.`label_type`,`l`.`label_id`) `label_asset` left join (select `sys_user`.`id` AS `id`,`sys_user`.`name` AS `name`,`sys_user`.`nick_name` AS `nick_name`,`sys_user`.`avatar` AS `avatar`,`sys_user`.`password` AS `password`,`sys_user`.`salt` AS `salt`,`sys_user`.`email` AS `email`,`sys_user`.`mobile` AS `mobile`,`sys_user`.`status` AS `status`,`sys_user`.`dept_id` AS `dept_id`,`sys_user`.`create_by` AS `create_by`,`sys_user`.`create_time` AS `create_time`,`sys_user`.`last_update_by` AS `last_update_by`,`sys_user`.`last_update_time` AS `last_update_time`,`sys_user`.`del_flag` AS `del_flag`,`A`.`id` AS `label_global_id`,`A`.`label_id` AS `label_id`,`A`.`resource_en` AS `resource_en`,`A`.`binding_id` AS `binding_id`,`A`.`project_id` AS `project_id` from (`sys_user` left join (select `label_global`.`id` AS `id`,`label_global`.`label_id` AS `label_id`,`label_global`.`resource_en` AS `resource_en`,`label_global`.`binding_id` AS `binding_id`,`label_global`.`project_id` AS `project_id` from `label_global` where (`label_global`.`resource_en` = 'sys_user')) `A` on((`sys_user`.`id` = `A`.`binding_id`))) where (`sys_user`.`del_flag` = 0)) `label_user` on((`label_asset`.`label_id` = `label_user`.`label_id`)))) `A` group by `A`.`binding_id` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `view_user_strategy`
--

/*!50001 DROP VIEW IF EXISTS `view_user_strategy`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_user_strategy` (`view_sys_user_id`,`view_sys_user_name`,`view_stgroup_st_json`) AS select `A`.`sys_user_id` AS `sys_user_id`,`A`.`sys_user_name` AS `sys_user_name`,concat('[',group_concat(`A`.`stgroup_st_json` separator ','),']') AS `stgroup_st_json` from (select `ug`.`sys_user_id` AS `sys_user_id`,`ug`.`sys_user_name` AS `sys_user_name`,`ug`.`sys_user_nick_name` AS `sys_user_nick_name`,`ug`.`sys_ugroup_id` AS `sys_ugroup_id`,`ug`.`sys_ugroup_name` AS `sys_ugroup_name`,concat(ifnull(`ug_st`.`sys_stgroup_st_name`,''),if((length(trim(ifnull(`u_st`.`sys_stgroup_st_name`,''))) = 0),'',if((length(trim(ifnull(`ug_st`.`sys_stgroup_st_name`,''))) = 0),`u_st`.`sys_stgroup_st_name`,concat('+',`u_st`.`sys_stgroup_st_name`)))) AS `stgroup_st_name`,concat(ifnull(`ug_st`.`sys_stgroup_st_json`,''),if((length(trim(ifnull(`u_st`.`sys_stgroup_st_json`,''))) = 0),'',if((length(trim(ifnull(`ug_st`.`sys_stgroup_st_json`,''))) = 0),`u_st`.`sys_stgroup_st_json`,concat(',',`u_st`.`sys_stgroup_st_json`)))) AS `stgroup_st_json`,concat(ifnull(`ug_st`.`st_identify`,''),if((length(trim(ifnull(`u_st`.`st_identify`,''))) = 0),'',if((length(trim(ifnull(`ug_st`.`st_identify`,''))) = 0),`u_st`.`st_identify`,concat('+',`u_st`.`st_identify`)))) AS `st_identify` from (((select `sys_user`.`id` AS `sys_user_id`,`sys_user`.`name` AS `sys_user_name`,`sys_user`.`nick_name` AS `sys_user_nick_name`,`sys_ugroup`.`id` AS `sys_ugroup_id`,`sys_ugroup`.`ug_name` AS `sys_ugroup_name` from ((`sys_user` left join `sys_user_ugroup` on((`sys_user`.`id` = `sys_user_ugroup`.`user_id`))) left join `sys_ugroup` on((`sys_user_ugroup`.`ugroup_id` = `sys_ugroup`.`id`)))) `ug` left join (select `sys_stgroup`.`id` AS `sys_stgroup_id`,`sys_stgroup`.`st_name` AS `sys_stgroup_st_name`,`sys_stgroup`.`st_json` AS `sys_stgroup_st_json`,`sys_stgroup`.`st_remark` AS `sys_stgroup_st_remark`,`sys_stgroup_ugroup`.`ugroup_id` AS `sys_stgroup_ugroup_id`,'ugroup' AS `st_identify` from (`sys_stgroup` left join `sys_stgroup_ugroup` on((`sys_stgroup`.`id` = `sys_stgroup_ugroup`.`stgroup_id`)))) `ug_st` on((`ug`.`sys_ugroup_id` = `ug_st`.`sys_stgroup_ugroup_id`))) left join (select `sys_stgroup`.`id` AS `sys_stgroup_id`,`sys_stgroup`.`st_name` AS `sys_stgroup_st_name`,`sys_stgroup`.`st_json` AS `sys_stgroup_st_json`,`sys_stgroup`.`st_remark` AS `sys_stgroup_st_remark`,`sys_stgroup_user`.`user_id` AS `sys_stgroup_user_id`,'user' AS `st_identify` from (`sys_stgroup` left join `sys_stgroup_user` on((`sys_stgroup`.`id` = `sys_stgroup_user`.`stgroup_id`)))) `u_st` on((`ug`.`sys_user_id` = `u_st`.`sys_stgroup_user_id`)))) `A` group by `A`.`sys_user_id` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!50112 SET @disable_bulk_load = IF (@is_rocksdb_supported, 'SET SESSION rocksdb_bulk_load = @old_rocksdb_bulk_load', 'SET @dummy_rocksdb_bulk_load = 0') */;
/*!50112 PREPARE s FROM @disable_bulk_load */;
/*!50112 EXECUTE s */;
/*!50112 DEALLOCATE PREPARE s */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-26 19:48:15
