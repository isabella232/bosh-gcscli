/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/bosh-gcscli/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

const PublicBucketEnv = "PUBLIC_BUCKET_NAME"

var _ = Describe("Integration", func() {
	Context("read-only configuration", func() {
		public := os.Getenv(PublicBucketEnv)

		var ctx AssertContext
		BeforeEach(func() {
			Expect(public).ToNot(BeEmpty(),
				fmt.Sprintf(NoBucketMsg, PublicBucketEnv))

			ctx = NewAssertContext()
		})
		AfterEach(func() {
			ctx.Cleanup()
		})

		configurations := []TableEntry{
			Entry("Public bucket", &config.GCSCli{
				BucketName: public,
			}),
		}

		DescribeTable("Read-only Get works",
			func(config *config.GCSCli) {
				ctx.AddConfig(config)
				AssertReadonlyGetWorks(gcsCLIPath, ctx)
			},
			configurations...)

		DescribeTable("Read-only invalid Get should fail",
			func(config *config.GCSCli) {
				ctx.AddConfig(config)
				AsserReadOnlytGetNonexistentFails(gcsCLIPath, ctx)
			},
			configurations...)

		DescribeTable("Read-only Put should fail",
			func(config *config.GCSCli) {
				ctx.AddConfig(config)
				AssertReadOnlyPutFails(gcsCLIPath, ctx)
			},
			configurations...)

		DescribeTable("Read-only Delete should fail",
			func(config *config.GCSCli) {
				ctx.AddConfig(config)
				AssertReadOnlyDeleteFails(gcsCLIPath, ctx)
			},
			configurations...)
	})
})
