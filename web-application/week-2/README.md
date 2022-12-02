# Web Application Mentoring Week-2

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
- Able to serve data using REST API.
- Able to use database as storage and source for REST API server.

## Pre-requisites

- Code editor (e.g. VS Code).
- Golang installed in local computer
- Postman installed in local computer
- Postgres installed in local computer

## Basic

### Quote App Server

- Concepts:
  - REST API
  - HTTP Client
  - Handling JSON format
  - HTTP Server
  - Database management
- Input:
  - Use Quote App with Database from Week-1
  - OR Find any free quote API and store some quotes to the database or choose from these two:
    - https://animechan.vercel.app/
    - https://zenquotes.io/
- Output:
  - API endpoints, at least 2
    - 1 endpoint for random quote
    - 1 endpoint for quote of the day
- Directions:
  - Store the data into Postgres if not already.
  - Create HTPP server with 2 endpoints
    - 1 for /random/ (endpoint can be changed)
    - 1 for /quote_of_the_day/  (endpoint can be changed)
  - Random endpoint should return random quote from database everytime the page is called.
  - Quote of the day should only change every 24 hours, so each call if still within same day, should not change the quote.
  - Style the quote as desired.

## Challenging

### N/A

- Concept
  - N/A
- Input:
  - N/A
- Output:
  - N/A
- Directions:
  - N/A
