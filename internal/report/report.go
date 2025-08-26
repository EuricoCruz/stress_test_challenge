package report

import (
	"fmt"
	"sort"
	"time"
)

type Result struct {
	StatusCode int
	Err        string
	Duration   time.Duration
}

type Report struct {
	Duration   time.Duration
	Total      int
	Success200 int
	Statuses   map[int]int    // 404,500,...
	Errors     map[string]int // mensagens de erro de rede/timeouts
	MinLatency time.Duration
	MaxLatency time.Duration
	AvgLatency time.Duration
}

func NewReport(results <-chan Result, duration time.Duration) Report {
	total := 0
	success := 0
	statuses := make(map[int]int)
	errors := make(map[string]int)

	var sum time.Duration
	var min time.Duration
	var max time.Duration
	first := true

	for r := range results {
		total++
		if r.Err != "" {
			errors[r.Err]++
		} else {
			if r.StatusCode == 200 {
				success++
			} else {
				statuses[r.StatusCode]++
			}
		}

		sum += r.Duration
		if first || r.Duration < min {
			min = r.Duration
		}
		if first || r.Duration > max {
			max = r.Duration
		}
		first = false
	}

	avg := time.Duration(0)
	if total > 0 {
		avg = sum / time.Duration(total)
	}

	return Report{
		Duration:   duration,
		Total:      total,
		Success200: success,
		Statuses:   statuses,
		Errors:     errors,
		MinLatency: min,
		MaxLatency: max,
		AvgLatency: avg,
	}
}

func (r Report) String() string {
	out := "===== Relatório =====\n"
	out += fmt.Sprintf("Tempo total: %s\n", r.Duration)
	out += fmt.Sprintf("Requests totais: %d\n", r.Total)
	out += fmt.Sprintf("Requests com status 200: %d\n", r.Success200)
	out += fmt.Sprintf("Latência (min/avg/max): %s / %s / %s\n", r.MinLatency, r.AvgLatency, r.MaxLatency)

	out += "\nOutros códigos de status:\n"
	if len(r.Statuses) == 0 {
		out += "  nenhum\n"
	} else {
		keys := make([]int, 0, len(r.Statuses))
		for k := range r.Statuses {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			out += fmt.Sprintf("  %d -> %d\n", k, r.Statuses[k])
		}
	}

	out += "\nErros de rede/timeouts:\n"
	if len(r.Errors) == 0 {
		out += "  nenhum\n"
	} else {
		keys := make([]string, 0, len(r.Errors))
		for k := range r.Errors {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			out += fmt.Sprintf("  %s -> %d\n", k, r.Errors[k])
		}
	}

	return out
}
