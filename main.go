package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewDevelopment()
	logger.Info("Hello zap", zap.String("key", "value"), zap.Time("now", time.Now()))
	logger.Debug("msg", zap.String("key", "string"), zap.Ints("ints", []int{10, 20}))
	sugar := zap.NewExample().Sugar()
	sugar.Info("one", "two", "three")
	sugar.Infof("one: %s, %d", "two", 10)
	sugar.Infow("msg", "key", "value", "intArray", []int{10, 100}, "duration", time.Second*200)
}
