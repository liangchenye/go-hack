/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e_node

import (
	"time"
"bytes"
"fmt"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	serviceCreateTimeout = 2 * time.Minute
)

var _ = Describe("Container Conformance Test", func() {
	var cl *client.Client

	BeforeEach(func() {
		// Setup the apiserver client
		cl = client.NewOrDie(&restclient.Config{Host: *apiServerAddress})
	})

	Describe("container conformance blackbox test", func() {
		Context("when running a container that terminates", func() {
			var terminateCase ConformanceContainer
			var terminateImage ConformanceImage
			BeforeEach(func() {
				var envs []api.EnvVar
				var env api.EnvVar

				env.Name = "USERNAME"
				env.Value = "tiesheng"
				envs = append(envs, env)

				env.Name = "PASSWORD"
				env.Value = "opensource"
				envs = append(envs, env)

				env.Name = "EMAIL"
				env.Value = "tieshengcheng@163.com"
				envs = append(envs, env)

				env.Name = "GIT_REPO"
				env.Value = "https://github.com/docker-library/hello-world.git"
				envs = append(envs, env)

				var sc api.SecurityContext
				pri := true
				sc.Privileged = &pri
				terminateCase = ConformanceContainer{
					Container: api.Container{
						Image:           "tiesheng/builder",
						Name:            "builder",
						Env: 		envs,
						SecurityContext: &sc,
						Args:	[]string{"tiesheng/20160315-ok"},
						ImagePullPolicy: api.PullNever,
					},
					Client:   cl,
					Phase:    api.PodSucceeded,
					NodeName: *nodeName,
				}
				terminateImage, _ = NewConformanceImage("docker", terminateCase.Container.Image)
			})
			It("it should start successfully [Conformance]", func() {
				err := terminateCase.Create()
				Expect(err).NotTo(HaveOccurred())
					Eventually(func() string {
				sinceTime := unversioned.NewTime(time.Now().Add(time.Duration(-1 * time.Hour)))
				rc, err := cl.Pods(api.NamespaceDefault).GetLogs("builder", &api.PodLogOptions{SinceTime: &sinceTime}).Stream()
				if err != nil {
					return ""
				}
				defer rc.Close()
				buf := new(bytes.Buffer)
				buf.ReadFrom(rc)
				fmt.Println(buf.String())
				return buf.String()
}, time.Second*200, time.Second*15).Should(Equal("'hello'"))
			})
		})
	})
})
