# game_server_client


## 如何创建项目
1. mkdir zinx &&  cd zinx  &&  go work init // 创建工作空间
2. mkdir zinx-lib && cd zinx-lib && go mod init zinx-lib // 创建lib 模块
3. 回到工作空间的根目录 执行 go work use zinx-lib // 使用lib 模块

重复第2和第3步骤，创建 zinx-client 和 zinx-server模块