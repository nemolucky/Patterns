package proxy

type Video struct {
	Name string
}

type Downloader interface {
	Download(name string) *Video
}

type YouTubeDownloader struct{}

func (y *YouTubeDownloader) Download(name string) *Video {
	return &Video{Name: name}
}

type ProxyDownloader struct {
	downloader Downloader
	downloads  map[string]*Video
}

func NewProxyDownloader(downloader Downloader) *ProxyDownloader {
	return &ProxyDownloader{
		downloader: downloader,
		downloads:  make(map[string]*Video),
	}
}

func (p *ProxyDownloader) Download(name string) *Video {
	if video, ok := p.downloads[name]; ok {
		return video
	} else {
		newVideo := p.downloader.Download(name)
		p.downloads[name] = newVideo
		return newVideo
	}
}
