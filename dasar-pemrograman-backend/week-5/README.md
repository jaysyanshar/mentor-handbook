# Dasar Pemrograman Backend Mentoring Week-5

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

- Understand how Concurrency works
- Understand how to use Authentication
- Understand how to use Custom Routing

## Pre-requisites

- Code editor (e.g. VS Code).
- Golang installed in local computer

## Basic

### Basic Login Page

- Concepts:
  - Authentication and Routing
- Input:
  - A simple HTML page showing login.
- Output:
  - Change to another page, deny access if not logged in
- Directions:
  - Create ListenAndServe
  - Create two routing, login and dashboard
  - Dashboard is protected, method of protection can use simple flag or use HTTP basic auth
  - Create login page with form.
  - Use static username and passwor for now.
  - If valid login, then show dashboard.
  - Content of dashboard can be a simple Hello World.

### Social Media Status Checker

- Concepts:
  - Concurrency
- Input:
  - A list of website in an array of string: e.g. https://facebook.com, https://twitter.com, https://linkedin.com
- Output:
  - A string showing HTTP Status Codes and website is up or down (one line for each website), e.g. "[200] https://facebook.com is up"
- Direction:
  - Create a function to output the website status.
  - Use simple http.Get() and format the output.
  - Show how the output is in order without using concurrency, if possible show the time taken to complete.
  - Update to use concurrency using go routine.
  - Show how the output is kind of random, if possible show the time taken to complete.
  - Update to use concurrency using channel
  - Reference: https://articles.wesionary.team/practical-example-of-concurrency-on-golang-fc4609ea8ed1
