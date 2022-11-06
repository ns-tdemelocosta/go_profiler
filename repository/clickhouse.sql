CREATE TABLE helloworld.my_first_table
(
    user_id UInt32,
    message String,
    timestamp DateTime,
    metric Float32
)
ENGINE = MergeTree()
PRIMARY KEY (user_id, timestamp);

-- Path: repository/clickhouse.sql

CREATE TABLE helloworld.processMessages(
    pid UInt32,
    cpu_usage Float32,
    memory_usage Float32,
    timestamp DateTime
    ctime DateTime
)
ENGINE = MergeTree()
PARTITION BY toYYYYMM(timestamp)
ORDER BY (pid, timestamp)
SAMPLE BY pid;


