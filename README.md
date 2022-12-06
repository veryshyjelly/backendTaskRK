# Task 2 RK hall open soft
### OVERVIEW
    Design a Room Booking System to allocate and view rooms of a student for RK Hall, IIT Kharagpur. All The Best :)
## Problem Statement
    Create two types of users, Student, and Admin. The users should be pre-defined in the database and don’t need to register.
    
### Student
    1. Can log in
    2. Can view the details of all rooms
    3. The Student model should consist of Name, Roll Number, Contact Number, Email id Block Number, and Room Number

### Admin
    1. Can log in
    2. Can view the details of all rooms and list of students
    3. Can allot rooms to the Students
    4. The Admin model should consist of Name, Contact Number, Email id, and Position (Gsec, Hall President, Warden)



## Instructions
1. When we open the web application, the login page should appear with an option to login as a student and login as an Admin.
    - Login as a Student - 
        - The homepage should contain a table showing all the rooms (room number) and the name of the student in the room allotted to.
        - On clicking the name of the student, a modal should appear containing the details of the student (Name, Roll Number, Contact Number, Email id)
        - There should be a separate page for the Profile containing the details of the user, including his name, roll number, contact number, email id, room number and an option to edit contact number.
    - Login as an Admin -
        - The homepage should again contain the table showing the details of the rooms.
        - There must be a separate page named Allot Rooms; the page should allow the Admin to allot rooms to the students
        - Allot Rooms Page - a dropdown to select rooms, a dropdown to select the roll number of student, after the student is selected, the complete information of him should be displayed along with the new allotted room number, the save button allots the room to the student.
2. Assume all rooms are single
NOTE: The frontend should be made using React.
