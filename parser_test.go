package nmonparser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestTmpMain(t *testing.T) {
	allSheetSlice := make([]string, 0, 32)
	allSheetMap := make(map[string]SeriesLine)

	name := "/Users/hero/Documents/temp/2020-01/to_guigui/LINUX_single_2563_144_50_10u10m.nmon"
	// name = "/Users/hero/Documents/temp/2020-01/to_guigui/AIX_single_2563_19_50_10u10m.nmon"
	// name = "/Users/hero/Documents/Pactera/CEC-Performance-Testing/2019-12/GoldenDB_2019-12-27/transaction3_1000w/Mixed_Node1/201912261503_Mixed_ip71_10.10.22.87.nmon"

	file, err := os.Open(name)
	if err != nil {
		log.Println(file)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	line, err := reader.ReadBytes('\n')
	for err == nil {
		line = line[:len(line)-1]
		split := bytes.Split(line, []byte(","))
		k := string(split[0])
		line = line[len(split[0])+1:]
		// fmt.Println("#", string(line[len(split[0])+1:]))
		if dl, ok := allSheetMap[k]; !ok {
			l := newDataLine()
			l.push(line)
			allSheetMap[k] = l
			allSheetSlice = append(allSheetSlice, k)
		} else {
			dl.push(line)
		}
		// if bytes.HasPrefix(line, []byte("AAA")) {
		// 	fmt.Println(string(line))
		// }
		line, err = reader.ReadBytes('\n')
	}
	if err == io.EOF {
		log.Println("文件读取 EOF")
	} else {
		log.Println("文件读取错误", err)
		return
	}

	fmt.Println(len(allSheetSlice))
	fmt.Println(allSheetSlice)

	sort.Strings(allSheetSlice)
	fmt.Println(allSheetSlice)
	startCPU, endCPU := -1, -1
	for i, s := range allSheetSlice {
		if strings.HasPrefix(s, "CPU") && !strings.HasSuffix(s, "_ALL") {
			if startCPU == -1 {
				startCPU = i
			} else if startCPU != -1 {
				endCPU = i
			}
		}
	}
	fmt.Println(startCPU, endCPU)
	allSheetSlice = append(allSheetSlice[:0],
		append(allSheetSlice[:startCPU],
			append(allSheetSlice[endCPU+1:], allSheetSlice[startCPU:endCPU+1]...
			)...
		)...
	)
	fmt.Println(allSheetSlice)

	for _, k := range allSheetSlice {
		dl := allSheetMap[k]
		fmt.Println(k, dl.Len())
		// fmt.Println(k, dl.Len(), sort.IsSorted(dl))
	}

	if _, ok := allSheetMap["ZZZZ"]; !ok {
		log.Println("无 ZZZZ 数据")
		return
	}

	// zzzz := allSheetMap["ZZZZ"]

	if dl, ok := allSheetMap["CPU_ALL"]; ok {
		for i, count := 0, dl.Len(); i < count; i++ {
			bsLine := dl.Get(i).([]byte)
			// if i > 0 {
			// 	split := bytes.Split(bsLine, []byte(","))
			// 	bsZ := zzzz.Get(i - 1).([]byte)
			// 	fmt.Println(string(split[1]), string(bsZ))
			// 	fmt.Println(string(bsLine))
			// 	break
			// }
			fmt.Println(string(bsLine))
		}
	}
}

func TestParseNmon(t *testing.T) {
	name := "/Users/hero/Documents/temp/2020-01/to_guigui/LINUX_single_2563_144_50_10u10m.nmon"
	// name = "/Users/hero/Documents/temp/2020-01/to_guigui/AIX_single_2563_19_50_10u10m.nmon"
	// name = "/Users/hero/Documents/Pactera/CEC-Performance-Testing/2019-12/GoldenDB_2019-12-27/transaction3_1000w/Mixed_Node1/201912261503_Mixed_ip71_10.10.22.87.nmon"

	nmon, err := ParseNmonByFilename(name)
	if err != nil {
		t.Fatal(err)
	}

	// nmon.showSeriesData("ZZZZ")
	// nmon.showSeriesData("CPU_ALL")
	// fmt.Println(nmon.GetSeriesClass())
	sl := nmon.GetSeriesLine("CPU_ALL")
	count := sl.Len()
	for i := 0; i < count; i++ {
		fmt.Println(sl.Get(i))
	}
	fmt.Println(sl.Len())
	// fmt.Println(http.ParseTime("08:47:01 08-JAN-2020"))
}
