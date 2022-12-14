// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !amd64 || noasm || gccgo || safe
// +build !amd64 noasm gccgo safe

package c128

import "math/cmplx"

// AxpyUnitary is
//
//	for i, v := range x {
//		y[i] += alpha * v
//	}
func AxpyUnitary(alpha complex128, x, y []complex128) {
	for i, v := range x {
		y[i] += alpha * v
	}
}

// AxpyUnitaryTo is
//
//	for i, v := range x {
//		dst[i] = alpha*v + y[i]
//	}
func AxpyUnitaryTo(dst []complex128, alpha complex128, x, y []complex128) {
	for i, v := range x {
		dst[i] = alpha*v + y[i]
	}
}

// AxpyInc is
//
//	for i := 0; i < int(n); i++ {
//		y[iy] += alpha * x[ix]
//		ix += incX
//		iy += incY
//	}
func AxpyInc(alpha complex128, x, y []complex128, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		y[iy] += alpha * x[ix]
		ix += incX
		iy += incY
	}
}

// AxpyIncTo is
//
//	for i := 0; i < int(n); i++ {
//		dst[idst] = alpha*x[ix] + y[iy]
//		ix += incX
//		iy += incY
//		idst += incDst
//	}
func AxpyIncTo(dst []complex128, incDst, idst uintptr, alpha complex128, x, y []complex128, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		dst[idst] = alpha*x[ix] + y[iy]
		ix += incX
		iy += incY
		idst += incDst
	}
}

// DscalUnitary is
//
//	for i, v := range x {
//		x[i] = complex(real(v)*alpha, imag(v)*alpha)
//	}
func DscalUnitary(alpha float64, x []complex128) {
	for i, v := range x {
		x[i] = complex(real(v)*alpha, imag(v)*alpha)
	}
}

// DscalInc is
//
//	var ix uintptr
//	for i := 0; i < int(n); i++ {
//		x[ix] = complex(real(x[ix])*alpha, imag(x[ix])*alpha)
//		ix += inc
//	}
func DscalInc(alpha float64, x []complex128, n, inc uintptr) {
	var ix uintptr
	for i := 0; i < int(n); i++ {
		x[ix] = complex(real(x[ix])*alpha, imag(x[ix])*alpha)
		ix += inc
	}
}

// ScalInc is
//
//	var ix uintptr
//	for i := 0; i < int(n); i++ {
//		x[ix] *= alpha
//		ix += incX
//	}
func ScalInc(alpha complex128, x []complex128, n, inc uintptr) {
	var ix uintptr
	for i := 0; i < int(n); i++ {
		x[ix] *= alpha
		ix += inc
	}
}

// ScalUnitary is
//
//	for i := range x {
//		x[i] *= alpha
//	}
func ScalUnitary(alpha complex128, x []complex128) {
	for i := range x {
		x[i] *= alpha
	}
}

// DotcUnitary is
//
//	for i, v := range x {
//		sum += y[i] * cmplx.Conj(v)
//	}
//	return sum
func DotcUnitary(x, y []complex128) (sum complex128) {
	for i, v := range x {
		sum += y[i] * cmplx.Conj(v)
	}
	return sum
}

// DotcInc is
//
//	for i := 0; i < int(n); i++ {
//		sum += y[iy] * cmplx.Conj(x[ix])
//		ix += incX
//		iy += incY
//	}
//	return sum
func DotcInc(x, y []complex128, n, incX, incY, ix, iy uintptr) (sum complex128) {
	for i := 0; i < int(n); i++ {
		sum += y[iy] * cmplx.Conj(x[ix])
		ix += incX
		iy += incY
	}
	return sum
}

// DotuUnitary is
//
//	for i, v := range x {
//		sum += y[i] * v
//	}
//	return sum
func DotuUnitary(x, y []complex128) (sum complex128) {
	for i, v := range x {
		sum += y[i] * v
	}
	return sum
}

// DotuInc is
//
//	for i := 0; i < int(n); i++ {
//		sum += y[iy] * x[ix]
//		ix += incX
//		iy += incY
//	}
//	return sum
func DotuInc(x, y []complex128, n, incX, incY, ix, iy uintptr) (sum complex128) {
	for i := 0; i < int(n); i++ {
		sum += y[iy] * x[ix]
		ix += incX
		iy += incY
	}
	return sum
}
