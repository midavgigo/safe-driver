build:
	sudo docker build -t opego .
run:
	sudo docker run -p 8080:8080 --name opego -it --rm opego
