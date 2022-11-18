# Frontend Web Development with React Mentoring Week-2

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

- Understand how rerendering works in ReactJS.
- Understand and able to implement React Hooks for Functional Component.
- Lifting State Up in ReactJS.
- Event handling in ReactJS.
- Render list with keys in ReactJS.
- Conditional Rendering.

## Pre-requisites

- Code editor (e.g. VS Code).
- Web Browser
- Nodejs and PNPM installed in local computer

## Basic

### HRApp - Employee Data

- Concepts:
  - Render data using list and keys.
  - Conditional rendering.
- Input:
  - A list of employee data in JSON: [{"id": 1,"name": "Jason"},{"id": 2,"name":"Edwin"},{"id": 3,"name": "Jack"},{"id": 4,"name": "Jane"},{"id": 5,"name":"Jill"}]
- Output:
  - A website with box of names of employee with a button that when clicked will increase the view count. The box of names will have alternating background color.
- Directions:
  - See folder "/employee" for showcase.
  - Feel free to style the EmployeeCard beforehand.
  - Use Create React App to create a new React project.
  - Replace the template with empty "div" in App.js
  - Create employee data in App.js, just copy the JSON data and create a variable for it.
  - Create folder "/cards" inside "/src"
  - Create file "employee.js" inside folder "/cards"
  - Create and export function EmployeeCard() receiving Props.
  - Create JSX with "div" and content is "props.name"
  - Import EmployeeCard into App.js
  - Use Component EmployeeCard in main function App() as part of a List, using map()
  - Congrats we have successfully created a list in ReactJS. Play around with the list and keys.
  - Next we are going to alternate the colors between odd and even rows.
  - There are multiple ways to do this, the main point is to use "conditional rendering" here:
    - activate the index from map, we'll use that as the condition for alternating color.
    - pass the index to the EmployeeCard
    - use it to create a conditional for background color, for example, white for even, light gray for odd.
    - That's it, feel free to experiment and play some more.

### HRApp - Employee Data Interaction

- Concepts:
  - Continue from HRApp - Employee Data
  - Event handling
  - Hooks
  - Component Data Flow
- Input:
  - Continue from HRApp - Employee Data
- Output:
  - Continue from HRApp - Employee Data
  - Add Present and Absent status which can be clicked and will affect the text on the button.
- Directions:
  - Continue from HRApp - Employee Data
  - See folder "/employee-interaction" for showcase.
  - Create new file "attendance.js" in folder "cards"
  - Return a button element with text "Present"
  - Add this Component to EmployeeCard
  - Create a state in EmployeeCard that will hold status of "absent" or "present".
  - Create a function to changeState from Present to Absent and Absent to Present.
  - Pass the state and it's function to AttendanceButton
  - Change the button text from "Present" to use Props."state"
  - Add event onClick to "button" inside "attendance.js" that will change the state.
  - For clarity, add the "state" next to "name".

## Challenging

### Calculator

- Concepts:
  - Same as Week-1 Calculator
  - Event handling
- Input:
  - Basic calculator template (create one before the session)
- Output:
  - A working calculator
- Directions:
  - This time the student will create their own event handler.
  - This is not tested, so make sure you create one yourself before using this example to make sure it fits with the material covered for week-1 and week-2.