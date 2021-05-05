package config

import "fmt"

type BaseBucketConfig struct {
	// Unique Code of Bucket For Identification
	Id               string
	Name             string      `yaml:"name"`
	BucketType       string      `yaml:"type"`
	Path             string      `yaml:"path"`
	Config           interface{} `yaml:"config"`
	ImgMaxResolution string      `yaml:"img_max_res"`
	IgnoreMime       []string    `yaml:",flow"`
}

func InitBucketConfig(bucketConfig *BaseBucketConfig) {
	switch bucketConfig.BucketType {
	case "gcs":
		{
			fmt.Printf("[Config] %s Bucket sets to firebase", bucketConfig.Name)
			break
		}
	case "s3":
		{
			fmt.Printf("[Config] %s Bucket sets to s3", bucketConfig.Name)
			break
		}
	default:
		{
			fmt.Printf("[Config] %s Bucket sets to locally", bucketConfig.Name)
			break
		}
	}
}
