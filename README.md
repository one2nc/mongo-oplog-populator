# Mongo-oplog-populator

![mongo-oplog-flow](assets/MongoDb-oplog-populator.png)


 System supports populating the data in Mongo Database. The operations can be performed in a bulk form or in stream form.

> *Out of total Operations:*
  > - 85% Insert
  > - 10% updates
  > - 5% deletes

## Features
1. Bulk Insert: Perform fixed number of operations provided by the user and terminates the program

2. Stream Insert: Perform fixed number of operations provided by user per second for indefinite time. The program stops only when the user gives a stop signal (ctrl+c) 

### Setup 
1. Setup mongo: Spins up 3 container of mongo
  
   `make setup`

2. Build the application: Builds a binary of mongo-oplog-populator
   
   `make build`

3. Create a `.env` file. You can refer [.env.example](.env.example) file

4.  Run `./mongopop 10` for bulk operations  and 
        `./mongopop -s 10` for stream operations
     *  *./mongopop* is the binary file
     *  *-s* is the flag for stream operation
     *  *10* is the total number of operations to be performed in case of bulk insert and per second in case of stream insert
    

5. To connect to mongo cluster `make connect`

6. To check oplogs in mongo execute the following commands:
    - Change the database to local `use local` since all the oplogs are located in a file oplog.rs which is in local database
    - To see the oplogs generated:
      *  `db.oplog.rs.find({})` to see all oplogs
      *  `db.oplog.rs.find({op:"i"})` to see the insertions
      *  `db.oplog.rs.find({op:"u"})` to see the updates
      *  `db.oplog.rs.find({op:"d"})` to see the deletions

7. Tear down mongo
    `make setup-down`


NOTE: If you want to insert more records in mongo, run multiple instances of mongo-populator.