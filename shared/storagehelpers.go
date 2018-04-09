package shared

import (
	"log"

	storage "github.com/Azure/azure-sdk-for-go/storage"
)

//good samples here: https://github.com/luigialves/sample-golang-with-azure-table-storage/blob/master/sample.go

// UpsertEntity upserts entity
func UpsertEntity(name string, ip string, node string, status string) {
	storageclient := GetStorageClient()

	tableservice := storageclient.GetTableService()

	table := tableservice.GetTableReference(TableName)
	table.Create(Timeout, storage.MinimalMetadata, nil)

	//partition key is the same as row key (pod's name)
	entity := table.GetEntityReference(name, name)

	props := make(map[string]interface{})

	if ip != "" {
		props["PublicIP"] = ip
	}

	if node != "" {
		props["NodeName"] = node
	}

	if status != "" {
		props["Status"] = status
	}

	entity.Properties = props

	if err := entity.InsertOrMerge(nil); err != nil {
		log.Fatalf("Cannot insert or merge entity due to %s", err)
	}
}

// GetEntity gets table entity
func GetEntity(name string) *storage.Entity {
	storageclient := GetStorageClient()

	tableservice := storageclient.GetTableService()

	table := tableservice.GetTableReference(TableName)
	table.Create(Timeout, storage.MinimalMetadata, nil)

	//partition key is the same as row key (pod's name)
	entity := table.GetEntityReference(name, name)

	if err := entity.Get(Timeout, storage.MinimalMetadata, nil); err != nil {
		log.Fatalf("Cannot get entity due to %s", err)
	}

	return entity
}

// DeleteEntity deletes table entity
func DeleteEntity(name string) {
	storageclient := GetStorageClient()

	tableservice := storageclient.GetTableService()

	table := tableservice.GetTableReference(TableName)
	table.Create(Timeout, storage.MinimalMetadata, nil)

	//partition key is the same as row key (pod's name)
	entity := table.GetEntityReference(name, name)

	if err := entity.Delete(true, nil); err != nil {
		log.Fatalf("Cannot delete entity due to %s", err)
	}
}