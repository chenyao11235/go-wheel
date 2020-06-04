/*
Navicat MySQL Data Transfer

Source Server         : mysql57
Source Server Version : 50721
Source Host           : localhost:3307
Source Database       : swoft2

Target Server Type    : MYSQL
Target Server Version : 50721
File Encoding         : 65001

Date: 2019-11-11 15:55:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for book_kinds
-- ----------------------------
DROP TABLE IF EXISTS `book_kinds`;
CREATE TABLE `book_kinds` (
  `kind_id` int(11) NOT NULL AUTO_INCREMENT,
  `kind_name` varchar(50) DEFAULT NULL,
  `pid` int(11) DEFAULT NULL,
  PRIMARY KEY (`kind_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of book_kinds
-- ----------------------------
INSERT INTO `book_kinds` VALUES ('1', '编程语言', '0');
INSERT INTO `book_kinds` VALUES ('2', 'PHP', '1');
INSERT INTO `book_kinds` VALUES ('3', 'Java', '1');
INSERT INTO `book_kinds` VALUES ('4', 'Go', '1');
INSERT INTO `book_kinds` VALUES ('5', '网站开发', '1');
INSERT INTO `book_kinds` VALUES ('6', 'Python', '1');
INSERT INTO `book_kinds` VALUES ('7', '其他', '1');
INSERT INTO `book_kinds` VALUES ('8', 'c/c++', '1');
