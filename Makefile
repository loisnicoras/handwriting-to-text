ENV := $(PWD)/.env
include $(ENV)

mysql:
	sudo docker run -d --name mysql-container \
	-e MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) \
	-e MYSQL_DATABASE=$(MYSQL_DATABASE) \
	-e MYSQL_USER=$(MYSQL_USER) \
	-e MYSQL_PASSWORD=$(MYSQL_PASSWORD) \
	-p 3306:3306 \
	mysql:latest
	
dump:
	mysql -u lois -p < db/dump.sql

dropdb:
	mysqladmin -u lois -p drop my_database

run:
	go run main.go --apiKey=AIzaSyBWomSOPNh-6Xxzg3aUTX7nyr0e5v91TsQ --addr=8081 \
	--dbUser=$(MYSQL_USER) \
	--dbPass=$(MYSQL_PASSWORD) \
	--dbHost=$(MYSQL_HOST) \
	--dbName=$(MYSQL_DATABASE)

.PHONY: image mysql dump dropdb run