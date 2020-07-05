package lod

/*
这个例子实现了简化版的搜索引擎爬取网页的功能。
代码中包含三个主要的类。其中，NetworkTransporter 类负责底层网络通信，根据请求获取数据；
HtmlDownloader 类用来通过 URL 获取网页；
Document 表示网页文档，后续的网页内容抽取、分词、索引都是以此为处理对象
*/

//HTMLRequest 请求
type HTMLRequest struct {
	url string
}

//NetworkTransporter 传输
type NetworkTransporter struct {
}

func (n *NetworkTransporter) send(r *HTMLRequest) {
}

//HTMLDownloader 获取网页
type HTMLDownloader struct {
	transporter NetworkTransporter
}

func (hd *HTMLDownloader) download(url string) {
	r := &HTMLRequest{
		url: url,
	}
	hd.transporter.send(r)
}
