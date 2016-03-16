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

	"k8s.io/kubernetes/pkg/api"
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
		Context("when testing images that exist", func() {
			var conformImages []ConformanceImage
			BeforeEach(func() {
				images := []string{
					"gcr.io/google_containers/busybox",
					"gcr.io/google_containers/eptest",
					"gcr.io/google_containers/mounttest",
					"gcr.io/google_containers/nettest",
					"gcr.io/google_containers/ubuntu",
				}
				for _, image := range images {
					conformImages = append(conformImages, NewConformanceImage("docker", image))
				}
			})
			It("it should pull successfully [Conformance]", func() {
				for _, image := range conformImages {
					if err := image.Pull(); err != nil {
						Expect(err).NotTo(HaveOccurred())
						break
					}
				}
			})
			It("it should list pulled images [Conformance]", func() {
			})
			It("it should remove successfully [Conformance]", func() {
				for _, image := range conformImages {
					if err := image.Remove(); err != nil {
						Expect(err).NotTo(HaveOccurred())
						break
					}
				}
			})
			It("it should not list removed images [Conformance]", func() {
			})
		})
	})
})
