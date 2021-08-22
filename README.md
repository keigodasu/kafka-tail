# kafka-tail
simple tail command to consume Kafka topic

## Usage

```
Usage of ./kafka-tail:
  -bootstrap-servers string (required)
        bootstrap servers
  -group-id string (option)
        consumer group id (if not specified, 'kafka-tail-<UUID>' is set as Consumer Group ID
  -topic string (required)
        topic (currently doesn't support regular expressions and multiple topics)
```

``` example
./kafka-tail -bootstrap-servers kafka:9092 -group-id testconsumer -topic test-topic
```

## Licence
MIT