
execveFile = builder

version = v0.0.1
model = builds
APIversion = v0.0.2

path = github.com/thinkgos/sharp/builder

# TIPS: 如果值内容中含有空格，可以用单引号！
flags = "-X '${path}.BuildTime=`date "+%F %T %z"`' \
	-X '${path}.GitCommit=`git rev-parse --short=8 HEAD`' \
	-X '${path}.GitFullCommit=`git rev-parse HEAD`' \
	-X '${path}.Version=${version}' \
	-X '${path}.Model=${model}' \
	-X '${path}.APIVersion=${APIversion}'"

system:
	@go build -ldflags ${flags} -o ${execveFile} .

clean:
	@echo "----> cleanning"
	@rm ${execveFile}
	@echo "----> clean successful"

.PHONY: system clean