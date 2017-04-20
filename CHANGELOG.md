# Release 0.4.0

This release added support to work with managed disks.
It added a dependency on the new azure cli tool **az**
since manages disks can only be created using it.

## Cleanup

We have new tooling on the project to cleanup test
resources that sometimes leak. Infrastructure is not free ;-).

## Login

We need set new var **AZURE_SERVICE_PRINCIPAL** to authenticate with **Azure CLI 2.0**

## VM Create

By default now KLB use "Managed Disks" when create Virtual Machines.

We need remove **azure_vm_set_storageaccount()** or add **azure_vm_set_unmanageddisk()**

When specifying an existing NIC, do not specify NSG, public IP, VNet or subnet.

We need create Availset before and pass to **azure_vm_create()** 

### Deprecated

#### azure_vm_set_nic()

We need change **azure_vm_set_nic()** to **azure_vm_set_nics()**.

#### azure_vm_set_osdiskvhd()

### Removed

#### azure_vm_set_datadiskvhd

We need remove **azure_vm_set_datadiskvhd()** because this function doesn't exist anymore.

#### azure_vm_set_disablebootdiagnostics()

We need remove **azure_vm_set_disablebootdiagnostics()** because this function doesn't exist anymore.

#### azure_vm_set_bootdiagnosticsstorage()

We need remove **azure_vm_set_bootdiagnosticsstorage()** because this function doesn't exist anymore.

### Added

#### azure_vm_set_storagesku()

**azure_vm_set_storagesku(sku)** sets the SKU storage account of "Virtual Machine".

- `instance` is the name of the instance.
- `storagesku` is the the sku of storage account to persist VM. By default, only Standard_LRS and Premium_LRS are allowed.

### azure_disk_*

**azure_disk_new(name, group, location)**  creates a new instance of "managed disk".

- `name` is the name of the managed disk.
- `group` is name of resource group.
- `location` is the Azure Region.

**azure_disk_set_size(instance, size)** sets the size of "managed disk".

- `instance` is the name of the instance.
- `size` is the size in Gb of managed disk.

**azure_disk_set_sku(instance, sku)** sets the kind of "managed disk".

- `instance` is the name of the instance.
- `sku` is the underlying storage sku. 

**azure_disk_set_source(instance, source)** sets the source of "managed disk".
- `instance` is the name of the instance.
- `source` is the source to create the disk from, including a sas uri to a blob, managed disk id or name, or snapshot id or name.

**azure_disk_create(instance)** creates a new "managed disk".
- `instance` is the name of the instance.

### azure_vm_availset_*

**azure_vm_availset_new(name, group, location)** creates a new instance of Availset.
- `name` is the name of the Availset.
- `group` is name of resource group.
- `location` is the Azure Region.

**azure_vm_availset_set_faultdomain(instance, count)** sets Fault Domain of Availset.
- `instance` is the name of the instance.
- `count` is the Fault Domain count. Example: 2.

**azure_vm_availset_set_updatedomain(instance, count)** sets Update Domain of Availset.
- `instance` is the name of the instance.
- `count` is the Update Domain count. Example: 2.

**azure_vm_availset_set_unmanaged(instance)** sets Contained VMs should use unmanaged disks.
- `instance` is the name of the instance.

**azure_vm_availset_create(instance)** creates a Availset.
- `instance` is the name of the instance.

**azure_vm_availset_delete(name, group)** deletes a Availset.
- `name` is the name of the Availset.
- `group` is name of resource group.