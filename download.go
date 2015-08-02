package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Download(uri, dest string) error {
	res, err := compareMd5(uri, dest)
	if err != nil {
		return err
	}
	if res == true {
		return download(uri, dest)
	}
	fmt.Println("Download Skipped (up-to-date) -> " + dest)
	return nil
}

func compareMd5(uri, dest string) (bool, error) {
	if _, err := os.Stat(dest); err == nil {
		repoMd5, err := downloadMd5(uri + ".md5")
		if err != nil {
			return false, err
		}
		localMd5, err := ComputeMd5(dest)
		if err != nil {
			return false, err
		}
		return strings.EqualFold(localMd5, repoMd5), nil
	}
	return true, nil
}

func downloadMd5(uri string) (string, error) {
	c, err := downloadAsString(uri)
	if err != nil {
		return "", err
	}
	return strings.Fields(c)[0], nil
}

func DownloadAsString(uri string) (string, error) {
	return downloadAsString(uri)
}

func downloadAsString(uri string) (string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fatal("Download failed (" + uri + "): " + strconv.Itoa(resp.StatusCode))
	}
	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(c), nil
}

func download(uri, dest string) error {
	fmt.Println("Download: " + uri + " -> " + dest)
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("Download failed (" + uri + "): " + strconv.Itoa(resp.StatusCode))
	}
	tmpDest := dest + "__temp"
	out, err := os.Create(tmpDest)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err == nil {
		err = os.Rename(tmpDest, dest)
		if err != nil {
			os.Remove(tmpDest)	
		}
	} else {
		os.Remove(tmpDest)
	}
	return err
}

func ComputeMd5(filePath string) (string, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return string(result), err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return string(result), err
	}
	return string(hash.Sum(result)), nil
}
