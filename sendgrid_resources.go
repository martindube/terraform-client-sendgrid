package sendgrid_client

import (
    "strconv"
	"encoding/json"
    "log"
	"fmt"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

// Sendgrid  Template
type Template struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Sendgrid  Template Version
type TemplateVersion struct {
	Id           string `json:"id,omitempty"`
	TemplateId   string `json:"template_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Subject      string `json:"subject,omitempty"`
	HtmlContent  string `json:"html_content,omitempty"`
	PlainContent string `json:"plain_content,omitempty"`
	Active       int    `json:"active,omitempty"`
}

// Sendgrid DNS Entry from whitelabel domain responses.
type DNSRecord struct {
	Host    string  `json:"host,omitempty"`
	Type    string  `json:"type,omitempty"`
	Data    string  `json:"data,omitempty"`
	Valid   bool    `json:"valid,omitempty"`
}

type WhitelabelDomainDNS struct {
     // Fields returned for Create method
     MailCname   DNSRecord `json:"mail_cname,omitempty"`
     Spf         DNSRecord `json:"spf,omitempty"`
     Dkim1       DNSRecord `json:"dkim1,omitempty"`
     Dkim2       DNSRecord `json:"dkim2,omitempty"`
     // Fields returned for other methods
     MailServer   DNSRecord `json:"mail_server,omitempty"`
     SubdomainSpf DNSRecord `json:"subdomain_spf,omitempty"`
     DomainSpf    DNSRecord `json:"domain_spf,omitempty"`
     Dkim         DNSRecord `json:"dkim,omitempty"`
 }

// Sendgrid Whitelabel Domain
type WhitelabelDomain struct {
	Id                  int      `json:"id,omitempty"`
	Domain              string   `json:"domain,omitempty"`
	Subdomain           string   `json:"subdomain,omitempty"`
	Username            string   `json:"username,omitempty"`
	UserId              int      `json:"user_id,omitempty"`
	AutomaticSecurity   bool     `json:"automatic_security,omitempty"`
	CustomSpf           bool     `json:"custom_spf,omitempty"`
	Default             bool     `json:"default,omitempty"`
    Dns                 WhitelabelDomainDNS `json:"dns,omitempty"`
}

type WhitelabelDomainList []WhitelabelDomain

///////////////////////////////////////////////////
// Create a transactional template.
// POST /templates

func (client *Client) CreateTemplate(template *Template) (*Template, error) {
	request := sendgrid.GetRequest(client.apiKey, "/v3/templates", "")
	request.Method = "POST"
	var err error
	request.Body, err = client.GetBody(template)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateResponse(response)
	}
}

///////////////////////////////////////////////////
// Get all transactional templates.
// GET /templates

func (client *Client) GetAllTemplates() (bool, error) {
	request := sendgrid.GetRequest(client.apiKey, "/v3/templates", "")
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return false, err
	} else {
		err := processEmptyResponse(response)
		if err != nil {
			return false, nil
		} else {
			return true, nil
		}
	}
}

///////////////////////////////////////////////////
// Update a transactional template.
// PATCH /templates/{template_id}

func (client *Client) UpdateTemplate(id string, template *Template) error {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s", id), "")
	request.Method = "PATCH"
	var err error
	request.Body, err = client.GetBody(template)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return processEmptyResponse(response)
	}
}

///////////////////////////////////////////////////
// Get a single transactional template.
// GET /templates/{template_id}

func (client *Client) GetTemplate(id string) (*Template, error) {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s", id), "")
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateResponse(response)
	}
}

///////////////////////////////////////////////////
// Delete a template.
// DELETE /templates/{template_id}

func (client *Client) DeleteTemplate(id string) error {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s", id), "")
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return processEmptyResponse(response)
	}
}

///////////////////////////////////////////////////
// Create a new transactional template version.
// POST /templates/{template_id}/versions

func (client *Client) CreateTemplateVersion(version *TemplateVersion) (*TemplateVersion, error) {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s/versions", version.TemplateId), "")
	request.Method = "POST"
	var err error
	request.Body, err = client.GetBody(version)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateVersionResponse(response)
	}
}

///////////////////////////////////////////////////
// Update a transactional template version.
// PATCH /templates/{template_id}/versions/{version_id}

func (client *Client) UpdateTemplateVersion(id string, version *TemplateVersion) error {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s/versions/%s", version.TemplateId, id), "")
	request.Method = "PATCH"
	var err error
	request.Body, err = client.GetBody(version)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return processEmptyResponse(response)
	}
}

///////////////////////////////////////////////////
// Get a specific transactional template version.
// GET /templates/{template_id}/versions/{version_id}

func (client *Client) GetTemplateVersion(templateId, versionId string) (*TemplateVersion, error) {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s/versions/%s", templateId, versionId), "")
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateVersionResponse(response)
	}
}

///////////////////////////////////////////////////
// Delete a transactional template version.
// DELETE /templates/{template_id}/versions/{version_id}

func (client *Client) DeleteTemplateVersion(templateId, versionId string) error {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s/versions/%s", templateId, versionId), "")
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return processEmptyResponse(response)
	}
}

///////////////////////////////////////////////////
// Activate a transactional template version.
// POST /templates/{template_id}/versions/{version_id}/activate

func (client *Client) ActivateTemplateVersion(templateId, versionId string) (*TemplateVersion, error) {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s/versions/%s/activate", templateId, versionId), "")
	request.Method = "POST"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateVersionResponse(response)
	}
}

func processTemplateResponse(response *rest.Response) (*Template, error) {
	err := processEmptyResponse(response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var out Template
	err = json.Unmarshal([]byte(response.Body), &out)
	if err != nil {
		fmt.Println(response)
		fmt.Println("Unmarshal Template error: ", err)
		return nil, err
	}
	return &out, nil
}

func processTemplateVersionResponse(response *rest.Response) (*TemplateVersion, error) {
	err := processEmptyResponse(response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var out TemplateVersion
	err = json.Unmarshal([]byte(response.Body), &out)
	if err != nil {
		fmt.Println("Unmarshal TemplateVersion error: ", err)
		return nil, err
	}
	return &out, nil
}

func processEmptyResponse(response *rest.Response) error {
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	if response.StatusCode >= 300 {
		return fmt.Errorf("Error calling API: status code: %d", response.StatusCode)
	}
	return nil
}

///////////////////////////////////////////////////
// Create a whitelabel domain
// POST /whitelabel/domains

func (client *Client) CreateWhitelabelDomain(whitelabeldomain *WhitelabelDomain) (*WhitelabelDomain, error) {
	request := sendgrid.GetRequest(client.apiKey, "/v3/whitelabel/domains", "")
	request.Method = "POST"
	var err error
	request.Body, err = client.GetBody(whitelabeldomain)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
    prettyPrint(request)
	response, err := sendgrid.API(request)
    prettyPrint(response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processWhitelabelDomainResponse(response)
	}
}

///////////////////////////////////////////////////
// Get all whitelabel domains.
// GET /whitelabel/domains

func (client *Client) GetAllWhitelabelDomains() (bool, error) {
	request := sendgrid.GetRequest(client.apiKey, "/v3/whitelabel/domains", "")
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return false, err
	} else {
		err := processEmptyResponse(response)
		if err != nil {
			return false, nil
		} else {
			return true, nil
		}
	}
}

///////////////////////////////////////////////////
// Update a whitelabel domain.
// PATCH /whitelabel/domains/{domain_id}

func (client *Client) UpdateWhitelabelDomain(id string, whitelabeldomain *WhitelabelDomain) error {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/whitelabel/domains/%s", id), "")
	request.Method = "PATCH"
	var err error
	request.Body, err = client.GetBody(whitelabeldomain)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return processEmptyResponse(response)
	}
}

///////////////////////////////////////////////////
// Get a single whitelabel domain.
// GET /whitelabel/domains/{domain_id}

func (client *Client) GetWhitelabelDomain(id string) (*WhitelabelDomain, error) {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/whitelabel/domains/%s", id), "")
	request.Method = "GET"
    //prettyPrint(request)
	response, err := sendgrid.API(request)
    //prettyPrint(response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processWhitelabelDomainResponse(response)
	}
}

///////////////////////////////////////////////////
// Get whitelabel domains from a domain name.
// GET /whitelabel/domains?domain=string

func (client *Client) GetWhitelabelDomainFromName(domain string) (*WhitelabelDomain, error) {
	request := sendgrid.GetRequest(client.apiKey, "/v3/whitelabel/domains", "")
	request.Method = "GET"
    queryParams := make(map[string]string)
    queryParams["domain"] = domain
    request.QueryParams = queryParams
    //prettyPrint(request)
	response, err := sendgrid.API(request)
    //prettyPrint(response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		wld_list,err := processWhitelabelDomainListResponse(response)
        if len(wld_list) == 0 {
            fmt.Errorf("Whitelabel Domain not found.")
            return nil, err
        }
        if len(wld_list) > 1 {
            log.Println("Warning: Got multiple Whitelabel Domain which share the same domain name")
        }
        return &wld_list[0], err
	}
}

///////////////////////////////////////////////////
// Delete a whitelabel domain.
// DELETE /whitelabel/domains/{domain_id}

func (client *Client) DeleteWhitelabelDomain(id string) error {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/whitelabel/domains/%s", id), "")
	request.Method = "DELETE"
    log.Printf("[DEBUG] Deleting whitelabel domain #%s", id)
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return processEmptyResponse(response)
	}
}

///////////////////////////////////////////////////
// Delete a whitelabel domain from domain name.
// DELETE /whitelabel/domains/{domain_id}

func (client *Client) DeleteAllWhitelabelDomainFromName(domain string) error {
	request := sendgrid.GetRequest(client.apiKey, "/v3/whitelabel/domains", "")
	request.Method = "GET"
    queryParams := make(map[string]string)
    queryParams["domain"] = domain
    request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	} 

    wld_list,err := processWhitelabelDomainListResponse(response)

    if len(wld_list) == 0 {
        fmt.Errorf("Whitelabel Domain not found.")
        return err
	}

    log.Printf("[DEBUG] Deleting %d entries", len(wld_list))

    // For each wld
    for i := 0; i < len(wld_list); i++ {
        err := client.DeleteWhitelabelDomain(strconv.Itoa(wld_list[i].Id))
    	if err != nil {
    		fmt.Println(err)
    		return err
    	}
    }
    return nil
}

func processWhitelabelDomainResponse(response *rest.Response) (*WhitelabelDomain, error) {

	err := processEmptyResponse(response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var out WhitelabelDomain
    //prettyPrint(response.Body)
	err = json.Unmarshal([]byte(response.Body), &out)
	if err != nil {
		fmt.Println(response)
		fmt.Errorf("Unmarshal whitelabel domain error: ", err)
		return nil, err
	}
	return &out, nil
}

// When the output is a list
func processWhitelabelDomainListResponse(response *rest.Response) (WhitelabelDomainList, error) {

	err := processEmptyResponse(response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var out WhitelabelDomainList
	err = json.Unmarshal([]byte(response.Body), &out)
    //prettyPrint(out)
	if err != nil {
		fmt.Println(response)
		fmt.Errorf("Unmarshal whitelabel domain error: ", err)
		return nil, err
	}
	return out, nil
}

