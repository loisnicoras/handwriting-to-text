# Use official MySQL image from Docker Hub
FROM mysql:latest

# Environment variables
ENV MYSQL_ROOT_PASSWORD=emanuel
ENV MYSQL_DATABASE=my_database
ENV MYSQL_USER=lois
ENV MYSQL_PASSWORD=emanuel

# Expose MySQL default port
EXPOSE 3306