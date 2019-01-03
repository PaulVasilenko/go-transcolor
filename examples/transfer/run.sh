#!/bin/bash

go run ./main.go -target ./source/minsk.jpg  -source ./source/bladerunnergray.jpg -out res_gray.png
go run ./main.go -target ./source/minsk.jpg  -source ./source/bladerunneryellow.jpg -out res_yellow.png
