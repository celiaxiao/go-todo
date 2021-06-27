mongo: 
	docker-compose down
	docker-compose up -d

# no need to create database as go-mongo-driver handled this
# database:
# 	docker exec -i mongodb -u root -p root -c 'use test'
# 	docker exec -i mongodb mongo -u root -p root -c 'db.todolist.insert()'

dev-server:
	cd server && go run main.go
# ps -ef | grep "go run" | grep -v grep | awk '{print $2}' 

dev-client:
	cd client && yarn start

build-client:
	cd client && yarn build

clean:
	

   