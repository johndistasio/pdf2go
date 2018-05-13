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

5) Copy the `invoke_url` output into a `curl` invocation:

```
$ curl -X POST -d 'World' -i -H 'Accept: application/pdf' '<invoke_url>' -o test.pdf
```

6) Open the PDF.
