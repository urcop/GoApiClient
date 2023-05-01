# Brawl Stars API Client

Module to get list of brawlers or detail about brawler

### Preparing

1. API KEY:
To get `API_KEY` you need to go to https://developer.brawlstars.com/#/
2. Get module from github: 
```bash
go get github.com/urcop/GoApiClient
```
3. Export your API KEY
```bash
export API_KEY=<YOUR_API_KEY>
```

### Example `main.go`

```go
package main

import (
	"fmt"
	"github.com/urcop/GoApiClient"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	client, err := GoAPIClient.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	brawlers, err := client.GetBrawlers()
	if err != nil {
		log.Fatal(err)
	}

	for _, brawler := range brawlers {
		fmt.Println(brawler.Info())
	}

	brawler, err := client.GetBrawler("16000069")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(brawler.Info())
}

```

### Output
```bash
[ID] 16000000 | [NAME] SHELLY
[ID] 16000001 | [NAME] COLT
[ID] 16000002 | [NAME] BULL
[ID] 16000003 | [NAME] BROCK
[ID] 16000004 | [NAME] RICO
[ID] 16000005 | [NAME] SPIKE
[ID] 16000006 | [NAME] BARLEY
[ID] 16000007 | [NAME] JESSIE
[ID] 16000008 | [NAME] NITA
[ID] 16000009 | [NAME] DYNAMIKE
[ID] 16000010 | [NAME] EL PRIMO
[ID] 16000011 | [NAME] MORTIS
[ID] 16000012 | [NAME] CROW
[ID] 16000013 | [NAME] POCO
[ID] 16000014 | [NAME] BO
[ID] 16000015 | [NAME] PIPER
[ID] 16000016 | [NAME] PAM
[ID] 16000017 | [NAME] TARA
[ID] 16000018 | [NAME] DARRYL
[ID] 16000019 | [NAME] PENNY
[ID] 16000020 | [NAME] FRANK
[ID] 16000021 | [NAME] GENE
[ID] 16000022 | [NAME] TICK
[ID] 16000023 | [NAME] LEON
[ID] 16000024 | [NAME] ROSA
[ID] 16000025 | [NAME] CARL
[ID] 16000026 | [NAME] BIBI
[ID] 16000027 | [NAME] 8-BIT
[ID] 16000028 | [NAME] SANDY
[ID] 16000029 | [NAME] BEA
[ID] 16000030 | [NAME] EMZ
[ID] 16000031 | [NAME] MR. P
[ID] 16000032 | [NAME] MAX
[ID] 16000034 | [NAME] JACKY
[ID] 16000035 | [NAME] GALE
[ID] 16000036 | [NAME] NANI
[ID] 16000037 | [NAME] SPROUT
[ID] 16000038 | [NAME] SURGE
[ID] 16000039 | [NAME] COLETTE
[ID] 16000040 | [NAME] AMBER
[ID] 16000041 | [NAME] LOU
[ID] 16000042 | [NAME] BYRON
[ID] 16000043 | [NAME] EDGAR
[ID] 16000044 | [NAME] RUFFS
[ID] 16000045 | [NAME] STU
[ID] 16000046 | [NAME] BELLE
[ID] 16000047 | [NAME] SQUEAK
[ID] 16000048 | [NAME] GROM
[ID] 16000049 | [NAME] BUZZ
[ID] 16000050 | [NAME] GRIFF
[ID] 16000051 | [NAME] ASH
[ID] 16000052 | [NAME] MEG
[ID] 16000053 | [NAME] LOLA
[ID] 16000054 | [NAME] FANG
[ID] 16000056 | [NAME] EVE
[ID] 16000057 | [NAME] JANET
[ID] 16000058 | [NAME] BONNIE
[ID] 16000059 | [NAME] OTIS
[ID] 16000060 | [NAME] SAM
[ID] 16000061 | [NAME] GUS
[ID] 16000062 | [NAME] BUSTER
[ID] 16000063 | [NAME] CHESTER
[ID] 16000064 | [NAME] GRAY
[ID] 16000065 | [NAME] MANDY
[ID] 16000066 | [NAME] R-T
[ID] 16000067 | [NAME] WILLOW
[ID] 16000068 | [NAME] MAISIE
[ID] 16000069 | [NAME] HANK
```