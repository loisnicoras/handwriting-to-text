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

.PHONY: image mysql dump dropdb