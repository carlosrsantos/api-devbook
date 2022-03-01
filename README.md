# api-devbook

First step: 
  - Execute the following commands:
    go mod init api
    go mod tidy


The commands will initialize the project with the correct configuration and download its dependencies.

This repository contains files that build an api to connect for the database and disposes resources for the web app of Devbook social media.
It was built to learn deep concepts of programming language Go and how a social media working at the back end.
Remembering that it is a simple project, social medias use too much resources to working good.


Directory structure:

src: it contains directories based on Clean Architecture, since controllers until security.

sql: it contains SQL files for building the application database. You can use docker with a MySQL ou MariaDB image or the own Database server, the choose is yours.
