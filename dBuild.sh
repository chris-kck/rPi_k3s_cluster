docker build -t kckateera/lemi011b-client:latest -f build/docker/client/Dockerfile . &&
docker build -t kckateera/lemi011b-server:latest -f build/docker/server/Dockerfile . &&
docker push kckateera/lemi011b-client:latest &&
docker push kckateera/lemi011b-server:latest
