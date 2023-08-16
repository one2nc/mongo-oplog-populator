MONGO_CONTAINER = mongo-oplog-populator_mongoCluster

setup: 
	docker-compose up -d
	sleep 3
	docker run -t -i --rm --network $(MONGO_CONTAINER) mongo mongosh --host mongo1 --port 30001 --eval 'config = {"_id": "my-replica-set", "members": [{"_id": 0,"host": "mongo1:30001" },{"_id": 1,"host": "mongo2:30002" },{"_id": 2,"host": "mongo3:30003" }]}; rs.initiate(config); rs.status(); sleep(15000); rs.status();'

get-primary:
	$(eval PRIMARY=$(shell sh -c "docker run -t -i --rm --network $(MONGO_CONTAINER) mongo mongosh --quiet --host mongo1 --port 30001 --eval 'rs.isMaster().primary'" | awk '{print $$3}' | sed 's/\x1b//g; s/\[1G//; s/\r//g'))

	$(eval PRIMARY_HOST=$(shell sh -c "echo $(PRIMARY)" | awk -F':' '{print $$1}' | tee host.txt))
	$(eval PRIMARY_PORT=$(shell sh -c "echo $(PRIMARY)" | awk -F':' '{print $$2}' | tee port.txt))
	@echo $(PRIMARY_HOST) 
	@echo $(PRIMARY_PORT)

connect: get-primary
	@echo Connecting to primary host $(PRIMARY_HOST)
	@docker run -t -i --rm --network $(MONGO_CONTAINER) mongo mongosh --host $(PRIMARY_HOST) --port $(PRIMARY_PORT)

setup-down:
	docker-compose down
	rm -rf ./data

build:
	go build -o mongopop ./cmd/mongo-populator
