package cacheTest

import "fmt"
import "module/cache"

func  Test1()  {
	cacheStore := cache.Cache()
	if err := cacheStore.Add("testKey","testValue", 10000); err != nil {
		fmt.Println(err)
	}

	var myValue string
	if err := cacheStore.Get("testKey", &myValue); err != nil {
		fmt.Println(err)
	}
	fmt.Println(myValue)
}