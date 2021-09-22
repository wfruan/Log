package main

import (
	"fmt"
	"github.com/guonaihong/gout"
	"golang.org/x/sync/singleflight"
	"sync"
)

func main() {
	var group singleflight.Group
	var wg sync.WaitGroup

	wg.Add(10)
	defer wg.Wait()
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()

			res, err, _ := group.Do("reqinfo", func() (interface{}, error) {
				s := ""
				err := gout.GET("https://http2.golang.org/reqinfo").Debug(true).BindBody(&s).Do()
				if err != nil {
					return nil, err
				}

				return s, nil
			})

			if err != nil {
				fmt.Printf("fail:%s\n", err)
			}

			fmt.Printf("id(%d) ------>%s\n", id, res.(string))
		}(i)
	}
}
