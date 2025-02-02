# Shinko

Welcome to Shinko!

This repository holds the portfolio project for a habit formation application.

![CI Status](https://github.com/Joeavaikath/shinko/actions/workflows/ci.yml/badge.svg)
![CD Status](https://github.com/Joeavaikath/shinko/actions/workflows/cd.yml/badge.svg)

## Description

Shinko is a habit formation application designed to help users build and maintain good habits through tracking and reminders. Streaks help users persevere!

## Features

- Track daily habits
- View progress over time
- Get recurrence goals for streaks

## Installation - Local Run

### Requirements

1. **Install Go dependencies**: `go install` your way through the present go.mod
2. **.env file**: The entry point for the app is `cmd/shinko/main.go`. It requires a .env file in the same directory. Key variables:
    - `DB_URL`: This is your database's connection string. PostgreSQL is required, as the queries were written with that dependency.
    - `JWT_SECRET`: Used for your JWT tokens. Create one using `openssl rand -base64 64`.
    - `PLATFORM`: Not used too much, but can qualify your logs behind this value.

3. **Set up PostgreSQL**
    - **Installation**
        - Mac: `brew install postgresql@15`
        - Linux: `sudo apt install postgresql postgresql-contrib`
        - Verify it worked with `psql --version`.

    - **Run it as a server in the background**
        - Mac: `brew services start postgresql@15`
        - Linux: `sudo service postgresql start`

    - **Enter the psql shell**
        - Mac: `psql postgres`
        - Linux: `sudo -u postgres psql`

    - **Create the database**
        ```sql
        CREATE DATABASE shinko;
        ```

    - **Connect to it**
        ```sql
        \c shinko
        ```

    - **Alter the user password to something you can remember. Used in our connection string**
        ```sql
        ALTER USER postgres PASSWORD 'postgres';
        ```

    - **Create your connection string using the following format**
        ```
        protocol://username:password@host:port/database
        ```
        For example:
        ```
        postgres://postgres:postgres@localhost:5432/shinko
        ```

    - **Test your connection string**
        ```
        psql "postgres://postgres:postgres@localhost:5432/shinko"
        ```

4. **Set up the schema for our app**
    - Use goose.
        ```
        go install github.com/pressly/goose/v3/cmd/goose@latest
        goose -version
        ```
        to ensure it's installed.
    - Navigate to `sql/schema`
    - Run
        ```
        goose postgres <connection_string> up
        ```
    - If all goes well, run
        ```
        psql shinko
        \dt
        ```
        to see your tables!

## Usage - Local

Recommend using something like Postman or Thunder Client to run your HTTP requests.

## Installation and Usage for Operator

Coming soon™️

## License

This project is licensed under the MIT License.
