package learnSelect

import (
	"errors"
	"net/http"
	"time"
)
var ErrTimeOut = errors.New("Sever took long to respond")
var tenSecondTimeout = 10 * time.Second

func Racer(url1, url2 string) (winner string, error error){
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", ErrTimeOut
	}

}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}