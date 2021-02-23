# Go Assessment Project

### 1.1 Description
A Go project with some API which retrieves/modifies some files containing the balance of accounts. 
The project should include a config file, which stores the directory of these files. 

#### 2.1 Requirements
- An API to create a number of accounts (balance files) in the directory
    - The balance of each account should be generated randomly, ranging from 1,000 to 100,000
    - These files (accounts' balances) could be created only once
- An API to get the balance of one/all account(s)
- An API to add number x to the balance of one/all account(s)
- Specify the pieces of code which have applied SOLID principle

##### 3.1 Guide
- Comment
  - Search for #SOLID to see the examples of applied SOLID principles
  - Search for #TODO to see the possible improvements which I could not implement in the specified time period
  - Search for #Desc to see my reasons for making some specific decisions

##### 3.2 Project structure
from the highest level to lowest:
1. API layer (./api)
2. The business layer (./account)
3. The internal layer (./internal)

#### 4.1 How to use
#### 4.2 Sample API call
- **Create accounts** URL: `localhost:8080/accounts?number=1000` Verb: `POST`
- **Get a balance** URL: `localhost:8080/accounts?id=12` Verb: `GET`
- **Get sum of all balances** URL: `localhost:8080/accounts?result=aggregate` Verb: `GET`
- **increase a balance** URL: `localhost:8080/accounts?id=12&increase=1000` Verb: `PUT`
- **increase all balances** URL: `localhost:8080/accounts?increase=1000` Verb: `PUT`
  
#### 5.1 TODO
- Add more tests
- Implement delete API
- Define more explicit error types 
  - assert them in tests 
  - use them to return more explicit HTTP error codes
- return error in get/update methods when there is no balance file
- Rollback changes when batch update is failed
- Wrap zap(logger) by an interface
- Put data validation on config