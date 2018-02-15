go-video-thumb
========================

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[license]: https://github.com/vural/go-video-thumb/blob/master/LICENSE
[godocs]: https://godoc.org/github.com/vural/go-video-thumb

Get the thumbnails from Youtube and Vimeo videos for Golang.

## install

```sh
go get -u github.com/vural/go-video-thumb
```

## usage
```go
package main

import "github.com/vural/go-video-thumb"

func main() {
	image := thumb.Get("https://www.youtube.com/watch?v=qUsE49lUQUY", thumb.LARGE)
	print(image)
}
```
