# cloudeye-exporter

Prometheus cloudeye exporter for [Huaweicloud](https://www.huaweicloud.com/).



## Building The Discovery with Exact steps on clean Ubuntu 16.04
```
$ wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.12.5.linux-amd64.tar.gz
$ export PATH=$PATH:/usr/local/go/bin # You should put in your .profile or .bashrc
$ go version # to verify it runs and version #

$ go get github.com/huaweicloud/cloudeye-exporter
$ cd ~/go/src/github.com/huaweicloud/cloudeye-exporter
$ go build
```

## Usage
```
 ./cloudeye-exporter  -config=clouds.yml
```

## Help
```
Usage of ./cloudeye-exporter:
  -config string
        Path to the cloud configuration file (default "./clouds.yml")
  -debug
        If debug the code.
 
```

## Example of config file(clouds.yml)
```
global:
  prefix: "huaweicloud"
  port: ":8087"
  metric_path: "/metrics"

auth:
  auth_url: "https://iam.cn-north-1.myhwclouds.com/v3"
  project_name: "cn-north-1"
  access_key: "xdfsdfsdfsdfsdfsdf"
  secret_key: "xsdfsddfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfdsfsd"
  region: "cn-north-1"

info_metrics:
  - SYS.ELB
  - SYS.VPC
```
or

```
auth:
  auth_url: "https://iam.cn-north-1.myhwclouds.com/v3"
  project_name: "cn-north-1"
  user_name: "username"
  password: "password"
  region: "cn-north-1"
  domain_name: "domain_name"


info_metrics:
  - SYS.RDS
```
