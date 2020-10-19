package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Identify file or directory
const (
	MFile      = iota
	MDirectory = iota
)

func insertItem(db *sql.DB, name string, size int64, modifiedDate time.Time, createdDate time.Time, t int, path string) (int64, error) {
	stmt, err := db.Prepare("INSERT item SET name=?, size=?, modified_date=?, created_date=?, type=?, path=?")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, size, modifiedDate, createdDate, t, path)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func scanDirectory(paths []string, db *sql.DB) (int, int) {
	v := ""
	fileCount := 0
	directoryCount := 0
	for i := 0; i < len(paths); i++ {
		v = paths[i]
		files, err := ioutil.ReadDir(v)
		if err != nil {
			log.Fatal(err)
		}

		for _, fileinfo := range files {
			if fileinfo.IsDir() {
				fmt.Printf("DIRECTORY:\nfull path:%v\nmod time:%v\nname:%v\nsize:%v\nsize string:%v\nmode:%v\n\n", v, fileinfo.ModTime(), fileinfo.Name(), fileinfo.Size(), fileinfo.Mode())
				paths = append(paths, path.Join(v, fileinfo.Name()))
				insertItem(db, fileinfo.Name(), fileinfo.Size(), fileinfo.ModTime(), time.Now(), MDirectory, v)
				directoryCount++
			} else {
				fmt.Printf("FILE:\nfull path:%v\nmod time:%v\nname:%v\nsize:%v\nsize string:%v\nmode:%v\n\n", path.Join(v, fileinfo.Name()), fileinfo.ModTime(), fileinfo.Name(), fileinfo.Size(), fileinfo.Mode())
				insertItem(db, fileinfo.Name(), fileinfo.Size(), fileinfo.ModTime(), time.Now(), MFile, v)
				fileCount++
			}
		}
	}

	return fileCount, directoryCount
}

func main() {
	fmt.Println("process started")
	overallStart := time.Now()

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ufiles")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	f, d := scanDirectory([]string{"/Users/marmoreno/Downloads"}, db)

	overallElapsed := time.Since(overallStart)
	fmt.Printf("process finished: %s\ntotal %v files in %v directories\n", overallElapsed, f, d)
}
