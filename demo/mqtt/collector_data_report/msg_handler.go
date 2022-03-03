package collector_data_report

import (
	"demo/model/tdengine"
	"github.com/stevenyao001/edgeCommon/logger"
	mqtt2 "github.com/stevenyao001/edgeCommon/mqtt"
)

/**
消息处理器
*/
type msgHandler struct {
	deviceId string
	msgQueue chan mqtt2.Msg
	close    chan struct{}
	count    int64
}

//消息出队
func (msgH *msgHandler) msgOutQueue() {
	//退出后删除程序标识
	defer func() {
		if r := recover(); r != nil {
			logger.ErrorLog("msgHandler-msgOutQueue", "异常退出", msgH.deviceId, r)
		}
		go msgH.msgOutQueue()
	}()

	for {
		select {
		case msg, ok := <-msgH.msgQueue:
			if !ok {
				return
			}

			//todo do some thing

			if msg.Content == nil {
				continue
			}

			uid, ok := msg.Content["uid"]
			if !ok {
				continue
			}

			name, ok := msg.Content["name"]
			if !ok {
				continue
			}

			_, _ = tdengine.NewUserTd().Insert(uid.(int), name.(string))
		case <-msgH.close:
			return
		}
	}
}
