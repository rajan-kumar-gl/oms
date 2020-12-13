package repo

import (
	"sync"

	"github.com/oms/order-service/src/proto"
)

type implimantaion struct {
	UserMap map[int]map[int]int
	mutex   *sync.Mutex
}

func New() proto.ProductLocker {
	return &implimantaion{
		UserMap: make(map[int]map[int]int),
		mutex:   &sync.Mutex{},
	}
}

func (o *implimantaion) AllocateProductToUser(userID int, productID int, stock int) {
	if o.UserMap == nil {
		o.UserMap = make(map[int]map[int]int)
	}

	o.mutex.Lock()
	defer o.mutex.Unlock()

	userMap, found := o.UserMap[userID]
	if !found {
		productMap := make(map[int]int)
		productMap[productID] = stock
		o.UserMap[userID] = productMap
		return
	}

	productMap, found := userMap[productID]
	if !found {
		userMap[productID] = stock
		o.UserMap[userID] = userMap
		return
	}

	newStock := productMap + stock
	userMap[productID] = newStock
	o.UserMap[userID] = userMap
	return
}
