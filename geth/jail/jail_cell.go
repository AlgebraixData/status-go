package jail

import (
	"sync"
	"time"

	"fknsrs.biz/p/ottoext/fetch"
	"fknsrs.biz/p/ottoext/loop"
	"fknsrs.biz/p/ottoext/timers"
	"github.com/eapache/go-resiliency/semaphore"
	"github.com/robertkrimen/otto"
)

const (
	// JailCellRequestTimeout seconds before jailed request times out.
	JailCellRequestTimeout = 60
)

// JailCell represents single jail cell, which is basically a JavaScript VM.
type JailCell struct {
	sync.Mutex

	id string
	vm *otto.Otto
	lo *loop.Loop

	// FIXME(tiabc): It's never used. Is it a mistake?
	sem *semaphore.Semaphore
}

// newJailCell encapsulates what we need to create a new jailCell from the
// provided vm and eventloop instance.
func newJailCell(id string, vm *otto.Otto, lo *loop.Loop) (*JailCell, error) {
	// Register fetch provider from ottoext.
	if err := fetch.Define(vm, lo); err != nil {
		return nil, err
	}

	// Register event loop for timers.
	if err := timers.Define(vm, lo); err != nil {
		return nil, err
	}

	return &JailCell{
		id:  id,
		vm:  vm,
		lo:  lo,
		sem: semaphore.New(1, JailCellRequestTimeout*time.Second),
	}, nil
}

// Fetch attempts to call the underline Fetch API added through the
// ottoext package.
func (cell *JailCell) Fetch(url string, callback func(otto.Value)) (otto.Value, error) {
	cell.Lock()
	if err := cell.vm.Set("__captureFetch", callback); err != nil {
		defer cell.Unlock()

		return otto.UndefinedValue(), err
	}

	cell.Unlock()

	return cell.Exec(`fetch("` + url + `").then(function(response){
			__captureFetch({
				"url": response.url,
				"type": response.type,
				"body": response.text(),
				"status": response.status,
				"headers": response.headers,
			});
		});
	`)
}

// Exec evaluates the giving js string on the associated vm loop returning
// an error.
func (cell *JailCell) Exec(val string) (otto.Value, error) {
	cell.Lock()
	defer cell.Unlock()

	res, err := cell.vm.Run(val)
	if err != nil {
		return res, err
	}

	return res, cell.lo.Run()
}

// Run evaluates the giving js string on the associated vm llop.
func (cell *JailCell) Run(val string) (otto.Value, error) {
	cell.Lock()
	defer cell.Unlock()

	return cell.vm.Run(val)
}

// Loop returns the ottoext.Loop instance which provides underline timeout/setInternval
// event runtime for the Jail vm.
func (cell *JailCell) Loop() *loop.Loop {
	cell.Lock()
	defer cell.Unlock()

	return cell.lo
}

// VM returns the associated otto.Vm connect to the giving cell.
func (cell *JailCell) VM() *otto.Otto {
	cell.Lock()
	defer cell.Unlock()

	return cell.vm
}
