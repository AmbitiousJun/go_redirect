# go_redirect

一个简单的 302 重定向服务, 类似于短链接, 使用 Go 语言编写, Docker 部署

## 使用方法

1. 准备一个具有 Docker + DockerCompose 的系统环境

2. 拉取代码

```shell
git clone https://github.com/AmbitiousJun/go_redirect.git
```

3. 修改配置

在 `config.yml` 文件中只有一个 `groups` 配置, 可以配置多条 `http 模板`

在需要动态变化的分段中包裹上 `${}` 符号即可, 括号内部可填写默认值, 没有提供默认值时, 程序认为这个参数是必须提供的

例子：

```yaml
groups:
  - https://gitee.com/ambitiousjun/${iptv-test}/raw/master/${}
```

在这个例子中, 只配置了一条 http 模板, 该模板即为 `1 号模板`

在模板中, `${iptv-test}` 表示一个可以动态变化的分段, 默认值是 `iptv-test`; 最后的 `${}` 是一个不具备默认值的 动态变化片段

实际使用中, 假设 `go_redirect` 服务部署在 `127.0.0.1` 的 `5555` 端口中:

访问地址 `http://127.0.0.1:5555/1/a.txt`, 会被 302 重定向到 1 号模板, 并将 `a.txt` 替换到最后的动态变化片段, 第一个片段使用默认值, 也就是重定向到 `https://gitee.com/ambitiousjun/iptv-test/raw/master/a.txt`

默认值可以被覆盖, 当访问地址 `http://127.0.0.1:5555/1/b/a.txt` 时, 会被重定向到 `https://gitee.com/ambitiousjun/b/raw/master/a.txt`

4. 编译并运行容器

```shell
docker-compose up -d --build
```

较新版本的 Docker 需要使用以下命令:

```shell
docker compose up -d --build
```