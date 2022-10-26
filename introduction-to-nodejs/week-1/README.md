# Introduction to NodeJS Week-1

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

- Understand how to use PNPM
    - Able to init a package
    - Able to install a package
    - Able to install dev package
    - Able to instal global package
- Understand how to use CommonJS Modules in NodeJS
    - Able to use CommonJS export
    - Able to use CommonJS require
    - Able to use pass variables between two files using CommonJS module
- Understand how to use ES Modules in NodeJS
    - Able to use ESModules export default
    - Able to use ESModules import
    - Able to use ESModules named export
    - Able to use pass variables between two files using ESModules
- Understand how to use environment variables
    - Able to create an environment variable in his/her own computer
    - Able to use the environment variable using NodeJS process
    - Able to user environment variables using dotenv package

## Pre-requisites

- Code editor (e.g. VS Code).
- NodeJS with
  - PNPM installed

## Basic

### Simple PNPM
- Concepts: 
  - PNPM setup
  - PNPM Dev and Global
- Input: 
  - None
- Output:
  - Installation of nodemon in devDependencies and ifup in global 
- Directions: 
- Create a new folder
    - Run pnpm init
    - Show package.json
  - Install nodemon using --dev flag
    - Show package.json have nodemon as devDependencies
  - Install is-up using --global flag
    - Show package.json
    - Call is-up, see its README for details.
  - Move away from the package folder
    - Run is-up again, show that it works if we use --global
    - Run nodemon, show that it failed because it's installed as part of the package
  - Using Github with NodeJS
    - Create .gitignore for node_modules folder
    - Create git repo for the package folder
    - Link to Github
    - Upload to Github
    - Ask one mentee or all to download from Github
    - Then ask them to share screen and try to run nodemon which is a devDependencies
    - If it's not working, then ask them to run pnpm install
    - Run again, it should work now
    - Explain that this is why PNPM is powerful, it's good for cooperation with other people, having the same package, especially the version
- Rule
  - N/A

### Simple Calculator
- Concepts: 
  - Function and simple math calculation
  - PNPM setup
  - CommonJS Modules
- Input: 
  - None
- Output:
  - Two files 
- Directions: 
  - Create a new folder for the calculator files
    - Run pnpm init
    - Show package.json
  - Create a new file to hold simple calculation functions
    - Create addition function, followed by subtraction, multiplication and division.
    - Test the individual function works using console.log.
    - Export the functions using CommonJS, using modules.exports object.
  - Create a new file as the main process file
    - Use CommonJS import to get functions from the simple calculation file
    - Create console.log lines to do the calculations, use different numbers to differentiate between this and the testing done before
    - Run the process file and show that we're using the process file to run the calculator
- Rule
  - N/A

### Simple Calculator with ESModules
- Concepts: 
  - Function and simple math calculation
  - PNPM setup
  - ES Modules
- Input: 
  - None
- Output:
  - Two files 
- Directions: 
- Create a new folder for the calculator files
    - Run pnpm init
    - Show package.json
  - Use the file from Simple Calculator
    - Export the functions using ESModules, using named export in simple calculation file
    - Use ESModules import to get functions from the simple calculation file in process file
    - Run the process file and show that we're using the process file to run the calculator
- Rule
  - N/A

### Simple Environment Variables
- Concepts: 
  - PNPM
  - Environment variables
- Input: 
  - Environment variables
- Output: 
  - Can display the environment variables
- Directions:
  - Use the terminal/command prompt to EXPORT a variable, for example name with mentee name in his/her own computer, i.e. EXPORT nama=mentee_name
  - Create a folder and setup using PNPM init
  - Create a file and inside use process.env to grab the variable and output using console.log
  - It should work, it not check if perhaps the variable was not exported properly in his/her own computer.
  - Don't take too long, next
  - Install dotenv
  - Create a file callend .env
  - Fill the file with some variables, for example nama and favorite
  - Create a new file and use dotenv to console log the variables inside .env file
  - Experiment more
- Rule:
  - N/A

## Challenging

### Date Time Converter
- Concepts: 
  - Emulate https://www.timeanddate.com/
