## What is it?

Backend to my [Technik Informatyk](https://play.google.com/store/apps/details?id=jebok.itexam) application written in GO language. It is supposed to serve a snapshot of questions and exams to the user. Another functionality is to support multiplayer quizzes using websocket.

## What structure should the database have

You can create a database structure using these commands:

`CREATE TABLE `exam`( `uuid`char(36) COLLATE utf8_bin NOT NULL, `name`varchar(32) COLLATE utf8_bin NOT NULL, `description`varchar(255) COLLATE utf8_bin NOT NULL, `icon` varchar(32) COLLATE utf8_bin NOT NULL, PRIMARY KEY (`uuid`), KEY `uuid` (`uuid`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin`

`CREATE TABLE `question`( `uuid`char(36) COLLATE utf8_bin NOT NULL, `content`varchar(512) COLLATE utf8_bin NOT NULL, `answerA`varchar(255) COLLATE utf8_bin NOT NULL, `answerB`varchar(255) COLLATE utf8_bin NOT NULL, `answerC`varchar(255) COLLATE utf8_bin NOT NULL, `answerD`varchar(255) COLLATE utf8_bin NOT NULL, `correct`tinyint(4) NOT NULL, `image` mediumblob DEFAULT NULL, PRIMARY KEY (`uuid`), KEY `uuid` (`uuid`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin`
