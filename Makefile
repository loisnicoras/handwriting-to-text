image:
	docker build -t my-mysql .
	
db:
	docker run --name my-mysql-container -p 3306:3306 -e MYSQL_USER=lois -e MYSQL_PASSWORD=emanuel -e MYSQL_DATABASE=my_database -d my-mysql

dump:
	mysql -u lois -p < db/dump.sql

dropdb:
	mysqladmin -u lois -p drop my_database

.PHONY: image db dump dropdb