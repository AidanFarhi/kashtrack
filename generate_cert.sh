certbot certonly --standalone -d kash-track.com \
  --config-dir ./cert/config \
  --work-dir ./cert \
  --logs-dir ./cert/logs