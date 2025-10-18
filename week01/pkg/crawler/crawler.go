package crawler

import (
	"net/http"
	"sync"
	"time"
	"io"
)

type Result struct {
	Url string
	Code int
	Size int64
	Err error
}

type Options struct {
	Workers int
	Client *http.Client 
}

func Crawl(urls []string, opts Options) <-chan Result {
	if opts.Workers < 0 {
		opts.Workers = 100
	}

	if opts.Client == nil {
		opts.Client = &http.Client{Timeout: 10 * time.Second}
	}

	out := make(chan Result, len(urls))
	tokens := make(chan struct{}, opts.Workers)

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _,u := range urls {
		go func(url string){
			defer wg.Done()
			tokens <- struct{}{}
			out <- doFetch(url, opts.Client)
			<- tokens
		}(u)
	}

	go func(){
		wg.Wait()
		close(out)
	}()

	return out
}

func doFetch(url string, cli *http.Client) Result {
	resp, err := cli.Get(url)
	if err != nil {
		return Result{Url: url, Err: err}
	}
	defer resp.Body.Close()

	size, _ := io.Copy(io.Discard, resp.Body) 
	return Result{Url: url, Code: resp.StatusCode, Size: size}
}
