// In this session , we will understand implicit implementation in interface
package main

import "fmt"

type Storer interface {
	Save(filename string, data string) error
}

type S3Bucket struct {
	Region string
}

func (s S3Bucket) Save(file string, data string) error {
	fmt.Printf("[S3 - %s] Saving %s ...\n", s.Region, file)
	return nil
}

type LocalDisk struct {
	path string
}

func (id LocalDisk) Save(file string, data string) error {
	fmt.Printf("[LocalDisk] Writing  %s to %s...\n", file, id.path)
	return nil
}

func UploadConfig(s Storer) {
	fmt.Println("Starting Upload")
	err := s.Save("config.json", "{status:active}")
	if err != nil {
		fmt.Println("Upload failed")
	}

}

func main() {
	aws := S3Bucket{Region: "utc"}
	drive := LocalDisk{path: "/var/data"}

	UploadConfig(aws)
	UploadConfig(drive)

	DescribeAnything(aws)
}

func DescribeAnything(it interface{}) {
	fmt.Printf("\nArchitect Audit -> Value: %+v | Type: %T\n", it, it)
}
