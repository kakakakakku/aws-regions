package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// Region is AWS region
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html
type Region struct {
	Code   string
	EnName string
	JaName string
	Az     string
}

// Regions is a list of AWS region
type Regions []Region

func main() {
	app := cli.NewApp()
	app.Name = "aws-regions"
	app.Version = "v1.2.0"
	app.HideHelp = false

	app.Action = func(c *cli.Context) {
		rs := Regions{}
		rs = append(rs, Region{"us-east-1", "US East (N. Virginia)", "米国東部（バージニア北部）", "6 AZs"})
		rs = append(rs, Region{"us-east-2", "US East (Ohio)", "米国東部（オハイオ）", "3 AZs"})
		rs = append(rs, Region{"us-west-1", "US West (N. California)", "米国西部（北カリフォルニア）", "3 AZs"})
		rs = append(rs, Region{"us-west-2", "US West (Oregon)", "米国西部（オレゴン）", "3 AZs"})
		rs = append(rs, Region{"ca-central-1", "Canada (Central)", "カナダ（中部）", "2 AZs"})

		rs = append(rs, Region{"cn-north-1", "China (Beijing)", "中国 (北京) リージョン", "2 AZs"})
		rs = append(rs, Region{"cn-northwest-1", "China (Ningxia)", "中国 (寧夏) リージョン", "3 AZs"})

		rs = append(rs, Region{"eu-central-1", "EU (Frankfurt)", "欧州（フランクフルト）", "3 AZs"})
		rs = append(rs, Region{"eu-west-1", "EU (Ireland)", "欧州（アイルランド）", "3 AZs"})
		rs = append(rs, Region{"eu-west-2", "EU (London)", "欧州（ロンドン）", "3 AZs"})
		rs = append(rs, Region{"eu-west-3", "EU (Paris)", "EU（パリ）", "3 AZs"})
		rs = append(rs, Region{"ap-northeast-1", "Asia Pacific (Tokyo)", "アジアパシフィック（東京）", "4 AZs"})
		rs = append(rs, Region{"ap-northeast-2", "Asia Pacific (Seoul)", "アジアパシフィック（ソウル）", "2 AZs"})
		rs = append(rs, Region{"ap-northeast-3", "Asia Pacific (Osaka-Local)", "アジアパシフィック（大阪: ローカル）", "1 AZs"})
		rs = append(rs, Region{"ap-southeast-1", "Asia Pacific (Singapore)", "アジアパシフィック（シンガポール）", "3 AZs"})
		rs = append(rs, Region{"ap-southeast-2", "Asia Pacific (Sydney)", "アジアパシフィック（シドニー）", "3 AZs"})
		rs = append(rs, Region{"ap-south-1", "Asia Pacific (Mumbai)", "アジアパシフィック（ムンバイ）", "2 AZs"})
		rs = append(rs, Region{"sa-east-1", "South America (São Paulo)", "南米（サンパウロ）", "3 AZs"})
		rs = append(rs, Region{"us-gov-east-1", "AWS GovCloud (US-East)", "AWS GovCloud (米国東部)", "3 AZs"})
		rs = append(rs, Region{"us-gov-west-1", "AWS GovCloud (US)", "AWS GovCloud (米国)", "3 AZs"})

		for _, r := range rs {
			fmt.Printf("%-16s %-7s %-28s %s\n", r.Code, r.Az, r.EnName, r.JaName)
		}
	}

	app.Run(os.Args)
}
