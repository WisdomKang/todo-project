package main

import (
	"fmt"
	"os"
	"strings"
	"todo-project/config"
)

func main() {
	// err := godotenv.Load("application.env")

	// if err != nil {
	// 	fmt.Printf("끼얏호 %v", err)
	// }

	config.LoadEnv()

	envVars := os.Environ()

	// 환경 변수 출력
	for _, envVar := range envVars {
		// KEY=VALUE 형식에서 KEY와 VALUE 분리
		keyValue := strings.SplitN(envVar, "=", 2)
		fmt.Printf("Key: %s, Value: %s\n", keyValue[0], keyValue[1])
	}
}
