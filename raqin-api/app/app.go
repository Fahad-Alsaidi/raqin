package app

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

const (
	raqinAppPath = "raqin"
	booksPath    = "raqin/books"
	usersPath    = "raqin/user"
)

func Init() {
	CreateDirIfNotExist(raqinAppPath)
	CreateDirIfNotExist(booksPath)
	CreateDirIfNotExist(usersPath)
}

// CreateDirIfNotExist will create a directory in the current path
func CreateDirIfNotExist(dirname string) {
	_, err := os.Stat(dirname)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirname, 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
}

// CheckAndGetBooksPathIfExist will get a books directory in the defined path
func CheckAndGetBooksPathIfExist() (string, error) {
	var err error
	if _, err = os.Stat(booksPath); os.IsNotExist(err) {
		return "", err
	}
	return booksPath, err
}

// CheckAndGetUsersPathIfExist will get a users directory in the defined path
func CheckAndGetUsersPathIfExist() (string, error) {
	var err error
	if _, err = os.Stat(usersPath); os.IsNotExist(err) {
		return "", err
	}
	return usersPath, err
}

// CreateOrGetBookDir will create a book directory in a defined path
func CreateOrGetBookDir(dirname string) (string, error) {
	path := fmt.Sprintf("%s/book_%s", booksPath, dirname)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0700)
		if errDir != nil {
			return "", err
		}
		return path, nil

	}
	return path, nil
}

// CreateOrGetPagesDir will create a pages directory in a book path
func CreateOrGetPagesDir(bookPath string) (string, error) {
	path := fmt.Sprintf("%s/pages", bookPath)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0700)
		if errDir != nil {
			return "", err
		}
		return path, nil

	}
	return path, nil
}

// CreateUserDir will create a user directory in a defined path
func CreateUserDir(dirname string) (string, error) {
	path := fmt.Sprintf("%s/user_%s", usersPath, dirname)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0700)
		if errDir != nil {
			return "", err
		}
		return path, nil

	}
	return path, nil
}

func BookPagesToImages(srcFile, destDir string) (int, error) {

	doc, err := fitz.New(srcFile)
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(destDir, fmt.Sprintf("page_%d.jpg", n)))
		if err != nil {
			panic(err)
		}

		imgOpts := &jpeg.Options{Quality: jpeg.DefaultQuality}
		err = jpeg.Encode(f, img, imgOpts)
		if err != nil {
			panic(err)
		}

		f.Close()
	}

	return doc.NumPage(), nil
}
