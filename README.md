# kvlog
Cripplingly simple kv logging

## Logging with arbitrary map input
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
