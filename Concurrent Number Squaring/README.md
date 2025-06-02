# Concurrent Number Squaring

Description:
-

Given a list of integers, spawn goroutines to square each number and collect the results.

- Use a channel to send tasks
-  Use another channel to collect results
- Print the squared numbers in input order

Constraints
-

- Use up to 3 concurrent goroutines
- Preserve the original input order