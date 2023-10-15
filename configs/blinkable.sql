/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80100 (8.1.0)
 Source Host           : localhost:3306
 Source Schema         : blinkable

 Target Server Type    : MySQL
 Target Server Version : 80100 (8.1.0)
 File Encoding         : 65001

 Date: 15/10/2023 15:42:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `article_id` int NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `author_id` int NOT NULL COMMENT '作者id',
  `title` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章标题',
  `label` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章标签',
  `category` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章分类',
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章内容',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`article_id`),
  KEY `fk_author_id` (`author_id`),
  CONSTRAINT `fk_author_id` FOREIGN KEY (`author_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章表';

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `comment_id` int NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `user_id` int NOT NULL COMMENT '用户id',
  `article_id` int NOT NULL COMMENT '文章id',
  `content` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论内容',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`comment_id`),
  KEY `fk_user_id` (`user_id`),
  KEY `fk_article_id` (`article_id`),
  CONSTRAINT `fk_article_id` FOREIGN KEY (`article_id`) REFERENCES `article` (`article_id`),
  CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表';

-- ----------------------------
-- Table structure for comments_article
-- ----------------------------
DROP TABLE IF EXISTS `comments_article`;
CREATE TABLE `comments_article` (
  `comments_id` int NOT NULL COMMENT '评论id',
  `articles_id` int NOT NULL COMMENT '文章id',
  PRIMARY KEY (`comments_id`,`articles_id`),
  KEY `fk_comments_article__article_id` (`articles_id`),
  CONSTRAINT `fk_comments_article__article_id` FOREIGN KEY (`articles_id`) REFERENCES `article` (`article_id`),
  CONSTRAINT `fk_comments_article_comment_id` FOREIGN KEY (`comments_id`) REFERENCES `comments` (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论文章关联表';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_name` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `avatar` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '头像地址',
  `articles_num` int NOT NULL COMMENT '文章数量',
  `level` int NOT NULL DEFAULT (1) COMMENT '等级',
  `signature` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT (_utf8mb4'什么都没写') COMMENT '签名',
  `experience` int NOT NULL DEFAULT (0) COMMENT '经验值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

SET FOREIGN_KEY_CHECKS = 1;
