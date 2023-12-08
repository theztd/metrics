package main

import (
	"flag"
	"fmt"
	"metrics/volumes"
	"strings"
)

func labelsToString(labels map[string]string) string {
	var parts []string

	for k, v := range labels {
		parts = append(parts, fmt.Sprintf("%s=\"%s\"", strings.ReplaceAll(k, ".", "_"), v))
	}

	return strings.Join(parts, ", ")
}

var (
	prefix string
)

func main() {

	flag.StringVar(&prefix, "p", "", "Metric prefix string")

	flag.Parse()

	// Handle prefix delimiter
	if len(prefix) > 2 && prefix[len(prefix)-1:] != "_" {
		prefix = prefix + "_"
	}

	// Generate metrics for each volume
	if volumes, err := volumes.Size(); err != nil {
		fmt.Println("ERR [volume.Size]:", err)
	} else {
		fmt.Printf("# HELP %sdocker_volume_size Size of local docker volumes in Bytes\n", prefix)
		fmt.Printf("# TYPE %sdocker_volume_size gauge\n", prefix)
		for _, v := range volumes {
			fmt.Printf("%sdocker_volume_size{vol_name=\"%s\", created=\"%s\", driver=\"%s\", %s} %d\n", prefix, v.Name, v.Created, v.Driver, labelsToString(v.Labels), v.SizeB)
		}

	}
}
