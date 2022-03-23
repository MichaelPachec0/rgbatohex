package main

import (
	"fmt"
	"testing"
)

func TestPv_3Val(t *testing.T) {
	expected := "ffffff"
	input := []string{"255", "255", "255"}
	actual, err := procInt(input)
	if err != nil {
		t.Errorf("ERR: %s", err)
	}
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}
func TestPv_3ValBAD(t *testing.T) {
	input := []string{"256", "255", "256"}
	_, err := procInt(input)
	if err == nil {
		t.Errorf("Expected error: got pass")
	} else {
		fmt.Println(err)
	}
}
func TestPv_2ValBAD(t *testing.T) {
	input := []string{"255", "255"}
	_, err := preProcVal(input)
	if err == nil {
		t.Errorf("Expected error: got pass")
	} else {
		fmt.Println(err)
	}
}

func TestPpV_4Val(t *testing.T) {
	expected := "ffffffff"
	input := []string{"255", "255", "255", ".999"}
	actual, err := preProcVal(input)
	if err != nil {
		t.Errorf("ERR: %s", err)
	} else {
		if expected != actual {
			t.Errorf("Expected %s do not match actual %s", expected, actual)
		}
	}
}
func TestPpV_3Val(t *testing.T) {
	expected := "ffffff"
	input := []string{"255", "255", "255"}
	actual, err := preProcVal(input)
	if err != nil {
		t.Errorf("ERR: %s", err)
	}
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}

func TestLn(t *testing.T) {
	input := []string{"0", "255"}
	expected := []int64{0, 255}
	for i, v := range input {
		actual, err := limInt(v)
		if err != nil {
			t.Errorf("Expected error: got pass")
		} else {
			if actual != expected[i] {
				t.Errorf("Expected %d do not match actual %d", expected[i], actual)
			}
		}
	}
}
func TestLn_BAD(t *testing.T) {
	input := "32768"
	_, err := limInt(input)
	if err == nil {
		t.Errorf("Expected error: got pass")
	} else {
		fmt.Println(err)
	}
}
func TestPf(t *testing.T) {
	input := []string{".002", "1"}
	expected := "01ff"
	//expected := []int64{1, 255}
	actual, err := procFloat(input)
	if err != nil {
		t.Errorf("Err: %s", err)
	}
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}
func TestLf_BAD(t *testing.T) {
	input := "1.5"
	_, err := limFloat(input)
	if err == nil {
		t.Errorf("Expected error: got pass")
	} else {
		fmt.Println(err)
	}
}

func TestPf_BAD(t *testing.T) {
	input := []string{"n"}
	_, err := procFloat(input)
	if err == nil {
		t.Errorf("Expected error: got pass")
	} else {
		fmt.Println(err)
	}
}
func TestHelp(t *testing.T) {
	//mainCode([]string{})
	main()
}
