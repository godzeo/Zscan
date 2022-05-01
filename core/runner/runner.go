package runner

import (
	banner "Zscan/core/banner"
	"Zscan/core/httpx"
	"Zscan/core/logger"
	options "Zscan/core/option"
	"Zscan/core/util"
	"bufio"
	"github.com/projectdiscovery/hmap/store/hybrid"
	"go.uber.org/ratelimit"
	"os"
	"strings"
	"time"
)

type Runner struct {
	Options     *options.Options
	rateLimiter ratelimit.Limiter
	hm          *hybrid.HybridMap
	httpnew     *httpx.HTTPX
	banner      *banner.BannerPrints
}
type Zscan struct {
	Name     string `json:"name"`
	FingerId string
	Extra    map[string]string `json:"extra,omitempty"`
}

func New(options *options.Options) (*Runner, error) {

	logger.Infof("Runner New()")

	runner := &Runner{
		Options: options,
	}

	// 发包速度限制
	runner.rateLimiter = ratelimit.New(options.RateLimit)

	targetNum := 0
	// 处理单个目标
	if options.Target != nil {
		for _, t := range options.Target {
			if util.IsCidr(t) {
				cidrIps, err := IPAddresses(t)
				if err != nil {
					//runner.hm.Set(t, nil)
					targetNum++
				} else {
					for _, ip := range cidrIps {
						ip = ip
						//runner.hm.Set(ip, nil)
						targetNum++
					}
				}
			} else {
				//err := runner.hm.Set(t, nil)
				//if err != nil {
				//	return nil
				//}
				targetNum++
			}
		}
	}

	// 处理文件
	var f *os.File
	if options.FileTargets != "" {
		var err error
		// 打开文件
		f, err = os.Open(options.FileTargets)
		if err != nil {
			logger.Fatalf("Could not open targets file '%s': %s\n", options.FileTargets, err)
		}
		defer f.Close()
	} else if util.HasStdin() && !options.Stdin {
		// 标准输入
		f = os.Stdin
	}
	if f != nil {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			// 获取目标，逐行读取
			urlline := strings.TrimSpace(scanner.Text())
			// skip empty lines
			if urlline == "" {
				continue
			}
			// nolint:errcheck // ignoring error
			if util.IsCidr(urlline) {
				cidrIps, err := IPAddresses(urlline)
				if err != nil {
					//runner.hm.Set(urlline, nil)
					targetNum++
				} else {
					for _ = range cidrIps {
						//runner.hm.Set(ip, nil)
						targetNum++
					}
				}
			} else {
				//runner.hm.Set(urlline, nil)
				targetNum++
			}
		}
	}

	if targetNum == 0 {
		logger.Fatalf("没有指定输入，合法url数量为 0 ，请检查输入")
	}

	logger.Infof("加载目标总数:%d", targetNum)

	// 借助HTTPX，初始化
	httpOptions := &httpx.Options{
		Timeout:          time.Duration(options.TimeOut) * time.Second,
		RetryMax:         3,
		FollowRedirects:  true,
		HTTPProxy:        options.ProxyURL,
		Unsafe:           false,
		DefaultUserAgent: httpx.GetRadnomUserAgent(),
	}

	httpnew, err := httpx.New(httpOptions)
	if err != nil {
		return nil, err
	}
	runner.httpnew = httpnew

	// 初始化指纹库
	bannerprints, err := banner.InitBanner()
	if err != nil {
		return nil, err
	}
	runner.banner = &bannerprints
	logger.Infof("初始化Banner")

	return runner, nil

}
