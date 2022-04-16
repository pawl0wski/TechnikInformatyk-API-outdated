## What is it?

Backend to my [Technik Informatyk](https://play.google.com/store/apps/details?id=jebok.itexam) application written in GO language. It is supposed to serve a snapshot of questions and exams to the user. Another functionality is to support multiplayer quizzes using websocket.

## What structure should the database have

You can create a database structure using these commands:

`CREATE TABLE "exam" ( "name" TEXT, "uuid" TEXT, "description" TEXT, "icon" TEXT );`

`CREATE TABLE "question" ( "uuid" TEXT, "content" TEXT, "img" BLOB, "answers" TEXT, "examUuids" TEXT);`
