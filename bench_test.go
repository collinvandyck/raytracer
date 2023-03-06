package rt

import "testing"

func BenchmarkMatrix4x4Get(b *testing.B) {
	m1 := NewMatrixFromTable(`
		+---------------------------+
		| -2.0 | -8.0 |  3.0 |  5.0 |
		| -3.0 |  1.0 |  7.0 |  3.0 |
		|  1.0 |  2.0 | -9.0 |  6.0 |
		| -6.0 |  7.0 |  7.0 | -9.0 |
		+---------------------------+
	`)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = m1.Get(2, 3)
	}
}

func BenchmarkMatrix4x4Minor(b *testing.B) {
	m1 := NewMatrixFromTable(`
		+---------------------------+
		| -2.0 | -8.0 |  3.0 |  5.0 |
		| -3.0 |  1.0 |  7.0 |  3.0 |
		|  1.0 |  2.0 | -9.0 |  6.0 |
		| -6.0 |  7.0 |  7.0 | -9.0 |
		+---------------------------+
	`)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = m1.Minor(0, 0)
	}
}

func BenchmarkMatrix4x4Cofactor(b *testing.B) {
	m1 := NewMatrixFromTable(`
		+---------------------------+
		| -2.0 | -8.0 |  3.0 |  5.0 |
		| -3.0 |  1.0 |  7.0 |  3.0 |
		|  1.0 |  2.0 | -9.0 |  6.0 |
		| -6.0 |  7.0 |  7.0 | -9.0 |
		+---------------------------+
	`)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = m1.Cofactor(0, 0)
	}
}

func BenchmarkMatrix4x4Inverse(b *testing.B) {
	m1 := NewMatrixFromTable(`
		+---------------------------+
		| -5.0 |  2.0 |  6.0 | -8.0 |
		|  1.0 | -5.0 |  1.0 |  8.0 |
		|  7.0 |  7.0 | -6.0 | -7.0 |
		|  1.0 | -3.0 |  7.0 |  4.0 |
		+---------------------------+
	`)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = m1.Inverse()
	}
}

func BenchmarkMatrix4x4Determinant(b *testing.B) {
	m1 := NewMatrixFromTable(`
		+---------------------------+
		| -2.0 | -8.0 |  3.0 |  5.0 |
		| -3.0 |  1.0 |  7.0 |  3.0 |
		|  1.0 |  2.0 | -9.0 |  6.0 |
		| -6.0 |  7.0 |  7.0 | -9.0 |
		+---------------------------+
	`)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = m1.Determinant()
	}
}

func BenchmarkMatrix4x4Submatrix(b *testing.B) {
	m1 := NewMatrixFromTable(`
			+---------------+
			| 1 | 2 | 3 | 4 |
			| 5 | 6 | 7 | 8 |
			| 9 | 8 | 7 | 6 |
			| 5 | 4 | 3 | 2 |
			+---------------+
		`)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m1.Submatrix(1, 2)
	}
}

func BenchmarkMatrixMultiply(b *testing.B) {
	m1 := NewMatrixFromTable(`
			+---------------+
			| 1 | 2 | 3 | 4 |
			| 5 | 6 | 7 | 8 |
			| 9 | 8 | 7 | 6 |
			| 5 | 4 | 3 | 2 |
			+---------------+
		`)
	m2 := NewMatrixFromTable(`
			+-----------------+
			| -2 | 1 | 2 | 3  |
			| 3  | 2 | 1 | -1 |
			| 4  | 3 | 6 | 5  |
			| 1  | 2 | 7 | 8  |
			+-----------------+
		`)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m1.Multiply(m2)
	}
}

func BenchmarkSphere(b *testing.B) {
	const canvasPixels = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var (
			wallSize  = Value(7)                              // how big the wall is
			pixelSize = wallSize / Value(canvasPixels)        // pixel size in world coordinates
			half      = wallSize / 2                          // half the wall size
			color     = NewColor(1, 0, 0)                     // color of the sphere
			canvas    = NewCanvas(canvasPixels, canvasPixels) // size of the canvas
			rayOrigin = NewPoint(0, 0, -5)                    // ray origin
			sphere    = NewSphere()                           // unit sphere
		)

		for y := 0; y < canvasPixels; y++ {
			// compute worldY (top = +half, bottom = -half)
			worldY := half - (float64(y) * pixelSize)

			for x := 0; x < canvasPixels; x++ {
				worldX := half - (float64(x) * pixelSize)
				point := NewPoint(worldX, worldY, 10) // the wall lives at z=10
				vector := point.SubtractPoint(rayOrigin).Normalize()
				ray := NewRay(rayOrigin, vector)
				xs := IntersectSphere(sphere, ray)
				_, hit := Hit(xs)
				if hit {
					canvas.WritePixel(x, y, color)
				}

			}
		}
	}
}
