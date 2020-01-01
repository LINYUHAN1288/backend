DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user`(
    `userid` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `username` VARCHAR(32) NOT NULL DEFAULT '',
    `userpass` CHAR(32) NOT NULL DEFAULT '',
    `useremail` VARCHAR(100) NOT NULL DEFAULT '',
    `createtime` INT UNSIGNED NOT NULL DEFAULT '0',
    UNIQUE user_username_userpass(`username`,`userpass`),
    UNIQUE user_useremail_userpass(`useremail`,`userpass`),
    PRIMARY KEY(`userid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `user`(username,userpass,useremail,createtime) VALUES('admin', md5('123'), '1057000273@qq.com', UNIX_TIMESTAMP());

DROP TABLE IF EXISTS `profile`;
CREATE TABLE IF NOT EXISTS `profile`(
    `truename` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '真实姓名',
    `age` TINYINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '年龄',
    `sex` ENUM('0','1','2') NOT NULL DEFAULT '0' COMMENT '性别',
    `birthday` date NOT NULL DEFAULT '2020-01-01' COMMENT '生日',
    `nickname` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '昵称',
    `userid` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户的ID',
    `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY(`userid`),
    UNIQUE profile_userid(`userid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
