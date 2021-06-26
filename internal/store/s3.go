package store

import "github.com/philippgille/gokv/s3"

type s3store struct {
	opts Options
	s3   s3.Client
}

// NewS3Store creates a new s3 store and return store and a error if any
func NewS3Store(opts ...Option) (Store, error) {
	ss := &s3store{}
	for _, o := range opts {
		o(&ss.opts)
	}
	s3opt := s3.DefaultOptions
	if len(ss.opts.Directory) != 0 {
		s3opt.BucketName = ss.opts.Directory
	}
	if len(ss.opts.S3Endpoint) != 0 {
		s3opt.CustomEndpoint = ss.opts.S3Endpoint
	}
	if ss.opts.UsePathStyleAddressing {
		s3opt.UsePathStyleAddressing = true
	}
	if len(ss.opts.AWSaccessKeyID) != 0 {
		s3opt.AWSaccessKeyID = ss.opts.AWSaccessKeyID
	}
	if len(ss.opts.AWSsecretAccessKey) != 0 {
		s3opt.AWSsecretAccessKey = ss.opts.AWSsecretAccessKey
	}
	c, err := s3.NewClient(s3opt)
	if err != nil {
		return nil, err
	}
	ss.s3 = c
	return ss, nil
}

func (s *s3store) Set(key string, value interface{}) error {
	return s.s3.Set(key, value)
}

func (s *s3store) Get(key string, value interface{}) (bool, error) {
	return s.s3.Get(key, value)
}

func (s *s3store) Delete(key string) error {
	return s.s3.Delete(key)
}

func (s *s3store) Close() error {
	return s.s3.Close()
}

func (s *s3store) String() string {
	return "S3 Store"
}
