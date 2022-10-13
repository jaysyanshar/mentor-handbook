# Dasar Pemrograman Backend Mentoring Week-3

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

- Understand how to use Struct
- Understand how to use Interface
- Understand how to handle errors
- Understand how to import package

## Pre-requisites

- Code editor (e.g. VS Code).
- Golang installed in local computer

## Basic

### Simple Attendance Data with Struct and Interface

- Concepts:
  - Struct, interface
- Input:
  - A string of "Present, Absent, Sick" multiple times separated by names. E.g. "Budi, Present, Present, Sick, Present, Ani, Present, Present, Absent, Absent"
- Output:
  - A struct containing the key of name and rawData. Name is self explanatory, rawData is the string of Present, Sick, Absent for the name. If two names, then two struct. Example: { name: Budi, rawData: "Present, Present,  Sick, Present" } and { name: Ani, rawData: "Present, Present, Absent, Absent" }
- Directions:
  - Create an interface method populate()
  - Create a struct with key "name, rawData" and default value.
  - Create method struct populate()
  - Use the populate to fill the struct value.

### Simple Attendance with Error Handling

- Concepts:
  - Struct, Interface, Error handling
- Input:
  - Same as above, with the addition of string "Not Present" mixed in.
- Output:
  - Same as above
- Direction:
  - Same as above
  - Add a check in populate() that if the string rawData contains "Not Present", it will return an error "Not a valid value".

### Simple attendance with package

- Concepts:
  - Struct, internal package
- Input:
  - Use the struct above.
- Output:
  - Update the struct with date and time
- Directions:
  - Import package time
  - Using the struct above, add new key "date" and fill in the key, the first is filled with the first date of this month. E.g. Today is October 2022, so the first data is 1 Oct 2022, second data is 2 Oct 2022, and so on.

## Challenging

### Attendance Data

- Concepts:
  - Same as Simple Attendance Data with Struct and Interface
- Input:
  - Same as Simple Attendance Data with Struct and Interface
- Output:
  - Similar to Simple Attendance Data with Struct and Interface, but use array of struct, instead of 2 struct.
- Directions:
  - Same as Simple Attendance Data with Struct and Interface with the exception of end result should be array of struct.
