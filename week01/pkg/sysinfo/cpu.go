package sysinfo

import (
	"bufio"
	"os"
	"strings"
)

func CPUCores() (int, error){
	f,err := os.Open("/proc/cpuinfo")
	if err != nil{
		return 0,err
	}

	defer f.Close()

	count := 0
	s:= bufio.NewScanner(f)
	for s.Scan(){
		if strings.HasPrefix(s.Text(),"processor"){
			count++
		}
	}

	return count,nil

}