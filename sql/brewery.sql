/*
 Navicat Premium Data Transfer

 Source Server         : 101.126.153.89
 Source Server Type    : MySQL
 Source Server Version : 50729 (5.7.29)
 Source Host           : 101.126.153.89:3306
 Source Schema         : brewery

 Target Server Type    : MySQL
 Target Server Version : 50729 (5.7.29)
 File Encoding         : 65001

 Date: 22/10/2024 15:03:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for guide
-- ----------------------------
DROP TABLE IF EXISTS `guide`;
CREATE TABLE `guide`  (
                          `id` bigint(20) NOT NULL AUTO_INCREMENT,
                          `owner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                          `sales_status` tinyint(4) NULL DEFAULT NULL,
                          `price` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                          `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
                          `like_count` bigint(20) NULL DEFAULT NULL,
                          `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                          `create_time` bigint(20) NULL DEFAULT NULL COMMENT '秒级时间戳',
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of guide
-- ----------------------------
INSERT INTO `guide` VALUES (2, '122213', 0, '', '上海三日游计划', 11, NULL, 1729341896);
INSERT INTO `guide` VALUES (3, '122213', 0, '', '北京游玩计划', 2, NULL, 1729341911);
INSERT INTO `guide` VALUES (4, '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266', 0, '', '11111', 0, NULL, 1729428696);
INSERT INTO `guide` VALUES (5, '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266', 0, '', '早上前往天安门广场看升旗，感受庄严氛围。接着参观故宫，领略古代宫殿建筑的宏伟。中午在附近品尝烤鸭。下午漫步景山公园，登顶可俯瞰故宫全景。然后去南锣鼓巷，体验老北京胡同风情，品尝特色小吃。这一天既能感受历史文化底蕴，又能享受美食，花费可控制在150元左右（不包含交通费用）。', 0, NULL, 1729577275);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                         `fans_count` bigint(20) NULL DEFAULT NULL,
                         `avatar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
                         `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
                         `create_time` bigint(20) NULL DEFAULT NULL,
                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('122213', 2, 'https://avatars.githubusercontent.com/u/13827299', '报周进院发值织电。达流并之量红。因层京路联消历用。难白京。矿道意性结。劳究统报市思出将大。', 1729346577);

SET FOREIGN_KEY_CHECKS = 1;
