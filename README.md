# Mongo-oplog-populator

![mongo-oplog-flow](assets/MongoDb-oplog-populator.png)


 System supports populating the data in Mongo Database. The operations can be performed in a bulk form or in stream form.

> *Out of total Operations:*
  > - 85% Insert
  > - 10% updates
  > - 5% deletes
  > - 10% of total number of operations Alters (for each database)
## Features
1. Bulk Insert: Perform fixed number of operations provided by the user and terminates the program

2. Stream Insert: Perform fixed number of operations provided by user per second for indefinite time. The program stops only when the user gives a stop signal (ctrl+c) 

### Setup 
1. Setup mongo: Spins up 3 container of mongo
  
   `make setup`

2. Build the application: Builds a binary of mongo-oplog-populator
   
   `make build`

2.  Run `./mongopop 10` for bulk operations  and 
        `./mongopop -s 10` for stream operations
     *  *./mongopop* is the binary file
     *  *-s* is the boolean flag for stream operation
     *  *10* is the total number of operations to be performed in case of bulk insert and per second in case of stream insert
    

3. To connect to mongo cluster `make connect`
4. To enter into student database `use student` and for employee datatbase `use employee`
   - to count number of records in students table in student database `db.students.count()` 
   - to count number of records in employees table in employee database `db.employees.count()`

5. To check oplogs in mongo execute the following commands:
    - Change the database to local `use local` since all the oplogs are located in a file oplog.rs which is in local database
    - To see the oplogs generated:
      *  `db.oplog.rs.find({})` to see all oplogs
      *  `db.oplog.rs.find({op:"i"})` to see the insertions
      *  `db.oplog.rs.find({op:"u"})` to see the updates
      *  `db.oplog.rs.find({op:"d"})` to see the deletions

6. Tear down mongo
    `make setup-down`


NOTE: If you want to insert more records in mongo, run multiple instances of mongo-populator.