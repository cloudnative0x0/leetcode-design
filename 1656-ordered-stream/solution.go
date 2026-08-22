package _1656_ordered_stream

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		stream: make([]string, n+2),
		ptr:    1,
	}
}

func (odr *OrderedStream) Insert(idKey int, value string) []string {
	odr.stream[idKey] = value

	var chunk []string

	for odr.ptr < len(odr.stream) && odr.stream[odr.ptr] != "" {
		chunk = append(chunk, odr.stream[odr.ptr])
		odr.ptr++
	}

	return chunk
}
