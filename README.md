# go-interval

**BE WARNED:** This is work in progress. The library is not completely
implemented!
The library is ready for production if I tag the library with 1.0.0.

[![Build Status][travis-svg]][travis-url]
[![Go Report Card][go-report-badge]][go-report-url]
[![Maintainability][codeclimate-badge]][codeclimate-url]


This library implements the [Allen's interval algebra][allens-url] for
go lang time intervals.

The go-interval library supports answering questions concerning time
intervals.
You can check if intervals overlap, seamlessly connect, there is a gap
between the intervals, and more.

## Documentation

[godoc.org documentation][go-doc-url]

## Installation

### go get

``` bash

go get github.com/retailify/go-interval
```

### glide

``` bash

glide get github.com/retailify/go-interval
```


## Usage

``` go
import "github.com/retailify/go-interval"
```

copyright 2018 by [Retailify GmbH](https://retailify.de)

[travis-svg]: https://travis-ci.org/retailify/go-interval.svg?branch=master
[travis-url]: https://travis-ci.org/retailify/go-interval
[go-report-badge]: https://goreportcard.com/badge/github.com/retailify/go-interval
[go-report-url]: https://goreportcard.com/report/github.com/retailify/go-interval
[codeclimate-badge]: https://api.codeclimate.com/v1/badges/19b3641a71295105f000/maintainability
[codeclimate-url]: https://codeclimate.com/github/retailify/go-interval/maintainability
[allens-url]: http://www.ics.uci.edu/~alspaugh/cls/shr/allen.html
[go-doc-url]: https://godoc.org/github.com/retailify/go-interval