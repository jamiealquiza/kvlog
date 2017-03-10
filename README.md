# kvlog
Cripplingly simple kv logging. It's not the _fastest_ (good for > 600k logs/s. on modern laptop hardware), but has essentially no features/options. Which is a fucking awesome antithesis to bloated libraries.

```
% go test -bench=Log -benchmem | grep -v ts
  500000              2849 ns/op             856 B/op         14 allocs/op
 1000000              1587 ns/op              80 B/op          4 allocs/op
 1000000              1578 ns/op              80 B/op          4 allocs/op
PASS
ok      github.com/jamiealquiza/kvlog   4.661s
```

Kvlog writes [logfmt](https://brandur.org/logfmt)'ish key/value entries with 3 pre-defined top-level fields:

- `ts`: RFC 3339 timestamp
- `type`: a log type identifier
- `msg`: a log info or error string, or, a JSON string using a map input


### Logging with arbitrary map input
Input:
```go
kvlog.Log("web-request", map[string]string{"method": "GET", "page": "about"})
```
Output:
```
ts='2017-03-08T11:48:54-0700' type='web-request' msg='{"method":"GET","page":"about"}'
```

### Info log
Input:
```go
kvlog.LogInfo("database", "something happened")
```
Output:
```
ts='2017-03-08T11:48:48-0700' type='database' msg='{"info":"something happened"}'
```

### Error log
Input:
```go
kvlog.LogErr("database", "something broke")
```
Output:
```
ts='2017-03-08T11:48:48-0700' type='database' msg='{"error":"something broke"}'
```

##
