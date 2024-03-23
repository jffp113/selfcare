package selfcare

type Option func(*Selfcare) error

func WithConfigFilePath(path string) Option {
	return func(s *Selfcare) error {
		s.path = path
		return nil
	}
}

func WithBaseUrl(url string) Option {
	return func(s *Selfcare) error {
		s.baseUrl = url
		return nil
	}
}
