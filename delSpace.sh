#!/bin/bash

# 指定要遍历的文件夹路径
# sed命令可以作用于视频文件 使用时需谨慎
folder_path="/f/Telegram/en"

# 遍历文件夹下所有扩展名为txt的文件
for file in "$folder_path"/*.txt; do
    # 使用sed命令删除每行第一个多余的空格
    sed 's/^[[:space:]]\+//' "$file" > "${file}_temp"
    # 用处理后的文件替换原文件
    mv "${file}_temp" "$file"
done
