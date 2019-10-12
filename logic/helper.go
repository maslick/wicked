package logic

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func listAllTextFiles() []string {
	dirname := "." + string(filepath.Separator) + "txt" + string(filepath.Separator)

	d, _ := os.Open(dirname)
	defer d.Close()

	files, _ := d.Readdir(-1)

	var filez []string
	for _, file := range files {
		if file.Mode().IsRegular() {
			var ext = filepath.Ext(file.Name())
			if ext == ".txt" {
				filez = append(filez, file.Name())
			}
		}
	}

	return filez
}

func getPageTitles() []string {
	var files = listAllTextFiles()
	var titles []string
	for _, file := range files {
		titles = append(titles, file[:len(file)-4])
	}
	return titles
}

func loadPage(title string) (*Page, error) {
	filename := "txt/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func (p *Page) Save() error {
	_ = os.Mkdir("txt", os.ModePerm)
	filename := "txt/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
