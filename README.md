# Xi-radio

[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/radio/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/radio?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/radio.svg?branch=master)](https://travis-ci.org/Konstantin8105/radio)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/radio)](https://goreportcard.com/report/github.com/Konstantin8105/radio)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/radio/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/Konstantin8105/radio?status.svg)](https://godoc.org/github.com/Konstantin8105/radio)


Xi-radio is terminal radio

![logo](https://github.com/Konstantin8105/Xi-radio/blob/master/pic/logo.svg.png)

(Used font **Go Mono**)

Xi is greek letter(see https://en.wikipedia.org/wiki/Xi_(letter))

## Example

![terminal](https://github.com/Konstantin8105/Xi-radio/blob/master/pic/radio.png)

## Run console radio

1. Commands to get source:
```
go get -u github.com/Konstantin8105/radio
```
2. Run:
```
go run ./cmd/main.go
```

## Commands of terminal radio

```
Start : Ξ (Xi-radio)
Enter 'help' for show all commands
Found : 500 stations
     clear	Clear playlist in player
      exit	Exit from terminal radio
      help	Show all commands
      info	Return information from player about current stream
      list	List of all allowable radio stations
      play	Play [station], is [station] is empty, then playing random station
    resume	Continue playing stopped stream
    search	Search [query] a specific name
      stop	Stop stream in player
     title	Title of the current stream
```
