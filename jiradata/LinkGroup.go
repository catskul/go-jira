package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata -overwrite schemas/SearchResults.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// LinkGroup defined from schema:
// {
//   "title": "Link Group",
//   "type": "object",
//   "properties": {
//     "groups": {
//       "type": "array",
//       "items": {
//         "$ref": "#/definitions/link-group"
//       }
//     },
//     "header": {
//       "$ref": "#/definitions/simple-link"
//     },
//     "id": {
//       "type": "string"
//     },
//     "links": {
//       "type": "array",
//       "items": {
//         "$ref": "#/definitions/simple-link"
//       }
//     },
//     "styleClass": {
//       "type": "string"
//     },
//     "weight": {
//       "type": "integer"
//     }
//   }
// }
type LinkGroup struct {
	Groups     []interface{} `json:"groups,omitempty" yaml:"groups,omitempty"`
	Header     interface{}   `json:"header,omitempty" yaml:"header,omitempty"`
	ID         string        `json:"id,omitempty" yaml:"id,omitempty"`
	Links      []interface{} `json:"links,omitempty" yaml:"links,omitempty"`
	StyleClass string        `json:"styleClass,omitempty" yaml:"styleClass,omitempty"`
	Weight     int           `json:"weight,omitempty" yaml:"weight,omitempty"`
}