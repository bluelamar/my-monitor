
.PHONY: build clean deploy

build:
	env GOOS=linux go build -o my-mon-demultiplexor  github.com/bluelamar/my-monitor/frontend
	chmod +x my-mon-demultiplexor
	#dep ensure -v
	#env GOOS=linux go build -ldflags="-s -w" -o my-mon-demultiplexor main.go

local-deploy: build
	sam local start-lambda -t my-mon-demultiplexor.yml --profile bluelamar

# depends on local-deploy to be run first in other window
test:
	aws lambda invokde --function-name "my-mon-demultiplexor" --endpoint-url "http://127.0.0.1:3001" --profile bluelamar --no-verify-ssl out.txt

package:
	zip handler.zip my-mon-demultiplexor

sam-deploy:
	sam package --template-file template.yaml --s3-bucket bluelamar-my-mon-demultiplexor --output-template-file packaged.yaml

clean:
	rm my-mon-demultiplexor handler.zip

