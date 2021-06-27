instruction:
	@echo "run <make dev> to start server and UI on development mode"
	@echo "run <make dev-server> <make dev-client> on two separate terminals to run server and UI in production mode"
	@echo "run <make prod-client> <make prod-server> to run server and UI on two terminals to run server and UI in production mode"

mongo: 
	docker-compose down
	docker-compose up -d

# no need to create database as go-mongo-driver handled this
# database:
# 	docker exec -i mongodb -u root -p root -c 'use test'
# 	docker exec -i mongodb mongo -u root -p root -c 'db.todolist.insert()'
.PHONY: clean dev-server dev-client prod-server prod-client dev

dev: clean dev-server dev-client

dev-server: clean
	cd server && go run main.go &
	cd client && yarn start

dev-client:
	cd client && yarn start

prod-server: clean
	go build && ./go-server

prod-client:
	cd client && yarn build

docker-server: clean
	cd server && docker-compose down
	cd server && docker-compose up

clean:
	./clean.sh
	

   