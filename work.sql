/*
 Navicat Premium Data Transfer

 Source Server         : xiaofuzi
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost:3306
 Source Schema         : work

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 14/12/2019 16:36:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for work
-- ----------------------------
DROP TABLE IF EXISTS `work`;
CREATE TABLE `work`  (
  `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `messages` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `date` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of work
-- ----------------------------
INSERT INTO `work` VALUES ('1576306771', 'sdfsdfsdfs', 'sdfsdxzxxxxxxxx', '123');
INSERT INTO `work` VALUES ('1576306786', 'sdfsdfsdfs', 'sdfsdxzxxxxxxxx', '123124');
INSERT INTO `work` VALUES ('1576312507', 'sdfsdfsdfs', 'sdfsdxzxxxxxxxx', '343');
INSERT INTO `work` VALUES ('1576312508', 'sdfsdfsdfs', 'sdfsdxzxxxxxxxx', '34');
INSERT INTO `work` VALUES ('1576312509', 'sdfsdfsdfs', 'sdfsdxzxxxxxxxx', '');
INSERT INTO `work` VALUES ('1576312510', 'sdfsdfsdfs', 'sdfsdxzxxxxxxxx', '');

SET FOREIGN_KEY_CHECKS = 1;
