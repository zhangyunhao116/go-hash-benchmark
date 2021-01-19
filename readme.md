# go-hash-benchmark



rand1 -> github.com/zhangyunhao116/wyhash

rand2 -> github.com/zeebo/xxh3

rand3 -> github.com/zeebo/wyhash

rand4 -> github.com/cespare/xxhash/v2



### How to run this benchmark

```go
$ go test -bench=All/rand[1234]
```



### Result

Go version: go1.15.6 linux/amd64

CPU: AMD 3700x(8C16T), running at 3.6GHz

OS: ubuntu 18.04

MEMORY: 16G x 2 (3200MHz)

rand1 and rand2 is faster than others

```
name                 time/op
All/rand1-3-16         3.88ns ± 0%
All/rand2-3-16         2.79ns ± 0%
All/rand1-6-16         3.65ns ± 0%
All/rand2-6-16         3.63ns ± 0%
All/rand1-9-16         3.65ns ± 0%
All/rand2-9-16         3.35ns ± 0%
All/rand1-12-16        3.66ns ± 0%
All/rand2-12-16        3.38ns ± 0%
All/rand1-15-16        3.65ns ± 0%
All/rand2-15-16        3.34ns ± 0%
All/rand1-21-16        4.75ns ± 0%
All/rand2-21-16        5.21ns ± 0%
All/rand1-24-16        4.37ns ± 0%
All/rand2-24-16        5.20ns ± 0%
All/rand1-29-16        4.70ns ± 0%
All/rand2-29-16        5.17ns ± 0%
All/rand1-32-16        4.67ns ± 0%
All/rand2-32-16        5.19ns ± 0%
All/rand1-35-16        6.00ns ± 0%
All/rand2-35-16        7.21ns ± 0%
All/rand1-64-16        7.16ns ± 0%
All/rand2-64-16        7.16ns ± 0%
All/rand1-69-16        7.38ns ± 0%
All/rand2-69-16        8.67ns ± 0%
All/rand1-96-16        8.34ns ± 0%
All/rand2-96-16        8.58ns ± 0%
All/rand1-97-16        9.18ns ± 0%
All/rand2-97-16        10.4ns ± 0%
All/rand1-128-16       10.3ns ± 1%
All/rand2-128-16       10.4ns ± 0%
All/rand1-129-16       10.1ns ± 0%
All/rand2-129-16       11.9ns ± 0%
All/rand1-240-16       15.4ns ± 0%
All/rand2-240-16       21.6ns ± 0%
All/rand1-241-16       16.4ns ± 1%
All/rand2-241-16       24.1ns ± 0%
All/rand1-512-16       28.1ns ± 0%
All/rand2-512-16       27.3ns ± 0%
All/rand1-1024-16      50.5ns ± 0%
All/rand2-1024-16      39.7ns ± 0%
All/rand1-102400-16    4.29µs ± 0%
All/rand2-102400-16    2.07µs ± 0%

name                 speed
All/rand1-3-16        773MB/s ± 0%
All/rand2-3-16       1.07GB/s ± 0%
All/rand1-6-16       1.65GB/s ± 0%
All/rand2-6-16       1.65GB/s ± 0%
All/rand1-9-16       2.47GB/s ± 0%
All/rand2-9-16       2.68GB/s ± 0%
All/rand1-12-16      3.29GB/s ± 0%
All/rand2-12-16      3.56GB/s ± 0%
All/rand1-15-16      4.11GB/s ± 0%
All/rand2-15-16      4.49GB/s ± 0%
All/rand1-21-16      4.43GB/s ± 0%
All/rand2-21-16      4.03GB/s ± 0%
All/rand1-24-16      5.49GB/s ± 0%
All/rand2-24-16      4.62GB/s ± 0%
All/rand1-29-16      6.17GB/s ± 0%
All/rand2-29-16      5.60GB/s ± 0%
All/rand1-32-16      6.85GB/s ± 0%
All/rand2-32-16      6.17GB/s ± 0%
All/rand1-35-16      5.83GB/s ± 0%
All/rand2-35-16      4.86GB/s ± 0%
All/rand1-64-16      8.95GB/s ± 0%
All/rand2-64-16      8.94GB/s ± 0%
All/rand1-69-16      9.34GB/s ± 0%
All/rand2-69-16      7.96GB/s ± 0%
All/rand1-96-16      11.5GB/s ± 0%
All/rand2-96-16      11.2GB/s ± 0%
All/rand1-97-16      10.6GB/s ± 0%
All/rand2-97-16      9.36GB/s ± 0%
All/rand1-128-16     12.4GB/s ± 1%
All/rand2-128-16     12.3GB/s ± 0%
All/rand1-129-16     12.8GB/s ± 0%
All/rand2-129-16     10.8GB/s ± 0%
All/rand1-240-16     15.6GB/s ± 0%
All/rand2-240-16     11.1GB/s ± 0%
All/rand1-241-16     14.7GB/s ± 1%
All/rand2-241-16     10.0GB/s ± 0%
All/rand1-512-16     18.3GB/s ± 0%
All/rand2-512-16     18.8GB/s ± 0%
All/rand1-1024-16    20.3GB/s ± 0%
All/rand2-1024-16    25.8GB/s ± 0%
All/rand1-102400-16  23.9GB/s ± 0%
All/rand2-102400-16  49.6GB/s ± 0%
```

