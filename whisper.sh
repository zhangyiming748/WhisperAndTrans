#!/usr/bin/env bash
# 用于查找指定路径下和子文件夹中所有扩展名为 MP4 的文件,并分离出当前视频的绝对路径(不包含文件名)。如果当前路径下不存在扩展名为 srt 的文件,则打印这个路径
# 设置要搜索的根目录
root_dir="/path/to/your/directory"
# shellcheck disable=SC2034
# model=large-v3
model=medium
# shellcheck disable=SC2034
# language=Japanese
language=English
# shellcheck disable=SC2034
model_dir="/path/to/your/directory"
# 遍历根目录及其子目录
find "$root_dir" -type f -name "*.mp4" | while read file; do
    # 获取文件的绝对路径
    file_path=$(dirname "$file")
    # 检查当前路径下是否存在扩展名为 srt 的文件
    if ! find "$file_path" -type f -name "*.srt" | read; then
        echo "$file_path"
        # whisper "$file_path" --model $model --language $language --model_dir $model_dir --output_format srt
    fi
done
