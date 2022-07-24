# Go Simple and Minimal Todolist API Follow Clean Arch

Challenge I did for technical analysis of GL Tech team. Simple todolist in edutech field helps children and teacher to manage their homework or mission that assigned for children. There are many things that is currently in development. Contribution is always welcomed and very precious to me.

## How to run ?

Firstly, create a config.env file to hold the environment variables by:
```makefile
make 
```
Secondly, edit your DB_USER and DB_PASS for the username and password of the database config. Also install necessary go pakags by:

``` go
go mod tidy
```

Finally, run the application by:
```makefile
make run
```

## Api Routes
- [POST] /login (missing)
- [POST] /user/register
- [POST] /assign?userId=?&todoId=?
- [POST] /revoke?userId=?&todoId=?
- [GET] /user/:id
- [PATCH] /user/:id
- [DELETE] /user/:id
- [POST] /todo
- [GET] /todo/:id
- [PATCH] /todo/:todoid
- [DELETE] /todo/:todoid
- [GET] /todo/:todoid/user
- [POST] /user/:userId/todo/:todoid/done

## References
https://eminetto.medium.com/clean-architecture-using-golang-b63587aa5e3f
https://github.com/percybolmer/ddd-go/tree/clean-architecture
https://github.com/ruslantsyganok/clean_arcitecture_golang_example
https://github.com/vinigracindo/fiber-gorm-clean-architecture
https://github.com/gofiber/fiber
https://gorm.io