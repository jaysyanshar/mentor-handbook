# Programming with Javascript Mentoring Week-3

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

- OOP (Theoritical)
    - Explain basic OOP theory
        - class as a blueprint
        - property as variable in OOP
        - method as function in OOP
        - instantiate using new
        - this as reference to self in OOP
    - Explain method chaining theory
- Callback
    - Able to use create callback function
- Promise
    - Able to create a promise
    - Able to use a promise
- Async Await
    - Able to create an async function
    - Able to use await inside the function

## Pre-requisites

- Code editor (e.g. VS Code).
- Web based playground
    - Web Browser with Developer Tools.
        - Bisa dengan membuat index.html berisi DOM yang akan refresh ketika kita mengetik kode JS (perlu effort dari mentor)
        - Bisa dengan extension LiveShare atau lainnya.
        - Atau metode lainnya.
    - Online playground such as jsbin.com

## Basic

### Simple Shopping using Callback
- Concepts: 
  - Create and use callback
- Input: 
  - One number with integer
  - One list of array of object with two keys: "product" and "price", 3 objects should be enough
- Output:
  - A series of console.log, one for each callback
- Directions: 
  - Create a number with integer which will be the money in wallet
  - Create the array of object as above
  - Create callback function so, each time the callback function is called, something will be bought and money in wallet will be deducted.
  - Output will be product bought and money leftover
  - Console.log the output from inside the callback function.
  - Repeat until money is not enough.
- Rule
  - N/A

### Simple Shopping using Promise
- Concepts: 
  - Create and use Promise
- Input: 
  - Same as above
- Output: 
  - Same as above
- Directions:
  - Same as above, but using Promise
- Rule:
  - N/A

### Simple Shopping using Async Await
- Concepts: 
  - Create and use Async Await
- Input: 
  - Same as above
- Output: 
  - Same as above
- Directions:
  - Same as above but using Async Await
- Rule:
  - N/A

## Challenging

### Advanced Shopping
- Concepts: 
  - Able to use either Callback, Promise or Async Await
- Input:
  - Same as above
- Output:
  - An array of string containing the repeated string 'I bought "product_name" for "product_price"' in the middle.
  - The beginning of the array will have the string 'I have "start_money_in_wallet"'
  - The end of the array will have the string 'I still have "end_money_in_wallet left"'
- Directions:
  - Create callback/promise/async await
  - Output should be stored in a string
  - Create global array to store the string
  - Create the start string
  - Create the end string
  - Combine the strings into an array
- Rule:
  - N/A
