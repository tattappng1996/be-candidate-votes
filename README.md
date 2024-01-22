# be-candidate-votes

After setup postgres docker

docker build -t be-candidate-votes .
docker run -d -p 8081:8081 -e CONFIG_PATH=./env/config.yaml be-candidate-votes