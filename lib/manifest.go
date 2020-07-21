// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    manifest, err := UnmarshalManifest(bytes)
//    bytes, err = manifest.Marshal()

package main

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalManifest(data []byte) (Manifest, error) {
	var r Manifest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Manifest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Manifest struct {
	Empty        *StringOrArrayOfStrings                          `json:"##"`                    // A comment.
	Comment      *StringOrArrayOfStrings                          `json:"_comment"`              // Deprecated. Use ## instead.
	Architecture *ManifestArchitecture                            `json:"architecture,omitempty"`
	Autoupdate   *Autoupdate                                      `json:"autoupdate,omitempty"`  
	Bin          *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings `json:"bin"`                   
	Checkver     *Checkver                                        `json:"checkver"`              
	Cookie       map[string]interface{}                           `json:"cookie,omitempty"`      // Undocumented: Found at https://github.com/se35710/scoop-java/search?l=JSON&q=cookie
	Depends      *StringOrArrayOfStrings                          `json:"depends"`               
	Description  *string                                          `json:"description,omitempty"` 
	EnvAddPath   *StringOrArrayOfStrings                          `json:"env_add_path"`          
	EnvSet       map[string]interface{}                           `json:"env_set,omitempty"`     
	ExtractDir   *StringOrArrayOfStrings                          `json:"extract_dir"`           
	ExtractTo    *StringOrArrayOfStrings                          `json:"extract_to"`            
	Hash         *Hash                                            `json:"hash"`                  
	Homepage     *string                                          `json:"homepage,omitempty"`    
	Innosetup    *bool                                            `json:"innosetup,omitempty"`   // True if the installer InnoSetup based. Found in; https://github.com/lukesampson/scoop/search?l=JSON&q=innosetup
	Installer    *Installer                                       `json:"installer,omitempty"`   
	License      *LicenseUnion                                    `json:"license"`               
	MSI          *StringOrArrayOfStrings                          `json:"msi"`                   // Deprecated
	Notes        *StringOrArrayOfStrings                          `json:"notes"`                 
	Persist      *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings `json:"persist"`               
	PostInstall  *StringOrArrayOfStrings                          `json:"post_install"`          
	PreInstall   *StringOrArrayOfStrings                          `json:"pre_install"`           
	Psmodule     *Psmodule                                        `json:"psmodule,omitempty"`    
	Shortcuts    [][]string                                       `json:"shortcuts,omitempty"`   
	Suggest      *Suggest                                         `json:"suggest,omitempty"`     
	Uninstaller  *Uninstaller                                     `json:"uninstaller,omitempty"` 
	URL          *StringOrArrayOfStrings                          `json:"url"`                   
	Version      string                                           `json:"version"`               
}

type ManifestArchitecture struct {
	The32Bit *The32BitClass `json:"32bit,omitempty"`
	The64Bit *The32BitClass `json:"64bit,omitempty"`
}

type The32BitClass struct {
	Bin         *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings `json:"bin"`                  
	Checkver    *Checkver                                        `json:"checkver"`             
	EnvAddPath  *StringOrArrayOfStrings                          `json:"env_add_path"`         
	EnvSet      map[string]interface{}                           `json:"env_set,omitempty"`    
	ExtractDir  *StringOrArrayOfStrings                          `json:"extract_dir"`          
	Hash        *Hash                                            `json:"hash"`                 
	Installer   *Installer                                       `json:"installer,omitempty"`  
	MSI         *StringOrArrayOfStrings                          `json:"msi"`                  // Deprecated
	PostInstall *StringOrArrayOfStrings                          `json:"post_install"`         
	PreInstall  *StringOrArrayOfStrings                          `json:"pre_install"`          
	Shortcuts   [][]string                                       `json:"shortcuts,omitempty"`  
	Uninstaller *Uninstaller                                     `json:"uninstaller,omitempty"`
	URL         *StringOrArrayOfStrings                          `json:"url"`                  
}

type CheckverClass struct {
	Github    *string `json:"github,omitempty"`   
	Jp        *string `json:"jp,omitempty"`       // Same as 'jsonpath'
	Jsonpath  *string `json:"jsonpath,omitempty"` 
	Re        *string `json:"re,omitempty"`       // Same as 'regex'
	Regex     *string `json:"regex,omitempty"`    
	Replace   *string `json:"replace,omitempty"`  // Allows rearrange the regexp matches
	Reverse   *bool   `json:"reverse,omitempty"`  // Reverse the order of regex matches
	URL       *string `json:"url,omitempty"`      
	Useragent *string `json:"useragent,omitempty"`
	Xpath     *string `json:"xpath,omitempty"`    
}

type Installer struct {
	Comment *string                 `json:"_comment,omitempty"`// Undocumented: only used in scoop-extras/oraclejdk* and scoop-extras/appengine-go
	Args    *StringOrArrayOfStrings `json:"args"`              
	File    *string                 `json:"file,omitempty"`    
	Keep    *bool                   `json:"keep,omitempty"`    
	Script  *StringOrArrayOfStrings `json:"script"`            
}

type Uninstaller struct {
	Args   *StringOrArrayOfStrings `json:"args"`          
	File   *string                 `json:"file,omitempty"`
	Script *StringOrArrayOfStrings `json:"script"`        
}

type Autoupdate struct {
	Architecture *AutoupdateArchitecture `json:"architecture,omitempty"`
	ExtractDir   *string                 `json:"extract_dir,omitempty"` 
	Hash         *HashExtraction         `json:"hash,omitempty"`        
	Note         *StringOrArrayOfStrings `json:"note"`                  
	URL          *string                 `json:"url,omitempty"`         
}

type AutoupdateArchitecture struct {
	The32Bit *The32Bit `json:"32bit,omitempty"`
	The64Bit *The64Bit `json:"64bit,omitempty"`
}

type The32Bit struct {
	ExtractDir *string         `json:"extract_dir,omitempty"`
	Hash       *HashExtraction `json:"hash,omitempty"`       
	URL        *string         `json:"url,omitempty"`        
}

type HashExtraction struct {
	Find     *string `json:"find,omitempty"`    // Same as 'regex'
	Jp       *string `json:"jp,omitempty"`      // Same as 'jsonpath'
	Jsonpath *string `json:"jsonpath,omitempty"`
	Mode     *Mode   `json:"mode,omitempty"`    
	Regex    *string `json:"regex,omitempty"`   
	Type     *Type   `json:"type,omitempty"`    // Deprecated, hash type is determined automatically
	URL      *string `json:"url,omitempty"`     
	Xpath    *string `json:"xpath,omitempty"`   
}

type The64Bit struct {
	ExtractDir *string         `json:"extract_dir,omitempty"`
	Hash       *HashExtraction `json:"hash,omitempty"`       
	URL        *string         `json:"url,omitempty"`        
}

type LicenseClass struct {
	Identifier string  `json:"identifier"`   
	URL        *string `json:"url,omitempty"`
}

type Psmodule struct {
	Name *string `json:"name,omitempty"`
}

type Suggest struct {
}

type Mode string
const (
	Download Mode = "download"
	Extract Mode = "extract"
	Fosshub Mode = "fosshub"
	JSON Mode = "json"
	Metalink Mode = "metalink"
	RDF Mode = "rdf"
	Sourceforge Mode = "sourceforge"
	Xpath Mode = "xpath"
)

// Deprecated, hash type is determined automatically
type Type string
const (
	Md5 Type = "md5"
	Sha1 Type = "sha1"
	Sha256 Type = "sha256"
	Sha512 Type = "sha512"
)

type StringOrArrayOfStringsOrAnArrayOfArrayOfStrings struct {
	String     *string
	UnionArray []StringOrArrayOfStrings
}

func (x *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.UnionArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.UnionArray != nil, x.UnionArray, false, nil, false, nil, false, nil, false)
}

// A comment.
//
// Deprecated. Use ## instead.
//
// Deprecated
type StringOrArrayOfStrings struct {
	String      *string
	StringArray []string
}

func (x *StringOrArrayOfStrings) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StringOrArrayOfStrings) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

type Checkver struct {
	CheckverClass *CheckverClass
	String        *string
}

func (x *Checkver) UnmarshalJSON(data []byte) error {
	x.CheckverClass = nil
	var c CheckverClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CheckverClass = &c
	}
	return nil
}

func (x *Checkver) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.CheckverClass != nil, x.CheckverClass, false, nil, false, nil, false)
}

type Hash struct {
	String      *string
	StringArray []string
}

func (x *Hash) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Hash) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

type LicenseUnion struct {
	LicenseClass *LicenseClass
	String       *string
}

func (x *LicenseUnion) UnmarshalJSON(data []byte) error {
	x.LicenseClass = nil
	var c LicenseClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.LicenseClass = &c
	}
	return nil
}

func (x *LicenseUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.LicenseClass != nil, x.LicenseClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
