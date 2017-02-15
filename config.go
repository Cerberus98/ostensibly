package air

import (
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// Config is a global set of configs that for an instance of the `Air` for customization.
type Config struct {
	// AppName represents the name of the `Air` instance.
	//
	// The default Value is "air".
	//
	// It's called "app_name" in the config file.
	AppName string

	// DebugMode indicates whether to enable the debug mode when the HTTP server is started. It
	// works only with the default `Logger`.
	//
	// The default value is false.
	//
	// It's called "debug_mode" in the config file.
	DebugMode bool

	// LogEnabled indicates whether to enable the `Logger` when the HTTP server is started. It
	// works only with the default `Logger`.
	//
	// It will be forced to the true if the `DebugMode` is true.
	//
	// The default value is false.
	//
	// It's called "log_enabled" in the config file.
	LogEnabled bool

	// LogFormat represents the format of the output content of the `Logger`. It works only with
	// the default `Logger`.
	//
	// The default value is:
	// `{"app_name":"{{.app_name}}","time":"{{.time_rfc3339}}","level":"{{.level}}",` +
	// `"file":"{{.short_file}}","line":"{{.line}}"}`
	//
	// It's called "log_format" in the config file.
	LogFormat string

	// Address represents the TCP address that the HTTP server to listen on.
	//
	// The default value is "localhost:2333".
	//
	// It's called "address" in the config file.
	Address string

	// ReadTimeout represents the maximum duration before timing out read of the HTTP request.
	//
	// The default value is 0.
	//
	// It's called "read_timeout" in the config file.
	//
	// **It's unit in the config file is MILLISECONDS.**
	ReadTimeout time.Duration

	// WriteTimeout represents the maximum duration before timing out write of the HTTP
	// response.
	//
	// The default value is 0.
	//
	// It's called "write_timeout" in the config file.
	//
	// **It's unit in the config file is MILLISECONDS.**
	WriteTimeout time.Duration

	// MaxHeaderBytes represents the maximum number of bytes the HTTP server will read parsing
	// the HTTP request header's keys and values, including the HTTP request line. It does not
	// limit the size of the HTTP request body.
	//
	// The default value is 1048576.
	//
	// It's called "max_header_bytes" in the config file.
	MaxHeaderBytes int

	// TLSCertFile represents the path of the TLS certificate file.
	//
	// The default value is "".
	//
	// It's called "tls_cert_file" in the config file.
	TLSCertFile string

	// TLSKeyFile represents the path of the TLS key file.
	//
	// The default value is "".
	//
	// It's called "tls_key_file" in the config file.
	TLSKeyFile string

	// TemplateRoot represents the root directory of the HTML templates. It will be parsed into
	// the `Renderer`. It works only with the default `Renderer`.
	//
	// The default value is "templates" that means a subdirectory of the runtime directory.
	//
	// It's called "template_root" in the config file.
	TemplateRoot string

	// TemplateExt represents the file name extension of the HTML templates. It will be used
	// when parsing the HTML templates. It works only with the default `Renderer`.
	//
	// The default value is ".html".
	//
	// It's called "template_ext" in the config file.
	TemplateExt string

	// TemplateLeftDelim represents the left side of the HTML template delimiter. It will be
	// used when parsing the HTML templates. It works only with the default `Renderer`.
	//
	// The default value is "{{".
	//
	// It's called "template_left_delim" in the config file.
	TemplateLeftDelim string

	// TemplateRightDelim represents the right side of the HTML template delimiter. It will be
	// used when parsing the HTML templates. It works only with the default `Renderer`.
	//
	// The default value is "}}".
	//
	// It's called "template_right_delim" in the config file.
	TemplateRightDelim string

	// TemplateMinified indicates whether to minify the HTML templates before they being parsed
	// into the `Renderer`. It works only with the default `Renderer`.
	//
	// The default value is false.
	//
	// It's called "template_minified" in the config file.
	TemplateMinified bool

	// TemplateWatched indicates whether to watch the changing of the HTML templates after they
	// are parsed into the `Renderer`. It works only with the default `Renderer`.
	//
	// It will be forced to the true if the `DebugMode` is true.
	//
	// The default value is false.
	//
	// It's called "template_watched" in the config file.
	TemplateWatched bool

	// Data represents the data that parsing from the config file. You can use it to access the
	// values in the config file.
	//
	// e.g. Data["foobar"] will accesses the value in the config file called "foobar".
	Data Map
}

// DefaultConfig is the default instance of the `Config`.
var DefaultConfig = Config{
	AppName: "air",
	LogFormat: `{"app_name":"{{.app_name}}","time":"{{.time_rfc3339}}","level":"{{.level}}",` +
		`"file":"{{.short_file}}","line":"{{.line}}"}`,
	Address:            "localhost:2333",
	MaxHeaderBytes:     1 << 20,
	TemplateRoot:       "templates",
	TemplateExt:        ".html",
	TemplateLeftDelim:  "{{",
	TemplateRightDelim: "}}",
}

// NewConfig returns a pointer of a new instance of the `Config` by parsing the config file found in
// the filename path. It returns a copy of the DefaultConfig if the config file does not exist.
func NewConfig(filename string) *Config {
	c := DefaultConfig
	c.ParseFile(filename)
	return &c
}

// Parse parses the src into the c.
func (c *Config) Parse(src string) error {
	if err := yaml.Unmarshal([]byte(src), &c.Data); err != nil {
		return err
	}
	c.fillData()
	return nil
}

// ParseFile parses the config file found in the filename path into the c.
func (c *Config) ParseFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	c.Parse(string(b))
	return nil
}

// fillData fills the values of the fields from the field `Data` of the c.
func (c *Config) fillData() {
	if an, ok := c.Data["app_name"]; ok {
		c.AppName = an.(string)
	}
	if dm, ok := c.Data["debug_mode"]; ok {
		c.DebugMode = dm.(bool)
	}
	if le, ok := c.Data["log_enabled"]; ok {
		c.LogEnabled = le.(bool)
	}
	if lf, ok := c.Data["log_format"]; ok {
		c.LogFormat = lf.(string)
	}
	if addr, ok := c.Data["address"]; ok {
		c.Address = addr.(string)
	}
	if rt, ok := c.Data["read_timeout"]; ok {
		c.ReadTimeout = time.Duration(rt.(int)) * time.Millisecond
	}
	if wt, ok := c.Data["write_timeout"]; ok {
		c.WriteTimeout = time.Duration(wt.(int)) * time.Millisecond
	}
	if mhb, ok := c.Data["max_header_bytes"]; ok {
		c.MaxHeaderBytes = mhb.(int)
	}
	if tlscf, ok := c.Data["tls_cert_file"]; ok {
		c.TLSCertFile = tlscf.(string)
	}
	if tlskf, ok := c.Data["tls_key_file"]; ok {
		c.TLSKeyFile = tlskf.(string)
	}
	if tr, ok := c.Data["template_root"]; ok {
		c.TemplateRoot = tr.(string)
	}
	if te, ok := c.Data["template_ext"]; ok {
		c.TemplateExt = te.(string)
	}
	if tld, ok := c.Data["template_left_delim"]; ok {
		c.TemplateLeftDelim = tld.(string)
	}
	if trd, ok := c.Data["template_right_delim"]; ok {
		c.TemplateRightDelim = trd.(string)
	}
	if tm, ok := c.Data["template_minified"]; ok {
		c.TemplateMinified = tm.(bool)
	}
	if tw, ok := c.Data["template_watched"]; ok {
		c.TemplateWatched = tw.(bool)
	}
}
