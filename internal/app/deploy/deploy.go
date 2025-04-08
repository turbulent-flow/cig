package main

import (
	"cig/internal/pkg/deploy"
	"log"
)

func main() {
	receipt := deploy.Run()

	if receipt.Status == 1 {
		log.Println("部署成功！")
	} else {
		log.Println("部署失败！")
	}
}
