# Banking Kafka

## How to run

1. Install kafka (macOS)
```
$ brew install kafka
```

2. Run docker compose
```
$ docker compose up -d
```

3. Test kafka
```
$ kafka-topics --bootstrap-server=localhost:9092 --list
```

4. Create Topic
```
$ kafka-topics --bootstrap-server=localhost:9092 --topic=songvut --create
```

## Producer and Consumer

### Consumer

- Consumer: Subscribe topic "songvut"
```
$ kafka-console-consumer --bootstrap-server=localhost:9092 --topic=songvut
```

- Consumer: Subscribe topic with group "notice"
```
$ kafka-console-consumer --bootstrap-server=localhost:9092 --topic=songvut --group=notice
```

- Consumer: Subscribe multi-topic "songvut" and "nakrong"
```
$ kafka-console-consumer --bootstrap-server=localhost:9092
 --include="songvut|nakrong" --group=not
```

### Producer

- Producer: Publish message "hello world"
```
$ kafka-console-producer --bootstrap-server=localhost:9092 --topic=songvut
> hello world
```