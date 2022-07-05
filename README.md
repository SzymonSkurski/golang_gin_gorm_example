# golang_gin_grom_example

gin + GORM REST API example

## run from CLI

`go run .`  
if u don't have gcc try  
`apt-get install build-essential` to install gcc/g++  
app demand mysql database  
u have to set DB ENV to connect

## run from CLI for windows users
gorm demand gcc (allows run C scripts in GO)  
u can see such error message `exec: "gcc": executable file not  found in %PATH%`
download gcc from here:  
https://jmeubank.github.io/tdm-gcc/

## docker run
`docker-compose up` will build and run containers

`localhost:8080` - go-app  
`localhost:8081` - PHPMyAdmin
