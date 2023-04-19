package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type resData struct {
	resp *http.Response
	err  error
}

func doCall(ctx context.Context) {
	transport := http.Transport{DisableKeepAlives: true} // 请求不频繁使用短链接
	client := http.Client{
		Transport: &transport,
	}
	dataChan := make(chan *resData, 1)
	var wg sync.WaitGroup
	defer wg.Wait()
	var i int32 = 0
	for {
		i++
		ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*10)
		req, err := http.NewRequest("GET", "http://localhost:9090", nil)
		if err != nil {
			fmt.Printf("%d new request failed, err = %v\n", i, err)
			continue
		}
		req = req.WithContext(ctx) // 使用带超时的context创建一个新的client request
		wg.Add(1)
		go func() {
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("%d clinet.do failed err: %v\n", i, err)
				return
			}
			aData := &resData{
				resp: resp,
				err:  err,
			}
			dataChan <- aData
			defer wg.Done()
		}()

		select {
		case result := <-dataChan:
			fmt.Printf("call api success\n")
			if result.err != nil {
				fmt.Printf("api response have error, err: %v\n", result.err)
			}
			fmt.Printf("response success, response = %v\n", result.resp)
		case <-ctx.Done():
			fmt.Println("call api timeout")
		}
		time.Sleep(time.Second)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10) // 设置10s过时的上下文对象
	defer cancel()
	doCall(ctx)
}
