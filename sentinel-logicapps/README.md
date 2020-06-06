# sentinel-logicapps

Sentinel-logicapps is a collection of Azure Logic Apps created to integrate TheHive with Azure Sentinel to be used together with [thehivesentinelhooks](../thehive-sentinel-hooks):

- sentinel-la-sendteamsmessage: Logic App to send a Teams message
- sentinel-la-closeincidentfromcase:  Logic App to close Sentinel incidents from a resolved TheHive case.
- sentinel-la-closeincidentfromalert: Logic App to close Sentinel incidents from a ignored or resolved TheHive alert
- sentinel-la-closeASCalert: Logic App to close Azure Security Center alerts
- sentinel-la-updateincident: Logic App to set a Sentinel incident to InProgress
