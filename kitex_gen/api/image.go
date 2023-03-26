// Code generated by thriftgo (0.2.8). DO NOT EDIT.

package api

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type ImageRequest struct {
	Prompt         string `thrift:"Prompt,1" frugal:"1,default,string" json:"Prompt"`
	N              int32  `thrift:"N,2" frugal:"2,default,i32" json:"N"`
	Size           string `thrift:"Size,3" frugal:"3,default,string" json:"Size"`
	ResponseFormat string `thrift:"ResponseFormat,4" frugal:"4,default,string" json:"ResponseFormat"`
	User           string `thrift:"User,5" frugal:"5,default,string" json:"User"`
}

func NewImageRequest() *ImageRequest {
	return &ImageRequest{}
}

func (p *ImageRequest) InitDefault() {
	*p = ImageRequest{}
}

func (p *ImageRequest) GetPrompt() (v string) {
	return p.Prompt
}

func (p *ImageRequest) GetN() (v int32) {
	return p.N
}

func (p *ImageRequest) GetSize() (v string) {
	return p.Size
}

func (p *ImageRequest) GetResponseFormat() (v string) {
	return p.ResponseFormat
}

func (p *ImageRequest) GetUser() (v string) {
	return p.User
}
func (p *ImageRequest) SetPrompt(val string) {
	p.Prompt = val
}
func (p *ImageRequest) SetN(val int32) {
	p.N = val
}
func (p *ImageRequest) SetSize(val string) {
	p.Size = val
}
func (p *ImageRequest) SetResponseFormat(val string) {
	p.ResponseFormat = val
}
func (p *ImageRequest) SetUser(val string) {
	p.User = val
}

var fieldIDToName_ImageRequest = map[int16]string{
	1: "Prompt",
	2: "N",
	3: "Size",
	4: "ResponseFormat",
	5: "User",
}

func (p *ImageRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ImageRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ImageRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Prompt = v
	}
	return nil
}

func (p *ImageRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.N = v
	}
	return nil
}

func (p *ImageRequest) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Size = v
	}
	return nil
}

func (p *ImageRequest) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.ResponseFormat = v
	}
	return nil
}

func (p *ImageRequest) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.User = v
	}
	return nil
}

func (p *ImageRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ImageRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ImageRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Prompt", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Prompt); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ImageRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("N", thrift.I32, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.N); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *ImageRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Size", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Size); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *ImageRequest) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("ResponseFormat", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.ResponseFormat); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *ImageRequest) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("User", thrift.STRING, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.User); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *ImageRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ImageRequest(%+v)", *p)
}

func (p *ImageRequest) DeepEqual(ano *ImageRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Prompt) {
		return false
	}
	if !p.Field2DeepEqual(ano.N) {
		return false
	}
	if !p.Field3DeepEqual(ano.Size) {
		return false
	}
	if !p.Field4DeepEqual(ano.ResponseFormat) {
		return false
	}
	if !p.Field5DeepEqual(ano.User) {
		return false
	}
	return true
}

func (p *ImageRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Prompt, src) != 0 {
		return false
	}
	return true
}
func (p *ImageRequest) Field2DeepEqual(src int32) bool {

	if p.N != src {
		return false
	}
	return true
}
func (p *ImageRequest) Field3DeepEqual(src string) bool {

	if strings.Compare(p.Size, src) != 0 {
		return false
	}
	return true
}
func (p *ImageRequest) Field4DeepEqual(src string) bool {

	if strings.Compare(p.ResponseFormat, src) != 0 {
		return false
	}
	return true
}
func (p *ImageRequest) Field5DeepEqual(src string) bool {

	if strings.Compare(p.User, src) != 0 {
		return false
	}
	return true
}

type ImageResponse struct {
	Created int64                     `thrift:"Created,1" frugal:"1,default,i64" json:"Created"`
	Data    []*ImageResponseDataInner `thrift:"Data,2" frugal:"2,default,list<ImageResponseDataInner>" json:"Data"`
}

func NewImageResponse() *ImageResponse {
	return &ImageResponse{}
}

func (p *ImageResponse) InitDefault() {
	*p = ImageResponse{}
}

func (p *ImageResponse) GetCreated() (v int64) {
	return p.Created
}

func (p *ImageResponse) GetData() (v []*ImageResponseDataInner) {
	return p.Data
}
func (p *ImageResponse) SetCreated(val int64) {
	p.Created = val
}
func (p *ImageResponse) SetData(val []*ImageResponseDataInner) {
	p.Data = val
}

var fieldIDToName_ImageResponse = map[int16]string{
	1: "Created",
	2: "Data",
}

func (p *ImageResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ImageResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ImageResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.Created = v
	}
	return nil
}

func (p *ImageResponse) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Data = make([]*ImageResponseDataInner, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewImageResponseDataInner()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.Data = append(p.Data, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *ImageResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ImageResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ImageResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Created", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Created); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ImageResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Data", thrift.LIST, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Data)); err != nil {
		return err
	}
	for _, v := range p.Data {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *ImageResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ImageResponse(%+v)", *p)
}

func (p *ImageResponse) DeepEqual(ano *ImageResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Created) {
		return false
	}
	if !p.Field2DeepEqual(ano.Data) {
		return false
	}
	return true
}

func (p *ImageResponse) Field1DeepEqual(src int64) bool {

	if p.Created != src {
		return false
	}
	return true
}
func (p *ImageResponse) Field2DeepEqual(src []*ImageResponseDataInner) bool {

	if len(p.Data) != len(src) {
		return false
	}
	for i, v := range p.Data {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}

type ImageResponseDataInner struct {
	URL     string `thrift:"URL,1" frugal:"1,default,string" json:"URL"`
	B64JSON string `thrift:"B64JSON,2" frugal:"2,default,string" json:"B64JSON"`
}

func NewImageResponseDataInner() *ImageResponseDataInner {
	return &ImageResponseDataInner{}
}

func (p *ImageResponseDataInner) InitDefault() {
	*p = ImageResponseDataInner{}
}

func (p *ImageResponseDataInner) GetURL() (v string) {
	return p.URL
}

func (p *ImageResponseDataInner) GetB64JSON() (v string) {
	return p.B64JSON
}
func (p *ImageResponseDataInner) SetURL(val string) {
	p.URL = val
}
func (p *ImageResponseDataInner) SetB64JSON(val string) {
	p.B64JSON = val
}

var fieldIDToName_ImageResponseDataInner = map[int16]string{
	1: "URL",
	2: "B64JSON",
}

func (p *ImageResponseDataInner) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ImageResponseDataInner[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ImageResponseDataInner) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.URL = v
	}
	return nil
}

func (p *ImageResponseDataInner) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.B64JSON = v
	}
	return nil
}

func (p *ImageResponseDataInner) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ImageResponseDataInner"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ImageResponseDataInner) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("URL", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.URL); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ImageResponseDataInner) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("B64JSON", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.B64JSON); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *ImageResponseDataInner) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ImageResponseDataInner(%+v)", *p)
}

func (p *ImageResponseDataInner) DeepEqual(ano *ImageResponseDataInner) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.URL) {
		return false
	}
	if !p.Field2DeepEqual(ano.B64JSON) {
		return false
	}
	return true
}

func (p *ImageResponseDataInner) Field1DeepEqual(src string) bool {

	if strings.Compare(p.URL, src) != 0 {
		return false
	}
	return true
}
func (p *ImageResponseDataInner) Field2DeepEqual(src string) bool {

	if strings.Compare(p.B64JSON, src) != 0 {
		return false
	}
	return true
}

type GPTService interface {
	CreateImage(ctx context.Context, request *ImageRequest) (r *ImageResponse, err error)
}

type GPTServiceClient struct {
	c thrift.TClient
}

func NewGPTServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *GPTServiceClient {
	return &GPTServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewGPTServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *GPTServiceClient {
	return &GPTServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewGPTServiceClient(c thrift.TClient) *GPTServiceClient {
	return &GPTServiceClient{
		c: c,
	}
}

func (p *GPTServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *GPTServiceClient) CreateImage(ctx context.Context, request *ImageRequest) (r *ImageResponse, err error) {
	var _args GPTServiceCreateImageArgs
	_args.Request = request
	var _result GPTServiceCreateImageResult
	if err = p.Client_().Call(ctx, "CreateImage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type GPTServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      GPTService
}

func (p *GPTServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *GPTServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *GPTServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewGPTServiceProcessor(handler GPTService) *GPTServiceProcessor {
	self := &GPTServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("CreateImage", &gPTServiceProcessorCreateImage{handler: handler})
	return self
}
func (p *GPTServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type gPTServiceProcessorCreateImage struct {
	handler GPTService
}

func (p *gPTServiceProcessorCreateImage) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GPTServiceCreateImageArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("CreateImage", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := GPTServiceCreateImageResult{}
	var retval *ImageResponse
	if retval, err2 = p.handler.CreateImage(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing CreateImage: "+err2.Error())
		oprot.WriteMessageBegin("CreateImage", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("CreateImage", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type GPTServiceCreateImageArgs struct {
	Request *ImageRequest `thrift:"request,1" frugal:"1,default,ImageRequest" json:"request"`
}

func NewGPTServiceCreateImageArgs() *GPTServiceCreateImageArgs {
	return &GPTServiceCreateImageArgs{}
}

func (p *GPTServiceCreateImageArgs) InitDefault() {
	*p = GPTServiceCreateImageArgs{}
}

var GPTServiceCreateImageArgs_Request_DEFAULT *ImageRequest

func (p *GPTServiceCreateImageArgs) GetRequest() (v *ImageRequest) {
	if !p.IsSetRequest() {
		return GPTServiceCreateImageArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *GPTServiceCreateImageArgs) SetRequest(val *ImageRequest) {
	p.Request = val
}

var fieldIDToName_GPTServiceCreateImageArgs = map[int16]string{
	1: "request",
}

func (p *GPTServiceCreateImageArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *GPTServiceCreateImageArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GPTServiceCreateImageArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GPTServiceCreateImageArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = NewImageRequest()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GPTServiceCreateImageArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("CreateImage_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GPTServiceCreateImageArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GPTServiceCreateImageArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GPTServiceCreateImageArgs(%+v)", *p)
}

func (p *GPTServiceCreateImageArgs) DeepEqual(ano *GPTServiceCreateImageArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *GPTServiceCreateImageArgs) Field1DeepEqual(src *ImageRequest) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

type GPTServiceCreateImageResult struct {
	Success *ImageResponse `thrift:"success,0,optional" frugal:"0,optional,ImageResponse" json:"success,omitempty"`
}

func NewGPTServiceCreateImageResult() *GPTServiceCreateImageResult {
	return &GPTServiceCreateImageResult{}
}

func (p *GPTServiceCreateImageResult) InitDefault() {
	*p = GPTServiceCreateImageResult{}
}

var GPTServiceCreateImageResult_Success_DEFAULT *ImageResponse

func (p *GPTServiceCreateImageResult) GetSuccess() (v *ImageResponse) {
	if !p.IsSetSuccess() {
		return GPTServiceCreateImageResult_Success_DEFAULT
	}
	return p.Success
}
func (p *GPTServiceCreateImageResult) SetSuccess(x interface{}) {
	p.Success = x.(*ImageResponse)
}

var fieldIDToName_GPTServiceCreateImageResult = map[int16]string{
	0: "success",
}

func (p *GPTServiceCreateImageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GPTServiceCreateImageResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GPTServiceCreateImageResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GPTServiceCreateImageResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewImageResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GPTServiceCreateImageResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("CreateImage_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GPTServiceCreateImageResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *GPTServiceCreateImageResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GPTServiceCreateImageResult(%+v)", *p)
}

func (p *GPTServiceCreateImageResult) DeepEqual(ano *GPTServiceCreateImageResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *GPTServiceCreateImageResult) Field0DeepEqual(src *ImageResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}