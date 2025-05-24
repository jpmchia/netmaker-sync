# PeerConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedIPs** | [**[]IpNet**](IPNet.md) | AllowedIPs specifies a list of allowed IP addresses in CIDR notation for this peer. | [optional] [default to null]
**Endpoint** | [***UdpAddr**](UDPAddr.md) |  | [optional] [default to null]
**PersistentKeepaliveInterval** | [***Duration**](Duration.md) |  | [optional] [default to null]
**PresharedKey** | [***Key**](Key.md) |  | [optional] [default to null]
**PublicKey** | [***Key**](Key.md) |  | [optional] [default to null]
**Remove** | **bool** | Remove specifies if the peer with this public key should be removed from a device&#39;s peer list. | [optional] [default to null]
**ReplaceAllowedIPs** | **bool** | ReplaceAllowedIPs specifies if the allowed IPs specified in this peer configuration should replace any existing ones, instead of appending them to the allowed IPs list. | [optional] [default to null]
**UpdateOnly** | **bool** | UpdateOnly specifies that an operation will only occur on this peer if the peer already exists as part of the interface. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


