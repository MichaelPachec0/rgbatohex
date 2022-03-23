/* Copyright (c) 2022, Michael Pacheco. All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary until re-licensed
 * Written by Michael Pacheco <michaelpacheco@protonmail.com>, 2022 February
 */
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func procInt(vRange []string) (string, error) {
	var ret string
	for _, v := range vRange {
		// ParseInt spits out signed integers
		pValue, err := limInt(v)
		if err != nil {
			return "", err
		}
		ret += fmt.Sprintf("%02x", pValue)
	}
	return ret, nil
}
func limInt(sv string) (int64, error) {
	pValue, err := strconv.ParseInt(sv, 10, 16)
	if err != nil {
		return 0, err
	} else if pValue >= 256 {
		return 0, fmt.Errorf("EVAL: %s is higher than 256", sv)
	} else {
		return pValue, nil
	}
}

func limFloat(sv string) (int64, error) {
	sVal, err := strconv.ParseFloat(sv, 64)
	if err != nil {
		return 0, err
	}
	// converting alpha in % to hex works out to % * 255
	pv, err := limInt(strconv.FormatInt(int64(math.Round(sVal*255)), 10))
	if err != nil {
		return 0, err
	}
	return pv, nil
	//pv, error := limInt()
}

func procFloat(vRange []string) (string, error) {
	var ret string
	for _, v := range vRange {
		pv, err := limFloat(v)
		if err != nil {
			return "", err
		}
		ret += fmt.Sprintf("%02x", pv)
	}
	return ret, nil
}

func preProcVal(vRange []string) (string, error) {
	switch len(vRange) {
	case 3:
		return procInt(vRange)
	case 4:
		{
			var ret string
			var err error
			ret, err = procInt(vRange[:len(vRange)-1])
			if err != nil {
				return "", err
			}
			ret2, err := procFloat([]string{vRange[len(vRange)-1]})
			if err != nil {
				return "", err
			}
			ret += ret2
			return ret, nil
		}
	default:
		return "", fmt.Errorf("EVAL")
	}
}

func help() {
	//command := strings.Split(os.Args[0], string(os.PathSeparator))
	command := os.Args[0][strings.LastIndex(os.Args[0], string(os.PathSeparator))+1:]
	ret := fmt.Sprintf("Usage:  %s r,g,b[,a]\n\t%s r g b [a]", command, command)
	fmt.Println(ret)
}

func mainCode(args []string) {
	switch len(args) {
	case 1:
		{
			help()
		}
	case 2:
		{
			switch strings.Count(os.Args[1], ",") {
			case 2:
				fallthrough
			case 3:
				{
					//r,g,b,a
					vRange := strings.Split(os.Args[1], ",")
					pv, err := preProcVal(vRange)
					if err != nil {
						if pErr, ok := err.(*strconv.NumError); ok {
							fmt.Printf("%s  \"%s\"\n", pErr.Err, os.Args[1])
						} else {
							fmt.Println(err)
						}
						help()
					}
					fmt.Println(pv)
				}
			default:
				{
					//help menu
					help()
				}
			}
		}
	// command r g b
	case 4:
		fallthrough
	// command r g b a
	case 5:
		{
			sv, err := preProcVal(os.Args[1:])
			if err != nil {
				if pErr, ok := err.(*strconv.NumError); ok {
					fmt.Printf("%s  \"%s\"\n", pErr.Err, os.Args[1])
				} else {
					fmt.Println(err)
				}
				help()
			}
			fmt.Println(sv)
		}
	default:
		help()
	}

}

func main() {
	mainCode(os.Args)
}
