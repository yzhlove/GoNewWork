# chat01

```text
部署好本机的 docker 环境，使用 ppt 中的 dockerfile build 自己的环境
使用 readelf 工具，查看编译后的进程入口地址
在 dlv 调试工具中，使用断点功能找到代码位置
使用断点调试功能，查看 Go 的 runtime 的下列函数执行流程，使用 IDE 查看函数的调用方：
必做：runqput，runqget，globrunqput，globrunqget
选做：schedule，findrunnable，sysmon
难度++课外作业：跟踪进程启动流程中的关键函数，rt0_go，需要汇编知识，可以暂时不做，只给有兴趣的同学

```

# centos 下

```text
FROM centos
RUN yum install golang -y \
&& yum install dlv -y \
&& yum install binutils -y \ && yum install vim -y \ && yum install gdb -y
docker build -t test .
docker run -it --rm test bash
```

> go build -x main.go 编译与链接

```text
WORK=/var/folders/jd/drfndq197ql6s3hmslgmjd700000gn/T/go-build2728402243
mkdir -p $WORK/b001/
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=/Users/yurisa/Library/Caches/go-build/7c/7cc2e0fcf8e30418b854b3ad52e8d21d08bb6118bae03a5c83f91cf24e7a1fb7-d
packagefile fmt=/usr/local/go/pkg/darwin_amd64/fmt.a
packagefile runtime=/usr/local/go/pkg/darwin_amd64/runtime.a
packagefile errors=/usr/local/go/pkg/darwin_amd64/errors.a
packagefile internal/fmtsort=/usr/local/go/pkg/darwin_amd64/internal/fmtsort.a
packagefile io=/usr/local/go/pkg/darwin_amd64/io.a
packagefile math=/usr/local/go/pkg/darwin_amd64/math.a
packagefile os=/usr/local/go/pkg/darwin_amd64/os.a
packagefile reflect=/usr/local/go/pkg/darwin_amd64/reflect.a
packagefile strconv=/usr/local/go/pkg/darwin_amd64/strconv.a
packagefile sync=/usr/local/go/pkg/darwin_amd64/sync.a
packagefile unicode/utf8=/usr/local/go/pkg/darwin_amd64/unicode/utf8.a
packagefile internal/abi=/usr/local/go/pkg/darwin_amd64/internal/abi.a
packagefile internal/bytealg=/usr/local/go/pkg/darwin_amd64/internal/bytealg.a
packagefile internal/cpu=/usr/local/go/pkg/darwin_amd64/internal/cpu.a
packagefile internal/goexperiment=/usr/local/go/pkg/darwin_amd64/internal/goexperiment.a
packagefile runtime/internal/atomic=/usr/local/go/pkg/darwin_amd64/runtime/internal/atomic.a
packagefile runtime/internal/math=/usr/local/go/pkg/darwin_amd64/runtime/internal/math.a
packagefile runtime/internal/sys=/usr/local/go/pkg/darwin_amd64/runtime/internal/sys.a
packagefile internal/reflectlite=/usr/local/go/pkg/darwin_amd64/internal/reflectlite.a
packagefile sort=/usr/local/go/pkg/darwin_amd64/sort.a
packagefile math/bits=/usr/local/go/pkg/darwin_amd64/math/bits.a
packagefile internal/itoa=/usr/local/go/pkg/darwin_amd64/internal/itoa.a
packagefile internal/oserror=/usr/local/go/pkg/darwin_amd64/internal/oserror.a
packagefile internal/poll=/usr/local/go/pkg/darwin_amd64/internal/poll.a
packagefile internal/syscall/execenv=/usr/local/go/pkg/darwin_amd64/internal/syscall/execenv.a
packagefile internal/syscall/unix=/usr/local/go/pkg/darwin_amd64/internal/syscall/unix.a
packagefile internal/testlog=/usr/local/go/pkg/darwin_amd64/internal/testlog.a
packagefile internal/unsafeheader=/usr/local/go/pkg/darwin_amd64/internal/unsafeheader.a
packagefile io/fs=/usr/local/go/pkg/darwin_amd64/io/fs.a
packagefile sync/atomic=/usr/local/go/pkg/darwin_amd64/sync/atomic.a
packagefile syscall=/usr/local/go/pkg/darwin_amd64/syscall.a
packagefile time=/usr/local/go/pkg/darwin_amd64/time.a
packagefile unicode=/usr/local/go/pkg/darwin_amd64/unicode.a
packagefile internal/race=/usr/local/go/pkg/darwin_amd64/internal/race.a
packagefile path=/usr/local/go/pkg/darwin_amd64/path.a
EOF
mkdir -p $WORK/b001/exe/
cd .
/usr/local/go/pkg/tool/darwin_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=Ty45rJoBKlopd0iHAXvI/8dTlYLCGMv7FkpSzO26J/Zck4lkgOJV8fiqJNyWXI/Ty45rJoBKlopd0iHAXvI -extld=clang /Users/yurisa/Library/Caches/go-build/7c/7cc2e0fcf8e30418b854b3ad52e8d21d08bb6118bae03a5c83f91cf24e7a1fb7-d
/usr/local/go/pkg/tool/darwin_amd64/buildid -w $WORK/b001/exe/a.out # internal
mv $WORK/b001/exe/a.out main
rm -r $WORK/b001/

```