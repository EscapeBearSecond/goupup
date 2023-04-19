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

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()
	doCall(ctx)
}

func doCall(ctx context.Context) {
	respChan := make(chan *resData)
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	clien := http.Client{
		Transport: &transport,
	}
	req, err := http.NewRequest("GET", "http://localhost:9099", nil)
	if err != nil {
		fmt.Printf("new request failed, err: %v", err)
		return
	}
	req = req.WithContext(ctx)
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func() {
		resp, err := clien.Do(req)
		fmt.Printf("client.do resp: %v err = %v\n", resp, err)
		rd := &resData{
			resp: resp,
			err:  err,
		}
		fmt.Println(rd)
		respChan <- rd
		wg.Done()
	}()
	select {
	case result := <-respChan:
		fmt.Printf("call server api success!")
		defer result.resp.Body.Close()
		if result.err != nil {
			fmt.Printf("call server api faield! err = %v\n", result.err)
			return
		}
		data, err := io.ReadAll(result.resp.Body)
		if err != nil {
			fmt.Printf("read response body failed! err = %v\n", err)
		}
		fmt.Printf("resp: %v\n", string(data))
	case <-ctx.Done():
		fmt.Println("call api timeout")
	}
}
