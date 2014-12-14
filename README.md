docker build -t img-wxapp

docker run -it -d -p 80:80 --name wxapp img-wxapp
