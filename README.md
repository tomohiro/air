Air
================================================================================

Command-line AirPlay video client for Apple TV

[![Build Status](https://travis-ci.org/Tomohiro/air.svg?branch=master)](https://travis-ci.org/Tomohiro/air)


Usage
--------------------------------------------------------------------------------

### Play video

```sh
$ air play '~/Movies/Trailers/007 SKYFALL.mp4'
```


Installation
--------------------------------------------------------------------------------

```sh
$ go get github.com/Tomohiro/air
```


Supported MIME types
--------------------------------------------------------------------------------

[AirPlay Overview - Configuring Your Server](http://developer.apple.com/library/ios/#documentation/AudioVideo/Conceptual/AirPlayGuide/PreparingYourMediaforAirPlay/PreparingYourMediaforAirPlay.html)

File extension | MIME type
-------------- | ----------------
.ts            | video/MP2T
.mov           | video/quicktime
.mp3           | audio/MPEG3
.aac           | audio/aac
.m4a           | audio/mpeg4
.m4v, mp4      | video/mpeg4


LICENSE
--------------------------------------------------------------------------------

&copy; 2014 Tomohiro TAIRA.
This project is licensed under the MIT license.
See LICENSE for details.
