package analyser

import (
	"fmt"
	"gopkg.in/eapache/queue.v1"
	"sort"
	"strconv"
	"strings"
)

type Nmon struct {
	seriesClass []string
	seriesMap   map[string]seriesLine
}

func newNmon() Nmon {
	return Nmon{
		seriesClass: make([]string, 0, 32),
		seriesMap:   make(map[string]seriesLine),
	}
}

func (n *Nmon) saveData(cls string, seriesData interface{}) {
	if sl, ok := n.seriesMap[cls]; !ok {
		s := newDataLine()
		s.push(seriesData)
		n.seriesMap[cls] = s
		n.seriesClass = append(n.seriesClass, cls)
	} else {
		sl.push(seriesData)
	}
}

func (n *Nmon) GetSeriesClass() []string {
	return n.seriesClass
}

func (n *Nmon) showSeriesData(cls string) {
	if sl, ok := n.seriesMap[cls]; ok {
		count := sl.length()
		for i := 0; i < count; i++ {
			fmt.Println(sl.get(i))
		}
	}
}

// 排序并将每一个 CPU 的使用情况分类放到最后
func (n *Nmon) sortSeriesClass() {
	sort.Slice(n.seriesClass, func(i, j int) bool {
		isPerCPU := func(cls string) bool {
			return strings.HasPrefix(cls, "CPU") && !strings.HasSuffix(cls, "_ALL")
		}
		left := n.seriesClass[i]
		right := n.seriesClass[j]
		if isPerCPU(left) && isPerCPU(right) {
			l, _ := strconv.Atoi(left[3:])
			r, _ := strconv.Atoi(right[3:])
			return l < r
		}
		if isPerCPU(left) {
			return false
		}
		if isPerCPU(right) {
			return true
		}
		return left < right
	})

	// sort.Strings(n.seriesClass)
	// fmt.Println(sort.StringsAreSorted(n.seriesClass))

	// start, end := -1, -1
	// for i, s := range n.seriesClass {
	// 	if strings.HasPrefix(s, "CPU") && !strings.HasSuffix(s, "_ALL") {
	// 		if start == -1 {
	// 			start = i
	// 		} else if start != -1 {
	// 			end = i
	// 		}
	// 	}
	// }
	// n.seriesClass = append(n.seriesClass[:0],
	// 	append(n.seriesClass[:start],
	// 		append(n.seriesClass[end+1:], n.seriesClass[start:end+1]...
	// 		)...
	// 	)...
	// )
}

// seriesLine 单并发使用，不考虑并发安全
type seriesLine struct {
	// mu sync.Mutex
	q *queue.Queue
}

func newDataLine() seriesLine {
	return seriesLine{q: queue.New()}
}

func (sl *seriesLine) push(v interface{}) {
	// sl.mu.Lock()
	// defer sl.mu.Unlock()
	sl.q.Add(v)
}

// func (sl *seriesLine) take() interface{} {
// 	// sl.mu.Lock()
// 	// defer sl.mu.Unlock()
// 	return sl.q.Remove()
// }

func (sl *seriesLine) get(i int) interface{} {
	// sl.mu.Lock()
	// defer sl.mu.Unlock()
	return sl.q.Get(i)
}

func (sl *seriesLine) length() int {
	return sl.q.Length()
}

// 暂未发现排序不正确的情况
// func (sl seriesLine) Len() int {
// 	return sl.q.Length()
// }
//
// func (sl seriesLine) Less(i, j int) bool {
// 	split1 := bytes.Split(sl.get(i).([]byte), []byte(","))
// 	split2 := bytes.Split(sl.get(j).([]byte), []byte(","))
// 	return bytes.Compare(split1[0], split2[0]) < 0
// 	return bytes.Compare(split1[1], split2[1]) < 0
// }
//
// func (sl seriesLine) Swap(i, j int) {
// 	panic("implement me")
// }
