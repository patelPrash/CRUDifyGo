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

## Basic Installation You need to have in your system - 
- GO language (https://go.dev/doc/install)
- docker (https://www.docker.com/products/docker-desktop/)

## How to start the Project - 
- start docker desktop or windows PowerShell
- then run this command - 
  > docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30
- run this command `go run main.go` on terminal

## Diagrams-

