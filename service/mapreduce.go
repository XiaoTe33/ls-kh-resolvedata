package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"ls-kh-resolvedata/log"
	"ls-kh-resolvedata/model"
)

var myLog = log.Log
var mq = model.NewListMQ(redis.NewClient(&redis.Options{
	Addr: viper.GetString("Redis.Addr"),
}))

func Serve() {
	for {
		task, err := mq.GetTask("mapreduce")
		if err != nil {
			myLog.Error(err)
			return
		}
		ctx := model.NewMyCtx(task)
		go resolveTask(ctx)
	}
}

func resolveTask(ctx *model.MyContext) {
	go func() {
		select {
		case res := <-ctx.Result:
			fmt.Println(res)
			//todo
		}
	}()
}
