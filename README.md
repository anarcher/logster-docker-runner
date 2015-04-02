logster-docker-runner
=====================

It is for running logster on docker container and to reading docker log files. 

fugu.yml

```
logster:
    image: anarcher/logster-docker-runner:0.0.1
    name: logster
    rm: true
    volume:
        - /var/lib/docker:/var/lib/docker:ro
    env:
        - CONTAINER_NAME=container_name
        - OUTPUT=stdout
        - PARSER=MetricLogster
        - LOGXI_FORMAT=text
        - LOGXI="*"
```

# State

- "WORK IN PROGRESS"

