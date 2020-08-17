package machine

import "sync"

type outlet struct {
	id			int
	isBusy		bool
	outletMutex 		sync.RWMutex
}

func newOutlet(id int) *outlet{
	return &outlet{
		id:          id,
		isBusy:      false,
		outletMutex: sync.RWMutex{},
	}
}

//This function returns false if outlet is busy
//and returns true if outlet is free, it also sets the outlet isBusy var to true in that case to ensure atomic operation
func (o *outlet)getOutlet() bool {
	o.outletMutex.Lock()
	defer o.outletMutex.Unlock()

	if o.isBusy {
		return false
	}
	o.isBusy = true
	return true
}

//This function changes to the outlet's isBusy status
func (o *outlet)changeOutletStatus(isBusy bool) {
	o.outletMutex.Lock()
	defer o.outletMutex.Unlock()

	o.isBusy = isBusy
}








