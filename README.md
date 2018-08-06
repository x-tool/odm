# xodm - sql & nosql odm

xodm is a pure Go odm for SQL or NoSQL. xodm is different from other orm or odm because it make for compatible SQL or NoSQL to simplify develop.

Five Seconds Learn
```go
client := NewClient("postgresql", connectionConf)
db := client.Database()
// sync cols to database
db.RegisterCols(new(myDoc))
db.SyncCols()

testInsert := new(myDoc)
testInsert.Name = "haha,I get"
_, err := db.Insert(testInsert)
```
## Percentage of progress
45%
## Features

* JSON support, You can get field in struct from database simple.
* Usual Mode.
* I will add more Features :)

## Support database

xodm only support postgresql now, it may support orther sql future.

* postgresql -- github.com/jackc/pgx

## Todo

* [ ] fix bugs

## Performance

Use reflect package. I don't think performance beautiful. but it write for develop simple.
