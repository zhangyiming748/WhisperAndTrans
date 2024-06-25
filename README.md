# WhisperAndTrans
使用openai-whisper生成字幕后自动翻译
# usage
```bash
docker run -dit --name wta -v /Users/zen/Github/FastYt-dlp:/data -e root=/data -e language=en -e pattern=mkv -e model=large -e location=/data -e 'proxy=192.168.1.20:8889' zhangyiming748/whisperandtrans:latest
```
or

```bash
docker compose up -d
```