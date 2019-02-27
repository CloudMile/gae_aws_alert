# gae_aws_alert
Use GAE to trace AWS CloudWatch, if find any log about `error` will send mail to users.

## Setup
```
$ vim app.yaml
```
Edit you env you need, like AWS keys and AWS log group name

## Deploy
```
$ gcloud app deploy app.yaml queue.yaml crom.yaml
```
