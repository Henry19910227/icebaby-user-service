run:
	docker-compose up --build -d
mysql:
	docker-compose up --build -d mysql
api:
	docker-compose up --build -d api