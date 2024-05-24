# Attendance-QR-Portal

This project is inspired by the idea of simplifying attendance tracking and access management. The core concept revolves around granting students or users permission to access certain locations, such as a room or building. Traditionally, this would involve signing a register or book with personal details like name and class.

With the Attendance-QR-Portal, this process is streamlined. Students or users first create an account, and upon logging in, they can easily check-in using the website. This can be done conveniently via a QR code placed at the location.

The system provides administrators or moderators with insights into the current access status, including the number of people accessing the location, their purpose, and their identities.

## Features
- User account creation and authentication
- Easy check-in using QR codes
- Real-time access tracking and monitoring
- Admin dashboard for managing access permissions and viewing access logs

## Technologies Used
- HTML/CSS
- JavaScript
- Golang (Backend)
- MySQL (Database)
- Tailwind CSS (Styling)
- QR Code Generator (For generating QR codes)


## Instructions
 `git clone https://github.com/Sreeram0045/Attendance-QR-Portal.git`

 `cd Attendance-QR-Portal/`

  make a file called db_connection and use this format to add your credentials
  (without the <, > characters)
  ```
  Username=<your mysql username>
  Password=<your mysql password>
  DBName=<database name>

  ```
 `go run cmd/backend/main.go`
   
