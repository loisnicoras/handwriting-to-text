mysql:
	sudo docker run -d --name mysql-container \
	-e MYSQL_ROOT_PASSWORD=emanuel \
	-e MYSQL_DATABASE=my_database \
	-e MYSQL_USER=lois \
	-e MYSQL_PASSWORD=emanuel \
	-p 3306:3306 \
	mysql:latest
	
dump:
	mysql -u lois -p < db/dump.sql

dropdb:
	mysqladmin -u lois -p drop my_database

run:
	go run main.go --apiKey=AIzaSyBWomSOPNh-6Xxzg3aUTX7nyr0e5v91TsQ --addr=8081 \
	--dbUser=lois \
	--dbPass=emanuel \
	--dbHost=localhost \
	--dbName=my_database

.PHONY: image mysql dump dropdb run