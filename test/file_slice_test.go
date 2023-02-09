package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

const chunkSize = 1024 * 1024

func TestGenerateChunkFile(t *testing.T) {
	// read file info
	fileInfo, err := os.Stat("./image/shiroko.png")
	if err != nil {
		t.Fatal(err)
	}
	// read file
	chunkNum := (fileInfo.Size() + chunkSize - 1) / chunkSize
	file, err := os.OpenFile("./image/shiroko.png", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	// partition file
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		_, err := file.Seek(int64(i*chunkSize), 0)
		if err != nil {
			t.Fatal(err)
		}
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			// if last partition size is less than chunk size
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		_, err = file.Read(b)
		if err != nil {
			t.Fatal(err)
		}
		// write partition
		f, err := os.OpenFile("./image/"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		_, err = f.Write(b)
		if err != nil {
			t.Fatal(err)
		}
		err = f.Close()
		if err != nil {
			t.Fatal(err)
		}
	}
	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMergeChunkFile(t *testing.T) {
	// new file
	file, err := os.OpenFile("./image/shiroko2.png", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	// chunk number
	fileInfo, err := os.Stat("./image/shiroko.png")
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := (fileInfo.Size() + chunkSize - 1) / chunkSize
	// merge file
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("./image/"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		_, err = file.Write(b)
		if err != nil {
			t.Fatal(err)
		}
		_ = f.Close()
	}
	_ = file.Close()
}

func TestFileConsistency(t *testing.T) {
	// file1
	file1, err := os.OpenFile("./image/shiroko.png", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := io.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}
	// file2
	file2, err := os.OpenFile("./image/shiroko2.png", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := io.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}
	// md5
	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))
	println(s1)
	println(s2)
	println(s1 == s2)
}
