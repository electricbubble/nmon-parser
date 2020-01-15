package analyser

import (
	"testing"
)

func TestDataLine(t *testing.T) {
	// var wg sync.WaitGroup
	// wg.Add(20)

	dl := newDataLine()
	for i := 0; i < 20; i++ {
		dl.push(i)
		// go func(v int) {
		// 	dl.push(v)
		// 	wg.Done()
		// }(i)
	}
	// time.Sleep(time.Millisecond * 500)
	// wg.Wait()
	t.Log("length:", dl.length())

	for i, count := 0, dl.length(); i < count; i++ {
		t.Log("get:", dl.get(i))
		// t.Log("take:", dl.take())
		// go func(no int) {
		// 	t.Log("gNo:", no, "take:", dl.take())
		// }(i)
	}
}
