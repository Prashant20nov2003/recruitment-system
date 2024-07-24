# Recruitment Management System

## Overview

The Recruitment Management System is a full-stack application designed to manage job openings and applicant resumes. It allows users to create profiles, upload resumes, view job openings, and apply for jobs. Admin users can manage job openings, view resumes, and review applicant data.

## Features

- **User Profile Management**: Users can create profiles and upload resumes.
- **Resume Parsing**: Uploaded resumes are processed using a third-party API to extract relevant information.
- **Job Management**: Admins can create and manage job openings.
- **Job Application**: Applicants can view job openings and apply for jobs.
- **Resume Review**: Admins can view all uploaded resumes and extracted data.

## Tech Stack

- **Backend**: Golang Gin
- **Database**: SQLite
- **Resume Parsing**: [Third-party API](https://api.apilayer.com/resume_parser/upload)

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.22 or later)
- [SQLite](https://www.sqlite.org/download.html)
- [GORM](https://gorm.io/) for ORM
- Environment variables for the API key

## Setup

### Clone the Repository

```bash
git clone https://github.com/yourusername/recruitment-management-system.git
cd recruitment-management-system
```
Install Dependencies

```bash
go mod tidy
Configure Environment Variables
Create a .env file in the root directory and add your API key:
```


```bash
RESUME_PARSER_API_KEY=your_api_key_here
Run Database Migrations
Run the migration script to set up the database schema:
```

```bash

go run migrate/migrate.go
``` 
Run the Application
Start the backend server


```bash

go run cmd/main.go

```

# API Endpoints
POST /signup: Create a user profile.

Body: { "Name": "string", "Email": "string", "Password": "string", "UserType": "string", "ProfileHeadline": "string", "Address": "string" }
POST /login: Authenticate users and return a JWT token.

Body: { "Email": "string", "Password": "string" }
POST /uploadResume: Upload resume (PDF or DOCX). Only accessible to applicants.

Header: Authorization: Bearer <token>
Body: Form-data with file field resume
POST /admin/job: Create a job opening. Only accessible to admins.

Header: Authorization: Bearer <token>
Body: { "Title": "string", "Description": "string", "PostedOn": "datetime", "CompanyName": "string" }
GET /admin/job/{job_id}: Fetch job details and list of applicants. Only accessible to admins.

Header: Authorization: Bearer <token>
GET /admin/applicants: Fetch all users. Only accessible to admins.

Header: Authorization: Bearer <token>
GET /admin/applicant/{applicant_id}: Fetch applicant's extracted data. Only accessible to admins.

Header: Authorization: Bearer <token>
GET /jobs: Fetch job openings. Accessible to all users.

Header: Authorization: Bearer <token>
GET /jobs/apply?job_id={job_id}: Apply to a job opening. Only accessible to applicants.

