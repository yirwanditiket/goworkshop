package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	st := time.Now()

	totalCar := 15
	cf := &CarFactory{PaintWorkerNum: 4}

	run(cf, totalCar)

	fmt.Println("All done, elapsed", time.Since(st))
	time.Sleep(2 * time.Second)
}

func run(cf *CarFactory, n int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chassisCh := cf.MakeChassis(ctx, n)
	assemblyCh := cf.Assembly(ctx, chassisCh)
	paintCh := cf.Paint(ctx, assemblyCh, "White")
	inspectErrCh := cf.Inspect(ctx, paintCh)

	for err := range inspectErrCh {
		if err != nil {
			fmt.Println("ERR happen : ", err)
		}
	}
}

// flow of process:
// 1: create chassis
// 2: assembly
// 3: paint
// 4: inspection
type CarFactory struct {
	PaintWorkerNum int
}

func (cf *CarFactory) MakeChassis(ctx context.Context, n int) <-chan Car {
	outCh := make(chan Car)
	go func() {
		defer close(outCh)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case outCh <- MakeChassis(i):
			}
		}
	}()
	return outCh
}

func (cf *CarFactory) Assembly(ctx context.Context, inCh <-chan Car) <-chan Car {
	outCh := make(chan Car)
	go func() {
		defer close(outCh)
		for c := range inCh {
			c.Assembly()
			select {
			case <-ctx.Done():
				return
			case outCh <- c:
			}
		}
	}()
	return outCh
}

func (cf *CarFactory) Paint(ctx context.Context, inCh <-chan Car, color string) <-chan Car {
	outCh := make(chan Car)
	go func() {
		defer close(outCh)
		for c := range inCh {
			c.Paint(color)
			select {
			case <-ctx.Done():
				return
			case outCh <- c:
			}
		}
	}()
	return outCh
}

// TODO: Create the implementation
func (cf *CarFactory) PaintFifo(ctx context.Context, inCh <-chan Car, color string, workerCount int) <-chan Car {
	return nil
}

func (cf *CarFactory) Inspect(ctx context.Context, inCh <-chan Car) <-chan error {
	outCh := make(chan error)
	go func() {
		defer close(outCh)
		for c := range inCh {
			select {
			case <-ctx.Done():
				return
			case outCh <- c.Inspect():
			}
		}
	}()
	return outCh
}

// car takes 1 second to assembly
// 2 seconds to paint
// 500ms to inspect
// total 3.5sec for all process
type Car struct {
	carId     int
	assembled bool
	color     string
}

func MakeChassis(id int) Car {
	return Car{
		carId: id,
	}
}

func (c *Car) Assembly() {
	time.Sleep(1 * time.Second)
	c.assembled = true
}

func (c *Car) Paint(color string) {
	if c.carId == 9 || c.carId == 11 {
		// forgot to paint car id 9
		// and 11
		return
	}

	time.Sleep(2 * time.Second)
	c.color = color
}

func (c *Car) Inspect() error {
	time.Sleep(500 * time.Millisecond)

	if !c.assembled {
		return fmt.Errorf("Car id %v is not assembled!", c.carId)
	}

	if c.color == "" {
		return fmt.Errorf("Car id %v is not painted!", c.carId)
	}

	return nil
}

func seq(cf *CarFactory, n int) {
	for i := 0; i < n; i++ {
		car := MakeChassis(i)
		car.Assembly()
		car.Paint("white")
		if err := car.Inspect(); err != nil {
			fmt.Println(err)
		}
	}
}
