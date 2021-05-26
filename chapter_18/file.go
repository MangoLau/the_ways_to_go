package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// （1）如何打开一个文件并读取：
	file, err := os.Open("../chapter_12/input.dat")
	if err != nil {
		log.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you  got acces to it?\n")
		return
	}
	defer file.Close()
	iReader := bufio.NewReader(file)
	for {
		str, err := iReader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Printf("The input was: %s", str)
	}



}

// cat （2）如何通过切片读写文件：
/*func cat(f *file.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, er := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n", f.String(), er.String)
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			if nw, ew := file.Stdout.Write(buf[:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n", f.String(), ew.String())
			}
		}
	}
}*/