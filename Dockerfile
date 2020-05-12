FROM golang
ADD . /go/src/Go_sample_microservice
ENV GOPATH=/go/src/Go_sample_microservice/
WORKDIR /go/src/Go_sample_microservice
RUN go get golang.org/x/crypto/pbkdf2
RUN go get golang.org/x/crypto/md4
RUN go get github.com/jcmturner/gofork/x/crypto/pbkdf2
RUN go get github.com/jcmturner/gofork/encoding/asn1
RUN go get github.com/hashicorp/go-uuid
RUN go get gopkg.in/alecthomas/kingpin.v2
RUN go get github.com/Shopify/sarama
RUN go get github.com/gorilla/mux    
RUN go get github.com/nanobox-io/golang-scribble      
RUN go get gopkg.in/alecthomas/kingpin.v2
RUN go get github.com/patrickmn/go-cache
RUN go build /go/src/Go_sample_microservice/server.go
