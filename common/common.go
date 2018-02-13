package common

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jenkins-x/godog-jenkins/jenkins"
	"github.com/jenkins-x/godog-jenkins/utils"
	"github.com/jenkins-x/golang-jenkins"
	cmdutil "github.com/jenkins-x/jx/pkg/jx/cmd/util"
)

type CommonTest struct {
	Factory       cmdutil.Factory
	JenkinsClient *gojenkins.Jenkins
	Interactive   bool
	Errors        *utils.ErrorSlice
	WorkDir       string
	AppName       string
}

// TheApplicationShouldBeBuiltAndPromotedViaCICD asserts that the project
// should be created in Jenkins and that the build should complete successfully
func (o *CommonTest) TheApplicationShouldBeBuiltAndPromotedViaCICD() error {
	appName := o.AppName
	if appName == "" {
		_, appName = filepath.Split(o.WorkDir)
	}
	f := o.Factory
	f.SetBatch(true)
	gitURL, err := o.GitProviderURL()
	if err != nil {
		return err
	}
	gitAuthSvc, err := f.CreateGitAuthConfigService()
	if err != nil {
		return err
	}
	gitConfig := gitAuthSvc.Config()
	server := gitConfig.GetServer(gitURL)
	if server == nil {
		return fmt.Errorf("Could not find a git auth user for git server URL %s", gitURL)
	}
	userName := server.CurrentUser
	if userName == "" {
		if len(server.Users) == 0 {
			return fmt.Errorf("No users are configured for authentication with git server URL %s", gitURL)
		}
		userName = server.Users[0].Username
	}
	if userName == "" {
		return fmt.Errorf("Could not detect username for git server URL %s", gitURL)
	}
	jobName := userName + "/" + appName + "/master"
	if o.JenkinsClient == nil {
		client, err := f.CreateJenkinsClient()
		if err != nil {
			return err
		}
		o.JenkinsClient = client
	}
	fmt.Printf("Checking that there is a job built successfully for %s\n", jobName)
	return jenkins.ThereShouldBeAJobThatCompletesSuccessfully(jobName, o.JenkinsClient)
}

// GitProviderURL returns the git provider to use
func (o *CommonTest) GitProviderURL() (string, error) {
	gitProviderURL := os.Getenv("GIT_PROVIDER_URL")
	if gitProviderURL != "" {
		return gitProviderURL, nil
	}
	// find the default  load the default one from the current ~/.jx/jenkinsAuth.yaml
	authConfigSvc, err := o.Factory.CreateGitAuthConfigService()
	if err != nil {
		return "", err
	}
	config := authConfigSvc.Config()
	url := config.CurrentServer
	if url != "" {
		return url, nil
	}
	servers := config.Servers
	if len(servers) == 0 {
		return "", fmt.Errorf("No servers in the ~/.jx/gitAuth.yaml file!")
	}
	return servers[0].URL, nil
}
