package down

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

type downer struct {
	concurrency int
}

func New(concurrency int) *downer {
	return &downer{concurrency: concurrency}
}

func (d *downer) Download(strUrl, filename string) error {
	if strUrl == "" {
		return errors.New("download url not empty")
	}
	if filename == "" {
		filename = path.Base(strUrl)
	}
	resp, err := http.Head(strUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		if resp.Header.Get("Accept-Ranges") == "bytes" {
			return d.multiDownload(strUrl, filename, int(resp.ContentLength))
		}
	}

	return d.signalDownload(strUrl, filename)
}

func (d *downer) signalDownload(strUrl, filename string) error {

	resp, err := http.Get(strUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return save(resp.Body, filename)
}

func (d *downer) multiDownload(strUrl, filename string, contentLen int) error {

	size := contentLen / d.concurrency

	root := obtFileName(filename)
	if err := os.MkdirAll(root, 0777); err != nil {
		return err
	}
	defer os.RemoveAll(root)

	var wg sync.WaitGroup

	start := 0

	for i := 0; i < d.concurrency; i++ {
		wg.Add(1)
		go func(start, idx int) {
			defer wg.Done()
			end := start + size
			if idx+1 == d.concurrency {
				end = contentLen
			}
			if err := d.down(strUrl, filename, start, end, idx); err != nil {
				log.Fatal(err)
			}
		}(start, i)

		start += size + 1
	}

	wg.Wait()
	return d.merge(filename)
}

func (d *downer) merge(filename string) error {

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := 0; i < d.concurrency; i++ {
		name := obtCacheFileName(filename, i)
		tf, err := os.Open(name)
		if err != nil {
			return err
		}
		if _, err := io.Copy(f, tf); err != nil {
			return err
		}
		tf.Close()
		if err := os.Remove(name); err != nil {
			return err
		}
	}

	return nil
}

func (d *downer) down(strUrl, name string, start, end, idx int) error {
	if start >= end {
		return errors.New("start must less then end")
	}

	req, err := http.NewRequest(http.MethodGet, strUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes:%d-%d", start, end))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return save(resp.Body, obtCacheFileName(name, idx))
}

func obtFileName(filename string) string {
	return strings.SplitN(filename, ".", 2)[0]
}

func obtCacheFileName(filename string, idx int) string {
	root := obtFileName(filename)
	return fmt.Sprintf("%s%s%s-%d", root, string(os.PathSeparator), filename, idx)
}

func save(reader io.ReadCloser, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 32*024)
	_, err = io.CopyBuffer(f, reader, buf)
	return err
}
