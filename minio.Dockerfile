
FROM minio/minio
ENTRYPOINT ["/bin/sh", "-c", "mdkir -p /tmp/data/s3test && minio ${MINIO_COMMAND:-server} ${MINIO_DIR:-/tmp/data}"] 