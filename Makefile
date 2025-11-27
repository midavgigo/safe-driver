build:
	sudo docker-compose build --no-cache
run:
	sudo docker-compose up 
stop:
	sudo docker-compose down
clear: 
	sudo docker-compose down -v
