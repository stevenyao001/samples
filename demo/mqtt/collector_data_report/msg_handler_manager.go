package collector_data_report

import (
	mqtt2 "github.com/stevenyao001/edgeCommon/mqtt"
	"sync"
)

/**
消息管理器管理者
*/
var msgHandlerM *msgHandlerManager

type msgHandlerManager struct {
	mutex    sync.RWMutex
	handlers map[string]msgHandler
}

func init() {
	msgHandlerM = &msgHandlerManager{
		mutex:    sync.RWMutex{},
		handlers: map[string]msgHandler{},
	}
}

//消息入队
func (mhm *msgHandlerManager) msgPutQueue(msg mqtt2.Msg) {
	if msg.DeviceId == "" {
		return
	}

	mhm.mutex.RLock()
	defer mhm.mutex.RUnlock()

	msgH, exists := mhm.handlers[msg.DeviceId]
	if !exists {
		return
	}

	msgH.msgQueue <- msg
}

//新增消息管理器
func (mhm *msgHandlerManager) newMsgHandler(deviceId string) {
	mhm.mutex.Lock()
	defer mhm.mutex.Unlock()

	if _, exists := mhm.handlers[deviceId]; exists {
		return
	}

	mh := msgHandler{
		deviceId: deviceId,
		msgQueue: make(chan mqtt2.Msg, 1000),
		close:    make(chan struct{}, 1),
	}

	go mh.msgOutQueue()

	mhm.handlers[deviceId] = mh

}

//删除消息管理器
func (mhm *msgHandlerManager) delMsgHandler(deviceId string) {
	mhm.mutex.Lock()
	defer mhm.mutex.Unlock()

	mh, exists := mhm.handlers[deviceId]

	if !exists {
		return
	}

	mh.close <- struct{}{}

	delete(mhm.handlers, deviceId)
}
