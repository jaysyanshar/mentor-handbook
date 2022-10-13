# Dasar Pemrograman Backend Mentoring Week-2

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

- Understand how to use first-class function
- Understand the use of pointer
- Understand how to use collection data type
    - Array
    - Slice
    - Map

## Pre-requisites

- Code editor (e.g. VS Code).
- Golang installed in local computer

## Basic

### Simple Attendance Data

- Concepts:
  - Create and make changes to array
- Input:
  - One string separated by comma containing 'Present/Absent/Sick'.
- Output:
  - An array of the data from the string
- Directions:
  - Create a string containing value of "Present, Absent, Sick" multiple times, e.g. "Present, Absent, Sick, Present, Present, Present"
  - Create an array from the string.
  - show the result
  - Change all the 'Absent' to 'Present'
  - show the result
  - Mix in pointer and first class function if possible.

### Simple Slice of Attendance Data

- Concepts:
  - Create and make changes to a slice
- Input:
  - A string containing 'Present/Absent/Sick' multiple times, mixed with names, can be 1 or 2 names. The name will the separator. So, the data is an attendance data in the following format: name, status, status, status, status, name, status, status, status, status.
- Output:
  - Slice of attendance data, if there are two name, then there are two slices.
- Direction:
  - Convert the string to array first is preferred.
  - Then slice the array.
  - Mix in pointer and first class function if possible.

### Simple Map of Attendance Data

- Concepts:
  - Create and make changes to a map
- Input:
  - A string of containing 'Present/Absent/Sick' multiple times. e.g. "Present, Present, Absent, Present"
- Output:
  - A map containing key 'Present, Absent, Sick' with the number of time the word is present in the string as the value. e.g. Present:3, Absent:1, Sick:0
- Directions:
  - Create the string first.
  - Create the map with the keys
  - Process and update the keys with the number of times the key shows up in the string.
  - This is not a test, so it's OK to simply type in the key manually into the map. No need to make it dynamic.

## Challenging

### Dynamic Simple Map of Attendance Data

- Concepts:
  - Follow Simple Map of Attendance Data
- Input:
  - Follow Simple Map of Attendance Data
- Output:
  - Follow Simple Map of Attendance Data
- Directions:
  - Follow Simple Map of Attendance Data, except the last directions, in this challenge, it must be dynamic.

### Multi-dimensional

- Concepts:
  - Slice of Map
- Input:
  - A string containing 'Present/Absent/Sick' multiple times, mixed with names, can be 1 or 2 names. The name will the separator. So, the data is an attendance data in the following format: name, status, status, status, status, name, status, status, status, status.
- Output:
  - Slice of map with name and attendance data, e.g.
    map[Budi: Present:3, Absent:1, Sick:0]
    map[Adi: Present:4, Absent:0, Sick:0]
- Directions:
  - Free flow
