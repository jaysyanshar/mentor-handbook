# Database Management Mentoring Week-2

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

- Understand how to implement SQL Commands using GORM
  - Able to Insert
  - Able to Update
  - Able to Delete
  - Able to Query

## Pre-requisites

- Code editor (e.g. VS Code).
- Golang installed in local computer
- Postgres installed in local computer

## Basic

### HRApp - Employee Data

- Concepts:
  - Same as Week-1 but using GORM
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
  - Data manipulated in database
- Directions:
  - Open Postgres pgAdmin / DBeaver / other tool
  - Use Database "hrApp"
  - Drop table using GORM.
  - Create table using GORM, make note of the difference using GORM.
  - Insert some data using GORM, make note of the difference using GORM.
  - Select some data using GORM, make note of the difference using GORM.
  - Do update and delete like insert, using GORM, make note of the difference using GORM.

### HRApp - Department Data

- Concepts:
  - Same as Week-1 but using GORM
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
  - Data manipulated in database
- Directions:
  - Open Postgres pgAdmin / DBeaver / other tool
  - Use Database "hrApp"
  - Drop table using GORM.
  - Create table using GORM.
  - Insert some data using GORM.
  - Select some data from the table using GORM.
  - Optional
    - Do update and delete like insert using GORM.
  - Select the name of the department head using SQL JOIN.

## Challenging

### Attendance Data

- Concepts:
  - Same as Week-1 but using GORM
    - All the above challenges
    - Database relationship / association
- Input:
  - As above
  - Add attendanceStatus with one field ( status ) which will be filled with present or absent.
- Output:
  - One association table called attendance
  - Data inserted into database
- Directions:
  - Use Database "hrApp"
  - Drop table attendanceStatus.
  - Drop association table attendance.
  - Insert data for attendanceStatus.
  - Insert data for attendace.
  - Select employee name and attendance status for one employee by name on one date from the database.
