package utils

import (
	"testing"
)

func TestIsEmpty1(t *testing.T) {
	if !IsEmpty("") {
		t.Error("Wrong return value, was expecting true but false")
	}
}

func TestIsEmpty3(t *testing.T) {
	if IsEmpty("a") {
		t.Error("Wrong return value, was expecting false but true")
	}
}

func TestIsStringLengthWithin1(t *testing.T) {
	if !IsStringLengthWithin("z", 1, 10) {
		t.Error("Wrong return value, was expecting true but false")
	}
}

func TestIsStringLengthWithin2(t *testing.T) {
	if !IsStringLengthWithin("abcABC1234", 1, 10) {
		t.Error("Wrong return value, was expecting true but false")
	}
}

func TestIsStringLengthWithin3(t *testing.T) {
	if !IsStringLengthWithin("", 1, 10) {
		t.Error("Wrong return value, was expecting true but false")
	}
}

func TestIsStringLengthWithin5(t *testing.T) {
	if IsStringLengthWithin("abcABC12345", 1, 10) {
		t.Error("Wrong return value, was expecting false but true")
	}
}

func TestIsStringLengthWithin6(t *testing.T) {
	if IsStringLengthWithin("a", 10, 1) {
		t.Error("Wrong return value, was expecting false but true")
	}
}

func TestIsHankakuAklphaNum1(t *testing.T) {
	if !IsHankakuAlphaNum("abcABC123") {
		t.Error("Wrong return value, was expecting true but false")
	}
}

func TestIsHankakuAklphaNum2(t *testing.T) {
	if !IsHankakuAlphaNum("") {
		t.Error("Wrong return value, was expecting true but false")
	}
}

func TestIsHankakuAklphaNum4(t *testing.T) {
	if IsHankakuAlphaNum("@") {
		t.Error("Wrong return value, was expecting false but true")
	}
}

func TestIsHankakuAklphaNum5(t *testing.T) {
	if IsHankakuAlphaNum("ａＡ１") {
		t.Error("Wrong return value, was expecting false but true")
	}
}
