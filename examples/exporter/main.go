package main

import (
	"fmt"
	"os"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/Gleipnir-Technology/arcgis-go/examples"
	"github.com/Gleipnir-Technology/arcgis-go/fieldseeker"
)

// Export all of the records
func main() {
	fs, err := examples.FieldSeekerFromFlags()
	if err != nil {
		fmt.Println("Failed to create FS:", err)
		os.Exit(1)
	}

	fs.EnsureHasServiceInfo()
	for _, layer := range fs.FeatureServer.Layers {
		if layer.Name != "TrapLocation" {
			continue
		}
		err := downloadAllRecords(fs, layer)
		if err != nil {
			fmt.Println("Failed: %v", err)
			os.Exit(2)
		}
	}
}

func downloadAllRecords(fs *fieldseeker.FieldSeeker, layer arcgis.Layer) error {
	fmt.Printf("%v %v\n", layer.ID, layer.Name)
	count, err := fs.Arcgis.QueryCount(fs.ServiceName, layer.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Need to get %v records\n", count.Count)

	output, err := os.Create(fmt.Sprintf("records/%v-%v.json", layer.Name, layer.ID))
	if err != nil {
		return err
	}
	offset := 0
	for {
		query := arcgis.NewQuery()
		query.ResultRecordCount = fs.FeatureServer.MaxRecordCount
		query.ResultOffset = offset
		query.OutFields = "*"
		query.Where = "1=1"
		qr, err := fs.Arcgis.QueryRaw(
			fs.ServiceName,
			layer.ID,
			query)
		if err != nil {
			fmt.Printf("Failure: %v", err)
			os.Exit(1)
		}
		b, err := output.Write(qr)
		if err != nil {
			return err
		}
		fmt.Printf("Wrote %v bytes\n", b)
		offset += query.ResultRecordCount
		if offset > count.Count {
			break
		}
	}

	return nil
}
