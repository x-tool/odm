# xodm - sql & nosql odm

xodm is a pure Go odm for SQL or NoSQL. xodm is different from other orm or odm because it make for compatible SQL or NoSQL to simplify develop.

Five Seconds Learn
```go
client := NewClient("postgresql", connectionConf)
db := client.Database("x")

db.SyncCols(new(myDoc))
col := db.GetCol(new(myDoc))

testInsert := new(myDoc)
testInsert.Name = "haha,I get"
_, err := col.Insert(testInsert)
```

## Features

* JSON support, You can get field in struct from database simple.
* Usual Mode.
* I will add more Features :)

## Support database

xodm only support postgresql now, it may support orther sql future.

* postgresql -- github.com/jackc/pgx

## Todo

* [ ] fix bugs
* [ ] add arangodb

## Performance

Use reflect package. I don't think performance beautiful. but it write for develop simple.
