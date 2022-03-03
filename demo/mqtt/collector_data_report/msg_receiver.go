package collector_data_report

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stevenyao001/edgeCommon/logger"
	mqtt2 "github.com/stevenyao001/edgeCommon/mqtt"
)

//消息接收者
var CollectorDataReport = func(client mqtt.Client, msg mqtt.Message) {
	msgEntity := mqtt2.MsgPool.Get().(mqtt2.Msg)
	defer mqtt2.MsgPool.Put(msgEntity)

	err := json.Unmarshal(msg.Payload(), &msgEntity)
	if err != nil {
		logger.ErrorLog("MsgReceiver-ReceiveMsg", "消息解析失败", "", err)
		return
	}
	if msgEntity.DeviceId == "" {
		logger.ErrorLog("MsgReceiver-ReceiveMsg", "设备id不能为空", "", err)
		return
	}

	if msgEntity.Cmd == mqtt2.CollectDeviceRegister {
		msgHandlerM.newMsgHandler(msgEntity.DeviceId)
		return
	}
	if msgEntity.Cmd == mqtt2.CollectDeviceDel {
		msgHandlerM.delMsgHandler(msgEntity.DeviceId)
		return
	}

	go msgHandlerM.msgPutQueue(msgEntity)
}
