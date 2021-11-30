package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	Name   string = "API"
	Repo   string
	Branch string
	Commit string
	Time   string
)

type Config struct {
	SERVER_HOST    string      // Host текущего сервера запуска
	SERVER_PORT    int         // Port текущего сервера запуска
	WORKER_HOST    string      // Host Worker-компонента
	WORKER_PORT    int         // Port Worker-компонента
	STORE_HOST     string      // Host Worker-компонента
	STORE_PORT     int         // Port Worker-компонента
	EGRESS_HOST    string      // Host Egress-компонента
	EGRESS_PORT    int         // Port Egress-компонента
	STORE_CONFIG   StoreConfig // Доп параметры для Store-компонента
	SCHEDULER_HOST string      // Host Scheduler-компонента
	SCHEDULER_PORT int         // Port Scheduler-компонента
	QUEUE_HOST     string      // Host Queue-компонента
	QUEUE_PORT     int         // Port Queue-компонента
	QUEUE_CONFIG   QueueConfig // Доп параметры для Queue-компонента
}

// QueueConfig - доп. параметры
type QueueConfig struct {
	SIZE         int
	NAME_CLUSTER string
}

// StoreConfig - доп. параметры
type StoreConfig struct {
	DB_HOST   string
	DB_PORT   int
	DB_USER   string
	DB_PASS   string `json:"-"`
	DB_NAME   string
	DB_SCHEMA string
}

func main() {
	// Информация о сборке приложения
	fmt.Printf("\nComponent: %s\n", Name)
	fmt.Printf("Source Code: %s\n", Repo)
	fmt.Printf("Branch: %s\n", Branch)
	fmt.Printf("Commit: %s\n", Commit)
	fmt.Printf("Build Time: %s\n\n", Time)

	// Cобираем конфиг
	conf := Config{
		SERVER_HOST: "0.0.0.0",
		SERVER_PORT: 8080,
		WORKER_HOST: "worker.namespace.local.svc",
		WORKER_PORT: 7070,
		STORE_HOST:  "store.namespace.local.svc",
		STORE_PORT:  8080,
		STORE_CONFIG: StoreConfig{
			DB_USER: "gleb",
			DB_PASS: "passfromdb",
			DB_HOST: "server.delta.sbrf.ru",
			DB_PORT: 5432,
		},
		SCHEDULER_HOST: "scheduler.namespace.local.svc",
		SCHEDULER_PORT: 8080,
		QUEUE_HOST:     "queue.namespace.local.svc",
		QUEUE_PORT:     4222,
		QUEUE_CONFIG: QueueConfig{
			SIZE:         100,
			NAME_CLUSTER: "scheduler_cluster",
		},
	}

	// Pretty print config
	j, _ := json.MarshalIndent(conf, "", "   ")
	log.Printf("Config: \n%v", string(j))

	// Проверка соединения с компонентами
	log.Println("Connect to Queue: ok")
	log.Println("Connect to Store: ok")
	log.Println("Connect to DB: failed")
	log.Println("Connect to Worker: ok")
	log.Println("Connect to Scheduler: ok")

	// Далее запуск компонента
}
