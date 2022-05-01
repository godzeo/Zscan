package options

import (
	"Zscan/core/logger"
	"errors"
	"flag"
	"net/url"
)

type Options struct {
	Target              multiStringFlag
	FileTargets         string
	Output              string
	ProxyURL            string
	TimeOut             int
	JSON                bool
	Verbose             bool
	OutputStatusCode    bool
	OutputWithNoColor   bool
	OutputContentLength bool
	OutputTitle         bool
	OutputIP            bool
	OutputFingerPrint   bool
	RateLimit           int
	Stdin               bool // goland调试程序stdin有问题，所以加上这个参数
	OutputCDN           bool
	NoCidr              bool
	Method              string
	Templates           multiStringFlag
}
type multiStringFlag []string

func (m *multiStringFlag) String() string {
	return ""
}

func (m *multiStringFlag) Set(value string) error {
	*m = append(*m, value)
	return nil
}
func ParseOptions() *Options {
	options := &Options{}
	flag.Var(&options.Target, "target", "目标,可以指定多个目标 eg:-target xx.com -target aa.com")
	flag.StringVar(&options.FileTargets, "l", "", "目标地址的列表")
	flag.StringVar(&options.Output, "o", "", "输出的文件")
	flag.IntVar(&options.TimeOut, "timeout", 30, "超时时间(s)")
	flag.StringVar(&options.ProxyURL, "proxy-url", "", "URL of the proxy server")
	//flag.BoolVar(&options.Silent, "silent", false, "静默模式，启用后将静默其他非结果类输出")
	flag.BoolVar(&options.Verbose, "verbose", false, "输出更多调试信息")
	flag.BoolVar(&options.OutputStatusCode, "status-code", false, "Extracts status code")
	flag.BoolVar(&options.OutputContentLength, "content-length", false, "Extracts content length")
	flag.BoolVar(&options.OutputTitle, "title", false, "Extracts title")
	flag.BoolVar(&options.OutputIP, "ip", false, "Extracts ip")
	flag.IntVar(&options.RateLimit, "limit", 200, "限制每秒的并发数量")
	flag.BoolVar(&options.OutputWithNoColor, "nocolor", false, "不输出颜色")
	flag.BoolVar(&options.Stdin, "stdin", false, "开启后程序将接受stdin的输入")
	flag.BoolVar(&options.OutputCDN, "cdn", false, "检测目标是否含有CDN")
	flag.BoolVar(&options.OutputFingerPrint, "fingerprint", false, "输出指纹识别结果")
	flag.BoolVar(&options.NoCidr, "nocidr", false, "开启后不解析cidr")
	flag.StringVar(&options.Method, "method", "scan", "模式")
	flag.Var(&options.Templates, "poc", "指定poc目录或单个poc文件")
	flag.Parse()

	options.configureOutput()

	err := options.validateOptions()
	if err != nil {
		logger.Fatalf("Program exiting: %s\n", err)
	}
	return options
}

// validateOptions validates the configuration options passed
func (options *Options) validateOptions() error {
	// Validate proxy options if provided
	err := validateProxyURL(
		options.ProxyURL,
		"invalid http proxy format (It should be http://username:password@host:port)",
	)
	if err != nil {
		return err
	}
	return nil
}
func validateProxyURL(proxyURL, message string) error {
	if proxyURL != "" && !isValidURL(proxyURL) {
		return errors.New(message)
	}

	return nil
}

func isValidURL(urlString string) bool {
	_, err := url.Parse(urlString)

	return err == nil
}

// configureOutput configures the output on the screen
func (options *Options) configureOutput() {
	// If the user desires verbose output, show verbose output
	if options.Verbose {
		logger.MaxLevel = logger.Verbose
	}
	//if options.Silent {
	//	gologger.MaxLevel = gologger.Silent
	//	options.OutputWithNoColor = true
	//}
}
