Fond of the Year 2017 lank board animation
==========================================

# Requirements

- GopherJS
- GopherJS WebGL

```
export GOPATH=~/gocode
export PATH=~/gocode/bin:$PATH
go get -u github.com/gopherjs/gopherjs/...
go get -u github.com/gopherjs/webgl/...
```

# Build

```
gopherjs build main.go
```

or

```apple js
./build.sh
```

# Run

```
open ./index.html
```
