package crawler

import (
	"testing"
)

func BenchmarkCrawl10(b *testing.B)   { benchN(b, 10) }
func BenchmarkCrawl50(b *testing.B)   { benchN(b, 50) }
func BenchmarkCrawl100(b *testing.B)  { benchN(b, 100) }

func benchN(b *testing.B, n int) {
	urls := make([]string, n)
	for i := 0; i < n; i++ {
		urls[i] = "https://goproxy.cn" // 稳定国内站点
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for r := range Crawl(urls, Options{Workers: n}) {
			_ = r.Code // 防止优化掉
		}
	}
}