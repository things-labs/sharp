
execvefile = builder

version = v0.0.1
model = builds
APIversion = v0.0.2

path = github.com/thinkgos/assist/builder

#完整 commit id git rev-parse HEAD
#短 commit id git rev-parse --short HEAD
# TIPS: 如果值内容中含有空格，可以用单引号！
flags = "-X '${path}.BuildTime=`date "+%F %T %z"`' \
	-X '${path}.GitCommit=`git rev-parse --short HEAD`' \
	-X '${path}.Version=${version}' \
	-X '${path}.Model=${model}' \
	-X '${path}.APIVersion=${APIversion}'"

system:
	@go build -ldflags ${flags} -o ${execvefile} .

clean:
	@echo "----> cleanning"
	@rm ${execvefile}
	@echo "----> clean successful"

.PHONY: system clean