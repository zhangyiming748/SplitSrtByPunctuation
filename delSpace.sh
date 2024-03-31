#!/bin/bash
# 删除每行第一个无效空格
# 遍历当前目录下的所有文件
for file in *; do
  # 检查是否为文件
  if [ -f "$file" ]; then
    # 使用sed命令删除每行第一个多余的空格
    sed 's/^[[:space:]]\+//' "$file" > "${file}_temp"
    # 用处理后的文件替换原文件
    mv "${file}_temp" "$file"
  fi
done
