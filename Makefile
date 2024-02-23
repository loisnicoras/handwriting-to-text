image:
	docker build -t my-mysql .
	
db:
	docker run --name my-mysql-container -p 3306:3306 -e MYSQL_USER=lois -e MYSQL_PASSWORD=emanuel -e MYSQL_DATABASE=my_database -d my-mysql

exec:
	docker exec -it my-mysql-container mysql -u lois -p

dropdb:
	mysqladmin -u lois -p drop my_database

.PHONY: image db exec dropdb