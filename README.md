# go-interval

[![Build Status](https://travis-ci.org/retailify/go-interval.svg?branch=master)](https://travis-ci.org/retailify/go-interval) [![Go Report Card](https://goreportcard.com/badge/github.com/retailify/go-interval)](https://goreportcard.com/report/github.com/retailify/go-interval) [![Maintainability](https://api.codeclimate.com/v1/badges/19b3641a71295105f000/maintainability)](https://codeclimate.com/github/retailify/go-interval/maintainability)


This library implements the [Allen's interval algebra](http://www.ics.uci.edu/~alspaugh/cls/shr/allen.html) for
go lang time intervals.

The go-interval library supports answering questions concerning time intervals.
You can check if intervals overlap, seamlessly connect, there is a gap between the intervals, and more.

## Documentation

[godoc.org documentation](https://godoc.org/github.com/retailify/go-interval)

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

