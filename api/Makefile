ENV := $(PWD)/.env
include $(ENV)

mysql:
	@sudo docker run -d --name mysql-container \
	-e MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) \
	-e MYSQL_DATABASE=$(MYSQL_DATABASE) \
	-e MYSQL_USER=$(MYSQL_USER) \
	-e MYSQL_PASSWORD=$(MYSQL_PASSWORD) \
	-p 3306:$(MYSQL_PORT) \
	-v $(PWD)/db/dump.sql:/docker-entrypoint-initdb.d/dump.sql \
	mysql:latest
	@sudo docker start mysql-container

dump:
	@docker cp ./db/dump.sql mysql-container:/dump.sql
	@docker exec -i mysql-container mysql -u $(MYSQL_USER) --password=$(MYSQL_PASSWORD) $(MYSQL_DATABASE) < ./db/dump.sql

dropdb:
	@docker exec -i mysql-container mysql -u $(MYSQL_USER) -p -e "DROP DATABASE $(MYSQL_DATABASE);"

removemysql:
	@sudo docker stop mysql-container
	@sudo docker rm mysql-container
	
run:
	@go run main.go --addr=$(SERVER_ADDR) \
	--apiKey=$(API_KEY) \
	--projectId=$(PROJECTID) \
	--region=$(REGION) 
	--dbUser=$(MYSQL_USER) \
	--dbPass=$(MYSQL_PASSWORD) \
	--dbHost=$(MYSQL_HOST) \
	--dbPort=$(MYSQL_PORT) \
	--dbName=$(MYSQL_USER) \

install:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b . v1.56.2

lint:
	@./golangci-lint run -v ./handlers/*
	@./golangci-lint run -v main.go

clean:
	@rm ./golangci-lint 

.PHONY: image mysql dump dropdb run install lint clean