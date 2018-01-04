# Test project for Jenkins-X using Godog

This test project connects to a running Jenkins instance and runs a number of feature level tests.  You will need to set some environment variables before running the tests.

## Prerequisits

- __golang__ https://golang.org/doc/install#install

## Setup

Set the following environment variables:
```
export BDD_JENKINS_URL=http://your.jenkins.io
export BDD_JENKINS_TOKEN=admin
export BDD_JENKINS_USERNAME=admin

export GITHUB_USER=rawlingsj
export GITHUB_PASSWORD=myPersonalAccessTokenGoesHere
```
Now run:
```
go get github.com/DATA-DOG/godog/cmd/godog
go get github.com/jenkins-x/godog-jenkins
cd $GOPATH/src/github.com/jenkins-x/godog-jenkins/jenkins/
```
Fork the sample springboot project into your own Github org that matches the `GITHUB_USER` env var above:

[Fork spring-boot-web-example](https://github.com/jenkins-x/spring-boot-web-example/fork)

## Run BDD tests

Run the __Jenkins X__ godog tests from this repo:
```
cd $GOPATH/src/github.com/jenkins-x/godog-jenkins/jenkins/
godog
```

# FAQ

## multibranch GitHub API rate limiting

Running these tests on minikube sometimes you will see github API rate limit errors when jobs are starting.

There seems to be an issue in the minikube VM where the date is 2 hours behind, a workaround is to run:

To check run:
```apple
minishift ssh date
```
To set it correctly:
```
minikube ssh -- docker run -i --rm --privileged --pid=host debian nsenter -t 1 -m -u -n -i date -u $(date -u +%m%d%H%M%Y)
```
The Jenkins admin console may be unavailable for a few seconds, when it returns retrigger the job.

## Generate a GitHub personal access token from cURL

To spead up getting started you can get a github personal access token using your terminal:

e.g. replace your username
```
curl https://api.github.com/authorizations \
--user "rawlingsj" \
--data '{"scopes":["repo:status"],"note":"jx"}'
``` 

You will be promted for your GitHub password then you will see a `"token":` returned.

We will wrap this in a CLI for an easier getting started experience.
