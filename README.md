# UserApi

Goa crud user api

### Info
App was built on goa framework + gorm,
server was generated from design.go

### Installation

#### For starting/stopping app use:
```shell
$ make start-cli
$ make start
$ male stop
```

##### ! Don't run it, only after changes in design.go you must run commands below:
```shell
$ goa gen UserApi
$ goa example UserApi
```

Goa documentation: https://goa.design/