package main

import "testing"

func TestCRC16(t *testing.T) {
	data := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	var expected uint16 = 39210
	var actual uint16 = CRC16(data)

	if actual != expected {
		t.Fatalf("Expected %d Actual %d\n", expected, actual)
	}
}

func TestCRC16Constant(t *testing.T) {
	data := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	var expected uint16 = 43803
	var actual uint16 = CRC16Constant(data, 13)

	if actual != expected {
		t.Fatalf("Expected %d Actual %d\n", expected, actual)
	}
}
