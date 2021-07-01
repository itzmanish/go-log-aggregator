package store

type Options struct {
	// Directory is the directory name for file store
	// For s3 Directory is used as Bucket name
	Directory string
	// S3Endpoint is used for custom s3 hosted server
	// or other aws s3 compatible solutions.
	// eg - for local minio server without https will be
	// http://localhost:9000
	S3Endpoint string
	// S3 differentiates between "virtual hosted bucket addressing" and "path-style addressing".
	// Example URL for the virtual host style: http://BUCKET.s3.amazonaws.com/KEY.
	// Example UL for the path style: http://s3.amazonaws.com/BUCKET/KEY.
	// Most S3-compatible servers work with both styles,
	// but some work only with the virtual host style (e.g. Alibaba Cloud Object Storage Service (OSS))
	// and some work only with the path style (especially self-hosted services like a Minio server running on localhost).
	// Optional (false by default).
	UsePathStyleAddressing bool
	// AWS access key ID (part of the credentials).
	// Optional (read from shared credentials file or environment variable if not set).
	// Environment variable: "AWS_ACCESS_KEY_ID".
	AWSaccessKeyID string
	// AWS secret access key (part of the credentials).
	// Optional (read from shared credentials file or environment variable if not set).
	// Environment variable: "AWS_SECRET_ACCESS_KEY".
	AWSsecretAccessKey string
	// Prefix
	Prefix string
}

type Option func(o *Options)

func WithDirectory(dir string) Option {
	return func(o *Options) {
		o.Directory = dir
	}
}

func WithS3Endpoint(ep string) Option {
	return func(o *Options) {
		o.S3Endpoint = ep
	}
}

func WithAWSAccessKey(ak string) Option {
	return func(o *Options) {
		o.AWSaccessKeyID = ak
	}
}

func WithAWSSecretAccessKey(sak string) Option {
	return func(o *Options) {
		o.AWSsecretAccessKey = sak
	}
}

func WithPathStyleAddressing(ps bool) Option {
	return func(o *Options) {
		o.UsePathStyleAddressing = ps
	}
}

func WithPrefix(prefix string) Option {
	return func(o *Options) {
		o.Prefix = prefix
	}
}
