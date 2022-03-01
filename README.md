# api-devbook

This repository contains files that build an api to connect for the database and disposes resources for the web app of Devbook social media.
It was built to learn deep concepts of programming language Go and how a social media working at the back end.
Remembering that it is a simple project, social medias use too much resources to working good.

Before you run the application, you should follow the steps below:

First step: 
  - Execute the following commands:
    go mod init api
    go mod tidy
These commands will initialize the project with the correct configuration and download its dependencies.

Second step:
  - It's necessary that you create a file .env in the root of your project with the following settings:

      DB_USUARIO=database_user

      DB_SENHA=database_password

      DB_NOME=database_name

      API_PORT=execution_port_of_your_application

      SECRET_KEY=the_secret_key_for_generate_web_token


#explain the api

Directory structure:

src: it contains directories based on Clean Architecture, since controllers until security.

sql: it contains SQL files for building the application database. You can use docker with a MySQL ou MariaDB image or the own Database server, the choose is yours.
