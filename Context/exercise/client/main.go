package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type resData struct {
	resp *http.Response
	err  error
}

func doCall(ctx context.Context) {
	// 请求频繁可定义全局的client对象，并启用长链接
	// 请求不频繁使用短链接
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}
	respChan := make(chan *resData, 1)
	req, err := http.NewRequest("GET", "http://localhost:9090", nil)
	if err != nil {
		fmt.Println("new request failed, err =", err)
		return
	}
	req = req.WithContext(ctx) // 使用带超时的ctx创建一个新的client request
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client.do resp: %v err = %v\n", resp, err)
		rd := &resData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
		wg.Done()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed, err: %v\n", result.err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := io.ReadAll(result.resp.Body)
		fmt.Printf("resp: %v\n", string(data))
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()
	doCall(ctx)
}
