package main

import "github.com/nucktwillieren/project-d/xlimit-grpc/internal"

func main() {
	layer := internal.NewXlimitRedisLayerFromEnv()
	internal.NewXLimitService(layer)
}
