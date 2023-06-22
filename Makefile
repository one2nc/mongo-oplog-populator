setup: 
	docker-compose up -d
	sleep 3
	docker run -t -i --rm --network mongo-oplog-populator_mongoCluster mongo mongosh --host mongo1 --port 30001 --eval 'config = {"_id": "my-replica-set", "members": [{"_id": 0,"host": "mongo1:30001" },{"_id": 1,"host": "mongo2:30002" },{"_id": 2,"host": "mongo3:30003" }]}; rs.initiate(config); rs.status(); sleep(15000); rs.status();'

connect:
	docker run -t -i --rm --network mongo-oplog-populator_mongoCluster mongo mongosh --host mongo1 --port 30001

setup-down:
	docker-compose down
	rm -rf ./data

build:
	go build -o mongopop ./cmd/mongo-populator
