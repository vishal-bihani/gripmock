// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode encodes AddStubOK as json.
func (s AddStubOK) Encode(e *jx.Encoder) {
	switch s.Type {
	case ListIDAddStubOK:
		s.ListID.Encode(e)
	}
}

// Decode decodes AddStubOK from json.
func (s *AddStubOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AddStubOK to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Array:
		if err := s.ListID.Decode(d); err != nil {
			return err
		}
		s.Type = ListIDAddStubOK
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AddStubOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AddStubOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AddStubReq as json.
func (s AddStubReq) Encode(e *jx.Encoder) {
	switch s.Type {
	case StubListAddStubReq:
		s.StubList.Encode(e)
	case StubAddStubReq:
		s.Stub.Encode(e)
	}
}

// Decode decodes AddStubReq from json.
func (s *AddStubReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AddStubReq to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Array:
		if err := s.StubList.Decode(d); err != nil {
			return err
		}
		s.Type = StubListAddStubReq
	case jx.Object:
		if err := s.Stub.Decode(d); err != nil {
			return err
		}
		s.Type = StubAddStubReq
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AddStubReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AddStubReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ID as json.
func (s ID) Encode(e *jx.Encoder) {
	unwrapped := uuid.UUID(s)

	json.EncodeUUID(e, unwrapped)
}

// Decode decodes ID from json.
func (s *ID) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ID to nil")
	}
	var unwrapped uuid.UUID
	if err := func() error {
		v, err := json.DecodeUUID(d)
		unwrapped = v
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ID(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ID) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ID) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ListID as json.
func (s ListID) Encode(e *jx.Encoder) {
	unwrapped := []ID(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes ListID from json.
func (s *ListID) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ListID to nil")
	}
	var unwrapped []ID
	if err := func() error {
		unwrapped = make([]ID, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem ID
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ListID(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ListID) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ListID) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ID as json.
func (o OptID) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes ID from json.
func (o *OptID) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptID to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptID) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptID) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes StubInput as json.
func (o OptStubInput) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes StubInput from json.
func (o *OptStubInput) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptStubInput to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptStubInput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptStubInput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes StubOutput as json.
func (o OptStubOutput) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes StubOutput from json.
func (o *OptStubOutput) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptStubOutput to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptStubOutput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptStubOutput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes uint32 as json.
func (o OptUint32) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.UInt32(uint32(o.Value))
}

// Decode decodes uint32 from json.
func (o *OptUint32) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUint32 to nil")
	}
	o.Set = true
	v, err := d.UInt32()
	if err != nil {
		return err
	}
	o.Value = uint32(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUint32) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUint32) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SearchStubsOK as json.
func (s SearchStubsOK) Encode(e *jx.Encoder) {
	switch s.Type {
	case StubSearchStubsOK:
		s.Stub.Encode(e)
	}
}

// Decode decodes SearchStubsOK from json.
func (s *SearchStubsOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchStubsOK to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Object:
		if err := s.Stub.Decode(d); err != nil {
			return err
		}
		s.Type = StubSearchStubsOK
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SearchStubsOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SearchStubsOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SearchStubsReq as json.
func (s SearchStubsReq) Encode(e *jx.Encoder) {
	switch s.Type {
	case StubListSearchStubsReq:
		s.StubList.Encode(e)
	case StubSearchStubsReq:
		s.Stub.Encode(e)
	}
}

// Decode decodes SearchStubsReq from json.
func (s *SearchStubsReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchStubsReq to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Array:
		if err := s.StubList.Decode(d); err != nil {
			return err
		}
		s.Type = StubListSearchStubsReq
	case jx.Object:
		if err := s.Stub.Decode(d); err != nil {
			return err
		}
		s.Type = StubSearchStubsReq
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SearchStubsReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SearchStubsReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Stub) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Stub) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		e.FieldStart("service")
		e.Str(s.Service)
	}
	{
		e.FieldStart("method")
		e.Str(s.Method)
	}
	{
		if s.Input.Set {
			e.FieldStart("input")
			s.Input.Encode(e)
		}
	}
	{
		if s.Output.Set {
			e.FieldStart("output")
			s.Output.Encode(e)
		}
	}
}

var jsonFieldsNameOfStub = [5]string{
	0: "id",
	1: "service",
	2: "method",
	3: "input",
	4: "output",
}

// Decode decodes Stub from json.
func (s *Stub) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Stub to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "service":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Service = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"service\"")
			}
		case "method":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Method = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"method\"")
			}
		case "input":
			if err := func() error {
				s.Input.Reset()
				if err := s.Input.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"input\"")
			}
		case "output":
			if err := func() error {
				s.Output.Reset()
				if err := s.Output.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"output\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Stub")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000110,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfStub) {
					name = jsonFieldsNameOfStub[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Stub) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Stub) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *StubInput) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *StubInput) encodeFields(e *jx.Encoder) {
	{
		if s.Equals != nil {
			e.FieldStart("equals")
			s.Equals.Encode(e)
		}
	}
	{
		if s.Contains != nil {
			e.FieldStart("contains")
			s.Contains.Encode(e)
		}
	}
	{
		if s.Matches != nil {
			e.FieldStart("matches")
			s.Matches.Encode(e)
		}
	}
}

var jsonFieldsNameOfStubInput = [3]string{
	0: "equals",
	1: "contains",
	2: "matches",
}

// Decode decodes StubInput from json.
func (s *StubInput) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StubInput to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "equals":
			if err := func() error {
				s.Equals = nil
				var elem StubInputEquals
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Equals = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"equals\"")
			}
		case "contains":
			if err := func() error {
				s.Contains = nil
				var elem StubInputContains
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Contains = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"contains\"")
			}
		case "matches":
			if err := func() error {
				s.Matches = nil
				var elem StubInputMatches
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Matches = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"matches\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode StubInput")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *StubInput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StubInput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *StubInputContains) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *StubInputContains) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfStubInputContains = [0]string{}

// Decode decodes StubInputContains from json.
func (s *StubInputContains) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StubInputContains to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode StubInputContains")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *StubInputContains) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StubInputContains) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *StubInputEquals) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *StubInputEquals) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfStubInputEquals = [0]string{}

// Decode decodes StubInputEquals from json.
func (s *StubInputEquals) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StubInputEquals to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode StubInputEquals")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *StubInputEquals) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StubInputEquals) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *StubInputMatches) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *StubInputMatches) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfStubInputMatches = [0]string{}

// Decode decodes StubInputMatches from json.
func (s *StubInputMatches) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StubInputMatches to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode StubInputMatches")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *StubInputMatches) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StubInputMatches) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes StubList as json.
func (s StubList) Encode(e *jx.Encoder) {
	unwrapped := []Stub(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes StubList from json.
func (s *StubList) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StubList to nil")
	}
	var unwrapped []Stub
	if err := func() error {
		unwrapped = make([]Stub, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem Stub
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = StubList(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s StubList) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StubList) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *StubOutput) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *StubOutput) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("data")
		s.Data.Encode(e)
	}
	{
		e.FieldStart("error")
		e.Str(s.Error)
	}
	{
		if s.Code.Set {
			e.FieldStart("code")
			s.Code.Encode(e)
		}
	}
}

var jsonFieldsNameOfStubOutput = [3]string{
	0: "data",
	1: "error",
	2: "code",
}

// Decode decodes StubOutput from json.
func (s *StubOutput) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StubOutput to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "data":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Data.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		case "error":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Error = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		case "code":
			if err := func() error {
				s.Code.Reset()
				if err := s.Code.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode StubOutput")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfStubOutput) {
					name = jsonFieldsNameOfStubOutput[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *StubOutput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StubOutput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *StubOutputData) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *StubOutputData) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfStubOutputData = [0]string{}

// Decode decodes StubOutputData from json.
func (s *StubOutputData) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StubOutputData to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode StubOutputData")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *StubOutputData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StubOutputData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
