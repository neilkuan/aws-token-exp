### Build Arch 
```bash
GOOS=linux GOARCH=arm64 go build .

GOOS=linux GOARCH=amd64 go build .

GOOS=darwin GOARCH=arm64 go build .

GOOS=darwin GOARCH=amd64 go build .
```

##$ To use 
```bash
aws-token-exp token --profile aad-xxx
2023/11/03 14:43:34 Try find profile_name: aad-xxx in the /Users/neil.xxx/.aws/credentials ...
2023-11-03 16:11:29 +0800 CST

```
![](./docs/1.png)