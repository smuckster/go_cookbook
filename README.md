# Go Cookbook

## Getting started

### .env
Your .env file should be in the project root and contain the following variables:

```
DB_USER: The username for the MySQL database
DB_PASSWORD: The user's password
DB_NAME: The name of the database for the application
```

### Compiling
When compiling the application on a Raspberry Pi, use the following command:
`env GOARCH=arm64 GOOS=linux go build .`
