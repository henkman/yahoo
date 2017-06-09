package yahoo

import (
	"io"
	"net/http"
	"net/http/cookiejar"
)

type Session struct {
	cli http.Client
}

func (s *Session) Init() error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	s.cli.Jar = jar
	return nil
}

func (s *Session) IsInitialized() bool {
	return s.cli.Jar != nil
}

func (s *Session) request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "yahoo api")
	return s.cli.Do(req)
}