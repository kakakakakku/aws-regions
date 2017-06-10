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
}

// Regions is a list of AWS region
type Regions []Region

func main() {
	app := cli.NewApp()
	app.Name = "aws-regions"
	app.Version = "0.1.0"
	app.HideHelp = false

	app.Action = func(c *cli.Context) {
		rs := Regions{}
		rs = append(rs, Region{"us-east-1", "US East (N. Virginia)", "米国東部（バージニア北部）"})
		rs = append(rs, Region{"us-east-2", "US East (Ohio)", "米国東部（オハイオ）"})
		rs = append(rs, Region{"us-west-1", "US West (N. California)", "米国西部（北カリフォルニア）"})
		rs = append(rs, Region{"us-west-2", "US West (Oregon)", "米国西部（オレゴン）"})
		rs = append(rs, Region{"ca-central-1", "Canada (Central)", "カナダ（中部）"})
		rs = append(rs, Region{"eu-west-1", "EU (Ireland)", "欧州（アイルランド）"})
		rs = append(rs, Region{"eu-central-1", "EU (Frankfurt)", "欧州（フランクフルト）"})
		rs = append(rs, Region{"eu-west-2", "EU (London)", "欧州（ロンドン）"})
		rs = append(rs, Region{"ap-northeast-1", "Asia Pacific (Tokyo)", "アジアパシフィック（東京）"})
		rs = append(rs, Region{"ap-northeast-2", "Asia Pacific (Seoul)", "アジアパシフィック（ソウル）"})
		rs = append(rs, Region{"ap-southeast-1", "Asia Pacific (Singapore)", "アジアパシフィック（シンガポール）"})
		rs = append(rs, Region{"ap-southeast-2", "Asia Pacific (Sydney)", "アジアパシフィック（シドニー）"})
		rs = append(rs, Region{"ap-south-1", "Asia Pacific (Mumbai)", "アジアパシフィック（ムンバイ）"})
		rs = append(rs, Region{"sa-east-1", "South America (São Paulo)", "南米（サンパウロ）"})

		for _, r := range rs {
			fmt.Printf("%-16s %-27s %s\n", r.Code, r.EnName, r.JaName)
		}
	}

	app.Run(os.Args)
}
