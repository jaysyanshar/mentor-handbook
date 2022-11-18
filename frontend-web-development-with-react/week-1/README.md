# Frontend Web Development with React Mentoring Week-1

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

- Able to create a React app using CRA.
- Able to create React Functional Component.
- Able to use JSX in React.
- Able to use State in React.
- Able to use Props in React.

## Pre-requisites

- Code editor (e.g. VS Code).
- Web Browser
- Nodejs and PNPM installed in local computer

## Basic

### My first React App

- Concepts:
  - Create React App (CRA)
  - JSX
  - Functional Component
- Input:
  - N/A
- Output:
  - A webpage showing "Hello from Indonesia"
- Directions:
  - See final result in folder "/hello", keep in mind that the React App in "/hello" is mixed with the next challenge.
  - Use CRA to create a React app template, project (folder) name is up to mentor.
  - Open src/App.js, remove the template "<div>"
  - Add paragraph "Hello from Indonesia"
  - Look in browser.
  - Play around, adding static HTML element(s).

### My first React App with Props and State

- Concepts:
  - Same with My first React App
  - Props
- Input:
  - React App from my first React App.
- Output:
  - See final result in folder "/hello"
  - A webpage showing "Hello from student_name",
  - with text "I like Javascript 0" below it
  - and a button that when clicked will update the number 0 on the screen, each click will increment by 1.
- Directions:
  - Continue from the previous challenge.
  - Create new folder "/comp" inside "/src" folder.
  - Create new file "hello.js".
  - Create new function called "Hello" that receives 1 parameter "props"
  - Inside create "h1" element with content "I am {props.name}", explain a little that for now it's a placeholder. Pay attention to "name". Remember to export the function.
  - Return to App.js
  - Import Hello from hello.js
  - Use it inside the div, creating a new "component", adding element "name" with student name or mentor name or the house name.
  - Check the browser
  - Explain what happens and how to pass parameter.

  - Add new element and update the props in hello.js to make sure student understand how props work.
  - Now go back to hello.js, add "p" element, note that since we are returning 2 elements, we need to create a container.
  - Put inside "p" element the text "I like Javascript {counter}", explain that counter is going to be used like a variable where we will change it.
  - Import useState hook.
  - Create a state called "counter" with initial value of "a little"
  - Add button with "onClick" event embedded and the text "Change" on the button where the onClick will trigger setCounter and change the value to "a lot".
  - Open browser and click on the button, it should change to "a lot"

  - Experiment more with state and props, keep it simple though, remember that we are only creating simple React website using state and props, nothing fancy yet. Keep it for next week.

## Challenging

### Calculator

- Concepts:
  - CRA
  - Functional Component
  - State
  - Props
- Input:
  - Basic calculator template (create one before the session)
- Output:
  - A working calculator
- Directions:
  - Make sure all the buttons already have the "onClick" event, so student only needs to pass the parameter. Event handling will be covered next week.
  - This is not tested, so make sure you create one yourself before using this example to make sure it fits with the material covered for week-1.
