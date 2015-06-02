# 百度音乐下载工具

百度音乐下载工具，是一款使用Go语言编写，用于自动获取百度音乐歌曲地址，并下载高品质歌曲的小工具。

![运行效果图](http://git.oschina.net/zhanglubing/baidu-music-downloader/raw/master/res/demo.png)

**特别申明** 该工具只用于学习和研究。

### 功能

1. 支持下载的歌曲列表有：
	+ 歌手 http://music.baidu.com/artist/{歌手ID}
	+ 专辑 http://music.baidu.com/album/{专辑ID}
	+ 榜单 http://music.baidu.com/top/{榜单ID}
	+ 风格 http://music.baidu.com/tag/{风格名称}
	+ 歌单 http://music.baidu.com/songlist/{歌单ID}
2. 支持歌曲过滤，可指定最小比特率。
3. 支持自定义文路径和文件名。
4. 已下过歌曲不会重复下载。
5. 采用JSON文件配置，方便修改。

### 编译和使用

编译代码需要安装Git和Go环境。

	git clone https://git.oschina.net/zhanglubing/baidu-music-downloader.git  
	cd baidu-music-downloader  
	# 修改配置文件（歌单和Cookie）  
	./install  
	./bin/baidu-music-downloader  

源码已附带编译好的可执行文件，见：  
$ baidu-music-downloader/bin/baidu-music-downloader  

附带的可执行文件的编译环境为：Ubuntu 12.04，Go版本：go version go1.1.2 linux/amd64

### 配置文件说明

配置文件名为：**config.json**，运行 **install** 时会自动拷贝到bin目录，和编译后的可执行文件位于同一文件夹下。

具体配置见示例文件：  
$ [baidu-music-downloader/res/config.sample.json](http://git.oschina.net/zhanglubing/baidu-music-downloader/blob/master/res/config.sample.json)  

### 关于Cookie

工具采用了Cookie欺骗的方式：首先网页端登陆百度账号，手动获取服务器返回的Cookie，并将Cookie用于工具去抓取下载链接。

Cookie获取方式如下：

1. 登陆百度账号。  
![登陆百度账号](http://git.oschina.net/zhanglubing/baidu-music-downloader/raw/master/res/login.png)
2. 这里以Chrome浏览器为例，其他的浏览器请自行想办法获取。  
登陆成功后，按F12打开“Developer Tools”，切换到“Network”标签，并按F5重新加载一次，然后选择左边第一个请求。  
![获取Cookie信息](http://git.oschina.net/zhanglubing/baidu-music-downloader/raw/master/res/cookie.png)
3. 拷贝上图中的Cookie值，粘贴到config.json中的cookie配置上。
4. 完成！

**NOTE** 登陆后的浏览器不要关闭，否则服务器销毁Session时会让Cookie失效。若下载工具提示Cookie过期，再按以上步骤获取一次Cookie。

### 后续改进？

这个工具目前可以工作，但还有许多问题，比如说：

1. 不支持分页列表的歌曲获取。
2. 登陆方式极不友好（需手动获取Cookie），不利于小白把玩。
3. 也不知道Windows下表现如何（Mac和Linux已验证过）。
4. 甚至没有测试代码！！

但这个小工具只用于学习和练手，可能不会有后续了，感兴趣的朋友请自便。代码上不到位的地方，无条件接受批评教育，无条件接受打脸和砸砖头！

### Enjoy！