package mqtt

import (
	"demo/mqtt/collector_data_report"
	"github.com/stevenyao001/edgeCommon/mqtt"
)

var Subscribes = map[string][]mqtt.SubscribeOpts{
	"rootcloud": {
		//新采集数据
		{
			Topic:    "test/collector/data/report",
			Qos:      0,
			Callback: collector_data_report.CollectorDataReport,
		},
	},
}
