# Go-tools


## Goget

wget-like application that takes n [n >=1 ] URL arguments as an input and downloads all of them simultaneously to the current directory

### Build

`go build -o goget.text cmd/goget/main.go`

### Test Run

`./goget.text http://212.183.159.230/200MB.zip http://212.183.159.230/50MB.zip https://speed.hetzner.de/1GB.bin`

Example
```bash
$ ls -1
README.md
cmd
go.mod
goget
$ go build -o goget.text cmd/goget/main.go
$ time ./goget.text http://212.183.159.230/200MB.zip http://212.183.159.230/50MB.zip https://speed.hetzner.de/1GB.bin
0. 200MB.zip 1. 50MB.zip 2. 1GB.bin 
[-----] [-----] [-----] 
[   0%] [   1%] [   0%] 
[   1%] [   2%] [   0%] 
[   1%] [   3%] [   3%] 
[   2%] [   5%] [   7%] 
[   3%] [   7%] [  10%] 
[   3%] [  10%] [  14%] 
[   5%] [  13%] [  17%] 
[   6%] [  17%] [  20%] 
[   8%] [  22%] [  23%] 
[   9%] [  26%] [  26%] 
[  11%] [  32%] [  30%] 
[  13%] [  39%] [  33%] 
[  16%] [  46%] [  35%] 
[  18%] [  54%] [  38%] 
[  21%] [  64%] [  41%] 
[  24%] [  73%] [  44%] 
[  27%] [  84%] [  46%] 
[  30%] [  96%] [  49%] 
[  33%] [ 100%] [  52%] 
[  37%] [ 100%] [  55%] 
[  41%] [ 100%] [  58%] 
[  45%] [ 100%] [  61%] 
[  49%] [ 100%] [  64%] 
[  54%] [ 100%] [  66%] 
[  58%] [ 100%] [  69%] 
[  63%] [ 100%] [  72%] 
[  68%] [ 100%] [  75%] 
[  73%] [ 100%] [  77%] 
[  78%] [ 100%] [  80%] 
[  84%] [ 100%] [  83%] 
[  89%] [ 100%] [  86%] 
[  95%] [ 100%] [  88%] 
[ 100%] [ 100%] [  91%] 
[ 100%] [ 100%] [  94%] 
[ 100%] [ 100%] [  98%] 
./goget.text http://212.183.159.230/200MB.zip http://212.183.159.230/50MB.zip  1.76s user 4.10s system 16% cpu 35.520 total
$ ls -1
1GB.bin
200MB.zip
50MB.zip
README.md
cmd
go.mod
goget
goget.text
```
