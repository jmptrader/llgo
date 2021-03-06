// RUN: llgo -o %t %s
// RUN: %t > %t1 2>&1
// RUN: go run %s > %t2 2>&1
// RUN: diff -u %t1 %t2

package main

type A struct {
	b1, b2 B
}

type B struct {
	a1, a2 *A
}

func main() {
	var a A
	_ = a
}
