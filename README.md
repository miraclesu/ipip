[17mon](http://www.ipip.net/) 批量查询程序
===

## 使用

### 下载相应平台下的执行文件

### 下载 17monipdb dat 数据放到执行文件目录

下载地址：
* http://www.ipip.net/
* https://github.com/wangtuanjie/ip17mon/blob/master/17monipdb.dat

### 把要查询的 ip 存到文件里，每个 ip 一行，命令为 ips.txt 放到执行文件目录

### 执行
./ipip

结果会输出到终端

| ip   | 国家  | 地区  |城市  |isp  |
| ------ | -----| ----- | ----- | ----- |
|1.180.84.230|中国|内蒙古|鄂尔多斯|N/A|

如果想指定路径

./ipip -path=/data/17monipdb.dat -ips=/data/ips.txt
