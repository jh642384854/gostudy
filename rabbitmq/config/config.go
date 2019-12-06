package config

const (
	MQ_URL                = "amqp://guest:guest@192.168.20.24:5672"
	FailToConnectMQ       = "Failed to connect to RabbitMQ"
	FailOpenChannel       = "Failed to open a channel"
	FailDeclareQueue      = "Failed to declare a queue"
	FailDeclareExchange   = "Failed to declare an exchange"
	FialDeclareAltemateExchange = "Failed to declare an Altemate exchange"
	FailPublishMessage    = "Failed to publish a message"
	FailSetQos            = "Failed to set QoS"
	FailRegisterConsumer  = "Failed to register a consumer"
	FailExchangeBindQueue = "Failed to bind a queue"
)
