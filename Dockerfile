FROM golang:1.22.4-bookworm
# docker run -dit -v /Users/zen/Github/WhisperAndTrans:/data --name test golang:1.22.4-bookworm bash
LABEL authors="zen"
COPY debian.sources /etc/apt/sources.list.d/
RUN apt update
RUN apt install -y python3 python3-pip translate-shell ffmpeg ca-certificates bsdmainutils sqlite3 gawk locales libfribidi-bin dos2unix
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
RUN rm /usr/lib/python3.11/EXTERNALLY-MANAGED
RUN pip install openai-whisper
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/go/bin

ENTRYPOINT ["go", "run","/app/main.go"]