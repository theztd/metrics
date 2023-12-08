package volumes

import (
	"context"
	"metrics/utils"

	"github.com/docker/docker/api/types/volume"

	"github.com/docker/docker/client"
)

type Volume struct {
	Name    string
	SizeB   int64
	Created string
	Driver  string
	// Status  map[string]interface{}
	Labels map[string]string
}

func Size() (volumes []Volume, err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return volumes, err
	}

	// List of volumes
	volumeList, err := cli.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		return volumes, err
	}

	// Generate metrics
	for _, v := range volumeList.Volumes {
		vol := Volume{
			Name:    v.Name,
			SizeB:   utils.DirSize(v.Mountpoint),
			Created: v.CreatedAt,
			// Status:  v.Status,
			Driver: v.Driver,
			Labels: v.Labels,
		}
		volumes = append(volumes, vol)
	}

	return volumes, nil
}
