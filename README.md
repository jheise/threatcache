#ThreatCache
---

Caching proxy for ThreatCrowd.org

Available routes are

- /ip/{addr}/

- /domain/{domain}/

- /av/{av}/

- /file/{file}/

- /email/{email}/

---

```
Usage of threatcache:
  -ipaddr string
        Address to bind on (default "0.0.0.0")
  -port string
        Port to bind on (default "8888")
  -redishost string
        Redis host to connect to (default "localhost")
  -redisport string
        Redis port to connect on (default "6379")
  -timeout string
        TTL for caching (default "14400")

```
