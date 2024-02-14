# Go Savings Account Banking Project

## Problem Statement

<p>
<b>Designing a Saving Account Banking System with REST API.</b>

<b>Requirements -</b>
  The system will have three roles: Customer, Admin, and Super-admin.

<b>Super-admin Privileges:</b><br/>
  Manage, view, and update Admin roles. <br/>
  Control all customer-related operations. <br/><br/>
<b>Admin Privileges:</b> <br/>
  Limited roles with specific branches to manage customers.<br/>
  Create, view, update, and delete customers of particular branches.<br/><br/>
<b>Customer Privileges:</b><br/>
  View their account statement.<br/>
  Create an account (Sign-Up).<br/>
  Individual login with valid details.<br/>
  Deposit money.<br/>
  Withdraw money.<br/>
  View balance.<br/><br/>
<b>General Features:</b><br/>
Both admin and customer can create their own account specifying their role.<br/>
Only one super-admin who monitors all admin and customer activities.<br/>

</p>


## Setup

This Project uses SQLite DB to handle database queries.
There are few records already seeded into database and whatever updations you make on database, it will persist even after you close the application. You can run the CleanUp command to start fresh.


Firstly, run the following command to download all dependencies
```bash
go mod download
```


1. Run following command to start e-commerce Application
```bash
make run
```

2. Run following command to run unit test cases
```bash
make test
```

3. Run following command to check test coverage
```bash
make test-cover

#you can also check code test coverage on top. Click on codeccov badge to check more about test coverage
```

4. Run following command to erase database to start fresh
```bash
make cleanDB
```


## APIs


1. <b>Signup API</b> : `POST http://localhost:1925/signup`
2. <b>Login API</b> : `POST http://localhost:1925/login`
3. <b>Update User API</b> : `PUT http://localhost:1925/update_user`
4. <b>Create Bank Account API</b> : `POST http://localhost:1925/account/create`
5. <b>Deposit Money API</b> : `PUT http://localhost:1925/account/deposit`
6. <b>Withdraw Money API</b> : `PUT http://localhost:1925/account/withdrawal`
7. <b>Delete Bank Account API</b> : `DEL localhost:1925/account/delete?acc_no=1`
8. <b>ADMIN => Get List of Users API</b> : `GET http://localhost:1925/admin/user_list`
9. <b>ADMIN => Get Users Info API</b> : `PUT http://localhost:1925/admin/update_user`

## Postman Collection


[here](postman_collection.json)


## Project Structure

```
jspnlp@unispab:~/go/Banking System$ tree
.
├── app
│   ├── account
│   │   ├── handler.go
│   │   ├── handler_test.go
│   │   └── service.go
│   ├── admin
│   │   ├── handler.go
│   │   └── service.go
│   ├── dependencies.go
│   ├── dto
│   │   ├── account.go
│   │   ├── admin.go
│   │   ├── errors.go
│   │   └── user.go
│   ├── enduser
│   │   ├── handler.go
│   │   ├── mocks
│   │   │   └── Service.go
│   │   ├── service.go
│   │   └── service_test.go
│   └── router.go
├── cmd
│   └── main.go
├── docs
│   └── usecase_ Saving Account System.pdf
├── go.mod
├── go.sum
├── Makefile
├── postman_collection.json
├── repository
│   ├── account.go
│   ├── admin.go
│   ├── bank.db
│   ├── init.go
│   ├── mocks
│   │   └── UserStorer.go
│   ├── repo.go
│   └── user.go
└── Sample Input Data for APIs.txt

10 directories, 29 files
```
