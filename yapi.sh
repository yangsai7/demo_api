#!/bin/bash

# YApi服务器地址和项目Token
yapi_server="https://yapi.rabbitgo.net"
project_token="19ec3bbdfcddd24312a41823c254ba5dbdeca7d70b0e006da5d8802a23551f9c"

# YApi CLI命令
yapi_command="yapi import --config"

# 遍历目录下所有swagger.json文件并执行导入命令
for file in $(find ./api -type f -name "*swagger.json"); do
  config_file="/tmp/yapi-import-$(basename "$file" .json).json"
  echo $config_file
  
  # 创建临时配置文件
  echo "{\"server\": \"$yapi_server\",\"merge\": \"mergin\", \"token\": \"$project_token\", \"type\": \"swagger\", \"file\": \"$file\"}" > "$config_file"
  
  # 执行导入命令
  $yapi_command "$config_file"
  
  # 删除临时配置文件
  rm "$config_file"
done
