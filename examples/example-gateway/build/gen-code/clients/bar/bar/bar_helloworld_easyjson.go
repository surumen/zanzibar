// Code generated by zanzibar
// @generated
// Checksum : L0w0pJM1eMig13h+mkQPNQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld(in *jlexer.Lexer, out *Bar_HelloWorld_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
			}
		case "barException":
			if in.IsNull() {
				in.Skip()
				out.BarException = nil
			} else {
				if out.BarException == nil {
					out.BarException = new(BarException)
				}
				easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in, &*out.BarException)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld(out *jwriter.Writer, in Bar_HelloWorld_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Success))
		}
	}
	if in.BarException != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"barException\":")
		if in.BarException == nil {
			out.RawString("null")
		} else {
			easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out, *in.BarException)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_HelloWorld_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_HelloWorld_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_HelloWorld_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_HelloWorld_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld(l, v)
}
func easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in *jlexer.Lexer, out *BarException) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
}
func easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out *jwriter.Writer, in BarException) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	out.RawByte('}')
}
func easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld1(in *jlexer.Lexer, out *Bar_HelloWorld_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld1(out *jwriter.Writer, in Bar_HelloWorld_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_HelloWorld_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_HelloWorld_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson687c569aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_HelloWorld_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_HelloWorld_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson687c569aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarHelloWorld1(l, v)
}
