package tayloring

import (
	"github.com/fogleman/gg"
)

type point struct {
	X, Y float64
}

func DrawGrid(dc *gg.Context, W int, H int) {
	const Minor = 10
	const Major = 100

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// minor grid

	// for x := Minor; x < W; x += Minor {
	// 	fx := float64(x) + 0.5
	// 	dc.DrawLine(fx, 0, fx, float64(H))
	// }
	// for y := Minor; y < H; y += Minor {
	// 	fy := float64(y) + 0.5
	// 	dc.DrawLine(0, fy, float64(W), fy)
	// }
	// dc.SetLineWidth(1)
	// dc.SetRGBA(0, 0, 0, 0.25)
	// dc.Stroke()

	// major grid
	for x := Major; x < W; x += Major {
		fx := float64(x) + 0.5
		dc.DrawLine(fx, 0, fx, float64(H))
	}
	for y := Major; y < H; y += Major {
		fy := float64(y) + 0.5
		dc.DrawLine(0, fy, float64(W), fy)
	}
	dc.SetLineWidth(1)
	dc.SetRGBA(0, 0, 0, 0.5)
	dc.Stroke()
}
