# AMPLIFY Central Agents SDK

The AMPLIFY Central Agent SDK provides APIs and utilities that developers can use to build Golang based applications to discover APIs hosted on remote API Gateway (for e.g. AWS, Azure, Axway API Manager etc.) and publish their representation in AMPLIFY Central as API server resources and Catalog items. The SDK can also be used to build applications that can monitor traffic events for discovered APIs and publish them to AMPLIFY Central API Observer.

The Agent SDK helps in reducing complexity in implementing against the direct AMPLIFY Central REST API interface and hides low level plumbing to provide discovery and traceability related features. 

## Installation 
Use the following command to install the AMPLIFY Central Agents SDK 

go get github.com/Axway/agent-sdk/

## Packages

| Name         | Description                                                                                                                                         |
| ------------ | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| agent        | This package holds the interface for agent initialization and managing discovered APIs                                                              |
| api          | This package provides client interface for making REST API calls                                                                                    |
| apic         | This package contains AMPLIFY Central service client                                                                                                |
| cache        | This package can be used to create an in-memory cache of items                                                                                      |
| cmd          | This package provides the implementation of the root command line processor                                                                         |
| config       | This package provides the base configuration required by Agent SDK to communicate with AMPLIFY Central                                              |
| filter       | This package provides the filter implementation to allow discovering APIs based on certain conditions                                               |
| notification | This package contains structs that can be used for creating notifications and subscribers to those notifications                                    |
| notify       | This package contains the subscription notification setup for the agents to send SMTP and/or webhook notification for subscription process outcomes |
| transaction  | This package holds definitions of event and interfaces to process them for traceability                                                             |
| traceability | This package provides the transport lumberjack/HTTP clients that can be used for building traceability agent                                        |
| util         | This package has SDK utility packages for use by all agents                                                                                         |


[Getting started to build discovery agent](./docs/discovery/index.md)

[Getting started to build traceability agent](./docs/traceability/index.md)

[Utilities](./docs/utilities/index.md)

## Sample projects
The developers can use the stubs packaged as zip file to build agents using the Agent SDK. The zip files contains code for sample discovery and traceability agent respectively, build scripts and instructions in README.md to make modifications to implement their own agents.

[Download the stub project with sample discovery agent - TBD]

[Download the stub project with sample traceability agent - TBD]
