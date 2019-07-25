[![BuildStatus](https://travis-ci.org/Yannick-S/mdhtml.svg?branch=master)](https://travis-ci.org/Yannick-S/mdhtml)
[![Codecov](https://img.shields.io/codecov/c/github/Yannick-S/mdhtml.svg)](https://codecov.io/github/Yannick-S/mdhtml?branch=master)
[![Go ReportCard](https://goreportcard.com/badge/github.com/Yannick-S/mdhtml)](https://goreportcard.com/report/github.com/Yannick-S/mdhtml)
[![Documentation](https://godoc.org/github.com/Yannick-S/mdhtml?status.svg)](http://godoc.org/github.com/Yannick-S/mdhtml)
![GitHub](https://img.shields.io/github/license/Yannick-S/mdhtml.svg)
![GitHub tag (latestbydate)](https://img.shields.io/github/tag-date/Yannick-S/mdhtml.svg)

# MDHTML : converting markdown to html files

This library provides a way for converting markdown files to html, to be view in
browsers and other. The main aim of this library is for me to learn how to
manage a library, write documentation, tests etc, do continuous integration and
other things along the way.

This project is inspired by a Podcast I heard
[here](https://testandcode.com/80), which creates the same library, however
written in python.

### Usage

```
$ ./mdhtml -in test.md -out test.html
$ [TODO] ./mdhtml test.md > test.html
$ [TODO] cat test.md | ./mdhtml > test.html
```

### History

- 0.1 Initial Tests
- 0.2 Getting Travis and codecov to work. Adding more badges while I was at it.
- 0.3 Starting on actual project. It also seems that I have misunderstood how
  tags worked before
- 0.4 At this stage, blocks work. Later I will need to split this package into
  two, one library and one cli
