# Dasar Pemrograman Backend Mentoring Week-4

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
- For 3rd party API, google with keyword "free api data"
  - Some sample: https://github.com/public-apis/public-apis

## Learning Objectives

- Understand how to do Unit Testing
- Understand how to interact with files (I/O)
- Understand how to create a HTTP server
- Understand how to use HTTP client

## Pre-requisites

- Code editor (e.g. VS Code).
- Golang installed in local computer

## Basic

### Repeating string testing

- Concepts:
  - Testing using Ginkgo
  - A simple function that given a repeating string will count the number of repeated string.
- Input:
  - A word repeated multiple times separated by space.
- Output:
  - The number of repeated word
- Directions:
  - Create the function that accepts a string. It is assumed it's the same word repeated multiple times separated by space.
  - Inside the function count the number of repeated word and return the value.
  - If there are two or more words, count the first word in the string.
  - Create a ginkgo scaffolding.
  - Create a unit test function to test the function.
  - Check if the input is empty.
  - Check if the input is two words, the result should be the first word.
  - Check if the input is repeated 5 times.
  - Check if the input is only 1.

### Read and write simple JSON file

- Concepts:
  - IO File, IO JSON
- Input:
  - A simple JSON structure inside a file, can be as simple as first_name, last_name, birthdate
- Output:
  - Will read the JSON file and output to terminal, then update the JSON file with a new data.
- Direction:
  - Create a simple JSON file containing one data.
  - Create a struct to hold the JSON data.
  - Read the file using io
  - Read and convert the JSON content to struct.
  - Output the JSON struct to terminal.
  - Append a new JSON data to struct
  - Write the new JSON to file, can use append or overwrite the file.

### Connect to 3rd party API

- Concepts:
  - Connect to a free 3rd party API
- Input:
  - Free 3rd party API data
- Output:
  - The data retrieved shown to terminal
- Directions:
  - Create the client
  - Connect to 3rd party API data
  - If processing is needed then process it.
  - Output the data to terminal.

### Simple web server

- Concepts:
  - Serve a simple calculator web server 
- Input:
  - None
- Output:
  - A HTML file that will have two input boxes and 4 buttons (+, -, *, /)
- Directions:
  - Setup ListenAndServe()
  - Can use template
  - Show a HTML file with form containing 2 input boxes and 4 buttons (+, -, *, /)
  - On click of the button will perform the calculation and output to terminal or redirect to new page with the calculation and the result, or on the same page, simply create a container.

## Challenging

### Full application for anime or pokemon or other API

- Concepts:
  - Create a web server that will retrieve data from 3rd party API and then display the title only with a link to a detail page. If clicked, then the detail page will have more data.
- Input:
  - 3rd party API data
- Output:
  - A web application that can display data and have a favorite button implemented.
- Directions:
  - Create a client first to get the data from 3rd party API
  - Create a server to serve the data
  - First show the list with link to a detail page using route
  - When the link is clicked then the detail page will show detail data based on the title.
  - Have a link to get back to the list.
