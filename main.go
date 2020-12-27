//go:generate goversioninfo -icon=assets/info.ico -manifest=assets/DurmonSysTray.exe.manifest
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/getlantern/systray"
	"github.com/matishsiao/goInfo"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	gi := goInfo.GetInfo()
	systray.SetIcon(getIcon("assets/info.ico"))
	systray.SetTitle("DurmonSysTray")
	systray.SetTooltip("Click me for info!")
	systray.AddMenuItem("Hostname: "+gi.Hostname, "Ignored")
	systray.AddMenuItem(GetOutboundIP(), "Ignored")
	systray.AddMenuItem(GetOutboundIPMac(), "Ignored")

	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	// Cleaning stuff here.
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}

// GetOutboundIP ...
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "169.254.169.254:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	outString := "IP: " + localAddr.IP.String()

	return outString
}

// GetOutboundIPMac ...
func GetOutboundIPMac() string {
	var currentNetworkHardwareName string
	conn, err := net.Dial("udp", "169.254.169.254:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr).IP.String()

	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {

		if addrs, err := interf.Addrs(); err == nil {
			for _, addr := range addrs {
				if strings.Contains(addr.String(), localAddr) {
					currentNetworkHardwareName = interf.Name
				}
			}
		}
	}

	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

	if err != nil {
		fmt.Println(err)
	}

	macAddress := netInterface.HardwareAddr
	outstring := "MAC : " + macAddress.String()
	return outstring
}
