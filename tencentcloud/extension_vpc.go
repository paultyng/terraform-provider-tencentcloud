package tencentcloud

/*
 all gate way types
 https://cloud.tencent.com/document/api/215/15824#Route
*/
const GATE_WAY_TYPE_CVM = "CVM"
const GATE_WAY_TYPE_VPN = "VPN"
const GATE_WAY_TYPE_DIRECTCONNECT = "DIRECTCONNECT"
const GATE_WAY_TYPE_PEERCONNECTION = "PEERCONNECTION"
const GATE_WAY_TYPE_SSLVPN = "SSLVPN"
const GATE_WAY_TYPE_HAVIP = "HAVIP"
const GATE_WAY_TYPE_NAT = "NAT"
const GATE_WAY_TYPE_NORMAL_CVM = "NORMAL_CVM"
const GATE_WAY_TYPE_EIP = "EIP"
const GATE_WAY_TYPE_CCN = "CCN"
const GATE_WAY_TYPE_LOCAL_GATEWAY = "LOCAL_GATEWAY"

var ALL_GATE_WAY_TYPES = []string{
	GATE_WAY_TYPE_CVM,
	GATE_WAY_TYPE_VPN,
	GATE_WAY_TYPE_DIRECTCONNECT,
	GATE_WAY_TYPE_PEERCONNECTION,
	GATE_WAY_TYPE_SSLVPN,
	GATE_WAY_TYPE_HAVIP,
	GATE_WAY_TYPE_NAT,
	GATE_WAY_TYPE_NORMAL_CVM,
	GATE_WAY_TYPE_EIP,
	GATE_WAY_TYPE_CCN,
	GATE_WAY_TYPE_LOCAL_GATEWAY,
}

const VPC_SERVICE_TYPE = "vpc"

/*
EIP
*/
const (
	EIP_STATUS_CREATING  = "CREATING"
	EIP_STATUS_BINDING   = "BINDING"
	EIP_STATUS_BIND      = "BIND"
	EIP_STATUS_UNBINDING = "UNBINDING"
	EIP_STATUS_UNBIND    = "UNBIND"
	EIP_STATUS_OFFLINING = "OFFLINING"
	EIP_STATUS_BIND_ENI  = "BIND_ENI"

	EIP_TYPE_EIP          = "EIP"
	EIP_TYPE_ANYCAST      = "AnycastEIP"
	EIP_TYPE_HIGH_QUALITY = "HighQualityEIP"
	EIP_TYPE_ANTI_DDOS    = "AntiDDoSEIP"

	EIP_ANYCAST_ZONE_GLOBAL   = "ANYCAST_ZONE_GLOBAL"
	EIP_ANYCAST_ZONE_OVERSEAS = "ANYCAST_ZONE_OVERSEAS"

	EIP_INTERNET_PROVIDER_BGP  = "BGP"
	EIP_INTERNET_PROVIDER_CMCC = "CMCC"
	EIP_INTERNET_PROVIDER_CTCC = "CTCC"
	EIP_INTERNET_PROVIDER_CUCC = "CUCC"

	EIP_RESOURCE_TYPE = "eip"

	EIP_TASK_STATUS_SUCCESS = "SUCCESS"
	EIP_TASK_STATUS_RUNNING = "RUNNING"
	EIP_TASK_STATUS_FAILED  = "FAILED"
)

var EIP_INTERNET_PROVIDER = []string{
	EIP_INTERNET_PROVIDER_BGP,
	EIP_INTERNET_PROVIDER_CMCC,
	EIP_INTERNET_PROVIDER_CTCC,
	EIP_INTERNET_PROVIDER_CUCC,
}

var EIP_TYPE = []string{
	EIP_TYPE_EIP,
	EIP_TYPE_ANYCAST,
	EIP_TYPE_HIGH_QUALITY,
	EIP_TYPE_ANTI_DDOS,
}

var EIP_ANYCAST_ZONE = []string{
	EIP_ANYCAST_ZONE_GLOBAL,
	EIP_ANYCAST_ZONE_OVERSEAS,
}

var EIP_AVAILABLE_PERIOD = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36}

// ENI
const (
	ENI_DESCRIBE_LIMIT = 100
)

const (
	ENI_STATE_PENDING   = "PENDING"
	ENI_STATE_AVAILABLE = "AVAILABLE"
	ENI_STATE_ATTACHING = "ATTACHING"
	ENI_STATE_DETACHING = "DETACHING"
	ENI_STATE_DELETING  = "DELETING"
)

const (
	ENI_IP_PENDING   = "PENDING"
	ENI_IP_AVAILABLE = "AVAILABLE"
	ENI_IP_ATTACHING = "ATTACHING"
	ENI_IP_DETACHING = "DETACHING"
	ENI_IP_DELETING  = "DELETING"
)

/*
NAT
*/

const (
	NAT_DESCRIBE_LIMIT = 100
	NAT_EIP_MAX_LIMIT  = 10
)

const (
	NAT_FAILED_STATE = "FAILED"
)

const (
	NAT_GATEWAY_TYPE_SUBNET            = "SUBNET"
	NAT_GATEWAY_TYPE_NETWORK_INTERFACE = "NETWORKINTERFACE"
)

/*
VPN
*/

const (
	VPN_DESCRIBE_LIMIT = 100
)

const (
	VPN_TASK_STATUS_SUCCESS = "SUCCESS"
	VPN_TASK_STATUS_RUNNING = "RUNNING"
	VPN_TASK_STATUS_FAILED  = "FAILED"
)

const (
	VPN_STATE_PENDING   = "PENDING"
	VPN_STATE_DELETING  = "DELETING"
	VPN_STATE_AVAILABLE = "AVAILABLE"
)

var VPN_STATE = []string{
	VPN_STATE_PENDING,
	VPN_STATE_DELETING,
	VPN_STATE_AVAILABLE,
}

const (
	VPN_PERIOD_PREPAID_RENEW_FLAG_AUTO_NOTIFY = "NOTIFY_AND_AUTO_RENEW"
	VPN_PERIOD_PREPAID_RENEW_FLAG_NOT         = "NOTIFY_AND_MANUAL_RENEW"
)

var VPN_PERIOD_PREPAID_RENEW_FLAG = []string{
	VPN_PERIOD_PREPAID_RENEW_FLAG_AUTO_NOTIFY,
	VPN_PERIOD_PREPAID_RENEW_FLAG_NOT,
}

const (
	VPN_CHARGE_TYPE_PREPAID          = "PREPAID"
	VPN_CHARGE_TYPE_POSTPAID_BY_HOUR = "POSTPAID_BY_HOUR"
)

var VPN_CHARGE_TYPE = []string{
	VPN_CHARGE_TYPE_PREPAID,
	VPN_CHARGE_TYPE_POSTPAID_BY_HOUR,
}

const (
	VPN_PURCHASE_PLAN_PRE_POST = "PREPAID_TO_POSTPAID"
)

var VPN_PURCHASE_PLAN = []string{
	VPN_PURCHASE_PLAN_PRE_POST,
}

const (
	VPN_RESTRICT_STATE_NORMAL  = "NORMAL"
	VPN_RESTRICT_STATE_ISOLATE = "PRETECIVELY_ISOLATED"
)

var VPN_RESTRICT_STATE = []string{
	VPN_RESTRICT_STATE_NORMAL,
	VPN_RESTRICT_STATE_ISOLATE,
}

const (
	VPN_IKE_PROPO_ENCRY_ALGORITHM_3DESCBC   = "3DES-CBC"
	VPN_IKE_PROPO_ENCRY_ALGORITHM_AESCBC128 = "AES-CBC-128"
	VPN_IKE_PROPO_ENCRY_ALGORITHM_AESCBC192 = "AES-CBS-192`"
	VPN_IKE_PROPO_ENCRY_ALGORITHM_AESCBC256 = "AES-CBC-256"
	VPN_IKE_PROPO_ENCRY_ALGORITHM_DESCBC    = "DES-CBC"
)

var VPN_IKE_PROPO_ENCRY_ALGORITHM = []string{
	VPN_IKE_PROPO_ENCRY_ALGORITHM_3DESCBC,
	VPN_IKE_PROPO_ENCRY_ALGORITHM_AESCBC128,
	VPN_IKE_PROPO_ENCRY_ALGORITHM_AESCBC192,
	VPN_IKE_PROPO_ENCRY_ALGORITHM_AESCBC256,
	VPN_IKE_PROPO_ENCRY_ALGORITHM_DESCBC,
}

const (
	VPN_IKE_PROPO_AUTHEN_ALGORITHM_SHA    = "SHA"
	VPN_IKE_PROPO_AUTHEN_ALGORITHM_MD5    = "MD5"
	VPN_IKE_PROPO_AUTHEN_ALGORITHM_SHA256 = "SHA-256"
)

var VPN_IKE_PROPO_AUTHEN_ALGORITHM = []string{
	VPN_IKE_PROPO_AUTHEN_ALGORITHM_SHA,
	VPN_IKE_PROPO_AUTHEN_ALGORITHM_MD5,
	VPN_IKE_PROPO_AUTHEN_ALGORITHM_SHA256,
}

const (
	VPN_IPSEC_INTEGRITY_ALGORITHM_SHA1   = "SHA1"
	VPN_IPSEC_INTEGRITY_ALGORITHM_MD5    = "MD5"
	VPN_IPSEC_INTEGRITY_ALGORITHM_SHA256 = "SHA-256"
)

var VPN_IPSEC_INTEGRITY_ALGORITHM = []string{
	VPN_IPSEC_INTEGRITY_ALGORITHM_SHA1,
	VPN_IPSEC_INTEGRITY_ALGORITHM_MD5,
	VPN_IPSEC_INTEGRITY_ALGORITHM_SHA256,
}

const (
	VPN_IKE_EXCHANGE_MODE_AGGRESSIVE = "AGGRESSIVE"
	VPN_IKE_EXCHANGE_MODE_MAIN       = "MAIN"
)

var VPN_IKE_EXCHANGE_MODE = []string{
	VPN_IKE_EXCHANGE_MODE_AGGRESSIVE,
	VPN_IKE_EXCHANGE_MODE_MAIN,
}

const (
	VPN_IKE_IDENTITY_ADDRESS = "ADDRESS"
	VPN_IKE_IDENTITY_FQDN    = "FQDN"
)

var VPN_IKE_IDENTITY = []string{
	VPN_IKE_IDENTITY_ADDRESS,
	VPN_IKE_IDENTITY_FQDN,
}

const (
	VPN_IKE_DH_GROUP_NAME_GROUP1  = "GROUP1"
	VPN_IKE_DH_GROUP_NAME_GROUP2  = "GROUP2"
	VPN_IKE_DH_GROUP_NAME_GROUP5  = "GROUP5"
	VPN_IKE_DH_GROUP_NAME_GROUP14 = "GROUP14"
	VPN_IKE_DH_GROUP_NAME_GROUP24 = "GROUP24"
)

var VPN_IKE_DH_GROUP_NAME = []string{
	VPN_IKE_DH_GROUP_NAME_GROUP1,
	VPN_IKE_DH_GROUP_NAME_GROUP2,
	VPN_IKE_DH_GROUP_NAME_GROUP5,
	VPN_IKE_DH_GROUP_NAME_GROUP14,
	VPN_IKE_DH_GROUP_NAME_GROUP24,
}

const (
	VPN_IPSEC_PFS_DH_GROUP_NAME_GROUP1  = "DH-GROUP1"
	VPN_IPSEC_PFS_DH_GROUP_NAME_GROUP2  = "DH-GROUP2"
	VPN_IPSEC_PFS_DH_GROUP_NAME_GROUP5  = "DH-GROUP5"
	VPN_IPSEC_PFS_DH_GROUP_NAME_GROUP14 = "DH-GROUP14"
	VPN_IPSEC_PFS_DH_GROUP_NAME_GROUP24 = "DH-GROUP24"
	VPN_IPSEC_PFS_DH_GROUP_NAME_NULL    = "NULL"
)

var VPN_IPSEC_PFS_DH_GROUP_NAME = []string{
	VPN_IPSEC_PFS_DH_GROUP_NAME_GROUP1,
	VPN_IPSEC_PFS_DH_GROUP_NAME_GROUP2,
	VPN_IPSEC_PFS_DH_GROUP_NAME_GROUP5,
	VPN_IPSEC_PFS_DH_GROUP_NAME_GROUP14,
	VPN_IPSEC_PFS_DH_GROUP_NAME_GROUP24,
	VPN_IPSEC_PFS_DH_GROUP_NAME_NULL,
}

const (
	VPN_IPSEC_ENCRY_ALGORITHM_3DESCBC   = "3DES-CBC"
	VPN_IPSEC_ENCRY_ALGORITHM_AESCBC128 = "AES-CBC-128"
	VPN_IPSEC_ENCRY_ALGORITHM_AESCBC192 = "AES-CBS-192`"
	VPN_IPSEC_ENCRY_ALGORITHM_AESCBC256 = "AES-CBC-256"
	VPN_IPSEC_ENCRY_ALGORITHM_DESCBC    = "DES-CBC"
	VPN_IPSEC_ENCRY_ALGORITHM_NULL      = "NULL"
)

var VPN_IPSEC_ENCRY_ALGORITHM = []string{
	VPN_IPSEC_ENCRY_ALGORITHM_3DESCBC,
	VPN_IPSEC_ENCRY_ALGORITHM_AESCBC128,
	VPN_IPSEC_ENCRY_ALGORITHM_AESCBC192,
	VPN_IPSEC_ENCRY_ALGORITHM_AESCBC256,
	VPN_IPSEC_ENCRY_ALGORITHM_DESCBC,
	VPN_IPSEC_ENCRY_ALGORITHM_NULL,
}

/*
HAVIP
*/

const (
	HAVIP_DESCRIBE_LIMIT = 100
)

/*
COMMON
*/
const (
	VPCNotFound             = "ResourceNotFound"
	VPCUnsupportedOperation = "UnsupportedOperation"
)

const (
	DPD_ACTION_CLEAR   = "clear"
	DPD_ACTION_RESTART = "restart"
)

var DPD_ACTIONS = []string{
	DPD_ACTION_CLEAR,
	DPD_ACTION_RESTART,
}
