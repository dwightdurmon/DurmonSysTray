all: linux windows

linux:
	rm -f DurmonSysTray
	env GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o DurmonSysTray main.go

windows:
	rm -f DurmonSysTray.exe
	rm -f resource.syso
	go generate
	env GO111MODULE=on GOOS=windows go build -ldflags -H=windowsgui -o DurmonSysTray.exe

clean:
	rm -f DurmonSysTray.exe
	rm -f resource.syso
	rm -f DurmonSysTray
	rm -f nohup.out