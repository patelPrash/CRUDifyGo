# CRUDifyGo

**Table of Contents**

-Description

-Functionalities

-Features

-Basic Installation You need to have in your system

-How to start the Project

-Diagrams

**Description**

This project is a Go-based application built using the GoFr framework, designed to facilitate CRUD (Create, Read, Update, Delete) operations on a specific dataset. The primary purpose is to provide a robust API for managing data with straightforward endpoints for performing CRUD actions.

**Functionalities**

_Create_: Enables the addition of new records or entries to the dataset.

_Read_: Allows retrieval of data from the dataset based on specified parameters.

_Update_: Facilitates modification of existing records within the dataset.

_Delete_: Enables the removal of specific records from the dataset.


**Features**

1.Add a product

2.Get a product

3.Fetch a product by id

4.Update Product information

5.Update Product information by id

6.Delete a product

## POSTMAN SCREENSHOTS - 

![image](https://github.com/patelPrash/CRUDifyGo/assets/105787742/e9aea184-cc6f-4e7f-ad05-c6560f07bd9f)

![image](https://github.com/patelPrash/CRUDifyGo/assets/105787742/a6463e1d-a28c-46bb-b470-e799ee88a6d2)

![image](https://github.com/patelPrash/CRUDifyGo/assets/105787742/30a757ed-e2c6-428b-bb78-7d7759f07f8b)



## Basic Installation You need to have in your system - 
- GO language (https://go.dev/doc/install)
- docker (https://www.docker.com/products/docker-desktop/)

## How to start the Project - 
- start docker desktop or windows PowerShell
- then run this command - 
  > docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30
  > docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE customers (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL);"

- run this command `go run main.go` on terminal


## Diagrams-
![UML diagram](https://github.com/patelPrash/CRUDifyGo/assets/105787742/b73c036a-06a6-424e-bb4e-fa371d0d7669)
![SequenceDiagram](https://github.com/patelPrash/CRUDifyGo/assets/105787742/21cf5889-3c4b-4bad-9054-a16ed3e1abcc)


## Authors-
Prashant Patel
(patel.pr0036@gmail.com)
