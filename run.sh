GOARM=6 GOARCH=arm GOOS=linux go build robocar.go
ssh pi@192.168.43.13 -C "sudo killall -9 robocar"
scp robocar pi@192.168.43.13:
ssh pi@192.168.43.13 -C "sudo ./robocar"
