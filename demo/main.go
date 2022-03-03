package main

import (
	"demo/global"
	http2 "demo/http"
	"github.com/stevenyao001/edgeCommon"
	"github.com/stevenyao001/edgeCommon/http"
	"github.com/stevenyao001/edgeCommon/mqtt"
	"github.com/stevenyao001/edgeCommon/pgsql"
	"github.com/stevenyao001/edgeCommon/tdengine"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	//core.Run()
	edge := edgeCommon.New()

	filePath, _ := os.Getwd()
	edge.RegisterConfig(filePath+"/conf/local.yaml", &global.Conf)

	edge.RegisterLogger(global.Conf.Log.MainPath)

	pgConf := make([]pgsql.Conf, 0)
	for _, conf := range global.Conf.Postgres {
		pgConf = append(pgConf, pgsql.Conf{
			InsName:      conf.InsName,
			Addr:         conf.Addr,
			Port:         conf.Port,
			Username:     conf.Username,
			Password:     conf.Password,
			Db:           conf.DbName,
			MaxIdleConns: conf.MaxIdleConns,
			MaxIdleTime:  conf.MaxIdleTime,
			MaxLifeTime:  conf.MaxLifetime,
			MaxOpenConns: conf.MaxOpenConns,
		})
	}
	edge.RegisterPgsql(pgConf)

	mqttConfs := make([]mqtt.Conf, 0)
	for k := range global.Conf.Mqtt {
		mqttConfs = append(mqttConfs, mqtt.Conf{
			InsName:  global.Conf.Mqtt[k].InsName,
			ClientId: global.Conf.Mqtt[k].ClientId,
			Username: global.Conf.Mqtt[k].Username,
			Password: global.Conf.Mqtt[k].Password,
			Addr:     global.Conf.Mqtt[k].Addr,
			Port:     global.Conf.Mqtt[k].Port,
		})
	}
	//edge.RegisterMqtt(mqttConfs, mqtt2.Subscribes)

	tdConfs := make([]tdengine.Conf, 0)
	for k := range global.Conf.Tdengine {
		tdConfs = append(tdConfs, tdengine.Conf{
			InsName:      global.Conf.Tdengine[k].InsName,
			Driver:       global.Conf.Tdengine[k].Driver,
			Network:      global.Conf.Tdengine[k].Network,
			Addr:         global.Conf.Tdengine[k].Fqdn,
			Port:         global.Conf.Tdengine[k].Port,
			Username:     global.Conf.Tdengine[k].Username,
			Password:     global.Conf.Tdengine[k].Password,
			Db:           global.Conf.Tdengine[k].DbName,
			MaxIdleConns: global.Conf.Tdengine[k].MaxIdleConns,
			MaxIdleTime:  global.Conf.Tdengine[k].MaxIdleTime,
			MaxLifeTime:  global.Conf.Tdengine[k].MaxLifeTime,
			MaxOpenConns: global.Conf.Tdengine[k].MaxOpenConns,
		})
	}
	//edge.RegisterTdEngine(tdConfs)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go edge.RegisterHttp(http.Conf{
		Addr:            global.Conf.App.ServerAddr,
		ShutdownTimeout: time.Second * 10,
		Router:          http2.RegisterRouter,
		Wg:              wg,
	})

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	_ = <-quit

	wg.Done()
	time.Sleep(time.Second * 10)
}
