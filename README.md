# Bing每日壁纸服务端
## 如何在Windows环境下编译Linux服务器可执行文件
- git clone源码到本地
- 在配置好Golang的环境下，双击build.bat

## 如何在Linux服务器上运行
- 上传编译好的文件夹(server)到你所需要运行的目录（例如：www/wwwroot/bing.gocos.cn/server）并配置server文件夹中的`config.yaml`
- 在服务器上安装screen，CentOS执行`yum install screen`，Debian/Ubuntu执行`apt-get install screen`
- 用screen新建一个会话`screen -S bing-server`
- 在新会话中cd到运行目录`cd www/wwwroot/bing.gocos.cn/server`
- 执行`./main`,运行程序，外网访问请打开防火墙9090端口
- 按ctrl+a ctrl+d 保存并退出此会话
- `screen -r bing-server`可以再次进入此会话

## 宝塔面板反向代理配置
> 此配置是为了配合前端页面，保证使用同一域名，无需放行9090端口
```
#PROXY-START/api

location ^~ /api
{
    proxy_pass http://127.0.0.1:9090/api;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header REMOTE-HOST $remote_addr;
    
    add_header X-Cache $upstream_cache_status;
    
    #Set Nginx Cache


    
    
    set $static_fileLpb3LN9D 0;
    if ( $uri ~* "\.(gif|png|jpg|css|js|woff|woff2)$" )
    {
    	set $static_fileLpb3LN9D 1;
    	expires 12h;
        }
    if ( $static_fileLpb3LN9D = 0 )
    {
    add_header Cache-Control no-cache;
    }
}

#PROXY-END/api
```
