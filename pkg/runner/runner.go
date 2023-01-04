package runner

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/executor"
	"github.com/kubeshop/testkube/pkg/executor/content"
	"github.com/kubeshop/testkube/pkg/executor/output"
	"github.com/kubeshop/testkube/pkg/executor/scraper"
)

// func NewRunner() *ExampleRunner *CypressRunner, error{
// 	var params Params
// 	err := envconfig.Process("runner", &params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	runner := &CypressRunner{
// 		Fetcher: content.NewFetcher(""),
// 		Scraper: scraper.NewMinioScraper(
// 			params.Endpoint,
// 			params.AccessKeyID,
// 			params.SecretAccessKey,
// 			params.Location,
// 			params.Token,
// 			params.Ssl,
// 		),
// 		Params:     params,
// 		dependency: dependency,
// 	}

// 	return runner, nil

// }
func NewRunner() (*ExampleRunner, error) {
	var params Params
	err := envconfig.Process("runner", &params)
	if err != nil {
		return nil, err
	}

	runner := &ExampleRunner{
		Fetcher: content.NewFetcher(""),
		Scraper: scraper.NewMinioScraper(
			params.Endpoint,
			params.AccessKeyID,
			params.SecretAccessKey,
			params.Location,
			params.Token,
			params.Ssl,
		),
		Params: params,
	}

	return runner, nil
}

type Params struct {
	Endpoint        string // RUNNER_ENDPOINT
	AccessKeyID     string // RUNNER_ACCESSKEYID
	SecretAccessKey string // RUNNER_SECRETACCESSKEY
	Location        string // RUNNER_LOCATION
	Token           string // RUNNER_TOKEN
	Ssl             bool   // RUNNER_SSL
	ScrapperEnabled bool   // RUNNER_SCRAPPERENABLED
	GitUsername     string // RUNNER_GITUSERNAME
	GitToken        string // RUNNER_GITTOKEN
	Datadir         string // RUNNER_DATADIR
}

// ExampleRunner for template - change me to some valid runner
type ExampleRunner struct {
	Params     Params
	Fetcher    content.ContentFetcher
	Scraper    scraper.Scraper
	dependency string
}

func (r *ExampleRunner) Run(execution testkube.Execution) (result testkube.ExecutionResult, err error) {
	// ScriptContent will have URI
	//
	err = r.Validate(execution)
	if err != nil {
		return result, err
	}
	path, err := r.Fetcher.Fetch(execution.Content)
	if err != nil {
		return result, err
	}
	output.PrintEvent("created content path", path)
	testDir, _ := filepath.Split(path)

	// convert executor env variables to os env variables
	for key, value := range execution.Envs {
		if err = os.Setenv(key, value); err != nil {
			return result, fmt.Errorf("setting env var: %w", err)
		}
	}
	// envManager := secret.NewEnvManagerWithVars(execution.Variables)
	// envManager.GetVars(execution.Variables)
	// envVars := make([]string, 0, len(execution.Variables))
	// for _, value := range execution.Variables {
	// 	envVars = append(envVars, fmt.Sprintf("%s=%s", value.Name, value.Value))
	// }
	// args := []string{"--env", strings.Join(envVars, ",")}
	// args = append(args, execution.Args...)
	args1 := []string{"install", "selenium-webdriver"}
	out1, err := executor.Run(testDir, "npm", args1...)
	if err != nil {
		err = fmt.Errorf("npm install error %w\n\n%s", err, out1)
	}
	args2 := []string{"install"}
	out2, err := executor.Run(testDir, "npm", args2...)
	if err != nil {
		err = fmt.Errorf("npm install error %w\n\n%s", err, out2)
	}
	// args2 := []string{"install", "chromedriver"}
	// out2, err := executor.Run(testDir, "npm", args2...)
	// if err != nil {
	// 	err = fmt.Errorf("npm install error %w\n\n%s", err, out2)
	// }
	args := []string{path, "-R", "json"}
	out, err := executor.Run(testDir, "mocha", args...)
	if err != nil {
		return result, fmt.Errorf("selenium execution error %w\n\n%s", err, out)
	}
	// out = envManager.Obfuscate(out)
	result.Output = string(out)
	result.Status = testkube.ExecutionStatusPassed
	return result, nil
	// return testkube.ExecutionResult{
	// 	Status: testkube.ExecutionStatusPassed,
	// 	Output: string(runPath) + " Error " + err.Error(),
	// }, nil
	// resp, err := http.Get(uri)
	// if err != nil {
	// 	return result, err
	// }
	// defer resp.Body.Close()

	// b, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return result, err
	// }

	// // if get is successful return success result
	// if resp.StatusCode == 200 {
	// 	return testkube.ExecutionResult{
	// 		Status: testkube.ExecutionStatusPassed,
	// 		Output: string(b),
	// 	}, nil
	// }

	// else we'll return error to simplify example
	// err = fmt.Errorf("invalid status code %d, (uri:%s)", resp.StatusCode, uri)
	// return result.Err(err), nil
}

func (r *ExampleRunner) Validate(execution testkube.Execution) error {

	if execution.Content == nil {
		return fmt.Errorf("can't find any content to run in execution data: %+v", execution)
	}

	return nil
}
