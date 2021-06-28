# GO-React web app

This is the web app using GO as server and React as frondend UI.

See Shubham Chadokar's [blog](https://levelup.gitconnected.com/build-a-todo-app-in-golang-mongodb-and-react-e1357b4690a6) for tutorial about this app.

This Git repo contains makefile for this project and its dockerfile for creating and running docker image.

run 'cd server && docker build -t go-todo .' to build image for go-server
run 'cd client && docker build -t react-client .' to build image for react-client

run 'make' to see instruction to use the make file
