# Plain Caching Implementation

> Given on an interview for the position of *Junior Golang Software Developer*

Description
-

Create a library for caching user information to increase database throughput and availability.<br />
Every request for use data should return the user information while simultaneously taking care of the database throughput.<br />
It means the library should take care of unnecessary requests for user data if some are pending.<br />

**Why**:

*Let's say we have a bunch of user information (idnetified by id) inside our main database.<br />
Simultaneously there's thousands of requets per second for user information.<br />
We can consider that database is a bottleneck right. To avoid that we need a cache mechanism.<br />*

**Notes**:

**You can use any language, library, database and framework - anything you want and think is the best for this case.**

*Example*:

*If within 1000 requests there are 100 unique user ids then there should be only a maximum 100 requests into the database but all 1000 requests should get a response with a user data.*
