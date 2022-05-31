![Technik Informatyk API](https://user-images.githubusercontent.com/59766830/169153125-16d783aa-9fef-4869-8421-aa007bffdc0f.jpg)

[![Go Report Card](https://goreportcard.com/badge/github.com/pawl0wski/technikinformatyk-backend)](https://goreportcard.com/report/github.com/pawl0wski/technikinformatyk-backend)

## Why archived

I feel more at ease in JavaScript than in Go, so I decided to rewrite the api to Javascript.

## What is it?

Backend to my [Technik Informatyk](https://play.google.com/store/apps/details?id=jebok.itexam) application written in GO language. It is supposed to serve a snapshot of questions and exams to the user. Another functionality is to support multiplayer quizzes using websocket.

## Database structure

You can find the commands with the database structure in the dbStruct.sql file.

## Env variables
| Name                  | What it does                                                                                                                                                                                                                                   |
|-----------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| SERVER_PORT           | Define on which port the application should listen. (Default. 3000)                                                                                                                                                                            |
| ENABLE_CDN            | Specify whether the API should generate a content delivery network folder. You should set the location in your HTTP server e.g. nginx to this [folder](https://docs.nginx.com/nginx/admin-guide/web-server/web-server/#configuring-locations). |
| CDN_PATH              | Folder where content will be generated for cdn.                                                                                                                                                                                                |
| MYSQL_CONNECTION_PATH | A path for the api to connect to the database. user:password@/dbname                                                                                                                                                                           |
