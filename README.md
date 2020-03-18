# simple-account
A simple account system
***

## ⚠️⚠️ DISCLAIMER ⚠️⚠️

It's my first code in GO, all advisement and help are welcome.


## API endpoints

This API's uses the OpenApi 3.0 standard for it's documentation. You can access it through a [Redocs's page](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/hbontempo-br/simple-account/master/api/endpoint_documentation.yaml) or, if you prefer, directly through it's yaml [file](api/endpoint_documentation.yaml). 

## API Install Guide

*This project is meant to run using **GO 1.14** in an Unix-like system , other versions and OSs may not work.*

### Clone project
```sh
$ git clone git@github.com:hbontempo-br/simple-account.git
```

#### Install dependencies

All dependencies are listed on bundled [go.mod](go.mod) file at file. Use the go mod command from within the project's folder to install then.

```sh
$ go mod download
```

#### Environment Variables

This project requires this environment variables:

| Variable | Definition | Example |
| - | - | - |
| DB_ADDRESS | The address of application MySQLs database | 11.111.111.111 |
| DB_DATABASE | Name of the database that the data is registered | account |
| DB_USER | Application's MongoDB user | CoolUser |
| DB_PASSWORD | Application's MongoDB password | CoolPassword |
| SERVER_PORT | Porth on witch the server is going to run | :8000 |


## Running the API locally

```sh
$ go run main.go
```

## Contributing

This is a personal study project. All contributions are VERY welcome. for direct contact please reach me in me@hbontempo.dev .

## Docker

Very simple build and run: very basic stuff. Just build it using the [DockerFile](Dockerfile) and run it however you 
want. Just make sure you have the required environment variables when running it and expose the selected application's
port adequately.

## 🚧 Tests

I'm still haven't figure it out yet. I will get there.


## TODO

- Better error handling (all application)
- Fix outgoing response on error
- Helper functions the handle parameters
- Add correct parameters to [OpenAPI file](api/endpoint_documentation.yaml) (and probably a error check)
- DTO to "clean" outgoing response
- Make the reference work in Gorm
- Probably a whole reorganization of the code (for a more idiomatic approach)
- Add a better logging packageMake created_at filed work on models
- Create GetTransactionList method on transaction_controller.go
- Probably controllers should Structs with the controllers package
- Add real filters to GetAccountList on account_controller.go
- Add method ToString to OperationModel
- All resources have basically no validation whatsoever
- Review DB middleware organization. (Probably move functions to methods or something like it)
- Transaction resource could use something like go routines to make queries simultaneously and save time
- For now the enumerator on the SQL query (OperationType) looks ok, but will have to change if more complex stuff appear (more columns?)
- mysql_connector could use a restructure. Move loose functions to methods of a "class" (struct)
- Put a path prefix on main.go regarding API's version for better version control of the API's client



<br/>
<br/>
<br/>

---

