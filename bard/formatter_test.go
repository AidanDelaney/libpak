/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bard_test

import (
	"fmt"
	"testing"

	"github.com/heroku/color"
	. "github.com/onsi/gomega"
	"github.com/paketoio/libpak/bard"
	"github.com/sclevine/spec"
)

func testFormatter(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	context("IdentityFormatter", func() {

		it("it formats name", func() {
			Expect(bard.IdentityFormatter{Name: "test-name"}.String()).
				To(Equal(color.New(color.Bold).Sprint("test-name")))
		})

		it("formats name and description", func() {
			Expect(bard.IdentityFormatter{Name: "test-name", Description: "test-description"}.String()).
				To(Equal(fmt.Sprintf("%s test-description", color.New(color.Bold).Sprint("test-name"))))
		})
	})

	context("LaunchConfigFormatter", func() {

		it("formats launch configuration", func() {
			l := bard.LaunchConfigFormatter{Name: "TEST_NAME", Default: "test-default"}
			Expect(l.String()).To(Equal(fmt.Sprintf("Set $TEST_NAME to configure. Default %s.", color.New(color.Italic).Sprint("test-default"))))
		})
	})

}