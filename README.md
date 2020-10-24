# go-compose
Go-Compose is an easy-to-use CLI that allows the user to create a basic Docker Compose YAML file for Mongo/MySQL 
database servers.


## Description
Go-Compose has one sub command: `generate` (alias: `gen`). This command and its children allow the `--output` 
(or `-o`) flag to be present. If specified, when the docker-compose.yml file is generated, it will output to 
this directory; if it is not specified, it will default to the user's desktop (`~/Desktop` on UNIX-like systems).


## Examples


### Mongo
Generate a file to the desktop:
```shell script
> go-compose generate mongo -d DATABASE_NAME
```


Generate a file to the current directory:
```shell script
> go-compose generate mongo -d DATABASE_NAME -o $(pwd)
```


#### Flags


##### `--database` (shorthand: `-d`)
Specifies the name of the database to be created when the container starts


### MySQL
Generate a file to the desktop:
```shell script
> go-compose generate mysql -d DATABASE_NAME -P ROOT_PASSWORD
```

Generate a file to the current directory:
```shell script
> go-compose generate mysql -d DATABASE_NAME -P ROOT_PASSWORD -o $(pwd)
```


#### Flags


##### `--database` (shorthand: `-d`)
Specifies the name of the database to be created when the container starts


##### `--root-password` (shorthand: `-P`)
Specifies the password for the root user when the container starts