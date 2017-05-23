package catalog

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/satori/uuid"
	"net/http"
)

// FileType enumerates the values for file type.
type FileType string

const (
	// Assembly specifies the assembly state for file type.
	Assembly FileType = "Assembly"
	// Nodeploy specifies the nodeploy state for file type.
	Nodeploy FileType = "Nodeploy"
	// Resource specifies the resource state for file type.
	Resource FileType = "Resource"
)

// DataLakeAnalyticsCatalogCredentialCreateParameters is data Lake Analytics
// catalog credential creation parameters.
type DataLakeAnalyticsCatalogCredentialCreateParameters struct {
	Password *string `json:"password,omitempty"`
	URI      *string `json:"uri,omitempty"`
	UserID   *string `json:"userId,omitempty"`
}

// DataLakeAnalyticsCatalogCredentialDeleteParameters is data Lake Analytics
// catalog credential deletion parameters.
type DataLakeAnalyticsCatalogCredentialDeleteParameters struct {
	Password *string `json:"password,omitempty"`
}

// DataLakeAnalyticsCatalogCredentialUpdateParameters is data Lake Analytics
// catalog credential update parameters.
type DataLakeAnalyticsCatalogCredentialUpdateParameters struct {
	Password    *string `json:"password,omitempty"`
	NewPassword *string `json:"newPassword,omitempty"`
	URI         *string `json:"uri,omitempty"`
	UserID      *string `json:"userId,omitempty"`
}

// DataLakeAnalyticsCatalogSecretCreateOrUpdateParameters is data Lake
// Analytics catalog secret creation and update parameters. This is deprecated
// and will be removed in the next release. Please use
// DataLakeAnalyticsCatalogCredentialCreateOrUpdateParameters instead.
type DataLakeAnalyticsCatalogSecretCreateOrUpdateParameters struct {
	Password *string `json:"password,omitempty"`
	URI      *string `json:"uri,omitempty"`
}

// DdlName is a Data Lake Analytics DDL name item.
type DdlName struct {
	FirstPart  *string `json:"firstPart,omitempty"`
	SecondPart *string `json:"secondPart,omitempty"`
	ThirdPart  *string `json:"thirdPart,omitempty"`
	Server     *string `json:"server,omitempty"`
}

// EntityID is a Data Lake Analytics catalog entity identifier object.
type EntityID struct {
	Name    *DdlName   `json:"name,omitempty"`
	Version *uuid.UUID `json:"version,omitempty"`
}

// ExternalTable is a Data Lake Analytics catalog external table item.
type ExternalTable struct {
	TableName  *string   `json:"tableName,omitempty"`
	DataSource *EntityID `json:"dataSource,omitempty"`
}

// Item is a Data Lake Analytics catalog item.
type Item struct {
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
}

// ItemList is a Data Lake Analytics catalog item list.
type ItemList struct {
	NextLink *string `json:"nextLink,omitempty"`
}

// TypeFieldInfo is a Data Lake Analytics catalog type field information item.
type TypeFieldInfo struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// USQLAssembly is a Data Lake Analytics catalog U-SQL Assembly.
type USQLAssembly struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string                       `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID                    `json:"version,omitempty"`
	DatabaseName       *string                       `json:"databaseName,omitempty"`
	Name               *string                       `json:"assemblyName,omitempty"`
	ClrName            *string                       `json:"clrName,omitempty"`
	IsVisible          *bool                         `json:"isVisible,omitempty"`
	IsUserDefined      *bool                         `json:"isUserDefined,omitempty"`
	Files              *[]USQLAssemblyFileInfo       `json:"files,omitempty"`
	Dependencies       *[]USQLAssemblyDependencyInfo `json:"dependencies,omitempty"`
}

// USQLAssemblyClr is a Data Lake Analytics catalog U-SQL assembly CLR item.
type USQLAssemblyClr struct {
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	Name               *string    `json:"assemblyClrName,omitempty"`
	ClrName            *string    `json:"clrName,omitempty"`
}

// USQLAssemblyDependencyInfo is a Data Lake Analytics catalog U-SQL dependency
// information item.
type USQLAssemblyDependencyInfo struct {
	EntityID *EntityID `json:"entityId,omitempty"`
}

// USQLAssemblyFileInfo is a Data Lake Analytics catalog U-SQL assembly file
// information item.
type USQLAssemblyFileInfo struct {
	Type         FileType `json:"type,omitempty"`
	OriginalPath *string  `json:"originalPath,omitempty"`
	ContentPath  *string  `json:"contentPath,omitempty"`
}

// USQLAssemblyList is a Data Lake Analytics catalog U-SQL assembly CLR item
// list.
type USQLAssemblyList struct {
	autorest.Response `json:"-"`
	NextLink          *string            `json:"nextLink,omitempty"`
	Value             *[]USQLAssemblyClr `json:"value,omitempty"`
}

// USQLAssemblyListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLAssemblyList) USQLAssemblyListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLCredential is a Data Lake Analytics catalog U-SQL credential item.
type USQLCredential struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	Name               *string    `json:"credentialName,omitempty"`
}

// USQLCredentialList is a Data Lake Analytics catalog U-SQL credential item
// list.
type USQLCredentialList struct {
	autorest.Response `json:"-"`
	NextLink          *string           `json:"nextLink,omitempty"`
	Value             *[]USQLCredential `json:"value,omitempty"`
}

// USQLCredentialListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLCredentialList) USQLCredentialListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLDatabase is a Data Lake Analytics catalog U-SQL database item.
type USQLDatabase struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	Name               *string    `json:"databaseName,omitempty"`
}

// USQLDatabaseList is a Data Lake Analytics catalog U-SQL database item list.
type USQLDatabaseList struct {
	autorest.Response `json:"-"`
	NextLink          *string         `json:"nextLink,omitempty"`
	Value             *[]USQLDatabase `json:"value,omitempty"`
}

// USQLDatabaseListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLDatabaseList) USQLDatabaseListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLDirectedColumn is a Data Lake Analytics catalog U-SQL directed column
// item.
type USQLDirectedColumn struct {
	Name       *string `json:"name,omitempty"`
	Descending *bool   `json:"descending,omitempty"`
}

// USQLDistributionInfo is a Data Lake Analytics catalog U-SQL distribution
// information object.
type USQLDistributionInfo struct {
	Type         *int32                `json:"type,omitempty"`
	Keys         *[]USQLDirectedColumn `json:"keys,omitempty"`
	Count        *int32                `json:"count,omitempty"`
	DynamicCount *int32                `json:"dynamicCount,omitempty"`
}

// USQLExternalDataSource is a Data Lake Analytics catalog U-SQL external
// datasource item.
type USQLExternalDataSource struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	Name               *string    `json:"externalDataSourceName,omitempty"`
	Provider           *string    `json:"provider,omitempty"`
	ProviderString     *string    `json:"providerString,omitempty"`
	PushdownTypes      *[]string  `json:"pushdownTypes,omitempty"`
}

// USQLExternalDataSourceList is a Data Lake Analytics catalog U-SQL external
// datasource item list.
type USQLExternalDataSourceList struct {
	autorest.Response `json:"-"`
	NextLink          *string                   `json:"nextLink,omitempty"`
	Value             *[]USQLExternalDataSource `json:"value,omitempty"`
}

// USQLExternalDataSourceListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLExternalDataSourceList) USQLExternalDataSourceListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLIndex is a Data Lake Analytics catalog U-SQL table index item.
type USQLIndex struct {
	Name              *string               `json:"name,omitempty"`
	IndexKeys         *[]USQLDirectedColumn `json:"indexKeys,omitempty"`
	Columns           *[]string             `json:"columns,omitempty"`
	DistributionInfo  *USQLDistributionInfo `json:"distributionInfo,omitempty"`
	PartitionFunction *uuid.UUID            `json:"partitionFunction,omitempty"`
	PartitionKeyList  *[]string             `json:"partitionKeyList,omitempty"`
	StreamNames       *[]string             `json:"streamNames,omitempty"`
	IsColumnstore     *bool                 `json:"isColumnstore,omitempty"`
	IndexID           *int32                `json:"indexId,omitempty"`
	IsUnique          *bool                 `json:"isUnique,omitempty"`
}

// USQLPackage is a Data Lake Analytics catalog U-SQL package item.
type USQLPackage struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	SchemaName         *string    `json:"schemaName,omitempty"`
	Name               *string    `json:"packageName,omitempty"`
	Definition         *string    `json:"definition,omitempty"`
}

// USQLPackageList is a Data Lake Analytics catalog U-SQL package item list.
type USQLPackageList struct {
	autorest.Response `json:"-"`
	NextLink          *string        `json:"nextLink,omitempty"`
	Value             *[]USQLPackage `json:"value,omitempty"`
}

// USQLPackageListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLPackageList) USQLPackageListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLProcedure is a Data Lake Analytics catalog U-SQL procedure item.
type USQLProcedure struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	SchemaName         *string    `json:"schemaName,omitempty"`
	Name               *string    `json:"procName,omitempty"`
	Definition         *string    `json:"definition,omitempty"`
}

// USQLProcedureList is a Data Lake Analytics catalog U-SQL procedure item
// list.
type USQLProcedureList struct {
	autorest.Response `json:"-"`
	NextLink          *string          `json:"nextLink,omitempty"`
	Value             *[]USQLProcedure `json:"value,omitempty"`
}

// USQLProcedureListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLProcedureList) USQLProcedureListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLSchema is a Data Lake Analytics catalog U-SQL schema item.
type USQLSchema struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	Name               *string    `json:"schemaName,omitempty"`
}

// USQLSchemaList is a Data Lake Analytics catalog U-SQL schema item list.
type USQLSchemaList struct {
	autorest.Response `json:"-"`
	NextLink          *string       `json:"nextLink,omitempty"`
	Value             *[]USQLSchema `json:"value,omitempty"`
}

// USQLSchemaListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLSchemaList) USQLSchemaListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLSecret is a Data Lake Analytics catalog U-SQL secret item.
type USQLSecret struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	Name               *string    `json:"secretName,omitempty"`
	CreationTime       *date.Time `json:"creationTime,omitempty"`
	URI                *string    `json:"uri,omitempty"`
	Password           *string    `json:"password,omitempty"`
}

// USQLTable is a Data Lake Analytics catalog U-SQL table item.
type USQLTable struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string               `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID            `json:"version,omitempty"`
	DatabaseName       *string               `json:"databaseName,omitempty"`
	SchemaName         *string               `json:"schemaName,omitempty"`
	Name               *string               `json:"tableName,omitempty"`
	ColumnList         *[]USQLTableColumn    `json:"columnList,omitempty"`
	IndexList          *[]USQLIndex          `json:"indexList,omitempty"`
	PartitionKeyList   *[]string             `json:"partitionKeyList,omitempty"`
	ExternalTable      *ExternalTable        `json:"externalTable,omitempty"`
	DistributionInfo   *USQLDistributionInfo `json:"distributionInfo,omitempty"`
}

// USQLTableColumn is a Data Lake Analytics catalog U-SQL table column item.
type USQLTableColumn struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// USQLTableList is a Data Lake Analytics catalog U-SQL table item list.
type USQLTableList struct {
	autorest.Response `json:"-"`
	NextLink          *string      `json:"nextLink,omitempty"`
	Value             *[]USQLTable `json:"value,omitempty"`
}

// USQLTableListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLTableList) USQLTableListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLTablePartition is a Data Lake Analytics catalog U-SQL table partition
// item.
type USQLTablePartition struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	SchemaName         *string    `json:"schemaName,omitempty"`
	Name               *string    `json:"partitionName,omitempty"`
	ParentName         *DdlName   `json:"parentName,omitempty"`
	IndexID            *int32     `json:"indexId,omitempty"`
	Label              *[]string  `json:"label,omitempty"`
	CreateDate         *date.Time `json:"createDate,omitempty"`
}

// USQLTablePartitionList is a Data Lake Analytics catalog U-SQL table
// partition item list.
type USQLTablePartitionList struct {
	autorest.Response `json:"-"`
	NextLink          *string               `json:"nextLink,omitempty"`
	Value             *[]USQLTablePartition `json:"value,omitempty"`
}

// USQLTablePartitionListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLTablePartitionList) USQLTablePartitionListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLTableStatistics is a Data Lake Analytics catalog U-SQL table statistics
// item.
type USQLTableStatistics struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	SchemaName         *string    `json:"schemaName,omitempty"`
	TableName          *string    `json:"tableName,omitempty"`
	Name               *string    `json:"statisticsName,omitempty"`
	UserStatName       *string    `json:"userStatName,omitempty"`
	StatDataPath       *string    `json:"statDataPath,omitempty"`
	CreateTime         *date.Time `json:"createTime,omitempty"`
	UpdateTime         *date.Time `json:"updateTime,omitempty"`
	IsUserCreated      *bool      `json:"isUserCreated,omitempty"`
	IsAutoCreated      *bool      `json:"isAutoCreated,omitempty"`
	HasFilter          *bool      `json:"hasFilter,omitempty"`
	FilterDefinition   *string    `json:"filterDefinition,omitempty"`
	ColNames           *[]string  `json:"colNames,omitempty"`
}

// USQLTableStatisticsList is a Data Lake Analytics catalog U-SQL table
// statistics item list.
type USQLTableStatisticsList struct {
	autorest.Response `json:"-"`
	NextLink          *string                `json:"nextLink,omitempty"`
	Value             *[]USQLTableStatistics `json:"value,omitempty"`
}

// USQLTableStatisticsListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLTableStatisticsList) USQLTableStatisticsListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLTableType is a Data Lake Analytics catalog U-SQL table type item.
type USQLTableType struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string          `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID       `json:"version,omitempty"`
	DatabaseName       *string          `json:"databaseName,omitempty"`
	SchemaName         *string          `json:"schemaName,omitempty"`
	Name               *string          `json:"typeName,omitempty"`
	TypeFamily         *string          `json:"typeFamily,omitempty"`
	CSharpName         *string          `json:"cSharpName,omitempty"`
	FullCSharpName     *string          `json:"fullCSharpName,omitempty"`
	SystemTypeID       *int32           `json:"systemTypeId,omitempty"`
	UserTypeID         *int32           `json:"userTypeId,omitempty"`
	SchemaID           *int32           `json:"schemaId,omitempty"`
	PrincipalID        *int32           `json:"principalId,omitempty"`
	IsNullable         *bool            `json:"isNullable,omitempty"`
	IsUserDefined      *bool            `json:"isUserDefined,omitempty"`
	IsAssemblyType     *bool            `json:"isAssemblyType,omitempty"`
	IsTableType        *bool            `json:"isTableType,omitempty"`
	IsComplexType      *bool            `json:"isComplexType,omitempty"`
	Columns            *[]TypeFieldInfo `json:"columns,omitempty"`
}

// USQLTableTypeList is a Data Lake Analytics catalog U-SQL table type item
// list.
type USQLTableTypeList struct {
	autorest.Response `json:"-"`
	NextLink          *string          `json:"nextLink,omitempty"`
	Value             *[]USQLTableType `json:"value,omitempty"`
}

// USQLTableTypeListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLTableTypeList) USQLTableTypeListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLTableValuedFunction is a Data Lake Analytics catalog U-SQL table valued
// function item.
type USQLTableValuedFunction struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	SchemaName         *string    `json:"schemaName,omitempty"`
	Name               *string    `json:"tvfName,omitempty"`
	Definition         *string    `json:"definition,omitempty"`
}

// USQLTableValuedFunctionList is a Data Lake Analytics catalog U-SQL table
// valued function item list.
type USQLTableValuedFunctionList struct {
	autorest.Response `json:"-"`
	NextLink          *string                    `json:"nextLink,omitempty"`
	Value             *[]USQLTableValuedFunction `json:"value,omitempty"`
}

// USQLTableValuedFunctionListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLTableValuedFunctionList) USQLTableValuedFunctionListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLType is a Data Lake Analytics catalog U-SQL type item.
type USQLType struct {
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	SchemaName         *string    `json:"schemaName,omitempty"`
	Name               *string    `json:"typeName,omitempty"`
	TypeFamily         *string    `json:"typeFamily,omitempty"`
	CSharpName         *string    `json:"cSharpName,omitempty"`
	FullCSharpName     *string    `json:"fullCSharpName,omitempty"`
	SystemTypeID       *int32     `json:"systemTypeId,omitempty"`
	UserTypeID         *int32     `json:"userTypeId,omitempty"`
	SchemaID           *int32     `json:"schemaId,omitempty"`
	PrincipalID        *int32     `json:"principalId,omitempty"`
	IsNullable         *bool      `json:"isNullable,omitempty"`
	IsUserDefined      *bool      `json:"isUserDefined,omitempty"`
	IsAssemblyType     *bool      `json:"isAssemblyType,omitempty"`
	IsTableType        *bool      `json:"isTableType,omitempty"`
	IsComplexType      *bool      `json:"isComplexType,omitempty"`
}

// USQLTypeList is a Data Lake Analytics catalog U-SQL type item list.
type USQLTypeList struct {
	autorest.Response `json:"-"`
	NextLink          *string     `json:"nextLink,omitempty"`
	Value             *[]USQLType `json:"value,omitempty"`
}

// USQLTypeListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLTypeList) USQLTypeListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// USQLView is a Data Lake Analytics catalog U-SQL view item.
type USQLView struct {
	autorest.Response  `json:"-"`
	ComputeAccountName *string    `json:"computeAccountName,omitempty"`
	Version            *uuid.UUID `json:"version,omitempty"`
	DatabaseName       *string    `json:"databaseName,omitempty"`
	SchemaName         *string    `json:"schemaName,omitempty"`
	Name               *string    `json:"viewName,omitempty"`
	Definition         *string    `json:"definition,omitempty"`
}

// USQLViewList is a Data Lake Analytics catalog U-SQL view item list.
type USQLViewList struct {
	autorest.Response `json:"-"`
	NextLink          *string     `json:"nextLink,omitempty"`
	Value             *[]USQLView `json:"value,omitempty"`
}

// USQLViewListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client USQLViewList) USQLViewListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}
