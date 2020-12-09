// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package probe

import (
	json "encoding/json"
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

func easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe(in *jlexer.Lexer, out *ProcessContextSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	out.ProcessCacheEntrySerializer = new(ProcessCacheEntrySerializer)
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
		case "parent":
			if in.IsNull() {
				in.Skip()
				out.Parent = nil
			} else {
				if out.Parent == nil {
					out.Parent = new(ProcessCacheEntrySerializer)
				}
				(*out.Parent).UnmarshalEasyJSON(in)
			}
		case "ancestors":
			if in.IsNull() {
				in.Skip()
				out.Ancestors = nil
			} else {
				in.Delim('[')
				if out.Ancestors == nil {
					if !in.IsDelim(']') {
						out.Ancestors = make([]*ProcessCacheEntrySerializer, 0, 8)
					} else {
						out.Ancestors = []*ProcessCacheEntrySerializer{}
					}
				} else {
					out.Ancestors = (out.Ancestors)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *ProcessCacheEntrySerializer
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(ProcessCacheEntrySerializer)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Ancestors = append(out.Ancestors, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pid":
			out.Pid = uint32(in.Uint32())
		case "tid":
			out.Tid = uint32(in.Uint32())
		case "uid":
			out.UID = uint32(in.Uint32())
		case "gid":
			out.GID = uint32(in.Uint32())
		case "user":
			out.User = string(in.String())
		case "group":
			out.Group = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "executable_container_path":
			out.ContainerPath = string(in.String())
		case "executable_path":
			out.Path = string(in.String())
		case "executable_inode":
			out.Inode = uint64(in.Uint64())
		case "executable_mount_id":
			out.MountID = uint32(in.Uint32())
		case "tty":
			out.TTY = string(in.String())
		case "fork_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ForkTime).UnmarshalJSON(data))
			}
		case "exec_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExecTime).UnmarshalJSON(data))
			}
		case "exit_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExitTime).UnmarshalJSON(data))
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
func easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe(out *jwriter.Writer, in ProcessContextSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"parent\":"
		out.RawString(prefix[1:])
		if in.Parent == nil {
			out.RawString("null")
		} else {
			(*in.Parent).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"ancestors\":"
		out.RawString(prefix)
		if in.Ancestors == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Ancestors {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"pid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Pid))
	}
	{
		const prefix string = ",\"tid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Tid))
	}
	{
		const prefix string = ",\"uid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.UID))
	}
	{
		const prefix string = ",\"gid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.GID))
	}
	if in.User != "" {
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	if in.Group != "" {
		const prefix string = ",\"group\":"
		out.RawString(prefix)
		out.String(string(in.Group))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if in.ContainerPath != "" {
		const prefix string = ",\"executable_container_path\":"
		out.RawString(prefix)
		out.String(string(in.ContainerPath))
	}
	{
		const prefix string = ",\"executable_path\":"
		out.RawString(prefix)
		out.String(string(in.Path))
	}
	{
		const prefix string = ",\"executable_inode\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Inode))
	}
	{
		const prefix string = ",\"executable_mount_id\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.MountID))
	}
	if in.TTY != "" {
		const prefix string = ",\"tty\":"
		out.RawString(prefix)
		out.String(string(in.TTY))
	}
	if true {
		const prefix string = ",\"fork_time\":"
		out.RawString(prefix)
		out.Raw((in.ForkTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"exec_time\":"
		out.RawString(prefix)
		out.Raw((in.ExecTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"exit_time\":"
		out.RawString(prefix)
		out.Raw((in.ExitTime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProcessContextSerializer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProcessContextSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProcessContextSerializer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProcessContextSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe(l, v)
}
func easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe1(in *jlexer.Lexer, out *ProcessCacheEntrySerializer) {
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
		case "pid":
			out.Pid = uint32(in.Uint32())
		case "tid":
			out.Tid = uint32(in.Uint32())
		case "uid":
			out.UID = uint32(in.Uint32())
		case "gid":
			out.GID = uint32(in.Uint32())
		case "user":
			out.User = string(in.String())
		case "group":
			out.Group = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "executable_container_path":
			out.ContainerPath = string(in.String())
		case "executable_path":
			out.Path = string(in.String())
		case "executable_inode":
			out.Inode = uint64(in.Uint64())
		case "executable_mount_id":
			out.MountID = uint32(in.Uint32())
		case "tty":
			out.TTY = string(in.String())
		case "fork_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ForkTime).UnmarshalJSON(data))
			}
		case "exec_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExecTime).UnmarshalJSON(data))
			}
		case "exit_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExitTime).UnmarshalJSON(data))
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
func easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe1(out *jwriter.Writer, in ProcessCacheEntrySerializer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"pid\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.Pid))
	}
	{
		const prefix string = ",\"tid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Tid))
	}
	{
		const prefix string = ",\"uid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.UID))
	}
	{
		const prefix string = ",\"gid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.GID))
	}
	if in.User != "" {
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	if in.Group != "" {
		const prefix string = ",\"group\":"
		out.RawString(prefix)
		out.String(string(in.Group))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if in.ContainerPath != "" {
		const prefix string = ",\"executable_container_path\":"
		out.RawString(prefix)
		out.String(string(in.ContainerPath))
	}
	{
		const prefix string = ",\"executable_path\":"
		out.RawString(prefix)
		out.String(string(in.Path))
	}
	{
		const prefix string = ",\"executable_inode\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Inode))
	}
	{
		const prefix string = ",\"executable_mount_id\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.MountID))
	}
	if in.TTY != "" {
		const prefix string = ",\"tty\":"
		out.RawString(prefix)
		out.String(string(in.TTY))
	}
	if true {
		const prefix string = ",\"fork_time\":"
		out.RawString(prefix)
		out.Raw((in.ForkTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"exec_time\":"
		out.RawString(prefix)
		out.Raw((in.ExecTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"exit_time\":"
		out.RawString(prefix)
		out.Raw((in.ExitTime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProcessCacheEntrySerializer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProcessCacheEntrySerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProcessCacheEntrySerializer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProcessCacheEntrySerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe1(l, v)
}
func easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe2(in *jlexer.Lexer, out *FileSerializer) {
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
		case "path":
			out.Path = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "container_path":
			out.ContainerPath = string(in.String())
		case "inode":
			out.Inode = uint64(in.Uint64())
		case "mode":
			out.Mode = uint32(in.Uint32())
		case "overlay_numlower":
			out.OverlayNumLower = int32(in.Int32())
		case "mount_id":
			out.MountID = uint32(in.Uint32())
		case "uid":
			out.UID = int32(in.Int32())
		case "gid":
			out.GID = int32(in.Int32())
		case "attribute_name":
			out.XAttrName = string(in.String())
		case "attribute_namespace":
			out.XAttrNamespace = string(in.String())
		case "flags":
			out.Flags = string(in.String())
		case "access_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Atime).UnmarshalJSON(data))
			}
		case "modification_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Mtime).UnmarshalJSON(data))
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
func easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe2(out *jwriter.Writer, in FileSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Path != "" {
		const prefix string = ",\"path\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Path))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.ContainerPath != "" {
		const prefix string = ",\"container_path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ContainerPath))
	}
	if in.Inode != 0 {
		const prefix string = ",\"inode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Inode))
	}
	if in.Mode != 0 {
		const prefix string = ",\"mode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.Mode))
	}
	if in.OverlayNumLower != 0 {
		const prefix string = ",\"overlay_numlower\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.OverlayNumLower))
	}
	{
		const prefix string = ",\"mount_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.MountID))
	}
	{
		const prefix string = ",\"uid\":"
		out.RawString(prefix)
		out.Int32(int32(in.UID))
	}
	{
		const prefix string = ",\"gid\":"
		out.RawString(prefix)
		out.Int32(int32(in.GID))
	}
	if in.XAttrName != "" {
		const prefix string = ",\"attribute_name\":"
		out.RawString(prefix)
		out.String(string(in.XAttrName))
	}
	if in.XAttrNamespace != "" {
		const prefix string = ",\"attribute_namespace\":"
		out.RawString(prefix)
		out.String(string(in.XAttrNamespace))
	}
	if in.Flags != "" {
		const prefix string = ",\"flags\":"
		out.RawString(prefix)
		out.String(string(in.Flags))
	}
	if true {
		const prefix string = ",\"access_time\":"
		out.RawString(prefix)
		out.Raw((in.Atime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"modification_time\":"
		out.RawString(prefix)
		out.Raw((in.Mtime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FileSerializer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FileSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FileSerializer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FileSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe2(l, v)
}
func easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe3(in *jlexer.Lexer, out *FileEventSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	out.FileSerializer = new(FileSerializer)
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
		case "destination":
			if in.IsNull() {
				in.Skip()
				out.Destination = nil
			} else {
				if out.Destination == nil {
					out.Destination = new(FileSerializer)
				}
				(*out.Destination).UnmarshalEasyJSON(in)
			}
		case "mount_id":
			out.MountID = uint32(in.Uint32())
		case "group_id":
			out.GroupID = uint32(in.Uint32())
		case "device":
			out.Device = uint32(in.Uint32())
		case "fstype":
			out.FSType = string(in.String())
		case "path":
			out.Path = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "container_path":
			out.ContainerPath = string(in.String())
		case "inode":
			out.Inode = uint64(in.Uint64())
		case "mode":
			out.Mode = uint32(in.Uint32())
		case "overlay_numlower":
			out.OverlayNumLower = int32(in.Int32())
		case "uid":
			out.UID = int32(in.Int32())
		case "gid":
			out.GID = int32(in.Int32())
		case "attribute_name":
			out.XAttrName = string(in.String())
		case "attribute_namespace":
			out.XAttrNamespace = string(in.String())
		case "flags":
			out.Flags = string(in.String())
		case "access_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Atime).UnmarshalJSON(data))
			}
		case "modification_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Mtime).UnmarshalJSON(data))
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
func easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe3(out *jwriter.Writer, in FileEventSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Destination != nil {
		const prefix string = ",\"destination\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Destination).MarshalEasyJSON(out)
	}
	if in.MountID != 0 {
		const prefix string = ",\"mount_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.MountID))
	}
	if in.GroupID != 0 {
		const prefix string = ",\"group_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.GroupID))
	}
	if in.Device != 0 {
		const prefix string = ",\"device\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.Device))
	}
	if in.FSType != "" {
		const prefix string = ",\"fstype\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FSType))
	}
	if in.Path != "" {
		const prefix string = ",\"path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Path))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.ContainerPath != "" {
		const prefix string = ",\"container_path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ContainerPath))
	}
	if in.Inode != 0 {
		const prefix string = ",\"inode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Inode))
	}
	if in.Mode != 0 {
		const prefix string = ",\"mode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.Mode))
	}
	if in.OverlayNumLower != 0 {
		const prefix string = ",\"overlay_numlower\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.OverlayNumLower))
	}
	{
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.UID))
	}
	{
		const prefix string = ",\"gid\":"
		out.RawString(prefix)
		out.Int32(int32(in.GID))
	}
	if in.XAttrName != "" {
		const prefix string = ",\"attribute_name\":"
		out.RawString(prefix)
		out.String(string(in.XAttrName))
	}
	if in.XAttrNamespace != "" {
		const prefix string = ",\"attribute_namespace\":"
		out.RawString(prefix)
		out.String(string(in.XAttrNamespace))
	}
	if in.Flags != "" {
		const prefix string = ",\"flags\":"
		out.RawString(prefix)
		out.String(string(in.Flags))
	}
	if true {
		const prefix string = ",\"access_time\":"
		out.RawString(prefix)
		out.Raw((in.Atime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"modification_time\":"
		out.RawString(prefix)
		out.Raw((in.Mtime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FileEventSerializer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FileEventSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FileEventSerializer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FileEventSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe3(l, v)
}
func easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe4(in *jlexer.Lexer, out *EventSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	out.EventContextSerializer = new(EventContextSerializer)
	out.FileEventSerializer = new(FileEventSerializer)
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
		case "evt":
			if in.IsNull() {
				in.Skip()
				out.EventContextSerializer = nil
			} else {
				if out.EventContextSerializer == nil {
					out.EventContextSerializer = new(EventContextSerializer)
				}
				(*out.EventContextSerializer).UnmarshalEasyJSON(in)
			}
		case "file":
			if in.IsNull() {
				in.Skip()
				out.FileEventSerializer = nil
			} else {
				if out.FileEventSerializer == nil {
					out.FileEventSerializer = new(FileEventSerializer)
				}
				(*out.FileEventSerializer).UnmarshalEasyJSON(in)
			}
		case "process":
			if in.IsNull() {
				in.Skip()
				out.ProcessContextSerializer = nil
			} else {
				if out.ProcessContextSerializer == nil {
					out.ProcessContextSerializer = new(ProcessContextSerializer)
				}
				(*out.ProcessContextSerializer).UnmarshalEasyJSON(in)
			}
		case "container":
			if in.IsNull() {
				in.Skip()
				out.ContainerContextSerializer = nil
			} else {
				if out.ContainerContextSerializer == nil {
					out.ContainerContextSerializer = new(ContainerContextSerializer)
				}
				(*out.ContainerContextSerializer).UnmarshalEasyJSON(in)
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
func easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe4(out *jwriter.Writer, in EventSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"evt\":"
		out.RawString(prefix[1:])
		if in.EventContextSerializer == nil {
			out.RawString("null")
		} else {
			(*in.EventContextSerializer).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"file\":"
		out.RawString(prefix)
		if in.FileEventSerializer == nil {
			out.RawString("null")
		} else {
			(*in.FileEventSerializer).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"process\":"
		out.RawString(prefix)
		if in.ProcessContextSerializer == nil {
			out.RawString("null")
		} else {
			(*in.ProcessContextSerializer).MarshalEasyJSON(out)
		}
	}
	if in.ContainerContextSerializer != nil {
		const prefix string = ",\"container\":"
		out.RawString(prefix)
		(*in.ContainerContextSerializer).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventSerializer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventSerializer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe4(l, v)
}
func easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe5(in *jlexer.Lexer, out *EventContextSerializer) {
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
		case "name":
			out.Name = string(in.String())
		case "category":
			out.Category = string(in.String())
		case "outcome":
			out.Outcome = string(in.String())
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
func easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe5(out *jwriter.Writer, in EventContextSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		out.String(string(in.Category))
	}
	{
		const prefix string = ",\"outcome\":"
		out.RawString(prefix)
		out.String(string(in.Outcome))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventContextSerializer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventContextSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventContextSerializer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventContextSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe5(l, v)
}
func easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe6(in *jlexer.Lexer, out *ContainerContextSerializer) {
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
		case "id":
			out.ID = string(in.String())
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
func easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe6(out *jwriter.Writer, in ContainerContextSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ContainerContextSerializer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ContainerContextSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA970e379EncodeGithubComDataDogDatadogAgentPkgSecurityProbe6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ContainerContextSerializer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ContainerContextSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA970e379DecodeGithubComDataDogDatadogAgentPkgSecurityProbe6(l, v)
}
