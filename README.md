# Loggers in Go - Test

## Zerolog
We add a programm that prints log messages of every allowed level every second.
The global setting for the log level is "debug" on startup.
This can be changed via the API, e.g.,
```bash
curl localhost:3000/loglevel/trace
```
The endpoint can be one of the well-known log levels, i.e.,
* trace
* debug
* info
* warn
* error

*fatal* and *panic* are not included, since using them while still running the code would make no sense.

## Benchmark Test of Logrus and Zerolog
100 logs per log level.
Run
```bash
go test ./... -bench=.
```