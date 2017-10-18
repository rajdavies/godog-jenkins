# Test project for fabric8 Jenkins using Godog

This test project connects to a running Jenkins instance and runs a number of feature level tests.  You will need to set some environment variables before running the tests.

Go to http://your.jenkins.io/me/configure, in the __API Token__ section click `Show API Token` and take note of the user id a token values.

![api token](images/api-token.png)

To run these tests against a [fabric8 installation](https://github.com/fabric8io/fabric8-platform/blob/master/INSTALL.md) then run the following command:

```shell
eval $(gofabric8 bdd-env)
```

Then set the following environment variables:
```
export BDD_JENKINS_URL=http://your.jenkins.io
export BDD_JENKINS_TOKEN=abcd1234
export BDD_JENKINS_USERNAME=jrawlings

export GITHUB_USER=rawlingsj
export GITHUB_PASSWORD=myPersonalAccessTokenGoesHere
```
Now run:
```
go get github.com/DATA-DOG/godog/cmd/godog
go get github.com/fabric8-jenkins/godog-jenkins
cd $GOPATH/src/github.com/fabric8-jenkins/godog-jenkins/jenkins/
```
And trigger the tests:
```
godog
```