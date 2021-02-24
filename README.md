DISCLAIMER: The readme looks better in dark mode.

![WAG](example/wag.png)

Renders text in whatever cheesy bitmap font you got.

![Usage](example/usage.png)

Run wag from your commandline, parameter:

```
  -h  --help     Print help information
  -f  --font     <font>.json you want to use. Default: font5
  -t  --text     text you want to write. Default: ABCDEFGHIJKLMNOPQRSTUVWXYZ 1234567890 ()!?:;.,"'
  -s  --spacing  Optional spacing . Default: 0
  -c  --color    Colour to be made transparent. Default: #000000FF
  -o  --output   Output file name (without ending). Default: text
```
Example `wag -f font3 -t  "HELLO WORLD" -o hello`


![Installation](example/installation.png)

* Grab the binary you need at the release tab. They are built using gox, and not really tested on every system.

![Build yourself](example/development.png)

1. **Be sure to run Go 1.16, because of the embed library.**
2. Clone the repo using `git clone --recursive https://github.com/morgulbrut/wag.git`
3. Cd into the dir get all the dependencies using `go get ./...`
4. Build using your normal go build setup.

![Add fonts](example/fonts.png)

1. Write a json file for your font, see `fontdesc/32X32-F1.json`
2. Add it to `fontdesc/fonts.json`
3. Make sure the path for the font png file is in `main.go` in some `//go:embed` comment, see example below

```
type imgSize image.Point

//go:embed fontdesc/*
//go:embed BitmapFonts/32X32-F*.png
var data embed.FS
```
