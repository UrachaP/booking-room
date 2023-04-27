package config

import (
	"github.com/go-redis/redis/v9"
)

type Person struct {
	Name string `redis:"name"`
	Age  int    `redis:"age"`
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//pong, err := rdb.Ping(c.Request().Context()).Result()
	//if err != nil {
	//	log.Fatal(err)
	//	return c.String(http.StatusInternalServerError, err.Error())
	//}
	//fmt.Println(pong, err)
	//
	//err = rdb.Set(c.Request().Context(), "name", "test", time.Hour).Err()
	//if err != nil {
	//	log.Fatal(err)
	//	return c.String(http.StatusInternalServerError, err.Error())
	//}
	//
	//val, err := rdb.Get(c.Request().Context(), "name").Result()
	//if err != nil {
	//	log.Fatal(err)
	//	return c.String(http.StatusInternalServerError, err.Error())
	//}
	//fmt.Println("name", val)
	//
	////get expired time
	//t := rdb.GetEx(c.Request().Context(), "name", time.Minute)
	//fmt.Println(t)
	//
	////hash
	//_, err = rdb.Pipelined(c.Request().Context(), func(rdb redis.Pipeliner) error {
	//	rdb.HSet(c.Request().Context(), "person1", "name", "uracha")
	//	rdb.HSet(c.Request().Context(), "person1", "age", 25)
	//	return nil
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//var person1 Person
	//err = rdb.HGetAll(c.Request().Context(), "person1").Scan(&person1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("person1", person1)
	//
	//var person2 Person
	//// Scan a subset of the fields.
	//err = rdb.HMGet(c.Request().Context(), "person1", "age", "int").Scan(&person2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("person2", person2)

	//close redis
	//defer func(rdb *redis.Client) {
	//	err := rdb.Close()
	//	if err != nil {
	//
	//	}
	//}(rdb)

	return rdb
}
