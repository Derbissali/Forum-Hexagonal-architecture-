docker image build -f dockerfile . -t forum
docker container run -p 9000:8080 -d --name forum forum