//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Helper function for maintaining backwards compatibility
func convertAuthorizeSecurityGroupIngressResponse(b []byte) ([]byte, error) {
	var raw struct {
		Ingressrule []interface{} `json:"ingressrule"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	if len(raw.Ingressrule) != 1 {
		return b, nil
	}

	return json.Marshal(raw.Ingressrule[0])
}

// Helper function for maintaining backwards compatibility
func convertAuthorizeSecurityGroupEgressResponse(b []byte) ([]byte, error) {
	var raw struct {
		Egressrule []interface{} `json:"egressrule"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	if len(raw.Egressrule) != 1 {
		return b, nil
	}

	return json.Marshal(raw.Egressrule[0])
}

type AuthorizeSecurityGroupEgressParams struct {
	p map[string]interface{}
}

func (p *AuthorizeSecurityGroupEgressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["securitygroupid"]; found {
		u.Set("securitygroupid", v.(string))
	}
	if v, found := p.p["securitygroupname"]; found {
		u.Set("securitygroupname", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["usersecuritygrouplist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].account", i), k)
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].group", i), m[k])
		}
	}
	return u
}

func (p *AuthorizeSecurityGroupEgressParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetSecuritygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupid"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetSecuritygroupname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupname"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) SetUsersecuritygrouplist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usersecuritygrouplist"] = v
}

// You should always use this function to get a new AuthorizeSecurityGroupEgressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewAuthorizeSecurityGroupEgressParams() *AuthorizeSecurityGroupEgressParams {
	p := &AuthorizeSecurityGroupEgressParams{}
	p.p = make(map[string]interface{})
	return p
}

// Authorizes a particular egress rule for this security group
func (s *SecurityGroupService) AuthorizeSecurityGroupEgress(p *AuthorizeSecurityGroupEgressParams) (*AuthorizeSecurityGroupEgressResponse, error) {
	resp, err := s.cs.newRequest("authorizeSecurityGroupEgress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AuthorizeSecurityGroupEgressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		b, err = convertAuthorizeSecurityGroupEgressResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AuthorizeSecurityGroupEgressResponse struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type AuthorizeSecurityGroupIngressParams struct {
	p map[string]interface{}
}

func (p *AuthorizeSecurityGroupIngressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["securitygroupid"]; found {
		u.Set("securitygroupid", v.(string))
	}
	if v, found := p.p["securitygroupname"]; found {
		u.Set("securitygroupname", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["usersecuritygrouplist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].account", i), k)
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].group", i), m[k])
		}
	}
	return u
}

func (p *AuthorizeSecurityGroupIngressParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetSecuritygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupid"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetSecuritygroupname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupname"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) SetUsersecuritygrouplist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usersecuritygrouplist"] = v
}

// You should always use this function to get a new AuthorizeSecurityGroupIngressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewAuthorizeSecurityGroupIngressParams() *AuthorizeSecurityGroupIngressParams {
	p := &AuthorizeSecurityGroupIngressParams{}
	p.p = make(map[string]interface{})
	return p
}

// Authorizes a particular ingress rule for this security group
func (s *SecurityGroupService) AuthorizeSecurityGroupIngress(p *AuthorizeSecurityGroupIngressParams) (*AuthorizeSecurityGroupIngressResponse, error) {
	resp, err := s.cs.newRequest("authorizeSecurityGroupIngress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AuthorizeSecurityGroupIngressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		b, err = convertAuthorizeSecurityGroupIngressResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AuthorizeSecurityGroupIngressResponse struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type CreateSecurityGroupParams struct {
	p map[string]interface{}
}

func (p *CreateSecurityGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *CreateSecurityGroupParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateSecurityGroupParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateSecurityGroupParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateSecurityGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateSecurityGroupParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

// You should always use this function to get a new CreateSecurityGroupParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewCreateSecurityGroupParams(name string) *CreateSecurityGroupParams {
	p := &CreateSecurityGroupParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Creates a security group
func (s *SecurityGroupService) CreateSecurityGroup(p *CreateSecurityGroupParams) (*CreateSecurityGroupResponse, error) {
	resp, err := s.cs.newRequest("createSecurityGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateSecurityGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateSecurityGroupResponse struct {
	Account             string                            `json:"account"`
	Description         string                            `json:"description"`
	Domain              string                            `json:"domain"`
	Domainid            string                            `json:"domainid"`
	Egressrule          []CreateSecurityGroupResponseRule `json:"egressrule"`
	Id                  string                            `json:"id"`
	Ingressrule         []CreateSecurityGroupResponseRule `json:"ingressrule"`
	JobID               string                            `json:"jobid"`
	Jobstatus           int                               `json:"jobstatus"`
	Name                string                            `json:"name"`
	Project             string                            `json:"project"`
	Projectid           string                            `json:"projectid"`
	Tags                []Tags                            `json:"tags"`
	Virtualmachinecount int                               `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                     `json:"virtualmachineids"`
}

type CreateSecurityGroupResponseRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type DeleteSecurityGroupParams struct {
	p map[string]interface{}
}

func (p *DeleteSecurityGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *DeleteSecurityGroupParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DeleteSecurityGroupParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DeleteSecurityGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteSecurityGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *DeleteSecurityGroupParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

// You should always use this function to get a new DeleteSecurityGroupParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewDeleteSecurityGroupParams() *DeleteSecurityGroupParams {
	p := &DeleteSecurityGroupParams{}
	p.p = make(map[string]interface{})
	return p
}

// Deletes security group
func (s *SecurityGroupService) DeleteSecurityGroup(p *DeleteSecurityGroupParams) (*DeleteSecurityGroupResponse, error) {
	resp, err := s.cs.newRequest("deleteSecurityGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSecurityGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSecurityGroupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSecurityGroupResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias DeleteSecurityGroupResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListSecurityGroupsParams struct {
	p map[string]interface{}
}

func (p *ListSecurityGroupsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["securitygroupname"]; found {
		u.Set("securitygroupname", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *ListSecurityGroupsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListSecurityGroupsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListSecurityGroupsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListSecurityGroupsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListSecurityGroupsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSecurityGroupsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListSecurityGroupsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSecurityGroupsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSecurityGroupsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListSecurityGroupsParams) SetSecuritygroupname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupname"] = v
}

func (p *ListSecurityGroupsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListSecurityGroupsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

// You should always use this function to get a new ListSecurityGroupsParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewListSecurityGroupsParams() *ListSecurityGroupsParams {
	p := &ListSecurityGroupsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SecurityGroupService) GetSecurityGroupID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListSecurityGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSecurityGroups(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.SecurityGroups[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SecurityGroups {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SecurityGroupService) GetSecurityGroupByName(name string, opts ...OptionFunc) (*SecurityGroup, int, error) {
	id, count, err := s.GetSecurityGroupID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSecurityGroupByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SecurityGroupService) GetSecurityGroupByID(id string, opts ...OptionFunc) (*SecurityGroup, int, error) {
	p := &ListSecurityGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSecurityGroups(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.SecurityGroups[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SecurityGroup UUID: %s!", id)
}

// Lists security groups
func (s *SecurityGroupService) ListSecurityGroups(p *ListSecurityGroupsParams) (*ListSecurityGroupsResponse, error) {
	resp, err := s.cs.newRequest("listSecurityGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSecurityGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSecurityGroupsResponse struct {
	Count          int              `json:"count"`
	SecurityGroups []*SecurityGroup `json:"securitygroup"`
}

type SecurityGroup struct {
	Account             string              `json:"account"`
	Description         string              `json:"description"`
	Domain              string              `json:"domain"`
	Domainid            string              `json:"domainid"`
	Egressrule          []SecurityGroupRule `json:"egressrule"`
	Id                  string              `json:"id"`
	Ingressrule         []SecurityGroupRule `json:"ingressrule"`
	JobID               string              `json:"jobid"`
	Jobstatus           int                 `json:"jobstatus"`
	Name                string              `json:"name"`
	Project             string              `json:"project"`
	Projectid           string              `json:"projectid"`
	Tags                []Tags              `json:"tags"`
	Virtualmachinecount int                 `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}       `json:"virtualmachineids"`
}

type SecurityGroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type RevokeSecurityGroupEgressParams struct {
	p map[string]interface{}
}

func (p *RevokeSecurityGroupEgressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RevokeSecurityGroupEgressParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new RevokeSecurityGroupEgressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewRevokeSecurityGroupEgressParams(id string) *RevokeSecurityGroupEgressParams {
	p := &RevokeSecurityGroupEgressParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a particular egress rule from this security group
func (s *SecurityGroupService) RevokeSecurityGroupEgress(p *RevokeSecurityGroupEgressParams) (*RevokeSecurityGroupEgressResponse, error) {
	resp, err := s.cs.newRequest("revokeSecurityGroupEgress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevokeSecurityGroupEgressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type RevokeSecurityGroupEgressResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RevokeSecurityGroupIngressParams struct {
	p map[string]interface{}
}

func (p *RevokeSecurityGroupIngressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RevokeSecurityGroupIngressParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new RevokeSecurityGroupIngressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewRevokeSecurityGroupIngressParams(id string) *RevokeSecurityGroupIngressParams {
	p := &RevokeSecurityGroupIngressParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a particular ingress rule from this security group
func (s *SecurityGroupService) RevokeSecurityGroupIngress(p *RevokeSecurityGroupIngressParams) (*RevokeSecurityGroupIngressResponse, error) {
	resp, err := s.cs.newRequest("revokeSecurityGroupIngress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevokeSecurityGroupIngressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type RevokeSecurityGroupIngressResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
