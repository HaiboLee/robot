# robot
golang lua robot

#### lua脚本实现电脑脚本自动化

####内置api
|API      | 参数 |描述|
| ----------- | ----------- | -----------|
| click(x int,y int,double bool)       | double为可选参数不填=false      | 鼠标左键点击|
| right(x int,y int,double bool)   | double为可选参数不填=false        |鼠标右键点击|
|keyTap(key string)| key 按键值| 键盘按键|
|keyTaps(keys table)|keys 多个按键 table结构 :{"ctrl","a"}| 多个按键|
|input(text string)| 字符串输入|文字输入|
|sleep(x int,ms bool)|x 睡眠时间,ms为可选参数 true=毫秒|程序睡眠等待|
|getRgb(x int,y int)|x,y 坐标|获取指定坐标颜色,返回string
|hasRgb(x int,y int,color string,timeout int,step int)|x,y 坐标,color指定颜色,timeout 超时时间可选,step检测频率可选ms 默认50ms|判断指定坐标是否为该颜色,返回bool|
|addEvent(key string)| key按键| 按键监听
|move(x,y int)| x,y坐标|鼠标移动到指定坐标
|eventSleep(x int)|x 单位ms|每次操作默认睡眠时间
|dragSmooth(x,y int)| x,y 坐标|点击拖动鼠标
|scroll(x,y int)|x,y 坐标|鼠标滚动