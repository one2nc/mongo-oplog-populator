# Mongo-oplog-populator

![mongo-oplog-flow](assets/MongoDb-oplog-populator.png)


> System supports populating the data in Mongo oplog. The operations can be performed in a bulk form or in stream form.
>
>> *Out of total Operations:*
  >> - 85% Insert
  >> - 10% updates
  >> - 5% deletes
>
### Setup 
1. Setup mongo 
  `make setup`

2.  Run `./mongopop -b 10` for bulk operations 
        `./mongopop -s 10` for stream operations
     *  *./mongopop* is the binary file
     *  *--op* is the flag for bulk operation
     *  *--b* is the flag for stream operation
     *  *10* is the total number of operations to be performed

3. Run docker `make connect` 

4. Change the database to local `use local`

5. To see the oplog generated `db.oplog.rs.find({})`
    *  `db.oplog.rs.find({})` to see all oplogs
    *  `db.oplog.rs.find({op:"i"})` to see the insertions
    *  `db.oplog.rs.find({op:"u"})` to see the updates
    *  `db.oplog.rs.find({op:"d"})` to see the deletions

6. Set down mongo
    `make setup-down`
