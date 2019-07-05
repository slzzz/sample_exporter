package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"

	"io"
	"log"
	"net/http"
	"strconv"
	// "reflect"
	"flag"
	"runtime"
	"time"
)

func metrics(w http.ResponseWriter, r *http.Request) {
	m, _ := mem.VirtualMemory()
	c, _ := cpu.Percent(time.Second, false)
	h, _ := host.Info()
	nv, _ := net.IOCounters(true)
	l, _ := load.Avg()
	count, _ := cpu.Counts(true)

	if runtime.GOOS == "windows" {
		d, _ := disk.Usage("c:/")
		fmt.Fprintf(w, "Mem_UsedPercent %v\nCpu_UsedPercent %v\nDisk_UsedPercent %v\nDisk_Inodes_UsedPercent %v\nUptime %d\nBootTime %d\nBytesRecv %v\nBytesSend %v\nLoad1 %v\nLoad5 %v\nLoad15 %v\nCpu_Count %d\n",
			m.UsedPercent, c[0], d.UsedPercent, d.InodesUsedPercent, h.Uptime, h.BootTime, nv[0].BytesRecv, nv[0].BytesSent, l.Load1, l.Load5, l.Load15, count)
	} else {
		d, _ := disk.Usage("/")
		fmt.Fprintf(w, "Mem_UsedPercent %v\nCpu_UsedPercent %v\nDisk_UsedPercent %v\nDisk_Inodes_UsedPercent %v\nUptime %d\nBootTime %d\nBytesRecv %v\nBytesSend %v\nLoad1 %v\nLoad5 %v\nLoad15 %v\nCpu_Count %d\n",
			m.UsedPercent, c[0], d.UsedPercent, d.InodesUsedPercent, h.Uptime, h.BootTime, nv[0].BytesRecv, nv[0].BytesSent, l.Load1, l.Load5, l.Load15, count)
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	port := flag.Int("port", 12345, "monitor listen port")
	flag.Parse()

	fmt.Println("Base Metrics Exports v1.0 is Running On Port :" + strconv.Itoa(*port) + "!!!!, Modify The Default Port,please use the -port args.\nAuthor: slzzz\nEmail: 21107689@qq.com ")
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/metrics", metrics)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
