# 用户表
CREATE TABLE `users` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `user_id` bigint(20) NOT NULL,
                        `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `email` varchar(64) COLLATE utf8mb4_general_ci,
                        `gender` tinyint(4) NOT NULL DEFAULT '0',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_username` (`username`) USING BTREE,
                        UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `communitys`  (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `community_id` int(10) UNSIGNED NOT NULL,
                              `community_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                              `introduction` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                              `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`) USING BTREE,
                              UNIQUE INDEX `idx_community_id`(`community_id`) USING BTREE,
                              UNIQUE INDEX `idx_community_name`(`community_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

INSERT INTO communities VALUES (1, 1, 'Go', 'Golang', '2022-02-12 17:25:26', '2022-02-12 17:25:28');
INSERT INTO communities VALUES (2, 2, 'leetcode', '刷题刷题刷题', '2022-02-12 17:25:38', '2022-02-12 17:25:40');
INSERT INTO communities VALUES (3, 3, 'Java', 'springboot', '2022-02-12 17:25:46', '2022-02-12 17:26:20');
INSERT INTO communities VALUES (4, 4, 'LOL', '欢迎来到英雄联盟!', '2022-02-12 17:25:53', '2022-02-12 17:25:55');


CREATE TABLE `posts`  (
                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                         `post_id` bigint(20) NOT NULL COMMENT '帖子id',
                         `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
                         `content` varchar(8192) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
                         `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
                         `community_id` bigint(20) NOT NULL COMMENT '所属社区',
                         `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '帖子状态',
                         `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `idx_post_id`(`post_id`) USING BTREE,
                         INDEX `idx_author_id`(`author_id`) USING BTREE,
                         INDEX `idx_community_id`(`community_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;