ec2search
======

Command line utility to easily execute commands across multiple EC2 instances.

## Installation

```
go get github.com/kurtmc/ec2search && go install github.com/kurtmc/ec2search
```

## Usage

```
$ ec2search my-app*dev.company.com
my-app-0e4f6d1a0c3b3a993.dev.company.com
my-app-041644d554e28eddf.dev.company.com
my-app-022729072724f77b3.dev.company.com
my-app-0891f7d1e9ab9fb5d.dev.company.com
my-app-0e13cfa85f70c876c.dev.company.com
my-app-07175b9538580400f.dev.company.com
```
