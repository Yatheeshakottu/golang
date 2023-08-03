package main

import (
	"fmt"
	"math"
	"time"
)

type Config struct {
	time_window_sec rune
	capacity        int
}

type Transaction struct {
	amount    float64
	timestamp string
	location  string
}

var (
	config_store   = make(map[string]Config)
	requests_store = make(map[string]map[time.Time][]Transaction)
)

func is_allowed(key string) bool {

	current_time := time.Now()

	config := get_rate_limit_config(key)

	start_time := current_time.Add(-time.Second * time.Duration(config.time_window_sec))
	no_of_requests := get_current_window_requests(key, start_time)

	return no_of_requests <= config.capacity
}

func get_rate_limit_config(key string) Config {
	return config_store[key]
}

func get_current_window_requests(key string, start_time time.Time) int {
	ts_data := requests_store[key]
	if key == "" {
		return 0
	}

	total_requests := 0
	for ts, trs := range ts_data {
		if ts.Before(start_time) {
			total_requests += len(trs)
		} else {
			delete(ts_data, ts)
		}
	}
	return total_requests
}

func get_current_window(key string, start_time time.Time) []Transaction {
	ts_data := requests_store[key]
	if key == "" {
		return []Transaction{}
	}

	transactions := []Transaction{}
	for ts, trs := range ts_data {
		if ts.Before(start_time) {
			transactions = append(transactions, trs...)
		} else {
			delete(ts_data, ts)
		}
	}
	return transactions
}

func register_request(key string, ts time.Time, tr Transaction) {
	if _, ok := requests_store[key]; ok {
		requests_store[key][ts.Round(1*time.Second)] = append(requests_store[key][ts.Round(1*time.Second)], tr)
	} else {
		requests_store[key] = make(map[time.Time][]Transaction)
		requests_store[key][ts.Round(1*time.Second)] = make([]Transaction, 0)
		requests_store[key][ts.Round(1*time.Second)] = append(requests_store[key][ts.Round(1*time.Second)], tr)
	}
}
func delete_transaction(key string, tr Transaction) bool {
	if trans, ok := requests_store[key]; ok {
		for duration, trs := range trans {
			for i, t := range trs {
				if t.amount == tr.amount && t.timestamp == tr.timestamp && t.location == tr.location {
					requests_store[key][duration] = append(requests_store[key][duration][:i], requests_store[key][duration][i+1:]...)
					return true
				}

			}
		}

	}
	return false
}

func main() {
	key := "Yatheesha_Cache"
	time_stamp := time.Now()
	config_store[key] = Config{
		time_window_sec: 60,
		capacity:        5,
	}
	if is_allowed(key) {
		register_request(key, time_stamp, Transaction{
			amount:    100.10,
			timestamp: time_stamp.GoString(),
			location:  "Bangalore",
		})

		register_request(key, time_stamp, Transaction{
			amount:    500.10,
			timestamp: time_stamp.GoString(),
			location:  "Chennai",
		})

	} else {
		// 204 – if the transaction is older than 60 seconds
	}
	time.Sleep(3 * time.Second)
	trans := get_current_window(key, time.Now())

	fmt.Println(trans)
	sum, avg, max, min, count := 0.0, 0.0, math.SmallestNonzeroFloat64, math.MaxFloat64, len(trans)

	for _, tr := range trans {
		sum += tr.amount
		if max < tr.amount {
			max = tr.amount
		}
		if min > tr.amount {
			min = tr.amount
		}
	}
	avg = sum / float64(count)

	fmt.Println("Sum :", sum)
	fmt.Println("Avg :", avg)
	fmt.Println("Max :", max)
	fmt.Println("Min :", min)
	fmt.Println("Count :", count)

	delete_transaction(key, Transaction{
		amount:    100.10,
		timestamp: time_stamp.GoString(),
		location:  "Bangalore",
	})
	fmt.Println("After Deleting")

	trans = get_current_window(key, time.Now())

	fmt.Println(trans)
	sum, avg, max, min, count = 0.0, 0.0, math.SmallestNonzeroFloat64, math.MaxFloat64, len(trans)

	for _, tr := range trans {
		sum += tr.amount
		if max < tr.amount {
			max = tr.amount
		}
		if min > tr.amount {
			min = tr.amount
		}
	}
	avg = sum / float64(count)

	fmt.Println("Sum :", sum)
	fmt.Println("Avg :", avg)
	fmt.Println("Max :", max)
	fmt.Println("Min :", min)
	fmt.Println("Count :", count)
}
