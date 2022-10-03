FROM golang:1.19


WORKDIR $GOPATH/src/github.com/201R/go_api_boilerplate
COPY . $GOPATH/src/github.com/201R/go_api_boilerplate
COPY .ini.exemple $GOPATH/src/github.com/201R/go_api_boilerplate/.ini

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN curl -LO https://release.ariga.io/atlas/atlas-linux-amd64-latest
RUN install -o root -g root -m 0755 ./atlas-linux-amd64-latest /usr/local/bin/atlas

# generate swagger documentation
RUN swag init
# RUN 

EXPOSE 8090
CMD go build . && atlas migrate apply  --dir 'file://migrations' --url 'postgres://postgres:password@postgres:5432/db?sslmode=disable' && ./go_api_boilerplate

