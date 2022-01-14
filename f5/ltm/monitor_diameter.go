// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/xxxlzj520/f5-bigip-rest-client/f5"

type MonitorDiameterConfigList struct {
	Items    []MonitorDiameterConfig `json:"items,omitempty"`
	Kind     string                  `json:"kind,omitempty"`
	SelfLink string                  `json:"selflink,omitempty"`
}

type MonitorDiameterConfig struct {
	AcctApplicationId               string `json:"acctApplicationId,omitempty"`
	AppService                      string `json:"appService,omitempty"`
	AuthApplicationId               string `json:"authApplicationId,omitempty"`
	DefaultsFrom                    string `json:"defaultsFrom,omitempty"`
	Description                     string `json:"description,omitempty"`
	Destination                     string `json:"destination,omitempty"`
	FullPath                        string `json:"fullPath,omitempty"`
	Generation                      int    `json:"generation,omitempty"`
	HostIpAddress                   string `json:"hostIpAddress,omitempty"`
	Interval                        int    `json:"interval,omitempty"`
	Kind                            string `json:"kind,omitempty"`
	ManualResume                    string `json:"manualResume,omitempty"`
	Name                            string `json:"name,omitempty"`
	OriginHost                      string `json:"originHost,omitempty"`
	OriginRealm                     string `json:"originRealm,omitempty"`
	Partition                       string `json:"partition,omitempty"`
	ProductName                     string `json:"productName,omitempty"`
	SelfLink                        string `json:"selfLink,omitempty"`
	TimeUntilUp                     int    `json:"timeUntilUp,omitempty"`
	Timeout                         int    `json:"timeout,omitempty"`
	UpInterval                      int    `json:"upInterval,omitempty"`
	VendorId                        string `json:"vendorId,omitempty"`
	VendorSpecificAcctApplicationId string `json:"vendorSpecificAcctApplicationId,omitempty"`
	VendorSpecificAuthApplicationId string `json:"vendorSpecificAuthApplicationId,omitempty"`
	VendorSpecificVendorId          string `json:"vendorSpecificVendorId,omitempty"`
}

const MonitorDiameterEndpoint = "/monitor/diameter"

type MonitorDiameterResource struct {
	c *f5.Client
}

func (r *MonitorDiameterResource) ListAll() (*MonitorDiameterConfigList, error) {
	var list MonitorDiameterConfigList
	if err := r.c.ReadQuery(BasePath+MonitorDiameterEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorDiameterResource) Get(id string) (*MonitorDiameterConfig, error) {
	var item MonitorDiameterConfig
	if err := r.c.ReadQuery(BasePath+MonitorDiameterEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorDiameterResource) Create(item MonitorDiameterConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorDiameterEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorDiameterResource) Edit(id string, item MonitorDiameterConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorDiameterEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorDiameterResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorDiameterEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
