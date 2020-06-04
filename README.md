# thehive-sentinel-integration

This repository contains code that can be used to integrate Azure Sentinel, a cloud native SIEM, with TheHive, a security incident management plaform.

It consists of the following components:
- [thehive-sentinel-incidents](thehive-sentinel-incidents): a Python script to pull Azure Sentinel incidents into TheHive as alerts
- [thehive-sentinel-hooks](thehive-sentinel-hooks): a Golang application packaged in a Docker container that calls Azure Logic Apps based on the the output from TheHive Webhooks
- [sentinel-logicapps](sentinel-logicapps): Azure Logic Apps to interact with Azure Sentinel
- [deploy-thehive-sentinel-integration](deploy-thehive-sentinel-integration): Deploy a TheHive/Cortex instance with the Azure Sentinel ingetration with Terraform

## License

thehive-sentinel-integration is released under the GNU GENERAL PUBLIC LICENSE v3 [GPL-3.0](LICENSE).