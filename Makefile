MONGO_CONTAINER = mongo-oplog-populator_mongoCluster

setup: 
	docker-compose up -d
	sleep 3
	docker run -t -i --rm --network $(MONGO_CONTAINER) mongo mongosh --host mongo1 --port 30001 --eval 'config = {"_id": "my-replica-set", "members": [{"_id": 0,"host": "mongo1:30001" },{"_id": 1,"host": "mongo2:30002" },{"_id": 2,"host": "mongo3:30003" }]}; rs.initiate(config); rs.status(); sleep(15000); rs.status();'

eval_primary:
	$(eval PRIMARY=$(shell sh -c "docker run -t -i --rm --network $(MONGO_CONTAINER) mongo mongosh --quiet --host mongo1 --port 30001 --eval 'rs.isMaster().primary'" | awk '{print $$3}' | tr -d '[1G'))
	
	$(eval PRIMARY_HOST=$(shell sh -c "echo $(PRIMARY)" | awk -F':' '{print $$1}'))
	$(eval PRIMARY_PORT=$(shell sh -c "echo $(PRIMARY)" | awk -F':' '{print $$2}'))
	@echo Primary host: $(PRIMARY_HOST) and port $(PRIMARY_PORT)

connect: eval_primary
	@echo Connecting to primary host: $(PRIMARY_HOST)
	@docker run -t -i --rm --network $(MONGO_CONTAINER) mongo mongosh --host $(PRIMARY_HOST) --port $(PRIMARY_PORT)

setup-down:
	docker-compose down
	rm -rf ./data

build:
	go build -o mongopop ./cmd/mongo-populator
