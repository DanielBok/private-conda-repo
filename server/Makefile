PROJECT_NAME = "PCR"

conda-image:
	docker image build -t danielbok/conda-repo-mgr -f conda.Dockerfile .
	docker image prune -f

start:
	docker-compose -p $(PROJECT_NAME) up -d

stop:
	docker-compose -p $(PROJECT_NAME) down

create-migration:
	migrate.exe create -dir store/migrations -ext sql -seq -digits 2 <message>