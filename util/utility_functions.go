package util

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

//read a file formatted as a hash table into a map and return the map
func GenFileHash(path string) (map[string]string, error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
		err    error
	)
	lines := make(map[string]string)
	if file, err = os.Open(path); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			re := regexp.MustCompile(" = ").Split(buffer.String(), 2)
			lines[re[0]] = re[1]
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return lines, err
}

//convert the values of a map from strings to ints
func Str2NumHash(str_hash map[string]string) map[string]int {
	num_hash := make(map[string]int)
	for key, val := range str_hash {
		num_hash[key], _ = strconv.Atoi(val)
	}
	return num_hash
}

//convert an int slice to a comma separated list
func SliceToCSL(input []int) string {
	if len(input) == 0 {
		return " "
	}
	var output string
	for _, i := range input {
		output += "," + strconv.Itoa(i)
	}
	return output[1:]
}

//remove duplicate entries from an int slice
//runs in linear time and preserves the order of the given slice - stable
func RemoveSliceDuplicates(input []int) []int {
	var output []int
	duplicates := make(map[int]int)
	for _, elt := range input {
		if _, err := duplicates[elt]; err {
			continue
		} else {
			duplicates[elt] = 1
			output = append(output, elt)
		}
	}
	return output
}
