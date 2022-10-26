# Interacting with Web APIs Mentoring Week-1

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

- Understand how DOM works
    - Able to select chosen HTML element
    - Able to add new HTML element
    - Able to change style for chosen HTML element
- Understand how to use Browser Events
    - Able to add event handler to chosen HTML element
    - Able to create interaction using event handler
- Understand how to use Web APIs
    - Able to use common Web API
    - Able to use Web Storage

## Pre-requisites

- Code editor (e.g. VS Code).
- Web based playground
    - Web Browser with Developer Tools.
        - Bisa dengan membuat index.html berisi DOM yang akan refresh ketika kita mengetik kode JS (perlu effort dari mentor)
        - Bisa dengan extension LiveShare atau lainnya.
        - Atau metode lainnya.
    - Online playground such as jsbin.com

## Basic

### Simple Calculator
- Concepts: 
  - HTML form, DOM, event handler
- Input: 
  - A form mimicking a calculator
- Output:
  - Correct calculation using the buttons on the calculator
- Directions: 
  - Create HTML form with the calculator
- Rule
  - N/A

### Simple Validator
- Concepts: 
  - HTML form, preventdefault, conditional
- Input: 
  - HTML form with a username and password inputs and login button
- Output: 
  - Correct username and password will redirect to a dashboard page. A simple Hello will do for the dashboard page.
  - Wrong username and password will result in the label text having a red color and a box with red background color showing the text "Wrong username and password"
- Directions:
  - Create a form with two input type and one button.
  - Create script at the bottom of the HTML page.
  - Create a dashboard page.
  - Set form action to the dashboard page
  - Write the script to preventdefault on login button click.
  - Create conditional where if username and password are correct, then form will be submitted.
  - If username and password are wrong, then change the label text color to red and create a box with red background and the text "Wrong username and password".
- Rule:
  - Username and password is hardcoded into the script.
  - Use one function for each button click.
  - The #result is a div, so use innerText.

## Challenging

### Online Calculator
- Concepts: 
  - Form, DOM, event handler, functions
- Input: 
  - Google "online calculator"
- Output:
  - Calculation result
- Directions:
  - Create something like Google calculator. Try the following text in Google: "5*5", a calculator will be shown, that's what we want to build.
- Rule:
  - Start with the numbers and basic arithmetic, before going out into mathematical function such as sin, cos, etc.
