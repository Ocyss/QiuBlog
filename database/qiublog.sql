/*
 Navicat Premium Data Transfer

 Source Server         : 本机MySQL8
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : localhost:3307
 Source Schema         : qiublog

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 29/10/2022 18:13:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标题',
  `img` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '首图',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '描述',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '内容',
  `cid` bigint UNSIGNED NULL DEFAULT NULL COMMENT '分类ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_article_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_category_articles`(`cid` ASC) USING BTREE,
  CONSTRAINT `fk_category_articles` FOREIGN KEY (`cid`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, '2022-10-02 15:36:05', '2022-10-02 15:36:05', NULL, '测试', '', '123', '<p>213</p>', 1);
INSERT INTO `article` VALUES (2, '2022-10-02 15:36:55', '2022-10-02 15:36:55', NULL, 'sssssss试试水', '', '312', '<p>333</p>', 1);
INSERT INTO `article` VALUES (3, '2022-10-03 18:09:12', '2022-10-03 18:09:12', NULL, '测试下超长描述', 'https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Article/1666158218039709800.png', '重启奥斯卡了的静安寺的阿萨德加速；‘两大类；’短手；垃圾袋；拉手机登陆奥斯卡辣鸡的嘎斯柯达嘎进口水果大号SV大V啥老司机还是的权威逾期未普洱有自行车安居客收到来自行程轨迹奥斯卡过街老鼠对嘎就开始感动哭了基安上周星驰不能满足小丑你们不在现', '<p>可以，很不错啊阿萨德</p>', 15);
INSERT INTO `article` VALUES (4, '2022-10-19 13:45:36', '2022-10-29 15:58:55', NULL, '一个好的博客就需要大量的测试信息,可以，完美', 'https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Article/1666158292184457000.png', '描述？里面有好看的😁', '<h1>H13333</h1><p>666asd</p><p><img src=\"https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Article/1666158292058500500.png\" alt=\"\" data-href=\"\" style=\"\"/><img src=\"https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Article/1666158292184457000.png\" alt=\"\" data-href=\"\" style=\"\"/><img src=\"https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Article/1666158292301009400.png\" alt=\"\" data-href=\"\" style=\"\"/></p><p><br></p><p><br></p><p><br></p><hr/><p><br></p><table style=\"width: auto;\"><tbody><tr><th colSpan=\"1\" rowSpan=\"1\" width=\"auto\">123</th><th colSpan=\"1\" rowSpan=\"1\" width=\"auto\">123</th><th colSpan=\"1\" rowSpan=\"1\" width=\"auto\">32</th><th colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></th><th colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></th><th colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></th><th colSpan=\"1\" rowSpan=\"1\" width=\"auto\">qwe</th></tr><tr><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\">jhk</td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td></tr><tr><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td></tr><tr><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\">123</td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\">qwe</td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\">sda</td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\">qwe</td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td></tr><tr><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td></tr><tr><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td></tr><tr><td colspan=\"1\" rowspan=\"1\" width=\"auto\" style=\"text-align: justify;\">wqe<br><br></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\">asd</td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td></tr><tr><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\">asd</td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td></tr><tr><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\">qwe</td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td></tr><tr><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td><td colSpan=\"1\" rowSpan=\"1\" width=\"auto\"></td></tr></tbody></table><p><br></p><p><br></p><p>😀</p>', 2);

-- ----------------------------
-- Table structure for article_tags
-- ----------------------------
DROP TABLE IF EXISTS `article_tags`;
CREATE TABLE `article_tags`  (
  `tags_id` bigint UNSIGNED NOT NULL,
  `article_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`tags_id`, `article_id`) USING BTREE,
  INDEX `fk_article_tags_article`(`article_id` ASC) USING BTREE,
  CONSTRAINT `fk_article_tags_article` FOREIGN KEY (`article_id`) REFERENCES `article` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_article_tags_tags` FOREIGN KEY (`tags_id`) REFERENCES `tags` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article_tags
-- ----------------------------
INSERT INTO `article_tags` VALUES (1, 1);
INSERT INTO `article_tags` VALUES (2, 1);
INSERT INTO `article_tags` VALUES (3, 1);
INSERT INTO `article_tags` VALUES (4, 1);
INSERT INTO `article_tags` VALUES (5, 1);
INSERT INTO `article_tags` VALUES (6, 1);
INSERT INTO `article_tags` VALUES (7, 1);
INSERT INTO `article_tags` VALUES (3, 2);
INSERT INTO `article_tags` VALUES (4, 2);
INSERT INTO `article_tags` VALUES (7, 2);
INSERT INTO `article_tags` VALUES (8, 2);
INSERT INTO `article_tags` VALUES (9, 2);
INSERT INTO `article_tags` VALUES (10, 2);
INSERT INTO `article_tags` VALUES (2, 4);
INSERT INTO `article_tags` VALUES (11, 4);
INSERT INTO `article_tags` VALUES (12, 4);
INSERT INTO `article_tags` VALUES (13, 4);

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名',
  `mid` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单子项ID',
  `homeshow` tinyint(1) NOT NULL DEFAULT 1 COMMENT '主页是否显示',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE,
  INDEX `fk_menuchild_cids`(`mid` ASC) USING BTREE,
  CONSTRAINT `fk_category_menuchild` FOREIGN KEY (`mid`) REFERENCES `menuchild` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_menuchild_cids` FOREIGN KEY (`mid`) REFERENCES `menuchild` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, '博客开发', 1, 0);
INSERT INTO `category` VALUES (2, '全栈', 1, 1);
INSERT INTO `category` VALUES (3, 'Magisk', 2, 1);
INSERT INTO `category` VALUES (4, 'LSPosed', 2, 1);
INSERT INTO `category` VALUES (5, 'Vue.js', 1, 1);
INSERT INTO `category` VALUES (6, 'Flask', 1, 1);
INSERT INTO `category` VALUES (7, 'Python', 1, 1);
INSERT INTO `category` VALUES (8, 'GoLang', 1, 1);
INSERT INTO `category` VALUES (9, 'MySql', 1, 1);
INSERT INTO `category` VALUES (10, 'MongoDB', 1, 1);
INSERT INTO `category` VALUES (11, '酷安', 2, 1);
INSERT INTO `category` VALUES (12, 'ROM推荐', 2, 1);
INSERT INTO `category` VALUES (13, '摸鱼', 3, 1);
INSERT INTO `category` VALUES (14, '敲代码', 3, 1);
INSERT INTO `category` VALUES (15, '散步', 3, 0);
INSERT INTO `category` VALUES (16, '见世面', 3, 0);
INSERT INTO `category` VALUES (17, '薅羊毛', 3, 0);
INSERT INTO `category` VALUES (18, '抓包', 1, 1);
INSERT INTO `category` VALUES (19, '爬虫', 1, 1);
INSERT INTO `category` VALUES (20, '她', 3, 0);
INSERT INTO `category` VALUES (21, '音乐', 3, 0);
INSERT INTO `category` VALUES (22, '好物安利', 3, 0);
INSERT INTO `category` VALUES (23, '捡垃圾', 3, 0);

-- ----------------------------
-- Table structure for menuchild
-- ----------------------------
DROP TABLE IF EXISTS `menuchild`;
CREATE TABLE `menuchild`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `sort` bigint UNSIGNED NULL DEFAULT NULL COMMENT '排序字段',
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单名',
  `ename` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '英文名',
  `logo` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图标名',
  `link` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '路由名',
  `parent_id` bigint UNSIGNED NOT NULL COMMENT '父级id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE,
  UNIQUE INDEX `ename`(`ename` ASC) USING BTREE,
  UNIQUE INDEX `link`(`link` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menuchild
-- ----------------------------
INSERT INTO `menuchild` VALUES (1, 1, '技术', 'BackEnd', '<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" viewBox=\"0 0 512 512\"><path d=\"M463.55 272.13H400v-48.2q0-4.32-.27-8.47c29.57-27.88 32.25-64.63 32.27-103c0-8.61-6.64-16-15.25-16.41A16 16 0 0 0 400 112c0 28-1.86 48.15-9.9 63.84c-19.22-41.15-65.78-63.91-134.1-63.91c-39.8 0-74.19 9.13-99.43 26.39c-14.9 10.19-26.2 22.91-33.7 37.72C114 160.65 112 141 112 112.46c0-8.61-6.6-16-15.2-16.44A16 16 0 0 0 80 112c0 37.63 2.61 73.73 32.44 101.63q-.43 5.06-.44 10.3v48.2H48.45c-8.61 0-16 6.62-16.43 15.23a16 16 0 0 0 16 16.77h64V320a143.32 143.32 0 0 0 10.39 53.69C96.74 396.64 80.18 422 80 463.34c0 8.74 6.62 16.3 15.36 16.65A16 16 0 0 0 112 464c0-27.66 9.1-44.71 26.17-61.32A144.37 144.37 0 0 0 220 459.42a16 16 0 0 0 20-15.49V192.45c0-8.61 6.62-16 15.23-16.43A16 16 0 0 1 272 192v251.93a16 16 0 0 0 20 15.49a144.4 144.4 0 0 0 81.82-56.74c17 16.54 26.09 33.52 26.17 60.95a16.27 16.27 0 0 0 15.1 16.37A16 16 0 0 0 432 464c0-41.68-16.6-67.23-42.39-90.31A143.32 143.32 0 0 0 400 320v-15.87h64a16 16 0 0 0 16-16.77c-.42-8.61-7.84-15.23-16.45-15.23z\" fill=\"currentColor\"></path><path d=\"M321.39 104l.32.09c13.57 3.8 25.07-10.55 18.2-22.85A95.86 95.86 0 0 0 256.21 32h-.42a95.87 95.87 0 0 0-84.19 50.13c-6.84 12.58 5.14 27 18.84 22.86c19.71-6 41.79-9.06 65.56-9.06c24.09 0 46.09 2.72 65.39 8.07z\" fill=\"currentColor\"></path></svg>', 'backend', 0);
INSERT INTO `menuchild` VALUES (2, 2, '玩机', 'PlayPhone', '<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" viewBox=\"0 0 512 512\"><path d=\"M408 48H104a72.08 72.08 0 0 0-72 72v192a72.08 72.08 0 0 0 72 72h24v64a16 16 0 0 0 26.25 12.29L245.74 384H408a72.08 72.08 0 0 0 72-72V120a72.08 72.08 0 0 0-72-72zM160 248a32 32 0 1 1 32-32a32 32 0 0 1-32 32zm96 0a32 32 0 1 1 32-32a32 32 0 0 1-32 32zm96 0a32 32 0 1 1 32-32a32 32 0 0 1-32 32z\" fill=\"currentColor\"></path></svg>', 'PlayPhone', 0);
INSERT INTO `menuchild` VALUES (3, 3, '生活', 'Life', '<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" viewBox=\"0 0 512 512\"><path d=\"M512 256c0-16.54-14.27-46.76-45.61-74a207.06 207.06 0 0 0-60.28-36.12a3.15 3.15 0 0 0-3.93 1.56c-.15.29-.3.57-.47.86l-9.59 15.9a183.24 183.24 0 0 0 .07 183.78l.23.39l8.74 16a4 4 0 0 0 4.94 1.82C479.63 337.42 512 281.49 512 256zm-93.92-.14a16 16 0 1 1 13.79-13.79a16 16 0 0 1-13.79 13.79z\" fill=\"currentColor\"></path><path d=\"M335.45 256a214.8 214.8 0 0 1 29.08-108l.12-.21l4.62-7.67a4 4 0 0 0-2.59-6a284.29 284.29 0 0 0-39.26-5.39a7.94 7.94 0 0 1-4.29-1.6c-19.28-14.66-57.5-40.3-96.46-46.89a16 16 0 0 0-18 20.18l10.62 37.17a4 4 0 0 1-2.42 4.84c-36.85 13.69-68.59 38.75-91.74 57.85a8 8 0 0 1-10.06.06q-4.72-3.75-9.69-7.39c-39.64-28.95-86.21-32.76-88.17-32.9a16 16 0 0 0-16.83 19.4c.42 1.93 9.19 40.69 31.7 71.61a8.09 8.09 0 0 1 0 9.55C9.57 291.52.8 330.29.38 332.22a16 16 0 0 0 16.83 19.4c2-.14 48.53-4 88.12-32.88q4.85-3.56 9.47-7.22a8 8 0 0 1 10.06.07c23.25 19.19 55.05 44.28 92 58a4 4 0 0 1 2.42 4.83l-10.66 37.18a16 16 0 0 0 18 20.18c17.16-2.9 51.88-12.86 96.05-46.83a8.15 8.15 0 0 1 4.36-1.65a287.36 287.36 0 0 0 39.22-5.3a4 4 0 0 0 2.69-5.83l-4.51-8.29A214.81 214.81 0 0 1 335.45 256z\" fill=\"currentColor\"></path></svg>', 'life', 0);

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标签名',
  `logo` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'LOGO',
  `color` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '颜色',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tags
-- ----------------------------
INSERT INTO `tags` VALUES (1, '我', '', '');
INSERT INTO `tags` VALUES (2, '测试', '', '');
INSERT INTO `tags` VALUES (3, '啊', '', '');
INSERT INTO `tags` VALUES (4, '复工后', '', '');
INSERT INTO `tags` VALUES (5, '电饭锅', '', '');
INSERT INTO `tags` VALUES (6, '日月潭', '', '');
INSERT INTO `tags` VALUES (7, '请问', '', '');
INSERT INTO `tags` VALUES (8, '朋友', '', '');
INSERT INTO `tags` VALUES (9, '气味儿', '', '');
INSERT INTO `tags` VALUES (10, '啊啊啊', '', '');
INSERT INTO `tags` VALUES (11, '涩涩涩', '', '');
INSERT INTO `tags` VALUES (12, 'cosplay', '', '');
INSERT INTO `tags` VALUES (13, '美女', '', '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `password` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `relation` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '关系',
  `role` bigint NULL DEFAULT NULL COMMENT '权限',
  `avatar` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '头像',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
