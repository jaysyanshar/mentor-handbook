# Web Application Mentoring Week-1

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

- Understand REST API.
- Able to retrieve data using REST API.
- Able to serve data using REST API.

## Pre-requisites

- Code editor (e.g. VS Code).
- Golang installed in local computer
- Postman installed in local computer
- Postgres installed in local computer

## Basic

### Quote App

- Concepts:
  - REST API
  - HTTP Client
  - Handling JSON format
  - HTTP Server
- Input:
  - Find any free quote API or choose from these two:
    - https://animechan.vercel.app/
    - https://zenquotes.io/
- Output:
  - Single qoute output to HTML page.
- Directions:
  - Read the documentation for the API.
  - Pick the random endpoint or endpoint that will output a single quote, e.g.
    - https://animechan.vercel.app/api/random
    - https://zenquotes.io/api/random
  - Use Postman to send the API and see what the output format is.
  - Create HTTP Client that will send the request.
  - Create HTTP Server that will output the single quote to the browser.
  - Style the quote as desired.
  - See example at "/quoteApp/" folder.

## Challenging

### Quote App with Database

- Concepts:
  - Same as Quote App
  - Database CRUD
    - Create database
    - Create Table
    - Insert Data into Table
    - Select Data from Table
- Input:
  - Same as Quote App
- Output:
  - Quote inserted to database
- Directions:
  - Best to continue using the Quote App.
  - Create database first to store the quotes.
  - Create table to store the quotes.
  - Add insert into database for the quotes, if there is endpoint that can get .more than 1 data, then use that endpoint.
  - Use select to return the count of quotes in the table.
  - And use select to return one quote of the day.
