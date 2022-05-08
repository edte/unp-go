// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-05-08 17:02
// @description:
package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"syscall"
)

// https://www.cems.uwe.ac.uk/~irjohnso/linsock/Book%20Notes/Appendices/Data%20Tables/General%20Socket%20Options.html

// 使用 getsockopt 获取 socket 选项的默认值

type args struct {
	pro string

	name     string
	comment  string
	optLevel int
	optName  int
}

var tests = []args{
	{
		pro:      "socket",
		name:     "SO_ACCEPTCONN",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_ACCEPTCONN,
	},
	{
		pro:      "socket",
		name:     "SO_ATTACH_FILTER",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_ATTACH_FILTER,
	},
	{
		pro:      "socket",
		name:     "SO_BINDTODEVICE",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_BINDTODEVICE,
	},
	{
		pro:      "socket",
		name:     "SO_BROADCAST",
		comment:  "支持广播",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_BROADCAST,
	},
	{
		pro:      "socket",
		name:     "SO_BSDCOMPAT",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_BSDCOMPAT,
	},
	{
		pro:      "socket",
		name:     "SO_DEBUG",
		comment:  "支持跟踪 tcp 报文",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_DEBUG,
	},
	{
		pro:      "socket",
		name:     "SO_DETACH_FILTER",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_DETACH_FILTER,
	},
	{
		pro:      "socket",
		name:     "SO_DOMAIN",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_DOMAIN,
	},
	{
		pro:      "socket",
		name:     "SO_DONTROUTE",
		comment:  "不走底层路由",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_DONTROUTE,
	},
	{
		pro:      "socket",
		name:     "SO_ERROR",
		comment:  "获取 socket 错误",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_ERROR,
	},
	{
		pro:      "socket",
		name:     "SO_KEEPALIVE",
		comment:  "采用心跳机制",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_KEEPALIVE,
	},
	{
		pro:      "socket",
		name:     "SO_LINGER",
		comment:  "close 时同时把缓冲区剩下的数据发送",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_LINGER,
	},
	{
		pro:      "socket",
		name:     "SO_MARK",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_MARK,
	},
	{
		pro:      "socket",
		name:     "SO_NO_CHECK",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_NO_CHECK,
	},
	{
		pro:      "socket",
		name:     "SO_OOBINLINE",
		comment:  "带外数据留在输入队列中",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_OOBINLINE,
	},
	{
		pro:      "socket",
		name:     "SO_PASSCRED",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_PASSCRED,
	},
	{
		pro:      "socket",
		name:     "SO_PASSSEC",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_PASSSEC,
	},
	{
		pro:      "socket",
		name:     "SO_PEERCRED",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_PEERCRED,
	},
	{
		pro:      "socket",
		name:     "SO_PEERNAME",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_PEERNAME,
	},
	{
		pro:      "socket",
		name:     "SO_PEERSEC",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_PEERSEC,
	},
	{
		pro:      "socket",
		name:     "SO_PRIORITY",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_PRIORITY,
	},
	{
		pro:      "socket",
		name:     "SO_PROTOCOL",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_PROTOCOL,
	},
	{
		pro:      "socket",
		name:     "SO_RCVBUF",
		comment:  "接收缓冲区大小",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_RCVBUF,
	},
	{
		pro:      "socket",
		name:     "SO_RCVBUFFORCE",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_RCVBUFFORCE,
	},
	{
		pro:      "socket",
		name:     "SO_RCVLOWAT",
		comment:  "接受低水位标记",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_RCVLOWAT,
	},
	{
		pro:      "socket",
		name:     "SO_RCVTIMEO",
		comment:  "接受超时",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_RCVTIMEO,
	},
	{
		pro:      "socket",
		name:     "SO_REUSEADDR",
		comment:  "复用地址",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_REUSEADDR,
	},
	{
		pro:      "socket",
		name:     "SO_REUSEPORT",
		comment:  "完全多路复用端口",
		optLevel: syscall.SOL_SOCKET,
		optName:  unix.SO_REUSEPORT,
	},
	{
		pro:      "socket",
		name:     "SO_RXQ_OVFL",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_RXQ_OVFL,
	},
	{
		pro:      "socket",
		name:     "SO_SECURITY_AUTHENTICATION",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_SECURITY_AUTHENTICATION,
	},
	{
		pro:      "socket",
		name:     "SO_SECURITY_ENCRYPTION_NETWORK",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_SECURITY_ENCRYPTION_NETWORK,
	},
	{
		pro:      "socket",
		name:     "SO_SECURITY_ENCRYPTION_TRANSPORT",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_SECURITY_ENCRYPTION_TRANSPORT,
	},
	{
		pro:      "socket",
		name:     "SO_SNDBUF",
		comment:  "发送缓冲区大小",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_SNDBUF,
	},
	{
		pro:      "socket",
		name:     "SO_SNDBUFFORCE",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_SNDBUFFORCE,
	},
	{
		pro:      "socket",
		name:     "SO_SNDLOWAT",
		comment:  "发送低水位标记",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_SNDLOWAT,
	},
	{
		pro:      "socket",
		name:     "SO_SNDTIMEO",
		comment:  "发送超时",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_SNDTIMEO,
	},
	{
		pro:      "socket",
		name:     "SO_TIMESTAMP",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_TIMESTAMP,
	},
	{
		pro:      "socket",
		name:     "SO_TIMESTAMPING",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_TIMESTAMPING,
	},
	{
		pro:      "socket",
		name:     "SO_TIMESTAMPNS",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_TIMESTAMPNS,
	},
	{
		pro:      "socket",
		name:     "SO_TYPE",
		comment:  "获取 socket 类型",
		optLevel: syscall.SOL_SOCKET,
		optName:  syscall.SO_TYPE,
	},
	{
		pro:      "ip",
		name:     "IP_ADD_MEMBERSHIP",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_ADD_MEMBERSHIP,
	},
	{
		pro:      "ip",
		name:     "IP_ADD_SOURCE_MEMBERSHIP",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_ADD_SOURCE_MEMBERSHIP,
	},
	{
		pro:      "ip",
		name:     "IP_BLOCK_SOURCE",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_BLOCK_SOURCE,
	},
	{
		pro:      "ip",
		name:     "IP_DEFAULT_MULTICAST_LOOP",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_DEFAULT_MULTICAST_LOOP,
	},
	{
		pro:      "ip",
		name:     "IP_DEFAULT_MULTICAST_TTL",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_DEFAULT_MULTICAST_TTL,
	},
	{
		pro:      "ip",
		name:     "IP_DF",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_DF,
	},
	{
		pro:      "ip",
		name:     "IP_DROP_MEMBERSHIP",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_DROP_MEMBERSHIP,
	},
	{
		pro:      "ip",
		name:     "IP_DROP_SOURCE_MEMBERSHIP",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_DROP_SOURCE_MEMBERSHIP,
	},
	{
		pro:      "ip",
		name:     "IP_FREEBIND",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_FREEBIND,
	},
	{
		pro:      "ip",
		name:     "IP_HDRINCL",
		comment:  "需要构造 ip 首部",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_HDRINCL,
	},
	{
		pro:      "ip",
		name:     "IP_IPSEC_POLICY",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_IPSEC_POLICY,
	},
	{
		pro:      "ip",
		name:     "IP_MAXPACKET",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MAXPACKET,
	},
	{
		pro:      "ip",
		name:     "IP_MAX_MEMBERSHIPS",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MAX_MEMBERSHIPS,
	},
	{
		pro:      "ip",
		name:     "IP_MF",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MF,
	},
	{
		pro:      "ip",
		name:     "IP_MINTTL",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MINTTL,
	},
	{
		pro:      "ip",
		name:     "IP_MSFILTER",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MSFILTER,
	},
	{
		pro:      "ip",
		name:     "IP_MSS",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MSS,
	},
	{
		pro:      "ip",
		name:     "IP_MTU",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MTU,
	},
	{
		pro:      "ip",
		name:     "IP_MTU_DISCOVER",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MTU_DISCOVER,
	},
	{
		pro:      "ip",
		name:     "IP_MULTICAST_IF",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MULTICAST_IF,
	},
	{
		pro:      "ip",
		name:     "IP_MULTICAST_LOOP",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MULTICAST_LOOP,
	},
	{
		pro:      "ip",
		name:     "IP_MULTICAST_TTL",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_MULTICAST_TTL,
	},
	{
		pro:      "ip",
		name:     "IP_OFFMASK",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_OFFMASK,
	},
	{
		pro:      "ip",
		name:     "IP_OPTIONS",
		comment:  "ip 选项",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_OPTIONS,
	},
	{
		pro:      "ip",
		name:     "IP_ORIGDSTADDR",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_ORIGDSTADDR,
	},
	{
		pro:      "ip",
		name:     "IP_PASSSEC",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_PASSSEC,
	},
	{
		pro:      "ip",
		name:     "IP_PKTINFO",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_PKTINFO,
	},
	{
		pro:      "ip",
		name:     "IP_PKTOPTIONS",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_PKTOPTIONS,
	},
	{
		pro:      "ip",
		name:     "IP_PMTUDISC",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_PMTUDISC,
	},
	{
		pro:      "ip",
		name:     "IP_PMTUDISC_DO",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_PMTUDISC_DO,
	},
	{
		pro:      "ip",
		name:     "IP_PMTUDISC_PROBE",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_PMTUDISC_PROBE,
	},
	{
		pro:      "ip",
		name:     "IP_PMTUDISC_DONT",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_PMTUDISC_DONT,
	},
	{
		pro:      "ip",
		name:     "IP_PMTUDISC_WANT",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_PMTUDISC_WANT,
	},
	{
		pro:      "ip",
		name:     "IP_RECVERR",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_RECVERR,
	},
	{
		pro:      "ip",
		name:     "IP_RECVOPTS",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_RECVOPTS,
	},
	{
		pro:      "ip",
		name:     "IP_RECVORIGDSTADDR",
		comment:  "recvmsg 获取 ip 目的地址",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_RECVORIGDSTADDR,
	},
	{
		pro:      "ip",
		name:     "IP_RECVRETOPTS",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_RECVRETOPTS,
	},
	{
		pro:      "ip",
		name:     "IP_RECVTOS",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_RECVTOS,
	},
	{
		pro:      "ip",
		name:     "IP_RECVTTL",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_RECVTTL,
	},
	{
		pro:      "ip",
		name:     "IP_RETOPTS",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_RETOPTS,
	},
	{
		pro:      "ip",
		name:     "IP_RF",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_RF,
	},
	{
		pro:      "ip",
		name:     "IP_ROUTER_ALERT",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_ROUTER_ALERT,
	},
	{
		pro:      "ip",
		name:     "IP_TOS",
		comment:  "设置 ip 首部服务字段",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_TOS,
	},
	{
		pro:      "ip",
		name:     "IP_TRANSPARENT",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_TRANSPARENT,
	},
	{
		pro:      "ip",
		name:     "IP_TTL",
		comment:  "获取 ttl",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_TTL,
	},
	{
		pro:      "ip",
		name:     "IP_UNBLOCK_SOURCE",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_UNBLOCK_SOURCE,
	},
	{
		pro:      "ip",
		name:     "IP_XFRM_POLICY",
		optLevel: syscall.IPPROTO_IP,
		optName:  syscall.IP_XFRM_POLICY,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_2292DSTOPTS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_2292DSTOPTS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_2292HOPLIMIT",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_2292HOPLIMIT,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_2292DSTOPTS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_2292DSTOPTS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_2292HOPOPTS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_2292HOPOPTS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_2292PKTINFO",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_2292PKTINFO,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_2292PKTOPTIONS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_2292PKTOPTIONS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_2292RTHDR",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_2292RTHDR,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_ADDRFORM",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_ADDRFORM,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_ADD_MEMBERSHIP",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_ADD_MEMBERSHIP,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_AUTHHDR",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_AUTHHDR,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_CHECKSUM",
		comment:  "指定校验和偏移",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_CHECKSUM,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_DROP_MEMBERSHIP",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_DROP_MEMBERSHIP,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_DSTOPTS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_DSTOPTS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_HOPLIMIT",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_HOPLIMIT,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_HOPOPTS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_HOPOPTS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_IPSEC_POLICY",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_IPSEC_POLICY,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_JOIN_ANYCAST",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_JOIN_ANYCAST,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_JOIN_GROUP",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_JOIN_GROUP,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_LEAVE_ANYCAST",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_LEAVE_ANYCAST,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_LEAVE_GROUP",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_LEAVE_GROUP,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_MTU",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_MTU,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_MTU_DISCOVER",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_MTU_DISCOVER,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_MULTICAST_HOPS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_MULTICAST_HOPS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_MULTICAST_IF",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_MULTICAST_IF,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_MULTICAST_LOOP",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_MULTICAST_LOOP,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_NEXTHOP",
		comment:  "指定下一跳套接字结构",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_NEXTHOP,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_PKTINFO",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_PKTINFO,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_PMTUDISC_DO",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_PMTUDISC_DO,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_PMTUDISC_DONT",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_PMTUDISC_DONT,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_PMTUDISC_PROBE",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_PMTUDISC_PROBE,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_PMTUDISC_WANT",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_PMTUDISC_WANT,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RECVDSTOPTS",
		comment:  "recvmsg 获取 ipv6 目的地址",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RECVDSTOPTS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RECVERR",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RECVERR,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RECVHOPLIMIT",
		comment:  "recvmsg 获取跳限字段",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RECVHOPLIMIT,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RECVHOPOPTS",
		comment:  "recvmsg 获取步跳选项",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RECVHOPOPTS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RECVPKTINFO",
		comment:  "recvmsg 获取目的 ipv6 地址和到达接口索引",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RECVPKTINFO,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RECVRTHDR",
		comment:  "recvmsg 获取接受路由首部",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RECVRTHDR,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RECVTCLASS",
		comment:  "recvmsg 获取接受流通类别",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RECVTCLASS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_ROUTER_ALERT",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_ROUTER_ALERT,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RTHDR",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RTHDR,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RTHDRDSTOPTS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RTHDRDSTOPTS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RTHDR_LOOSE",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RTHDR_LOOSE,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RTHDR_STRICT",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RTHDR_STRICT,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RTHDR_TYPE_0",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RTHDR_TYPE_0,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RXDSTOPTS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RXDSTOPTS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_RXHOPOPTS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_RXHOPOPTS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_TCLASS",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_TCLASS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_UNICAST_HOPS",
		comment:  "获取 ipv6 ttl",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_UNICAST_HOPS,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_V6ONLY",
		comment:  "只执行 ipv6",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_V6ONLY,
	},
	{
		pro:      "ipv6",
		name:     "IPV6_XFRM_POLICY",
		optLevel: syscall.IPPROTO_IPV6,
		optName:  syscall.IPV6_XFRM_POLICY,
	},
	{
		pro:      "tcp",
		name:     "TCP_CONGESTION",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_CONGESTION,
	},
	{
		pro:      "tcp",
		name:     "TCP_CORK",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_CORK,
	},
	{
		pro:      "tcp",
		name:     "TCP_DEFER_ACCEPT",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_DEFER_ACCEPT,
	},
	{
		pro:      "tcp",
		name:     "TCP_INFO",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_INFO,
	},
	{
		pro:      "tcp",
		name:     "TCP_KEEPCNT",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_KEEPCNT,
	},
	{
		pro:      "tcp",
		name:     "TCP_KEEPIDLE",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_KEEPIDLE,
	},
	{
		pro:      "tcp",
		name:     "TCP_KEEPINTVL",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_KEEPINTVL,
	},
	{
		pro:      "tcp",
		name:     "TCP_LINGER2",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_LINGER2,
	},
	{
		pro:      "tcp",
		name:     "TCP_MAXSEG",
		comment:  "获取或设置 mss",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_MAXSEG,
	},
	{
		pro:      "tcp",
		name:     "TCP_MAXWIN",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_MAXWIN,
	},
	{
		pro:      "tcp",
		name:     "TCP_MAX_WINSHIFT",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_MAX_WINSHIFT,
	},
	{
		pro:      "tcp",
		name:     "TCP_MD5SIG",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_MD5SIG,
	},
	{
		pro:      "tcp",
		name:     "TCP_MD5SIG_MAXKEYLEN",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_MD5SIG_MAXKEYLEN,
	},
	{
		pro:      "tcp",
		name:     "TCP_MSS",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_MSS,
	},
	{
		pro:      "tcp",
		name:     "TCP_NODELAY",
		comment:  "禁用 Nagle 算法",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_NODELAY,
	},
	{
		pro:      "tcp",
		name:     "TCP_QUICKACK",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_QUICKACK,
	},
	{
		pro:      "tcp",
		name:     "TCP_SYNCNT",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_SYNCNT,
	},
	{
		pro:      "tcp",
		name:     "TCP_WINDOW_CLAMP",
		optLevel: syscall.IPPROTO_TCP,
		optName:  syscall.TCP_WINDOW_CLAMP,
	},
}

func main() {
	for _, t := range tests {
		socket(t)
	}
}

func socket(t args) {
	var fd int
	var err error

	switch t.pro {
	case "socket":
		fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
		break
	case "tcp":
		fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
		break
	case "ip":
		fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
		break
	case "ipv6":
		fd, err = syscall.Socket(syscall.AF_INET6, syscall.SOCK_STREAM, 0)
		break
	}

	if err != nil {
		panic(err)
	}

	val, err := syscall.GetsockoptInt(fd, t.optLevel, t.optName)
	if err != nil {
		fmt.Printf("%33s: %10s  ->  %15s\n", t.name, "off", t.comment)

	} else {
		fmt.Printf("%33s: %10d  ->  %15s\n", t.name, val, t.comment)
	}

	if err = syscall.Close(fd); err != nil {
		panic(err)
	}
}
