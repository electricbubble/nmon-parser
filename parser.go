package nmonparser

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ParseNmon(name string) (nmon Nmon, err error) {
	nmonFile, err := os.Open(name)
	if err != nil {
		return Nmon{}, err
	}
	defer func() {
		_ = nmonFile.Close()
	}()

	nmon = newNmon()

	reader := bufio.NewReader(nmonFile)
	line, err := reader.ReadString('\n')
	for err == nil {
		indexSep := strings.Index(line, ",")
		cls := line[:indexSep]
		line = line[indexSep+1 : len(line)-1]
		// fmt.Println(cls, line)
		nmon.saveData(cls, line)
		line, err = reader.ReadString('\n')
	}
	if err != io.EOF {
		return Nmon{}, err
	}

	// fmt.Println(nmon.seriesClass)
	nmon.sortSeriesClass()
	// fmt.Println(nmon.seriesClass)

	return nmon, nil
}
