// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

// EncodeURI encodes WhereTodoInput as URI form.
func (s *WhereTodoInput) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("title", func(e uri.Encoder) error {
		if val, ok := s.Title.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"title\"")
	}
	if err := e.EncodeField("description", func(e uri.Encoder) error {
		if val, ok := s.Description.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"description\"")
	}
	if err := e.EncodeField("priorityID", func(e uri.Encoder) error {
		if val, ok := s.PriorityID.Get(); ok {
			return e.EncodeValue(conv.IntToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"priorityID\"")
	}
	if err := e.EncodeField("statusID", func(e uri.Encoder) error {
		if val, ok := s.StatusID.Get(); ok {
			return e.EncodeValue(conv.IntToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"statusID\"")
	}
	return nil
}

var uriFieldsNameOfWhereTodoInput = [4]string{
	0: "title",
	1: "description",
	2: "priorityID",
	3: "statusID",
}

// DecodeURI decodes WhereTodoInput from URI form.
func (s *WhereTodoInput) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode WhereTodoInput to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "title":
			if err := func() error {
				var sDotTitleVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotTitleVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Title.SetTo(sDotTitleVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				var sDotDescriptionVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotDescriptionVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Description.SetTo(sDotDescriptionVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "priorityID":
			if err := func() error {
				var sDotPriorityIDVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					sDotPriorityIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.PriorityID.SetTo(sDotPriorityIDVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"priorityID\"")
			}
		case "statusID":
			if err := func() error {
				var sDotStatusIDVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					sDotStatusIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.StatusID.SetTo(sDotStatusIDVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"statusID\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode WhereTodoInput")
	}

	return nil
}
