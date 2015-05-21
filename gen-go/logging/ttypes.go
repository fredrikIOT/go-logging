// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package logging

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type ThriftLogRecord struct {
	Name     string `thrift:"name,1" json:"name"`
	Level    int32  `thrift:"level,2" json:"level"`
	PathName string `thrift:"path_name,3" json:"path_name"`
	FileName string `thrift:"file_name,4" json:"file_name"`
	LineNo   int32  `thrift:"line_no,5" json:"line_no"`
	FuncName string `thrift:"func_name,6" json:"func_name"`
	Message  string `thrift:"message,7" json:"message"`
}

func NewThriftLogRecord() *ThriftLogRecord {
	return &ThriftLogRecord{}
}

func (p *ThriftLogRecord) GetName() string {
	return p.Name
}

func (p *ThriftLogRecord) GetLevel() int32 {
	return p.Level
}

func (p *ThriftLogRecord) GetPathName() string {
	return p.PathName
}

func (p *ThriftLogRecord) GetFileName() string {
	return p.FileName
}

func (p *ThriftLogRecord) GetLineNo() int32 {
	return p.LineNo
}

func (p *ThriftLogRecord) GetFuncName() string {
	return p.FuncName
}

func (p *ThriftLogRecord) GetMessage() string {
	return p.Message
}
func (p *ThriftLogRecord) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.ReadField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ThriftLogRecord) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *ThriftLogRecord) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Level = v
	}
	return nil
}

func (p *ThriftLogRecord) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.PathName = v
	}
	return nil
}

func (p *ThriftLogRecord) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.FileName = v
	}
	return nil
}

func (p *ThriftLogRecord) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		p.LineNo = v
	}
	return nil
}

func (p *ThriftLogRecord) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 6: %s", err)
	} else {
		p.FuncName = v
	}
	return nil
}

func (p *ThriftLogRecord) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 7: %s", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *ThriftLogRecord) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ThriftLogRecord"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ThriftLogRecord) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:name: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return fmt.Errorf("%T.name (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:name: %s", p, err)
	}
	return err
}

func (p *ThriftLogRecord) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("level", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:level: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Level)); err != nil {
		return fmt.Errorf("%T.level (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:level: %s", p, err)
	}
	return err
}

func (p *ThriftLogRecord) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("path_name", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:path_name: %s", p, err)
	}
	if err := oprot.WriteString(string(p.PathName)); err != nil {
		return fmt.Errorf("%T.path_name (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:path_name: %s", p, err)
	}
	return err
}

func (p *ThriftLogRecord) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("file_name", thrift.STRING, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:file_name: %s", p, err)
	}
	if err := oprot.WriteString(string(p.FileName)); err != nil {
		return fmt.Errorf("%T.file_name (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:file_name: %s", p, err)
	}
	return err
}

func (p *ThriftLogRecord) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("line_no", thrift.I32, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:line_no: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.LineNo)); err != nil {
		return fmt.Errorf("%T.line_no (5) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:line_no: %s", p, err)
	}
	return err
}

func (p *ThriftLogRecord) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("func_name", thrift.STRING, 6); err != nil {
		return fmt.Errorf("%T write field begin error 6:func_name: %s", p, err)
	}
	if err := oprot.WriteString(string(p.FuncName)); err != nil {
		return fmt.Errorf("%T.func_name (6) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 6:func_name: %s", p, err)
	}
	return err
}

func (p *ThriftLogRecord) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 7); err != nil {
		return fmt.Errorf("%T write field begin error 7:message: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return fmt.Errorf("%T.message (7) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 7:message: %s", p, err)
	}
	return err
}

func (p *ThriftLogRecord) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ThriftLogRecord(%+v)", *p)
}