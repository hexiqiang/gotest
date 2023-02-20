package main

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"title": "这是sayHello",
	})
}

func gets(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", "3242")
}

func posts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "提交了数据",
	})
}

func del(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "删除了数据",
	})
}

func failOnError1(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func main() {
	//r := gin.Default()
	//r.LoadHTMLGlob("templates/**/*")
	//r.GET("/hello", sayHello)
	//r.POST("/posts", posts)
	//
	////r.LoadHTMLFiles("./templates/book/index.tmpl")
	//r.GET("/posts", gets)
	//r.DELETE("/posts", del)
	//
	//r.Run(":80")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. 接下来，我们创建一个通道，大多数API都是用过该通道操作的。
	ch, err := conn.Channel()
	failOnError1(err, "Failed to open a channel")
	defer ch.Close()

	// 3. 声明消息要发送到的队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError1(err, "Failed to declare a queue")

	body := "Hello World!"
	// 4.将消息发布到声明的队列
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError1(err, "Failed to publish a message")
}
