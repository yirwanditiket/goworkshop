package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// toyolah has 30 available cars, can handle 5 concurrent request
var toyolah *CarDealer

func init() {
	toyolah = &CarDealer{
		MaxConcRequest: 5,
	}

	toyolah.AvailableCar.Store(30)
	toyolah.ConcRequest.Store(0)
	toyolah.SerialNum.Store(100)
}

func main() {
	orders := getOrders()

	st := time.Now()

	// fmt.Println("Sequential res err: ", processOrderSeq(orders))
	fmt.Println("Concurrent res err: ", processOrderConc(orders))

	fmt.Println("Elapsed ", time.Since(st))
}

func processOrderConc(orders []int) error {
	carCh := make(chan string)

	go func() {
		for _, orderId := range orders {
			go func(orderId int) {
				serial, err := toyolah.BuyCar(orderId)
				if err != nil {
					fmt.Println("Error when buying: ", err)
				}

				carCh <- serial

			}(orderId)
		}
	}()

	for serial := range carCh {
		fmt.Println("Bought a car! serial num: ", serial)
	}

	return nil
}

// ----------------------------------------------------
// ----- vvv-- do not modify the code below --vvv------
// ----------------------------------------------------

func processOrderSeq(orders []int) error {
	for _, orderId := range orders {
		serial, err := toyolah.BuyCar(orderId)
		if err != nil {
			return err
		}

		fmt.Println("Bought a car! serial num: ", serial)
	}

	return nil
}

func getOrders() []int {
	orders := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		orders[i] = i
	}
	return orders
}

type CarDealer struct {
	AvailableCar   atomic.Int64
	ConcRequest    atomic.Int64
	MaxConcRequest int64
	SerialNum      atomic.Int64
	muRequest      sync.Mutex
}

// accept orderid return the serial number of the car or error
// process takes 2 second to finish
// error may happen if: conc request greater than the MaxConcRequest
// or there is no available car anymore
func (c *CarDealer) BuyCar(orderId int) (string, error) {
	if err := c.canHandle(); err != nil {
		return "", err
	}
	defer c.ConcRequest.Add(-1)

	fmt.Printf("Processing order id %v\n", orderId)
	time.Sleep(2 * time.Second)

	return c.getSerialNum(), nil
}

func (c *CarDealer) canHandle() error {
	c.muRequest.Lock()
	defer c.muRequest.Unlock()

	if c.ConcRequest.Load() >= c.MaxConcRequest {
		return fmt.Errorf("TooMuchConcurrentRequest: can't handle more than %v at the same time", c.MaxConcRequest)
	}

	if c.AvailableCar.Load() <= 0 {
		return fmt.Errorf("No more car available")
	}

	c.ConcRequest.Add(1)
	c.AvailableCar.Add(-1)

	return nil
}

func (c *CarDealer) getSerialNum() string {
	return fmt.Sprintf("car-%04d", c.SerialNum.Add(1))
}
