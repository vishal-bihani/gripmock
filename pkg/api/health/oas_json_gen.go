// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Encode implements json.Marshaler.
func (s *MessageOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MessageOK) encodeFields(e *jx.Encoder) {
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfMessageOK = [1]string{
	0: "message",
}

// Decode decodes MessageOK from json.
func (s *MessageOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MessageOK to nil")
	}
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MessageOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MessageOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MessageOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
