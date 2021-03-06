// Copyright 2019 HuaweiCloud.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collector

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type CloudAuth struct {
	ProjectName       string `yaml:"project_name"`
	ProjectID         string `yaml:"project_id"`
	DomainName        string `yaml:"domain_name"`
	AccessKey         string `yaml:"access_key"`
	Region            string `yaml:"region"`
	SecretKey         string `yaml:"secret_key"`
	AuthURL           string `yaml:"auth_url"`
	UserName          string `yaml:"user_name"`
	Password          string `yaml:"password"`
}

type Global struct {
	Port       string `yaml:"port"`
	Prefix     string `yaml:"prefix"`
	MetricPath string `yaml:"metric_path"`
}

type CloudConfig struct {
	Auth               CloudAuth `yaml:"auth"`
	InfoMetrics        []string  `yaml:"info_metrics"`
	Global			   Global    `yaml:"global"`
}


func NewCloudConfigFromFile(file string) (*CloudConfig, error) {
	var config CloudConfig

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}


func SetDefaultConfigValues(config *CloudConfig)  {
	if config.Global.Port == "" {
		config.Global.Port = ":8087"
	}

	if config.Global.MetricPath == "" {
		config.Global.MetricPath = "/metrics"
	}

	if config.Global.Prefix == "" {
		config.Global.Prefix = "huaweicloud"
	}
}
