# Dockerize Demo - Wait for containers

Simple Docker Compose setup: front-end [API](https://github.com/mefellows/dockerize-compose-demo/blob/master/main.go) that takes 5 seconds to start up, and a [test](https://github.com/mefellows/dockerize-compose-demo/blob/master/test/web_test.go) harness that hits the API 100 times, failing if it receives any non-20x response codes.

## Running

### Failing Example

```
git co fail
docker-compose up
```

The test should fail with something like:

> test_1 | 	web_test.go:46: Error not expected: Get http://api/: dial tcp 172.17.0.3:80: getsockopt: connection refused

This is because Compose knows the container has started, but not the service within it.

### Passing Example

Using [Dockerize](https://github.com/jwilder/dockerize), we can ensure the tests don't run until the dependent API is up and running:

```
git co dockerize
docker-compose up
```

## Background

It is common when using tools like [Docker Compose](https://docs.docker.com/compose/) to depend on services in other linked containers, however oftentimes relying on [links](https://docs.docker.com/compose/compose-file/#links) is not enough - whilst the container itself may have _started_, the _service(s)_ within it may not yet be ready - resulting in shell script hacks to work around race conditions.

This [PR](https://github.com/jwilder/dockerize/pull/23/) gives `dockerize` the ability to wait for services on a specified protocol (`tcp`, `tcp4`, `tcp6`, `http`, and `https`) before starting the main application:

```
dockerize -wait tcp://web:80 -wait http://web:80
```

I've found this to be particularly useful when using Docker Compose as a test harness, where one of the containers needs to test another. Instead of `netcat`ing my way around the problem, I can just wrap the command using `dockerize`.



## References / Related reading

* https://github.com/docker/compose/issues/374#issuecomment-126312313
* https://github.com/docker/compose/issues/235
* https://github.com/jwilder/dockerize
