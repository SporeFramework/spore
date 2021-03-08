FROM golang

WORKDIR /spore
COPY . .
RUN go build -o spore .
ENTRYPOINT [ "/spore/spore" ]
