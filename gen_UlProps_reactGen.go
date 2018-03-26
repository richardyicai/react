// Code generated by reactGen. DO NOT EDIT.

package react

// UlProps defines the properties for the <ul> element
type UlProps struct {
	AriaSet
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Role  string
	Style *CSS
}

func (u *UlProps) assign(v *_UlProps) {

	if u.AriaSet != nil {
		for dk, dv := range u.AriaSet {
			v.o.Set("aria-"+dk, dv)
		}
	}

	v.ClassName = u.ClassName

	v.DangerouslySetInnerHTML = u.DangerouslySetInnerHTML

	if u.DataSet != nil {
		for dk, dv := range u.DataSet {
			v.o.Set("data-"+dk, dv)
		}
	}

	if u.ID != "" {
		v.ID = u.ID
	}

	if u.Key != "" {
		v.Key = u.Key
	}

	if u.OnChange != nil {
		v.o.Set("onChange", u.OnChange.OnChange)
	}

	if u.OnClick != nil {
		v.o.Set("onClick", u.OnClick.OnClick)
	}

	if u.Ref != nil {
		v.o.Set("ref", u.Ref.Ref)
	}

	v.Role = u.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = u.Style.hack()

}
