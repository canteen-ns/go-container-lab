package main

import (
	"flag"
	"runtime"
	"fmt"
	"github.com/canteen_ns/go-container-lab/week01/pkg/crawler"
	"time"
)

func main(){
	workers := flag.Int("w",100,"worker count")
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	urls := []string{
		"https://www.baidu.com",
		"https://www.sina.com.cn",
		"https://www.taobao.com",
		"https://www.bilibili.com",
	}

	for len(urls) < 100 {
		urls = append(urls, urls...)
	}

	urls = urls[:100]
	start := time.Now()
	results := crawler.Crawl(urls, crawler.Options{Workers: *workers})

	var ok, fail int
	for r := range results {
		if r.Err != nil {
			fail++
			fmt.Printf("FAIL %s : %v\n", r.Url, r.Err)
		} else {
			ok++
			fmt.Printf("OK %d %s %d bytes\n", r.Code, r.Url, r.Size)
		}
	}
	fmt.Printf("Done: %d ok, %d fail, elapsed %v\n", ok, fail, time.Since(start))
}