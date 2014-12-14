# why
A simple attempt to utilize the WeiXin open service, for learning GO and practising Docker. That's it.
# how
* docker build -t img-wxapp
* docker run -it -d -p 80:80 --name wxapp img-wxapp

or
* docker run  -v "$(pwd)":/myapp /myapp img-wxapp go build -v
# more ..
It's still under dev, based on my schedule

Author: Bruce.Lu
