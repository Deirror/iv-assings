# Parking Lot

Goals: Design a parking lot using object-oriented principles
-

### Here are a few methods that you should be able to run:

- Tell us how many spots are remaining
- Tell us how many total spots are in the parking lot
- Tell us when the parking lot is full
- Tell us when the parking lot is empty
- Tell us when certain spots are full e.g. when all motorcycle spots are taken
- Tell us how many spots vans are taking up

### Assumptions:

- The parking lot can hold motorcycles, cars and vans
- The parking lot has motorcycle spots, car spots and large spots
- A motorcycle can park in any spot
- A car can park in a single compact spot, or a regular spot
- A van can park, but it will take up 3 regular spots

Clarifications from me, ***Deirror***
--
- Motorcycle: Can park everywhere (in any spot)
- Car: Can only park in car spots (either regular or compact)
- Van: Needs 3 regular car spots
- We have to write just an exemplary solution which also includes a test file so that we can show that it actually solves the problem
- I will create a RESTful API where the server listens for specific routes (URLs) related to the vehicles
