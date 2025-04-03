# Plain Caching Implementation

> Given on an interview for the position of Junior Software Developer

Description
-

Create a library for caching user information to increase database throughput and availability.

Every request for use data should return the user information while simultaneously taking care of the database throughput.

It means the library should take care of unnecessary requests for user data if some are pending,

Why:

