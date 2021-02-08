DISCLAIMER: The readme looks better in dark mode.

![WAG](example/wag.png)

Renders text in whatever cheesy bitmap font you got.

![Usage](example/usage.png)

1. Get a bitmap font which is basically an image (https://github.com/ianhan/BitmapFonts has tons of them)
2. Write a json file with the positions of the characters (see example).
3. Write some text using `wag -f <YOURFONT.json> -t <YOUR TEXT>`
4. Text should be found in `text.png`

![Installation](example/installation.png)

Just use good old `go get github.com/morgulbrut/wag`
