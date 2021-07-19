package cmd

import "flag"

type params struct {
	DownURL     string //下载地址
	Concurrency int    //并发数量
	Filename    string //新文件名
}

type set func(pms *params)

func withStrUrl(strUrl string) set {
	return func(pms *params) {
		pms.DownURL = strUrl
	}
}

func withConcurrency(c int) set {
	return func(pms *params) {
		pms.Concurrency = c
	}
}

func withFilename(f string) set {
	return func(pms *params) {
		pms.Filename = f
	}
}

func obtained(opts ...set) *params {
	var p = &params{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func ShowText() *params {
	strUrl := flag.String("url", "", "download url")
	con := flag.Int("con", 5, "Concurrency number")
	filename := flag.String("name", "", "download obtained file name")
	flag.Parse()

	var opts []set
	if a := *strUrl; a != "" {
		opts = append(opts, withStrUrl(a))
	}
	if b := *con; b != 0 {
		opts = append(opts, withConcurrency(b))
	}
	if c := *filename; c != "" {
		opts = append(opts, withFilename(c))
	}
	return obtained(opts...)
}
