setup: 
	docker-compose up -d

setup-down:
	docker-compose down

build:
	go build -o mongopop main.go

populate: build
	./mongopop oplogpop $(ARGS)
