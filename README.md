# Go Assessment Project

### Description
A Go project with some API which retrieves/modifies some files containing the balance of accounts. 
The project should include a config file, which stores the directory of these files. 

#### Requirements
- An API to create a number of accounts (balance files) in the directory
    - The balance of each account should be generated randomly, ranging from 1,000 to 100,000
    - These files (accounts' balances) could be created only once
- An API to get the balance of one/all account(s)
- An API to add number x to the balance of one/all account(s)
- Specify the pieces of code which have applied SOLID principle

#### Guide
- Comment
  - Search for #SOLID to see the examples of applied SOLID principles
  - Search for #TODO to see the possible improvements which I could not implement in the specified time period
  - Search for #Desc to see my reasons for making some specific decisions
  

#### TODO
- Implement delete
- Define more explicit error types, and assert them in tests and also use them to return more explicit HTTP error codes
- Add error on get/update methods as there is no balance file
- Enhance project terminology
- Rollback changes when batch update is failed
- Wrap zap(logger) over an interface