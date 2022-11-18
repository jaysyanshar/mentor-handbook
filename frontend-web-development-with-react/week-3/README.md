# Frontend Web Development with React Mentoring Week-3

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

- Understand how routing works.
- Able to implement React Router DOM to create SPA with multiple page.
  - Able to use URL Params to pass parameter from URL.
- Able to implement ChakraUI as UI element.
  - Able to use Common Elements from ChakraUI.

## Pre-requisites

- Code editor (e.g. VS Code).
- Web Browser
- Nodejs and PNPM installed in local computer

## Basic

### Simple Book Collection

- Concepts:
  - React Router DOM.
    - Able to create multipage website.
- Input:
  - A list of books with ISBN in JSON format, available at "src/data/books.json"
- Output:
  - A website showing the books in a list and link for book details that will show in another page. And a search page to filter books by name.
- Directions:
  - Example is available at folder "/books-collection"
  - Use CRA to create React app.
  - Install using pnpm react-router-dom
  - Aktifkan React Router DOM mengikuti halaman berikut: https://neurons.ruangguru.com/course/frontend-development/frontend-web-development-with-react/react-router-dom/introduction-to-react-router-dom
  - Import Routes and Route ke App.js
  - Create three routes:
    - endpoint "/" as "Home"
    - endpoint "book" as "Book List"
    - endpoint "book/:isbn" as "Book Detail" as child of "book"
  - Create folder "/routes" in "/src"
  - Create Components for each route
    - "Home" contains "Hello, welcome to my Book Collection"
      - Add Link that will redirect to endpoint "book"
    - "Book List" contains List of Books
      - Setup the input booksData here, by import from JSON file.
      - Create Component to render the data in folder "/cards"
        - Book name
        - Link to book detail using ISBN as param.
    - "Book Detail" contain Information about the book. 
      - Grab the isbn from useParams(), pass it to BookDetail Component that we'll create
      - Create Component to render the data in folder "/cards"
        - This is not ideal, but we also setup the input booksData here, by import from JSON file.
        - We'll search by isbn the actual book.
        - Then we'll show
          - Book name
          - Book description
  - Play around some more creating or modifying new route
  - If you want, you can create a wrapper to simulate API call.
  - You can also use fetch() and actually use 3rd party API.
  - There is no styling, feel free to style it before hand.
  
### Simple Book Collection with Styling

- Concepts:
  - Same as Simple Book Collection
  - ChakraUI as styling component
- Input:
  - Simple Book Collection
- Output:
  - Simple Book Collection with ChakraUI styling
- Directions:
  - Re-use the Simple Book Collection
  - Style using ChakraUI, remember to install ChakraUI first using pnpm

## Challenging

### Use 3rd party API and showcase the data

- Concepts:
  - Free
- Input:
  - Free
- Output:
  - Free
- Directions:
  - Free
.