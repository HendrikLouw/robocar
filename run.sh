GOARM=6 GOARCH=arm GOOS=linux go build robocar.go
ssh pi@172.20.10.2 -C "sudo killall -9 robocar"
scp robocar pi@172.20.10.2:
ssh pi@172.20.10.2 -C "sudo ./robocar"
