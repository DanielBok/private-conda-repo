PROJECT_NAME = "PCR"

start: stop
	docker-compose -p $(PROJECT_NAME) up -d

stop:
	docker-compose -p $(PROJECT_NAME) down

start-ssl: stop-ssl
	docker-compose -p $(PROJECT_NAME) -f docker-compose.ssl.yml up -d

stop-ssl:
	docker-compose -p $(PROJECT_NAME) -f docker-compose.ssl.yml down
