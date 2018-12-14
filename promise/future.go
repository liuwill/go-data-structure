package promise

type Resolve func(interface{}) interface{}

type Reject func(interface{})

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
	future.nextTicket(func(response interface{}) interface{} {
		future.response = callback(response)
		return nil
	}, future.reject)

	return NewFuture(func(resolve Resolve, reject Reject) {
		resolve(future.response)
	})
}

func handler(source int) *Future {
	return NewFuture(func(resolve Resolve, reject Reject) {
		resolve(source)
	}).then(func(response interface{}) interface{} {
		return response
	})
}
