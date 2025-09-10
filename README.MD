# game_server_client


## 如何创建项目
1. mkdir zinx &&  cd zinx  &&  go work init // 创建工作空间
2. mkdir zinx-lib && cd zinx-lib && go mod init zinx-lib // 创建lib 模块
3. 回到工作空间的根目录 执行 go work use zinx-lib // 使用lib 模块

重复第2和第3步骤，创建 zinx-client 和 zinx-server模块

## 0.8 版本改进点
1. 读写分离, 读channel 和写channel 通过chan进行通信，写channel负责给客户端回写数据
   a. 读写分离，读写分离可以提高并发量，提高吞吐量
   b. 读写分离，读写分离可以降低读写锁的竞争，提高性能
   c. 读写分离，读写分离可以提高代码的可读性，降低耦合度

2. WorkerPool 池化协程