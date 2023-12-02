package main

import "testing"

// Part 1

func TestNumbersAtTheEnds(t *testing.T) {
	result := GetCalibrationValue("1abc2")
	expected := 12

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

func TestNumbersBetweenAlphaChars(t *testing.T) {
	result := GetCalibrationValue("pqr3stu8vwx")
	expected := 38

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

func TestMultipleNumbers(t *testing.T) {
	result := GetCalibrationValue("a1b2c3d4e5f")
	expected := 15

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

func TestOneNumber(t *testing.T) {
	result := GetCalibrationValue("treb7uchet")
	expected := 77

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

// Part 2

func TestTwoCharNumbersBetweenInt(t *testing.T) {
	result := GetCalibrationValueWithChars("two1nine")
	expected := 29

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

func TestThreeNumChars(t *testing.T) {
	result := GetCalibrationValueWithChars("eightwothree")
	expected := 83

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

func TestRandomCharsAtEnds(t *testing.T) {
	result := GetCalibrationValueWithChars("abcone2threexyz")
	expected := 13

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

func TestRandomCharsAtStart(t *testing.T) {
	result := GetCalibrationValueWithChars("xtwone3four")
	expected := 24

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

func TestIntegersAtBothEnds(t *testing.T) {
	result := GetCalibrationValueWithChars("4nineeightseven2")
	expected := 42

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

func TestStartWithCharsEndsWithInt(t *testing.T) {
	result := GetCalibrationValueWithChars("zoneight234")
	expected := 14

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

func TestStartWithIntEndsWithChar(t *testing.T) {
	result := GetCalibrationValueWithChars("7pqrstsixteen")
	expected := 76

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}
