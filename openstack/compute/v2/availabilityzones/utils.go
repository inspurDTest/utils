package availabilityzones

import (
	"github.com/inspurDTest/gophercloud"
	"github.com/inspurDTest/gophercloud/openstack/compute/v2/extensions/availabilityzones"
)

// ListAvailableAvailabilityZones is a convenience function that return a slice of available Availability Zones.
func ListAvailableAvailabilityZones(client *gophercloud.ServiceClient) ([]string, error) {
	var ret []string
	allPages, err := availabilityzones.List(client).AllPages()
	if err != nil {
		return ret, err
	}

	availabilityZoneInfo, err := availabilityzones.ExtractAvailabilityZones(allPages)
	if err != nil {
		return ret, err
	}

	for _, zoneInfo := range availabilityZoneInfo {
		if zoneInfo.ZoneState.Available {
			ret = append(ret, zoneInfo.ZoneName)
		}
	}
	return ret, nil
}
