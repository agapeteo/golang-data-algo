package arrays

type RingBuffer struct {
	container []interface{}
	readIdx   int
	writeIdx  int
	count     int
	max       int
}

func NewRingBuffer(maxSize int) *RingBuffer {
	buf := new(RingBuffer)
	buf.max = maxSize
	buf.container = make([]interface{}, buf.max)
	return buf
}

func (buf *RingBuffer) Dequeue() (elem interface{}, ok bool) {
	ok = true
	if buf.count == 0 {
		return nil, false
	}
	elem = buf.container[buf.readIdx]

	buf.readIdx = buf.nextIdx(buf.readIdx)
	buf.count--
	return
}

func (buf *RingBuffer) Enqueue(elem interface{}) {
	buf.container[buf.writeIdx] = elem
	buf.writeIdx = buf.nextIdx(buf.writeIdx)
	if buf.count == buf.max {
		buf.readIdx = buf.writeIdx
	} else {
		buf.count++
	}
}

func (buf *RingBuffer) Elements() (result []interface{}) {
	result = make([]interface{}, 0, buf.max)

	curReadIdx := buf.readIdx
	for i := 0; i < buf.count; i++ {
		result = append(result, buf.container[curReadIdx])
		curReadIdx = buf.nextIdx(curReadIdx)
	}
	return
}

func (buf *RingBuffer) nextIdx(curIdx int) int {
	return (curIdx + 1) % buf.max
}
