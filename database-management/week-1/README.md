# Database Management Mentoring Week-1

## Basic Rules
- **This is not a demo!**
- ### **The challenges are for reference, can be used as is, customized or completely use your own challenges. Keep in mind the Learning Objectives!**
- Mentor will be the facilitator, guiding the mentee to solve the tasks.
  - Share screen method
    - One person will share their screen, can be the mentor and one of the mentee.
    - Mentee will take turn in giving the correct syntax to be typed by person sharing his/her screen.
  - Using tools such as replit https://replit.com/
    - Everyone can login using their own account.
    - Mentor can have a tab for each mentee to monitor their work.
    - Details over at Discord: https://discord.com/channels/996138478817513632/996141512704401579/1019923694254043136
  
## Learning Objectives

- Understand data modeling
- Understand SQL Constraints
- Able to create and run SQL commands for DDL (Data Definition Language) using Go.
- Able to create and run SQL commands for DML (Data Manipulation Language) using Go.
- Able to create and run SQL commands for DQL (Data Query Language) using Go.
- Able to create and run basic JOIN between 2 or 3 tables using Go.

## Pre-requisites

- Code editor (e.g. VS Code).
- Golang installed in local computer
- Postgres installed in local computer

## Basic

### HRApp - Employee Data

- Concepts:
  - Data modeling
  - Primary Key
  - CREATE TABLE ...
  - INSERT INTO ...
  - SELECT ... FROM ... WHERE ...
  - UPDATE ... SET .. WHERE ...
  - DELETE ... FROM ... WHERE ...
- Input:
  - Employee Data ( name (text), city (text) )
- Output:
  - Data inserted to database
- Directions:
  - Given the input, draw table structure ( preferrably using draw.io or ERD tools )
    - make note of the primary key ( remember to create id field as needed )
  - Open Postgres pgAdmin / DBeaver / other tool
  - Create Database "hrApp"
  - Create table using SQL Commands, note the structure of the table.
  - Drop table using SQL Commands, note that table is gone.
  - Create table using Golang, make note that the SQL command is the same.
  - Insert some data into the table using SQL Commands.
  - Insert some data using Golang.
  - Select some data from the table using SQL Commands.
  - Select some data from the table using Golang.
  - Do update and delete like insert, first using SQL Commands, then using Golang.

### HRApp - Department Data

- Concepts:
  - Data modeling
  - Primary Key
  - Foreign Key
  - CREATE TABLE ...
  - INSERT INTO ...
  - SELECT ... FROM ... WHERE ...
  - UPDATE ... SET .. WHERE ...
  - DELETE ... FROM ... WHERE ...
  - JOIN two tables
- Input:
  - Database "hrApp" with Employee table
  - Department Data ( name (text), head_department (link to employee_id) )
- Output:
  - Data inserted to database
- Directions:
  - Given the input, draw table structure ( preferrably using draw.io or ERD tools )
    - make note of the primary key
    - make note of foreign key used by head_department
  - Open Postgres pgAdmin / DBeaver / other tool
  - Use Database "hrApp"
  - Create table using Golang.
  - Insert some data using Golang.
  - Select some data from the table using Golang.
  - Optional
    - Do update and delete like insert using Golang.
  - Select the name of the department head using SQL JOIN

## Challenging

### Attendance Data

- Concepts:
  - All the above challenges
  - Database relationship / association
- Input:
  - As above
  - Add attendanceStatus with one field ( status ) which will be filled with present or absent.
- Output:
  - One association table called attendance
  - Data inserted into database
- Directions:
- Given the input, draw table structure ( preferrably using draw.io or ERD tools )
    - make note of the primary key
    - make note of foreign key
    - make note of association table that need to be created which contains Primary Keys of two tables ( employee and attendanceStatus) and field date
  - Use Database "hrApp"
  - Create table attendanceStatus.
  - Create association table called attendance.
  - Insert data for attendanceStatus.
  - Insert data for attendace.
  - Select employee name and attendance status for one employee by name on one date from the database.
