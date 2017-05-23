package insightsclienteventcategories

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// EventCategoryCollection is a collection of event categories. Currently
// possible values are: Administrative, Security, ServiceHealth, Alert,
// Recommendation, Policy.
type EventCategoryCollection struct {
	autorest.Response `json:"-"`
	Value             *[]LocalizableString `json:"value,omitempty"`
}

// LocalizableString is the localizable string class.
type LocalizableString struct {
	Value          *string `json:"value,omitempty"`
	LocalizedValue *string `json:"localizedValue,omitempty"`
}
