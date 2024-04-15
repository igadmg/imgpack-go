package imgpack

/*
#cgo CFLAGS: -I${SRCDIR}/external -std=gnu99 -Wno-unused-result

#define STBRP_STATIC
#define STB_RECT_PACK_IMPLEMENTATION
#include "stb_rect_pack.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type PackContext struct {
	context C.stbrp_context
	nodes   []C.stbrp_node
}

type PackRect = C.stbrp_rect

func (r PackRect) Id() int {
	return int(r.id)
}

func (r PackRect) X() int {
	return int(r.x)
}

func (r PackRect) Y() int {
	return int(r.y)
}

func (r PackRect) W() int {
	return int(r.w)
}

func (r PackRect) H() int {
	return int(r.h)
}

func PackRects(width, height int, rects ...PackRect) ([]PackRect, error) {
	ctx := NewPackContext(width, height, len(rects))
	return ctx.PackRects(rects...)
}

func NewPackContext(width, height, size int) *PackContext {
	ctx := &PackContext{
		nodes: make([]C.stbrp_node, size),
	}
	C.stbrp_init_target(&ctx.context, C.int(width), C.int(height), unsafe.SliceData(ctx.nodes), C.int(len(ctx.nodes)))
	return ctx
}

func NewRect(id int, w, h int) PackRect {
	return PackRect{
		id: C.int(id),
		w:  C.stbrp_coord(w),
		h:  C.stbrp_coord(h),
	}
}

func (ctx *PackContext) PackRects(rects ...PackRect) ([]PackRect, error) {
	p := runtime.Pinner{}
	defer p.Unpin()

	crects := unsafe.SliceData(rects)
	p.Pin(ctx.context.active_head)
	p.Pin(ctx.context.free_head)
	res := C.stbrp_pack_rects(&ctx.context, crects, C.int(len(rects)))

	if res == 0 {
		return nil, fmt.Errorf("failed to pack rects")
	}
	return rects, nil
}
