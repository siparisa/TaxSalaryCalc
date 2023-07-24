# Tax Salary Calculation

The Tax Salary Calculation application provides an HTTP API with an endpoint that calculates the total income tax based on the annual income and tax year provided. The application returns the result of the calculation in JSON format.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Make Commands](#make-commands)
- [Project Layout](#project-layout)
- [Environment Variables](#Environment-Variables)

## Prerequisites

Before running the application, make sure you have the following prerequisites installed:

- Go 
- Docker

## Getting Started

To get started, follow these steps:

1. Pull the Docker image:

    ```shell
    docker pull ptsdocker16/interview-test-server
    ```

2. Run the Docker container:

    ```shell
    docker run --init -p 5000:5000 -it ptsdocker16/interview-test-server
    ```
      feel free to use any avaiable port on your system

3. Navigate to the root directory of the project in the command line.

4. Run the following command to start the application:

    ```shell
    make run
    ```

5. To run the API, use the following endpoint:

    ```
    http://localhost:8080/income-tax/calculate-tax?year=X&salary=Y
    ```

   Replace `X`, and `Y` with the appropriate values. For example:

    ```
    http://localhost:8080/income-tax/calculate-tax?year=2020&salary=78000
    ```
## API Documentation
   Endpoint: `/income-tax/calculate-tax`
   
   Calculates the total income tax based on the provided annual income and tax year.

   Request Method: `GET`

   Parameters:
   
   1.`year` (required): The tax year for which the calculation is performed. Format: YYYY (e.g., 2022).

   2.`salary` (required): The annual income amount for the calculation.

   Example Request:

  `GET /income-tax/calculate-tax?year=2022&salary=50000`
  
  Example Response: 
  ```
  HTTP/1.1 200 OK
  Content-Type: application/json
{
"taxAmount": 7500.0
}
```
`taxAmount`: The calculated total income tax amount.

Error Responses:

HTTP/1.1 400 Bad Request

Content-Type: application/json

Body:
```
{
  "error": "Invalid Salary",
  "details": "salary value is not numeric."
}

```

HTTP/1.1 500 Internal Server Error

Content-Type: application/json

Body:
```
{
  "error": "Internal Server Error",
  "details": "Failed to get tax bracket."
}

```
## Testing

To run the tests, navigate to the root directory of the project in the command line and run the following command:

    `make test`

## Make Commands
  
   |  Command       | Description                                                         |
   |----------------|---------------------------------------------------------------------|
   | run            | Starts the service and all necessary dependencies in the foreground |
   | test           | Starts tests                                                        |    

## Project Layout

This project roughly followed the layout of Go projects as described at
[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout).

| Directory              | Description                                                                                    |
|------------------------|------------------------------------------------------------------------------------------------|
| `cmd/`                 | This Go package is where `main` is used for the executables of the project                     |
| `internal/`            | Application specific Go packages, e.g., they cannot be shared and are specific to this service |
| `internal/repository`  | Any files relating to repository layer                                                         |
| `internal/database`    | Any files relating to database connection and configuration                                    |
| `internal/migrations/` | Any files relating to migration                                                                |
| `internal/tests/`      | tests for the service are located in here.                                                     |

### Layers and Folder Structure

There are 3 main layers in this repo, including `Controller`, `Service`, and `Repository`. The only way for these layers
to interact with each other should be through their interfaces. The lower layers do not have any knowledge about
the upper layers.

## Environment Variables

The following environment variables are [defined in Tax Calculator App](./.env), and can be used to
influence behaviour.

| Name                                    | Description                    |
|-----------------------------------------|--------------------------------|
| `PORT_TAX_YEAR`                         | Tax Bracket Service Port       |
| `PORT_APP`                              | Salary Tax Calculator App Port |
