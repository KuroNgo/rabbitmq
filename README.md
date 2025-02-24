## RabbitMQ
### Prerequisites
- Assumes RabbitMQ is installed and running on localhost on the standard port 5678
- In case you use a different host, port or credentials, connections settings would require adjusting

### Notion
- RabbitMQ is a message broker. It accepts and forwards messages
- RabbitMQ, and messaging in general, uses some jargon
* Producing means nothing more than sending. A program that sends messages is a producer
* A queue is them name for the post box in RabbitMQ. Although messages flow through RabbitMQ and your applications,
they can only be stored inside a queue. A queue is only bound by the host's memory & disk limits, it's essentially a large message buffer

Many producers can send messages that go to one queue, and many consumers can try to receive data from one queue
This is how we represent a queue
Consuming has a similar meaning to receiving. A consumer is a program that mostly waits to receive messages

**Note** that the producer, consumer, and broker do not have to reside on the same host; indeed in most applications they don't. An application can be both a producer and consumer, too

## New words
assume: cho rằng
producer: nhà sản xuất
jargon: từ ngữ khó hiểu
consumer: người tiêu dùng
reside: cư trú
indeed in: thật vậy