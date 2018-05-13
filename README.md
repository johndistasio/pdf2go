# pdf2go

Demo application for AWS Lambda golang support.

## Instructions

1) Install depdencies:

```
$ dep ensure
```

2) Build handler function and deployment bundle:

```
$ make
```

3) Set up Terraform environment:

```
$ export AWS_PROFILE=<profile>; export AWS_DEFAULT_REGION=<region>
```

4) Create infrastructure:

```
$ terraform apply
```
