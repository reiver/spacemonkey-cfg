package reg

import (
	"github.com/reiver/spacemonkey-cfg/lib/hdl"

	"sync"
)

type Registrar struct {
	values map[string]hdl.Handler
	mutex sync.Mutex
}

func (receiver *Registrar) Register(handler hdl.Handler, name string) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil == handler {
		return errNilHandler
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.values {
		receiver.values = map[string]hdl.Handler{}
	}

	_, found := receiver.values[name]
	if found {
		return errFound
	}

	receiver.values[name] = handler

	return nil
}

func (receiver *Registrar) Get(name string) (string, bool, error) {
	if nil == receiver {
		return "", false, errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	values := receiver.values
	if nil == values {
		return "", false, errNotFound
	}

	handler, found := values[name]
	if !found {
		return "", false, errNotFound
	}
	if nil == handler {
		return "", false, errNilHandler
	}

	{
		value, err := handler.Value()
		if nil != err {
			return "", false, err
		}

		return value, true, nil
	}
}
