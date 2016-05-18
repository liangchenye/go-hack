
##Configs
|Key|Type|Description|Example|
|------|----|------| ----- |
| Port | int | The port of the Scheduler.| 8001 |
| ServerListFile | string | For simple cases, we put all the server infos into a file.| ["servers.conf"](#servers) |
| Debug | bool | Print the debug information on the screen| true, default to false |

## [Filesystem bundle](https://github.com/opencontainers/runtime-spec/blob/master/bundle.md)
- [x] config.json : This REQUIRED file, which MUST be named config.json.When the bundle is packaged up for distribution, this file MUST be included
- [ ] When the bundle is packaged up for distribution, this directory MUST be included.
- [x] This directory MUST be referenced from within the config.json file.

## [Runtime and Lifecycle](https://github.com/opencontainers/runtime-spec/blob/master/runtime.md)
- [x] The container's runtime environment MUST be created according to the configuration in config.json.
- [x] Any updates to config.json after container is running MUST not affect the container.
- [ ] [WIP](https://github.com/opencontainers/ocitools/pull/61) The prestart hooks MUST be invoked by the runtime.
      If any prestart hook fails, then the container MUST be stopped and the lifecycle continues at step 8.
- [x] The user specified process MUST be executed in the container.
- [ ] [WIP](https://github.com/opencontainers/ocitools/pull/61) poststart hooks MUST be invoked by the runtime.
       If any poststart hook fails, then the container MUST be stopped and the lifecycle continues at step 8
- [ ] The container MUST be destroyed by undoing the steps performed during create phase (step 2)
- [ ] [WIP](https://github.com/opencontainers/ocitools/pull/61) The poststop hooks MUST be invoked by the runtime and errors, if any, MAY be logged.
- [ ] [WIP](https://github.com/opencontainers/ocitools/pull/61) Query State: state <container-id>
- [ ] [WIP](https://github.com/opencontainers/ocitools/pull/61) This operation MUST generate an error if it is not provided the ID of a container.
- [ ] [WIP](https://github.com/opencontainers/ocitools/pull/61) This operation MUST return the state of a container as specified in the State section. In particular, the state MUST be serialized as JSON.

### Start: start <container-id> <path-to-bundle>
- [ ] [WIP](https://github.com/opencontainers/ocitools/pull/61) This operation MUST generate an error if it is not provided a path to the bundle and the container ID to associate with the container.
- [ ] [WIP](https://github.com/opencontainers/ocitools/pull/61) If the ID provided is not unique across all containers within the scope of the runtime, or is not valid in any other way, the implementation MUST generate an error.
- [x] Using the data in config.json, that are in the bundle's directory, this operation MUST create a new container.
This includes creating the relevant namespaces, resource limits, etc and configuring the appropriate capabilities for the container.
- [x] A new process within the scope of the container MUST be created as specified by the config.json file otherwise an error MUST be generated.
- [ ] [WIP](https://github.com/opencontainers/ocitools/pull/61) Attempting to start an already running container MUST have no effect on the container and MUST generate an error.	

### Stop: stop <container-id>
- [ ] This operation MUST generate an error if it is not provided the container ID.
- [ ] This operation MUST stop and delete a running container.
- [ ] Stopping a container MUST stop all of the processes running within the scope of the container.
- [ ] Deleting a container MUST delete the associated namespaces and resources associated with the container.
- [ ] Once a container is deleted, its id MAY be used by subsequent containers.
- [ ] Attempting to stop a container that is not running MUST have no effect on the container and MUST generate an error.	
 
### Exec: exec <container-id> <path-to-json>
- [ ] This operation MUST generate an error if it is not provided the container ID and a path to the JSON describing the process to start.
- [ ] The JSON describing the new process MUST adhere to the Process configuration definition.
This operation MUST create a new process within the scope of the container.
- [ ] If the container is not running then this operation MUST have no effect on the container and MUST generate an error.
Executing this operation multiple times MUST result in a new process each time.	Planed
- [ ] The stopping, or exiting, of these secondary process MUST have no effect on the state of the container.
In other words, a container (and its PID 1 process) MUST NOT be stopped due to the exiting of a secondary process.	Planed

## [Configuration](https://github.com/opencontainers/runtime-spec/blob/master/config.md)
- [x] The container's top-level directory MUST contain a configuration file called config.json.
- [x] ociVersion (string, required) must be in SemVer v2.0.0 format and specifies the version of the OpenContainer specification with which the bundle complies.	Included
- [x] path (string, required) Specifies the path to the root filesystem for the container, relative to the path where the manifest is. A directory MUST exist at the relative path declared by the field.	Included
- [x] readonly (bool, optional) If true then the root filesystem MUST be read-only inside the container. Defaults to false.	Included
- [ ] The runtime MUST mount entries in the listed order.	
- [x] Destination of mount point: path inside container.
- [x] Configuration | Mounts	type (string, required) Linux, filesystemtype argument supported by the kernel are listed in /proc/filesystems (e.g., "minix", "ext2", "ext3", "jfs", "xfs", "reiserfs", "msdos", "proc", "nfs", "iso9660"). Windows: ntfs	
- [x] source (string, required) a device name, but can also be a directory name or a dummy. Windows, the volume name that is the target of the mount point. \?\Volume{GUID}\ (on Windows source is called target)
- [x] Mounts	options (list of strings, optional) in the fstab format https://wiki.archlinux.org/index.php/Fstab.
- [x] cwd (string, required) is the working directory that will be set for the executable. This value MUST be an absolute path.
- [x] Process configuration
- [x] Platform-specific configuration	os (string, required)  $GOOS.
- [x] Platform-specific configuration	arch (string, required)  $GOARCH.
Planed
- [ ] Hooks	Hooks MUST be called in the listed order.
- [x] Hooks	path is required for a hook.
- [x] Hooks	args and env are optional.
- [ ] Hooks	timeout is the number of seconds before aborting the hook.

## [Linux-specific Container Configuration](https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md)
- [ ] Default File Systems	The following filesystems MUST be made available in each application's filesystem 
```
Path	Type
/proc	procfs
/sys	sysfs
/dev/pts	devpts
/dev/shm	tmpfs
```
- [ ] Namespaces Also, when a path is specified, a runtime MUST assume that the setup for that particular namespace has already been done and error out if the config specifies anything else related to that namespace.
- [ ] devices is an array specifying the list of devices that MUST be available in the container.In addition to any devices configured with this setting, the runtime MUST also supply
- [ ] Control groups 	The runtime MUST apply entries in the listed order.
- [ ] Control groups	You must specify at least one of weight or leafWeight in a given entry, and can specify both.	Planed
- [x] JSON	All configuration JSON MUST be encoded in UTF-8.
