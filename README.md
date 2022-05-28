## LAPF - Overlapping Finder

> Version: v0.0.1

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

Retrieve the latest executable from Github releases and set up the execution permissions

```bash
curl -s https://api.github.com/repos/hatzelencio/lapf/releases/latest \                                                                                                     ──(Fri,May27)─┘
| grep "browser_download_url" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget -qi -
chmod u+x lapf 
```

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

------

## Futures improvements

- [ ] Ensuring input data
- [ ] Support for ipv6
- [ ] Adding other cloud providers like Azure, GCP
- [ ] Adding testing
- [ ] Adding the installation command   
