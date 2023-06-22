# Mongo-oplog-populator

MongoDB oplog populator (alias `mongopop`) allows you to simulate insert, update and delete traffic in MongoDB.
This repo is a companion repo for https://github.com/one2nc/mongo-oplog-to-sql
If you want to test your `mongo-oplog-to-sql` binary on large data, you should use `mongopop`.

### How it works?

When you run `mongopop`, it creates two collection - employees and students in MongoDB and populates some data in those tables.
The employees collection is created in employee database while the students collection is created in student database.

![mongo-oplog-flow](assets/MongoDb-oplog-populator.png)

`mongopop` supports two modes: bulk form or stream mode. 

1. Bulk mode: Perform fixed number of operations provided by the user and terminate the program.
2. Stream mode: Perform fixed number of operations provided by user per second for indefinite time. The program stops only when the user gives a stop signal (ctrl+c)

The number of operations you specify via CLI are divided into insert, update and delete operations in following proportion.

- 85% Insert
- 10% updates
- 5% deletes

`mongopop` performs insert, update and delete operations randomly. If you were to specify 100 operations, 85 of them will be inserts,
10 will be updates and 5 will be delete opertions. The operations are divided equally into the two collections - employees.employee and students.student.
i.e. In case of `$ ./mongopop 100`, 50 operations will be done on employee collection and remaining 50 will be done on student collection.

### Setup 
1. Setup mongo: Spins up 3 container of mongo
  
   `make setup`

2. Build the application: Builds a binary of mongo-oplog-populator
   
   `make build`

2.  Run `./mongopop 10` for bulk operations  and 
        `./mongopop -s 10` for stream operations
     -  *./mongopop* is the binary file
     -  *-s* is the flag for stream operation
     -  *10* is the total number of operations to be performed in case of bulk insert and per second in case of stream insert
     
3. To connect to mongo cluster `make connect`

4. To check oplogs in mongo execute the following commands:
    - Change the database to local `use local` since all the oplogs are located in a file oplog.rs which is in local database
    - To see the oplogs generated:
      -  `db.oplog.rs.find({})` to see all oplogs
      -  `db.oplog.rs.find({op:"i"})` to see the insertions
      -  `db.oplog.rs.find({op:"u"})` to see the updates
      -  `db.oplog.rs.find({op:"d"})` to see the deletions

6. Tear down mongo
    `make setup-down`

NOTE: If you want to insert more records in mongo, run multiple instances of `mongopop`.