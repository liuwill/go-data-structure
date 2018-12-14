package promise

import (
	"testing"
)

func Test_Future(t *testing.T) {
	source := 1
	future := handler(source)

	future.then(func(response interface{}) interface{} {
		if response != source {
			t.Error("Test_Future Fail", response)
		}
		return nil
	})

	t.Log("Test_Future Success")
}
