ENV := $(PWD)/.env
include $(ENV)

mysql:
	@sudo docker run -d --name mysql-container \
	-e MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) \
	-e MYSQL_DATABASE=$(MYSQL_DATABASE) \
	-e MYSQL_USER=$(MYSQL_USER) \
	-e MYSQL_PASSWORD=$(MYSQL_PASSWORD) \
	-p 3306:$(MYSQL_PORT) \
	-v $(PWD)/db/schema.sql:/docker-entrypoint-initdb.d/dump.sql \
	mysql:latest

dump:
	@docker cp ./db/dump.sql mysql-container:/dump.sql
	@docker exec -i mysql-container mysql -u lois --password=emanuel my_database < ./db/dump.sql

dropdb:
	@docker exec -i mysql-container mysql -ulois -p -e "DROP DATABASE my_database;"

removemysql:
	@sudo docker stop mysql-container
	@sudo docker rm mysql-container
	
run:
	@go run main.go --apiKey=$(API_KEY) --addr=8081 \
	--dbUser=$(MYSQL_USER) \
	--dbPass=$(MYSQL_PASSWORD) \
	--dbHost=$(MYSQL_HOST) \
	--dbPort=$(MYSQL_PORT)
	--dbName=$(MYSQL_DATABASE)
	--projectId=$(PROJECTID)
	--region=$(REGION)

install:
	@sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b . v1.56.2

lint:
	@./golangci-lint run -v ./handlers/
	@./golangci-lint run -v .

clean:
	@sudo rm ./golangci-lint 

.PHONY: image mysql dump dropdb run install lint clean