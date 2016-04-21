WARNING: Do not use this if you are not fully aware of the context.

This is a script for fixing LP: #1459033 on environments prior to 1.25.2

# Usage:

```shell
$ juju ssh 0
machine-0 $ sudo apt-get install git-core 
machine-0 $ curl https://github.com/niedbalski/raw/run-fix-1459033 && chmod 0755 
machine-0 $ sudo ./run-fix-1459033
```

Expected Output:

```shell
2016-04-21 15:55:51 DEBUG juju.agent agent.go:482 read agent config, format "1.18"
2016-04-21 15:55:51 INFO juju.state open.go:51 opening state, mongo addresses: ["127.0.0.1:37017"]; entity machine-0
2016-04-21 15:55:51 DEBUG juju.state open.go:52 dialing mongo
2016-04-21 15:55:51 INFO juju.mongo open.go:125 dialled mongo successfully on address "127.0.0.1:37017"
2016-04-21 15:55:51 DEBUG juju.state open.go:57 connection established
2016-04-21 15:55:51 INFO juju.mongo open.go:125 dialled mongo successfully on address "10.5.1.84:37017"
2016-04-21 15:55:51 DEBUG juju.state open.go:64 mongodb login successful
2016-04-21 15:55:51 INFO juju.state state.go:195 creating lease client as machine-0
2016-04-21 15:55:51 INFO juju.state state.go:208 starting leadership manager
2016-04-21 15:55:51 INFO juju.state state.go:219 creating cloud image metadata storage
2016-04-21 15:55:51 INFO juju.state state.go:222 starting presence watcher
2016-04-21 15:55:51 DEBUG juju.state.leadership manager.go:214 waking to check leases at 2016-04-21 15:56:24.016547094 +0000 UTC
2016-04-21 15:55:51 INFO juju.mongo open.go:125 dialled mongo successfully on address "10.5.1.84:37017"

2016-04-21 15:55:51 DEBUG juju.state.toolstorage tools.go:193 invalid tools version: {'\x02' "\x10\x00\x00\x001.25.3.1--amd64\x00"}

2016-04-21 15:55:51 INFO juju.mongo open.go:125 dialled mongo successfully on address "127.0.0.1:37017"
2016-04-21 15:55:51 INFO juju.mongo open.go:125 dialled mongo successfully on address "10.5.1.84:37017"
2016-04-21 15:55:51 DEBUG juju.state open.go:306 closed state without error
```
