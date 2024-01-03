# Install
### For Linux users

```bash
$ curl -fsSL https://github.com/neilkuan/aws-token-exp/releases/latest/download/aws-token-exp-Linux-$(uname -m).tar.gz -o - | tar zxvf -
$ mv ./aws-token-exp /usr/local/bin/aws-token-exp
```

### For macOS users

```bash
$ curl -fsSL https://github.com/neilkuan/aws-token-exp/releases/latest/download/aws-token-exp-Darwin-$(uname -m).tar.gz -o - | tar zxvf -
$ mv ./aws-token-exp /usr/local/bin/aws-token-exp
```

## Build Arch 
```bash
GOOS=linux GOARCH=arm64 go build .

GOOS=linux GOARCH=amd64 go build .

GOOS=darwin GOARCH=arm64 go build .

GOOS=darwin GOARCH=amd64 go build .
```

## To use 
```bash
# use profile
aws-token-exp token --profile aad-xxx
2023/11/03 14:43:34 Try find profile_name: aad-xxx in the /Users/neil.xxx/.aws/credentials ...
2023-11-03 16:11:29 +0800 CST


# Use AWS_PROFILE env.
AWS_PROFILE=aad-xxx; aws-token-exp token
2023/11/03 14:43:34 Try find profile_name: aad-xxx in the /Users/neil.xxx/.aws/credentials ...
2023-11-03 16:11:29 +0800 CST
```
![](./docs/1.png)


## `pc` profile change
```bash
curl -o ~/pc.sh https://raw.githubusercontent.com/neilkuan/aws-token-exp/main/pc.sh

source ~/pc.sh
```

### alias in to `~/.zshrc` or `~/.bashrc`
```bash
curl -o ~/pc.sh https://raw.githubusercontent.com/neilkuan/aws-token-exp/main/pc.sh

echo 'alias pc="source pc.sh"' >> ~/.zshrc
```

![](./docs/pc-demo.gif)


# License

[Apache-2.0](LICENSE)