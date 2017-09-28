Fond of the Year 2017 lank board animation
==========================================

# Requirements

- GopherJS
- GopherJS WebGL
- three (ThreeJS GopherJS bindings)

## Set GOPATH

```
export GOPATH=~/gocode:`pwd`
export PATH=~/gocode/bin:$PATH
```

## Get libraries

```
go get -u github.com/gopherjs/gopherjs/...
go get -u github.com/gopherjs/webgl/...
go get -u bitbucket.org/mikehouston/three/...
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
