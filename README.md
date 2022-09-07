## LAPF - Overlapping Finder

> Version: v0.0.7

------

## How does it work?

```bash
NAME:
   Overlapping Finder - Overlapping Finder (a.k.a "lapf") is a binary-tool made and built in golang to find possible overlapping CIDR Block notation through cloud providers (AWS)

USAGE:
   Overlapping Finder [global options] command [command options] [arguments...]

COMMANDS:
   ipv4     lapf ipv4 192.168.0.0/24
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h        show help (default: false)
   --output value    Output format (default: text, json)
   --provider value  Cloud Provider (default: aws)
```

------

## How to run it?

Retrieve the latest executable from GitHub releases and set up the execution permissions

```bash
curl -s https://api.github.com/repos/hatzelencio/lapf/releases/latest \
| grep "browser_download_url" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget -qi -
chmod u+x lapf*
```

## Overlap

### RAW Output

```bash
./lapf --provider aws ipv4 --region us-west-1 10.0.0.0/24
```

```text
[X][10.0.0.0/24]	is overlapping at: "aws" (vpc-xxa)
[ ][10.0.0.0/24]	is not overlapping at: "aws" (vpc-xxb)
[ ][10.0.0.0/24]	is not overlapping at: "aws" (vpc-xxc)
```

### JSON Output

```bash
./lapf --provider aws --output json \
  ipv4 --region us-west-1 \
  10.0.0.0/8
```

```json
[
  {
    "CloudNetwork": {
      "Name": "vpc-aaaa",
      "ProviderName": "aws",
      "CidrBlock": "10.0.0.0/24"
    },
    "CurrentCidr": "10.0.0.0/8",
    "IsOverlap": true
  }
]
```

### Streaming  pipeline results with other accounts

> [WARN] Bear in mind, flag `--output json` must be set up before pipelining your command, otherwise, you can not be able to performance it.


```bash
./lapf --output json --provider aws \
  ipv4 --profile blue --region us-west-1 10.0.0.0/24 | \
./lapf --provider aws \
  ipv4 --profile green --region us-east-1
```

```text
[ ][10.0.0.0/24]	is not overlapping at: "aws GreenAccount" (vpc-zyw = Default [172.31.0.0/16])
[ ][10.0.0.0/24]	is not overlapping at: "aws GreenAccount" (vpc-zyv = GreenVpc [10.120.0.0/16])
[ ][10.0.0.0/24]	is not overlapping at: "aws BlueAccount" (vpc-abd = Default [172.31.0.0/16])
[X][10.0.0.0/24]	is overlapping at:     "aws BlueAccount" (vpc-abc = BlueVpc [10.0.0.0/16])
```

------
## Is CIDR Block private?
### RAW Output

```bash
./lapf ensure cidr 10.0.0.0/16 172.10.0.0/22 192.168.1.0/24
```

```text
[✓][192.168.1.0/24] is private
[x][172.10.0.0/22] is not private
[✓][10.0.0.0/16] is private
```

### JSON Output

> Bear in mind, flag `--show-ip-list` only works over output json

```bash
./lapf --output json ensure cidr --show-ip-list 11.0.0.0/30 10.10.20.0/24
```

```json
[
  {
    "cidr": "10.10.20.0/24",
    "isPrivate": true
  },
  {
    "cidr": "11.0.0.0/30",
    "isPrivate": false,
    "publicIPList": [
      "11.0.0.0",
      "11.0.0.1",
      "11.0.0.2",
      "11.0.0.3"
    ]
  }
]
```

------

## Futures improvements

- [x] Ensuring input data
- [x] Fix default account credentials for AWS Provider
- [ ] Support for ipv6
- [ ] Adding other cloud providers like Azure, GCP
- [ ] Adding testing
- [ ] Adding the installation command   
