package astiav

//#include <libavfilter/avfilter.h>
import "C"
import (
	"math"
	"unsafe"
)

// https://ffmpeg.org/doxygen/7.0/structAVFilterGraphSegment.html
type FilterGraphSegment struct {
	c *C.AVFilterGraphSegment
}

func newFilterGraphSegmentFromC(c *C.AVFilterGraphSegment) *FilterGraphSegment {
	if c == nil {
		return nil
	}
	return &FilterGraphSegment{c: c}
}

// https://ffmpeg.org/doxygen/7.0/group__lavfi.html#ga51283edd8f3685e1f33239f360e14ae8
func (fgs *FilterGraphSegment) Free() {
	if fgs.c != nil {
		C.avfilter_graph_segment_free(&fgs.c)
	}
}

// https://ffmpeg.org/doxygen/7.0/structAVFilterGraphSegment.html#ad5a2779af221d1520490fe2719f9e39a
func (fgs *FilterGraphSegment) Chains() (cs []*FilterChain) {
	ccs := (*[(math.MaxInt32 - 1) / unsafe.Sizeof((*C.AVFilterChain)(nil))](*C.AVFilterChain))(unsafe.Pointer(fgs.c.chains))
	for i := 0; i < fgs.NbChains(); i++ {
		cs = append(cs, newFilterChainFromC(ccs[i]))
	}
	return
}

// https://ffmpeg.org/doxygen/7.0/structAVFilterGraphSegment.html#ab7563eca151d89e693f6258de5ce0214
func (fgs *FilterGraphSegment) NbChains() int {
	return int(fgs.c.nb_chains)
}

func (fgs *FilterGraphSegment) CreateFilters(flags int) error {
	return newError(C.avfilter_graph_segment_create_filters(fgs.c, C.int(flags)))
}

func (fgs *FilterGraphSegment) ApplyOpts(flags int) error {
	return newError(C.avfilter_graph_segment_apply_opts(fgs.c, C.int(flags)))
}

func (fgs *FilterGraphSegment) Init(flags int) error {
	return newError(C.avfilter_graph_segment_init(fgs.c, C.int(flags)))
}

func (fgs *FilterGraphSegment) Link(flags int, inputs, outputs *FilterInOut) error {
	var ic **C.AVFilterInOut
	if inputs != nil {
		ic = &inputs.c
	}
	var oc **C.AVFilterInOut
	if outputs != nil {
		oc = &outputs.c
	}
	return newError(C.avfilter_graph_segment_link(fgs.c, C.int(flags), ic, oc))
}
