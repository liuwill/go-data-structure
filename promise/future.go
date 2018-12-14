package promise

type Resolve func(interface{}) (interface{}, error)

type Reject func(error)

type Ticket func(resolve Resolve, reject Reject)

type Future struct {
	nextTicket Ticket
	reject     Reject
	response   interface{}
}

func NewFuture(init Ticket) *Future {
	return &Future{
		nextTicket: init,
	}
}

func (future *Future) then(callback Resolve) *Future {
	future.nextTicket(func(response interface{}) (interface{}, error) {
		res, err := callback(response)
		future.response = res
		return nil, err
	}, future.reject)

	return NewFuture(func(resolve Resolve, reject Reject) {
		resolve(future.response)
	})
}

func handlerFuture(source int) *Future {
	return NewFuture(func(resolve Resolve, reject Reject) {
		resolve(source)
	}).then(func(response interface{}) (interface{}, error) {
		return response, nil
	})
}
