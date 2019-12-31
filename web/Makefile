build: build-normal build-ssl

build-normal:
	docker image build -t danielbok/pcr-web .

build-ssl:
	docker image build -t danielbok/pcr-web-ssl -f ssl.Dockerfile .
