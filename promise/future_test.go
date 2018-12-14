package promise

import (
	"testing"
)

func Test_Future(t *testing.T) {
	source := 1
	future := handlerFuture(source)
	ch := make(chan interface{})

	go future.then(func(response interface{}) (interface{}, error) {

		ch <- response
		return nil, nil
	})

	target := <-ch
	if target != source {
		t.Error("Test_Future Fail", target)
	}

	t.Log("Test_Future Success")
}
