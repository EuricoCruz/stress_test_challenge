package tester

import (
	"io"
	"github.com/EuricoCruz/stress_test_challenge/internal/report"
	"net/http"
	"sync"
	"time"
)

type Service struct {
	client *http.Client
}

func NewService() *Service {
	return &Service{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *Service) Run(url string, totalRequests, concurrency int) report.Report {
	start := time.Now()
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)
	results := make(chan report.Result, totalRequests)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			reqStart := time.Now()
			resp, err := s.client.Get(url)
			lat := time.Since(reqStart)

			if err != nil {
				results <- report.Result{StatusCode: 0, Err: err.Error(), Duration: lat}
			} else {
				// drenar body para permitir reuse de conexões
				io.Copy(io.Discard, resp.Body)
				status := resp.StatusCode
				resp.Body.Close()
				results <- report.Result{StatusCode: status, Err: "", Duration: lat}
			}

			<-sem
		}()
	}

	wg.Wait()
	close(results)

	return report.NewReport(results, time.Since(start))
}
