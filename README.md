# depextract

`depextract` is a very simple tool that will parse a `Gopkg.lock` from
[`dep`](https://github.com/golang/dep) and dump all of the dependencies out on
the console for you.

The code has been derived from the actual conversion code used in `go mod init`,
but with all of the bits stripped out that don't pertain to output.

## Install

```
go get -u github.com/vancluever/depextract/cmd/depextract
```

## Usage

```
depextract FILE
```

Example: `depextract Gopkg.lock`

## License

```
Copyright 2019 Chris Marchesi. Licensed under the terms of the MPL, which can
be see in the LICENSE file.

Code here contains work derived from the Go project, Copyright 2018 The Go
Authors. All rights reserved. Use of this source code is governed by a BSD-style
license that can be found in the LICENSE-Go file.
```
