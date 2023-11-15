package rabbitmq

import (
	"encoding/json"
	"testing"
)

type TempVo struct {
	Name    string `json:"name"`
	OrderNo string `json:"order_no"`
}

func TestSendMq(t *testing.T) {
	bodyObj := TempVo{
		Name:    "hwq",
		OrderNo: "222323",
	}
	body, _ := json.Marshal(&bodyObj)
	rabbitClient := NewRabbitMQ("local.event", "new_send", "test")
	defer rabbitClient.Close()
	msg := Message{
		Body: string(body),
	}
	rabbitClient.SendMessage(msg)
}
