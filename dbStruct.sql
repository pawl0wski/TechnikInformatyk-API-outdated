CREATE DATABASE itexam;
USE itexam;
CREATE TABLE `exam`( `uuid`char(36) COLLATE utf8_bin NOT NULL, `name`varchar(32) COLLATE utf8_bin NOT NULL, `description`varchar(255) COLLATE utf8_bin NOT NULL, `icon` varchar(32) COLLATE utf8_bin NOT NULL, PRIMARY KEY (`uuid`), KEY `uuid` (`uuid`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
CREATE TABLE `question`( `uuid`char(36) COLLATE utf8_bin NOT NULL, `content`varchar(512) COLLATE utf8_bin NOT NULL, `answerA`varchar(255) COLLATE utf8_bin NOT NULL, `answerB`varchar(255) COLLATE utf8_bin NOT NULL, `answerC`varchar(255) COLLATE utf8_bin NOT NULL, `answerD`varchar(255) COLLATE utf8_bin NOT NULL, `correct`tinyint(4) NOT NULL, `image` mediumblob DEFAULT NULL, PRIMARY KEY (`uuid`), KEY `uuid` (`uuid`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
CREATE TABLE `apiKey` (`id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,`key` char(32) COLLATE utf8_bin NOT NULL UNIQUE ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
