# Basic database design

# 瀑布式开发 感觉前期写的有点水 见谅了

# 用户
CREATE TABLE `users`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `login_name` VARCHAR(64) NOT NULL COMMENT "用户登录名",
    `pwd` CHAR(32) NOT NULL COMMENT "密码",
    `salt` CHAR(32) NOT NULL COMMENT "盐",
    PRIMARY KEY (`id`),
    UNIQUE KEY `login_name`(`login_name`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;

# 视频
CREATE TABLE `video_info`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `author_id` INT UNSIGNED NOT NULL COMMENT "关联用户id",
    `name` VARCHAR(266) NOT NULL DEFAULT "" COMMENT "视频名称",
    `create_time` INT UNSIGNED NOT NULL COMMENT "创建时间",
    PRIMARY KEY (`id`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;

# 评论 (后面会取消,这个系统不会包含cms)
CREATE TABLE `comments` (
    `id` INT UNSIGNED AUTO_INCREMENT,
    `video_id` INT UNSIGNED NOT NULL COMMENT "关联视频表",
    `author_id` INT UNSIGNED NOT NULL COMMENT "关联用户id",
    `content` VARCHAR (666) NOT NULL DEFAULT "" COMMENT "评论",
    `create_time` INT UNSIGNED NOT NULL COMMENT "评论时间",
    PRIMARY KEY (`id`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;

# sessions
CREATE TABLE `sessions` (
    `session_id` VARCHAR(255) NOT NULL,
    `ttl` INT UNSIGNED NOT NULL,
    `login_name` VARCHAR(266) NOT NULL ,
    PRIMARY KEY(`session_id`)
)ENGINE=INNODB DEFAULT CHARSET=UTF8;
