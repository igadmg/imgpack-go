package imgpack

/*
#cgo CFLAGS: -I${SRCDIR}/external -std=gnu99 -Wno-unused-result

#define STB_RECT_PACK_IMPLEMENTATION
#include "stb_rect_pack.h"
*/
import "C"
import "unsafe"

/*
struct stbrp_node
{
   stbrp_coord  x,y;
   stbrp_node  *next;
};
*/
type PackContext struct {
	context C.stbrp_context
	nodes   []C.stbrp_node
	//width     int
	//height    int
	//align     int
	//init_mode int
	//heuristic int
	//num_nodes int
	//stbrp_node *active_head;
	//stbrp_node *free_head;
	//stbrp_node extra[2]; // we allocate two extra nodes so optimal user-node-count is 'width' not 'width+2'
}

func NewPackContext(width, height, size int) *PackContext {
	ctx := &PackContext{
		nodes: make([]C.stbrp_node, size),
	}
	C.stbrp_init_target(&ctx.context, C.int(width), C.int(height), unsafe.SliceData(ctx.nodes), C.int(len(ctx.nodes)))
	return ctx
}

func NewRect(id int, w, h int) C.stbrp_rect {
	return C.stbrp_rect{
		id: C.int(id),
		w:  C.stbrp_coord(w),
		h:  C.stbrp_coord(h),
		//x	:	_Ctype_stbrp_coord
		//y	:	_Ctype_stbrp_coord
	}
}

func (ctx *PackContext) PackRects(rect ...C.stbrp_rect) []C.stbrp_rect {
	C.stbrp_pack_rects(&ctx.context, unsafe.SliceData(rect), C.int(len(rect)))
	return rect
}
